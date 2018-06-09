package explore

import (
	"8tracks/playlists"
	"8tracks/tags"
	"context"
	"os"
	"testing"

	kitlog "github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

func TestExplore(t *testing.T) {
	logger := kitlog.NewLogfmtLogger(kitlog.NewSyncWriter(os.Stderr))
	logger = level.NewFilter(logger, level.AllowAll())
	logger = kitlog.With(logger, "ts", kitlog.DefaultTimestamp, "caller", kitlog.DefaultCaller)

	tagLogger := kitlog.With(logger, "service", "tag")
	tagDAL := tags.NewDAL()
	tagSvc := tags.NewService(tagLogger, tagDAL)

	playlistLogger := kitlog.With(logger, "service", "playlists")
	playlistDAL := playlists.NewDAL()
	playlistSvc := playlists.NewService(tagSvc, playlistLogger, playlistDAL)

	exploreLogger := kitlog.With(logger, "service", "explore")
	exploreSvc := NewService(exploreLogger, tagSvc, playlistSvc)

	ctx := context.Background()
	tagAResp, _ := tagSvc.CreateTag(ctx, &tags.CreateTagRequest{
		Tag: &tags.Tag{
			TagName: "tagA",
			TagType: 1,
		},
	})
	tagBResp, _ := tagSvc.CreateTag(ctx, &tags.CreateTagRequest{
		Tag: &tags.Tag{
			TagName: "tagB",
			TagType: 2,
		},
	})

	tags := []*tags.Tag{
		&tags.Tag{
			TagID:   tagAResp.TagID,
			TagName: "tagA",
			TagType: 1,
		},
		&tags.Tag{
			TagID:   tagBResp.TagID,
			TagName: "tagB",
			TagType: 2,
		},
	}

	playlistSvc.CreatePlaylist(ctx, &playlists.CreatePlaylistRequest{
		Playlist: &playlists.Playlist{
			PlaylistName: "A",
			Tags:         tags,
			Creator: playlists.Creator{
				ID:   "u_1",
				Name: "Tiger",
			},
		},
	},
	)
	playlistSvc.CreatePlaylist(ctx, &playlists.CreatePlaylistRequest{
		Playlist: &playlists.Playlist{
			PlaylistName: "B",
			Tags:         tags,
			Creator: playlists.Creator{
				ID:   "u_1",
				Name: "Tiger",
			},
		},
	})

	resp, _ := exploreSvc.Explore(ctx, &ExploreRequest{
		TagNames: []string{"tagA", "tagB"},
	})
	if len(resp.Playlists) != 2 {
		t.Errorf("playlist count do not match. expected %v actual %v", 2, len(resp.Playlists))
	}
}
