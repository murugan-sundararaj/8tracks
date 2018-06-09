package explore

import (
	"8tracks/playlists"
	"8tracks/tags"
	"context"

	kitlog "github.com/go-kit/kit/log"
	"github.com/pkg/errors"
)

// Service is the interface that prvoides explore methods
type Service interface {
	Explore(ctx context.Context, r *ExploreRequest) (*ExploreResponse, error)
}

type service struct {
	logger      kitlog.Logger
	tagSvc      tags.Service
	playlistSvc playlists.Service
}

// NewService creates and return a new explore service
func NewService(logger kitlog.Logger, tagSvc tags.Service, playlistSvc playlists.Service) Service {
	return &service{
		logger:      logger,
		tagSvc:      tagSvc,
		playlistSvc: playlistSvc,
	}
}

func (s *service) Explore(ctx context.Context, r *ExploreRequest) (*ExploreResponse, error) {
	playlistIDs, err := s.loadPlaylistID(ctx, r.TagNames)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't load playlist ids")
	}
	if len(playlistIDs) == 0 {
		return &ExploreResponse{}, nil
	}

	tagIDs, err := s.loadTagID(ctx, playlistIDs)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't load tag ids")
	}

	tagResp, err := s.tagSvc.LoadTag(ctx, &tags.LoadTagRequest{
		TagIDs: tagIDs,
	})
	if err != nil {
		return nil, errors.Wrap(err, "couldn't load tag")
	}

	playlistResp, err := s.playlistSvc.LoadPlaylist(ctx, &playlists.LoadPlaylistRequest{PlaylistIDs: playlistIDs})
	if err != nil {
		return nil, errors.Wrap(err, "couldn't load playlist")
	}

	return &ExploreResponse{
		Tags:      tagResp.Tags,
		Playlists: defaultRankingOrder(playlistResp.Playlists),
	}, nil
}
