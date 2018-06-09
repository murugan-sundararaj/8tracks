package playlists

import (
	"8tracks/tags"
	"context"
	"os"
	"testing"

	kitlog "github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

func TestCreatePlaylist(t *testing.T) {
	logger := kitlog.NewLogfmtLogger(kitlog.NewSyncWriter(os.Stderr))
	logger = level.NewFilter(logger, level.AllowAll())
	logger = kitlog.With(logger, "ts", kitlog.DefaultTimestamp, "caller", kitlog.DefaultCaller)

	tagLogger := kitlog.With(logger, "service", "tag")
	tagDAL := tags.NewDAL()
	tagSvc := tags.NewService(tagLogger, tagDAL)

	playlistLogger := kitlog.With(logger, "service", "playlists")
	playlistDAL := NewDAL()
	playlistSvc := NewService(tagSvc, playlistLogger, playlistDAL)

	tagA := tags.Tag{
		TagName: "tagA",
		TagType: 1,
	}
	ctx := context.Background()
	tagSvc.CreateTag(ctx, &tags.CreateTagRequest{
		Tag: &tagA,
	})

	playlistSvc.CreatePlaylist(ctx, &CreatePlaylistRequest{
		&Playlist{
			PlaylistName: "A",
			Tags: []*tags.Tag{
				&tagA,
			},
			Creator: Creator{
				"u_1",
				"Tiger",
			},
		},
	})

	resp, _ := playlistSvc.LoadPlaylist(ctx, &LoadPlaylistRequest{
		PlaylistNames: []string{"A"},
	})
	for _, p := range resp.Playlists {
		if p.PlaylistName != "A" {
			t.Errorf("playlist name did not match. expecting %s got %s", "A", p.PlaylistName)
		}
	}
}

func TestUpsertPlaylist(t *testing.T) {
	logger := kitlog.NewLogfmtLogger(kitlog.NewSyncWriter(os.Stderr))
	logger = level.NewFilter(logger, level.AllowAll())
	logger = kitlog.With(logger, "ts", kitlog.DefaultTimestamp, "caller", kitlog.DefaultCaller)

	tagLogger := kitlog.With(logger, "service", "tag")
	tagDAL := tags.NewDAL()
	tagSvc := tags.NewService(tagLogger, tagDAL)

	playlistLogger := kitlog.With(logger, "service", "playlists")
	playlistDAL := NewDAL()
	playlistSvc := NewService(tagSvc, playlistLogger, playlistDAL)

	tagA := tags.Tag{
		TagName: "tagA",
		TagType: 1,
	}
	ctx := context.Background()
	tagSvc.CreateTag(ctx, &tags.CreateTagRequest{
		Tag: &tagA,
	})

	playlistSvc.UpsertPlaylist(ctx, &UpsertPlaylistRequest{
		&Playlist{
			PlaylistName: "A",
			Tags: []*tags.Tag{
				&tagA,
			},
			Creator: Creator{
				"u_1",
				"Tiger",
			},
		},
	})

	resp, _ := playlistSvc.LoadPlaylist(ctx, &LoadPlaylistRequest{
		PlaylistNames: []string{"A"},
	})
	for _, p := range resp.Playlists {
		if p.PlaylistName != "A" {
			t.Errorf("playlist name did not match. expecting %s got %s", "A", p.PlaylistName)
		}
	}
}

func TestUpdatePlaylistName(t *testing.T) {
	logger := kitlog.NewLogfmtLogger(kitlog.NewSyncWriter(os.Stderr))
	logger = level.NewFilter(logger, level.AllowAll())
	logger = kitlog.With(logger, "ts", kitlog.DefaultTimestamp, "caller", kitlog.DefaultCaller)

	tagLogger := kitlog.With(logger, "service", "tag")
	tagDAL := tags.NewDAL()
	tagSvc := tags.NewService(tagLogger, tagDAL)

	playlistLogger := kitlog.With(logger, "service", "playlists")
	playlistDAL := NewDAL()
	playlistSvc := NewService(tagSvc, playlistLogger, playlistDAL)

	tagA := tags.Tag{
		TagName: "tagA",
		TagType: 1,
	}
	ctx := context.Background()
	tagSvc.CreateTag(ctx, &tags.CreateTagRequest{
		Tag: &tagA,
	})

	createResp, _ := playlistSvc.CreatePlaylist(ctx, &CreatePlaylistRequest{
		&Playlist{
			PlaylistName: "A",
			Tags: []*tags.Tag{
				&tagA,
			},
			Creator: Creator{
				"u_1",
				"Tiger",
			},
		},
	})

	playlistSvc.UpdatePlaylistName(ctx, &UpdatePlaylistNameRequest{
		PlaylistID:   createResp.PlaylistID,
		PlaylistName: "X",
	})

	resp, _ := playlistSvc.LoadPlaylist(ctx, &LoadPlaylistRequest{
		PlaylistNames: []string{"A"},
	})
	if resp.Err != ErrInvalidPlaylist.Error() {
		t.Error("playlist name is not update")
	}
	resp, _ = playlistSvc.LoadPlaylist(ctx, &LoadPlaylistRequest{
		PlaylistNames: []string{"X"},
	})
	for _, p := range resp.Playlists {
		if p.PlaylistName != "X" {
			t.Errorf("playlist name did not match. expecting %s got %s", "X", p.PlaylistName)
		}
	}
}

func TestRemovePlaylist(t *testing.T) {
	logger := kitlog.NewLogfmtLogger(kitlog.NewSyncWriter(os.Stderr))
	logger = level.NewFilter(logger, level.AllowAll())
	logger = kitlog.With(logger, "ts", kitlog.DefaultTimestamp, "caller", kitlog.DefaultCaller)

	tagLogger := kitlog.With(logger, "service", "tag")
	tagDAL := tags.NewDAL()
	tagSvc := tags.NewService(tagLogger, tagDAL)

	playlistLogger := kitlog.With(logger, "service", "playlists")
	playlistDAL := NewDAL()
	playlistSvc := NewService(tagSvc, playlistLogger, playlistDAL)

	tagA := tags.Tag{
		TagName: "tagA",
		TagType: 1,
	}
	ctx := context.Background()
	tagSvc.CreateTag(ctx, &tags.CreateTagRequest{
		Tag: &tagA,
	})

	createResp, _ := playlistSvc.CreatePlaylist(ctx, &CreatePlaylistRequest{
		&Playlist{
			PlaylistName: "A",
			Tags: []*tags.Tag{
				&tagA,
			},
			Creator: Creator{
				"u_1",
				"Tiger",
			},
		},
	})

	playlistSvc.RemovePlaylist(ctx, &RemovePlaylistRequest{
		PlaylistID: createResp.PlaylistID,
	})

	resp, _ := playlistSvc.LoadPlaylist(ctx, &LoadPlaylistRequest{
		PlaylistNames: []string{"A"},
	})
	if resp.Err != ErrInvalidPlaylist.Error() {
		t.Error("playlist name is not removed")
	}
}

func TestTrack(t *testing.T) {
	logger := kitlog.NewLogfmtLogger(kitlog.NewSyncWriter(os.Stderr))
	logger = level.NewFilter(logger, level.AllowAll())
	logger = kitlog.With(logger, "ts", kitlog.DefaultTimestamp, "caller", kitlog.DefaultCaller)

	tagLogger := kitlog.With(logger, "service", "tag")
	tagDAL := tags.NewDAL()
	tagSvc := tags.NewService(tagLogger, tagDAL)

	playlistLogger := kitlog.With(logger, "service", "playlists")
	playlistDAL := NewDAL()
	playlistSvc := NewService(tagSvc, playlistLogger, playlistDAL)

	tagA := tags.Tag{
		TagName: "tagA",
		TagType: 1,
	}
	ctx := context.Background()
	tagSvc.CreateTag(ctx, &tags.CreateTagRequest{
		Tag: &tagA,
	})

	createResp, _ := playlistSvc.CreatePlaylist(ctx, &CreatePlaylistRequest{
		&Playlist{
			PlaylistName: "A",
			Tags: []*tags.Tag{
				&tagA,
			},
			Creator: Creator{
				"u_1",
				"Tiger",
			},
		},
	})

	playlistSvc.AddTrack(ctx, &AddTrackRequest{
		PlaylistID: createResp.PlaylistID,
		Track: &Track{
			ID:   "trk_1",
			Name: "lala",
		},
	})
	resp, _ := playlistSvc.LoadPlaylist(ctx, &LoadPlaylistRequest{
		PlaylistNames: []string{"A"},
	})
	for _, p := range resp.Playlists {
		for _, trk := range p.Tracks {
			if trk.Name != "lala" {
				t.Errorf("track name did not match. expecting %s got %s", "lala", trk.Name)
			}
		}
	}

	playlistSvc.RemoveTrack(ctx, &RemoveTrackRequest{
		PlaylistID: createResp.PlaylistID,
		TrackID:    "trk_1",
	})
	resp, _ = playlistSvc.LoadPlaylist(ctx, &LoadPlaylistRequest{
		PlaylistNames: []string{"A"},
	})
	for _, p := range resp.Playlists {
		if len(p.Tracks) != 0 {
			t.Errorf("track not remove")
		}
	}
}
