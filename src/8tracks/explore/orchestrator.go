package explore

import (
	"context"

	"github.com/pkg/errors"
	set "gopkg.in/fatih/set.v0"
)

func (s *service) loadPlaylistID(ctx context.Context, tagNames []string) ([]string, error) {
	if len(tagNames) == 0 {
		return []string{}, nil
	}

	commonPlaylistIDSet, err := s.tagSvc.LoadTagPlaylistID(ctx, tagNames[0])
	if err != nil {
		return nil, errors.Wrap(err, "couldn't load tag playlist id")
	}

	for i := 1; i < len(tagNames); i++ {
		playlistIDSet, err := s.tagSvc.LoadTagPlaylistID(ctx, tagNames[i])
		if err != nil {
			return nil, errors.Wrap(err, "couldn't load tag playlist id")
		}
		commonPlaylistIDSet = set.Intersection(commonPlaylistIDSet, playlistIDSet).(*set.Set)
	}

	return set.StringSlice(commonPlaylistIDSet), nil
}

func (s *service) loadTagID(ctx context.Context, playlistIDs []string) ([]string, error) {
	if len(playlistIDs) == 0 {
		return []string{}, nil
	}

	allTagIDSet, err := s.tagSvc.LoadPlaylistTag(ctx, playlistIDs[0])
	if err != nil {
		return nil, errors.Wrap(err, "couldn't load playlist tag")
	}

	for i := 1; i < len(playlistIDs); i++ {
		tagIDSet, err := s.tagSvc.LoadPlaylistTag(ctx, playlistIDs[i])
		if err != nil {
			return nil, errors.Wrap(err, "couldn't load playlist tag")
		}
		allTagIDSet = set.Union(allTagIDSet, tagIDSet).(*set.Set)
	}

	return set.StringSlice(allTagIDSet), nil
}
