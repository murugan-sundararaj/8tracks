package playlists

import (
	"8tracks/lib/econst"
	"8tracks/tags"
	"context"

	"github.com/go-kit/kit/log/level"

	"gopkg.in/fatih/set.v0"

	kitlog "github.com/go-kit/kit/log"
	"github.com/pkg/errors"
)

// Service is the interface that prvoides playlist methods
type Service interface {
	CreatePlaylist(ctx context.Context, r *CreatePlaylistRequest) (*CreatePlaylistResponse, error)
	LoadPlaylist(ctx context.Context, r *LoadPlaylistRequest) (*LoadPlaylistResponse, error)
	UpsertPlaylist(ctx context.Context, r *UpsertPlaylistRequest) (*UpsertPlaylistResponse, error)
	UpdatePlaylistName(ctx context.Context, r *UpdatePlaylistNameRequest) (*UpdatePlaylistNameResponse, error)
	RemovePlaylist(ctx context.Context, r *RemovePlaylistRequest) (*RemovePlaylistResponse, error)
	AddTrack(ctx context.Context, r *AddTrackRequest) (*AddTrackResponse, error)
	RemoveTrack(ctx context.Context, r *RemoveTrackRequest) (*RemoveTrackResponse, error)
	Plays(ctx context.Context, r *PlaysRequest) (*PlaysResponse, error)
	Likes(ctx context.Context, r *LikesRequest) (*LikesResponse, error)
	Dislikes(ctx context.Context, r *DislikesRequest) (*DislikesResponse, error)
}

type service struct {
	logger kitlog.Logger
	d      *DAL
	tagSvc tags.Service
}

// NewService creates and return a new playlist service
func NewService(tagSvc tags.Service, logger kitlog.Logger, d *DAL) Service {
	return &service{
		tagSvc: tagSvc,
		logger: logger,
		d:      d,
	}
}

func (s *service) CreatePlaylist(ctx context.Context, r *CreatePlaylistRequest) (*CreatePlaylistResponse, error) {
	playlistID, err := s.d.createPlaylist(ctx, r.Playlist)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't create playlist")
	}
	// update the tag service
	for _, tag := range r.Tags {
		resp, err := s.tagSvc.AssignTagToPlaylist(ctx, &tags.AssignTagToPlaylistRequest{
			TagID:      tag.TagID,
			PlaylistID: playlistID,
		})
		if err != nil {
			return nil, errors.Wrap(err, "couldn't assign tag to playlist")
		}
		if resp.Err != "" {
			// can't ignore this error as the tags we are trying to add are invalid
			return &CreatePlaylistResponse{
				Err: resp.Err,
			}, nil
		}
	}

	return &CreatePlaylistResponse{
		PlaylistID: playlistID,
	}, nil
}

func (s *service) LoadPlaylist(ctx context.Context, r *LoadPlaylistRequest) (*LoadPlaylistResponse, error) {
	playlists, err := s.d.loadPlaylist(ctx, r.PlaylistIDs, r.PlaylistNames)
	if err != nil {
		if err == ErrInvalidPlaylist {
			return &LoadPlaylistResponse{
				Err: ErrInvalidPlaylist.Error(),
			}, nil
		}
		return nil, errors.Wrap(err, "couldn't load playlist")
	}
	// fetch the tag info
	for _, playlist := range playlists {
		tagIDSet, err := s.tagSvc.LoadPlaylistTag(ctx, playlist.PlaylistID)
		if err != nil {
			return nil, errors.Wrap(err, "couldn't load playlist tag")
		}

		tag, err := s.tagSvc.LoadTag(ctx, &tags.LoadTagRequest{
			TagIDs: set.StringSlice(tagIDSet),
		})
		if err != nil {
			return nil, errors.Wrap(err, "couldn't load tag")
		}
		playlist.Tags = tag.Tags
	}

	return &LoadPlaylistResponse{
		Playlists: playlists,
	}, nil
}

func (s *service) removeTagReference(ctx context.Context, playlistID string) error {
	playlistResp, err := s.LoadPlaylist(ctx, &LoadPlaylistRequest{
		PlaylistIDs: []string{playlistID},
	})
	if err != nil {
		return err
	}

	if len(playlistResp.Playlists) == 1 {
		for _, tag := range playlistResp.Playlists[0].Tags {
			resp, err := s.tagSvc.UnAssignTagFromPlaylist(ctx, &tags.UnAssignTagFromPlaylistRequest{
				TagID:      tag.TagID,
				PlaylistID: playlistID,
			})
			if err != nil {
				return errors.Wrap(err, "couldn't unassign tag from playlist")
			}
			if resp.Err != "" {
				level.Warn(s.logger).Log(
					"request_id", ctx.Value(econst.RequestID),
					"remove_tag_reference", "failed",
					"err", playlistResp.Err,
				)
				// ignoring the error here as we are trying to remove the tags are invalid
			}
		}
	}
	return nil
}

func (s *service) addTagReference(ctx context.Context, p *Playlist) error {
	for _, tag := range p.Tags {
		resp, err := s.tagSvc.AssignTagToPlaylist(ctx, &tags.AssignTagToPlaylistRequest{
			TagID:      tag.TagID,
			PlaylistID: p.PlaylistID,
		})
		if err != nil {
			return errors.Wrap(err, "couldn't assign tag to playlist")
		}
		if resp.Err != "" {
			return errors.New(resp.Err)
		}
	}
	return nil
}

func (s *service) UpsertPlaylist(ctx context.Context, r *UpsertPlaylistRequest) (*UpsertPlaylistResponse, error) {
	if err := s.removeTagReference(ctx, r.Playlist.PlaylistID); err != nil {
		return nil, err
	}

	if err := s.d.upsertPlaylist(ctx, r.Playlist); err != nil {
		return nil, errors.Wrap(err, "couldn't upsert playlist")
	}

	err := s.addTagReference(ctx, r.Playlist)
	if err != nil {
		if err.Error() == tags.ErrInvalidTag.Error() {
			return &UpsertPlaylistResponse{
				Err: err.Error(),
			}, nil
		}
		return nil, err
	}

	return &UpsertPlaylistResponse{}, nil
}

func (s *service) UpdatePlaylistName(ctx context.Context, r *UpdatePlaylistNameRequest) (*UpdatePlaylistNameResponse, error) {
	if err := s.d.updatePlaylistName(ctx, r.PlaylistID, r.PlaylistName); err != nil {
		if err == ErrInvalidPlaylist {
			return &UpdatePlaylistNameResponse{
				Err: ErrInvalidPlaylist.Error(),
			}, nil
		}
		return nil, errors.Wrap(err, "couldn't update playlist name")
	}
	return &UpdatePlaylistNameResponse{}, nil
}

func (s *service) RemovePlaylist(ctx context.Context, r *RemovePlaylistRequest) (*RemovePlaylistResponse, error) {
	if err := s.removeTagReference(ctx, r.PlaylistID); err != nil {
		return nil, err
	}

	if err := s.d.removePlaylist(ctx, r.PlaylistID); err != nil {
		if err == ErrInvalidPlaylist {
			return &RemovePlaylistResponse{
				Err: ErrInvalidPlaylist.Error(),
			}, nil
		}
		return nil, errors.Wrap(err, "couldn't remove playlist")
	}
	return &RemovePlaylistResponse{}, nil
}

func (s *service) AddTrack(ctx context.Context, r *AddTrackRequest) (*AddTrackResponse, error) {
	if err := s.d.addTrack(ctx, r.PlaylistID, r.Track); err != nil {
		if err == ErrInvalidPlaylist {
			return &AddTrackResponse{
				Err: ErrInvalidPlaylist.Error(),
			}, nil
		}
		return nil, errors.Wrap(err, "couldn't add track")
	}
	return &AddTrackResponse{}, nil
}

func (s *service) RemoveTrack(ctx context.Context, r *RemoveTrackRequest) (*RemoveTrackResponse, error) {
	if err := s.d.removeTrack(ctx, r.PlaylistID, r.TrackID); err != nil {
		if err == ErrInvalidPlaylist {
			return &RemoveTrackResponse{
				Err: ErrInvalidPlaylist.Error(),
			}, nil
		}
		return nil, errors.Wrap(err, "couldn't remove track")
	}
	return &RemoveTrackResponse{}, nil
}

func (s *service) Plays(ctx context.Context, r *PlaysRequest) (*PlaysResponse, error) {
	if err := s.d.plays(ctx, r.PlaylistID); err != nil {
		if err == ErrInvalidPlaylist {
			return &PlaysResponse{
				Err: ErrInvalidPlaylist.Error(),
			}, nil
		}
		return nil, errors.Wrap(err, "couldn't increment plays count")
	}
	return &PlaysResponse{}, nil
}

func (s *service) Likes(ctx context.Context, r *LikesRequest) (*LikesResponse, error) {
	if err := s.d.likes(ctx, r.PlaylistID); err != nil {
		if err == ErrInvalidPlaylist {
			return &LikesResponse{
				Err: ErrInvalidPlaylist.Error(),
			}, nil
		}
		return nil, errors.Wrap(err, "couldn't increment likes count")
	}
	return &LikesResponse{}, nil
}

func (s *service) Dislikes(ctx context.Context, r *DislikesRequest) (*DislikesResponse, error) {
	if err := s.d.dislikes(ctx, r.PlaylistID); err != nil {
		if err == ErrInvalidPlaylist {
			return &DislikesResponse{
				Err: ErrInvalidPlaylist.Error(),
			}, nil
		}
		return nil, errors.Wrap(err, "couldn't decrement likes count")
	}
	return &DislikesResponse{}, nil
}
