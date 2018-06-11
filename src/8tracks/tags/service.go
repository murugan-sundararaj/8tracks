package tags

import (
	"context"

	kitlog "github.com/go-kit/kit/log"
	"github.com/pkg/errors"
	set "gopkg.in/fatih/set.v0"
)

// Service is the interface that prvoides tag methods
type Service interface {
	CreateTag(ctx context.Context, r *CreateTagRequest) (*CreateTagResponse, error)
	LoadTag(ctx context.Context, r *LoadTagRequest) (*LoadTagResponse, error)
	UpsertTag(ctx context.Context, r *UpsertTagRequest) (*UpsertTagResponse, error)
	UpdateTag(ctx context.Context, r *UpdateTagRequest) (*UpdateTagResponse, error)
	RemoveTag(ctx context.Context, r *RemoveTagRequest) (*RemoveTagResponse, error)
	LoadTagTypes(ctx context.Context, r *LoadTagTypesRequest) (*LoadTagTypesResponse, error)
	AssignTagToPlaylist(ctx context.Context, r *AssignTagToPlaylistRequest) (*AssignTagToPlaylistResponse, error)
	UnAssignTagFromPlaylist(ctx context.Context, r *UnAssignTagFromPlaylistRequest) (*UnAssignTagFromPlaylistResponse, error)

	LoadTagPlaylistID(ctx context.Context, tagName string) (*set.Set, error)
	LoadPlaylistTag(ctx context.Context, playlistID string) (*set.Set, error)
}

type service struct {
	logger kitlog.Logger
	d      *DAL
}

// NewService creates and return a new tag service
func NewService(logger kitlog.Logger, d *DAL) Service {
	return &service{
		logger: logger,
		d:      d,
	}
}

func (s *service) CreateTag(ctx context.Context, r *CreateTagRequest) (*CreateTagResponse, error) {
	tagID, err := s.d.createTag(ctx, r.Tag)
	if err != nil {
		if err == ErrTagNameExist {
			return &CreateTagResponse{
				Err: err.Error(),
			}, nil
		}
		return nil, errors.Wrap(err, "couldn't create tag")
	}
	return &CreateTagResponse{
		TagID: tagID,
	}, nil
}

func (s *service) LoadTag(ctx context.Context, r *LoadTagRequest) (*LoadTagResponse, error) {
	tags, err := s.d.loadTag(ctx, r.TagIDs, r.TagNames)
	if err != nil {
		if err == ErrInvalidTag {
			return &LoadTagResponse{
				Err: ErrInvalidTag.Error(),
			}, nil
		}
		return nil, errors.Wrap(err, "couldn't load tag")
	}
	return &LoadTagResponse{
		Tags: tags,
	}, nil
}

func (s *service) UpsertTag(ctx context.Context, r *UpsertTagRequest) (*UpsertTagResponse, error) {
	if err := s.d.upsertTag(ctx, r.Tag); err != nil {
		return nil, errors.Wrap(err, "couldn't upsert tag")
	}
	return &UpsertTagResponse{}, nil
}

func (s *service) UpdateTag(ctx context.Context, r *UpdateTagRequest) (*UpdateTagResponse, error) {
	if err := s.d.updateTag(ctx, r.Tag); err != nil {
		if err == ErrInvalidTag {
			return &UpdateTagResponse{
				Err: ErrInvalidTag.Error(),
			}, nil
		}
		return nil, errors.Wrap(err, "couldn't update tag")
	}
	return &UpdateTagResponse{}, nil
}

func (s *service) RemoveTag(ctx context.Context, r *RemoveTagRequest) (*RemoveTagResponse, error) {
	if err := s.d.removeTag(ctx, r.TagID); err != nil {
		if err == ErrInvalidTag {
			return &RemoveTagResponse{
				Err: ErrInvalidTag.Error(),
			}, nil
		}
		return nil, errors.Wrap(err, "couldn't remove tag")
	}
	return &RemoveTagResponse{}, nil
}

func (s *service) LoadTagTypes(ctx context.Context, r *LoadTagTypesRequest) (*LoadTagTypesResponse, error) {
	return &LoadTagTypesResponse{
		TagTypes: loadTypes(),
	}, nil
}

func (s *service) AssignTagToPlaylist(ctx context.Context, r *AssignTagToPlaylistRequest) (*AssignTagToPlaylistResponse, error) {
	if err := s.d.tagPlayList(ctx, r.TagID, r.PlaylistID); err != nil {
		if err == ErrInvalidTag {
			return &AssignTagToPlaylistResponse{
				Err: ErrInvalidTag.Error(),
			}, nil
		}
		return nil, errors.Wrap(err, "couldn't tag playlist")
	}
	return &AssignTagToPlaylistResponse{}, nil
}

func (s *service) UnAssignTagFromPlaylist(ctx context.Context, r *UnAssignTagFromPlaylistRequest) (*UnAssignTagFromPlaylistResponse, error) {
	if err := s.d.unTagPlayList(ctx, r.TagID, r.PlaylistID); err != nil {
		if err == ErrInvalidTag {
			return &UnAssignTagFromPlaylistResponse{
				Err: ErrInvalidTag.Error(),
			}, nil
		}
		return nil, errors.Wrap(err, "couldn't untag playlist")
	}
	return &UnAssignTagFromPlaylistResponse{}, nil
}

func (s *service) LoadTagPlaylistID(ctx context.Context, tagName string) (*set.Set, error) {
	return s.d.loadTagPlaylistID(ctx, tagName)
}

func (s *service) LoadPlaylistTag(ctx context.Context, playlistID string) (*set.Set, error) {
	return s.d.loadPlaylistTag(ctx, playlistID)
}
