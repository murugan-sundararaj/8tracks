package tags

import (
	"context"
	"os"
	"testing"

	kitlog "github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

func TestCreateTag(t *testing.T) {
	logger := kitlog.NewLogfmtLogger(kitlog.NewSyncWriter(os.Stderr))
	logger = level.NewFilter(logger, level.AllowAll())
	logger = kitlog.With(logger, "ts", kitlog.DefaultTimestamp, "caller", kitlog.DefaultCaller)
	tagLogger := kitlog.With(logger, "service", "tag")
	tagDAL := NewDAL()
	tagSvc := NewService(tagLogger, tagDAL)
	tagA := Tag{
		TagName: "tagA",
		TagType: 1,
	}
	ctx := context.Background()
	tagSvc.CreateTag(ctx, &CreateTagRequest{
		&tagA,
	})

	tagResp, _ := tagSvc.LoadTag(ctx, &LoadTagRequest{
		TagNames: []string{"tagA"},
	})
	for _, tag := range tagResp.Tags {
		if tag.TagName != "tagA" {
			t.Errorf("tag name did not match. expecting %s got %s", "tagA", tag.TagName)
		}
	}
}

func TestUpsertTag(t *testing.T) {
	logger := kitlog.NewLogfmtLogger(kitlog.NewSyncWriter(os.Stderr))
	logger = level.NewFilter(logger, level.AllowAll())
	logger = kitlog.With(logger, "ts", kitlog.DefaultTimestamp, "caller", kitlog.DefaultCaller)
	tagLogger := kitlog.With(logger, "service", "tag")
	tagDAL := NewDAL()
	tagSvc := NewService(tagLogger, tagDAL)
	ctx := context.Background()
	tagX := Tag{
		TagName: "tagX",
		TagType: 1,
	}
	tagSvc.UpsertTag(ctx, &UpsertTagRequest{
		&tagX,
	})

	tagResp, _ := tagSvc.LoadTag(ctx, &LoadTagRequest{
		TagNames: []string{"tagX"},
	})
	if tagResp.Err == ErrInvalidTag.Error() {
		t.Errorf("tag not inserted")
	}
	for _, tag := range tagResp.Tags {
		if tag.TagName != "tagX" {
			t.Errorf("tag name did not match. expecting %s got %s", "tagX", tag.TagName)
		}
	}
}

func TestUpdateTag(t *testing.T) {
	logger := kitlog.NewLogfmtLogger(kitlog.NewSyncWriter(os.Stderr))
	logger = level.NewFilter(logger, level.AllowAll())
	logger = kitlog.With(logger, "ts", kitlog.DefaultTimestamp, "caller", kitlog.DefaultCaller)
	tagLogger := kitlog.With(logger, "service", "tag")
	tagDAL := NewDAL()
	tagSvc := NewService(tagLogger, tagDAL)
	tagA := Tag{
		TagName: "tagA",
		TagType: 1,
	}
	ctx := context.Background()
	resp, _ := tagSvc.CreateTag(ctx, &CreateTagRequest{
		&tagA,
	})

	tagX := Tag{
		TagID:   resp.TagID,
		TagName: "tagX",
		TagType: 2,
	}
	tagSvc.UpdateTag(ctx, &UpdateTagRequest{
		&tagX,
	})

	tagResp, _ := tagSvc.LoadTag(ctx, &LoadTagRequest{
		TagNames: []string{"tagA"},
	})
	if tagResp.Err != ErrInvalidTag.Error() {
		t.Errorf("tag not replaced")
	}
	for _, tag := range tagResp.Tags {
		if tag.TagName != "tagX" {
			t.Errorf("tag name did not match. expecting %s got %s", "tagX", tag.TagName)
		}
		if tag.TagType != 2 {
			t.Errorf("tag type did not match. expecting %v got %s", 2, tag.TagType)
		}
	}
}

func TestRemoveTag(t *testing.T) {
	logger := kitlog.NewLogfmtLogger(kitlog.NewSyncWriter(os.Stderr))
	logger = level.NewFilter(logger, level.AllowAll())
	logger = kitlog.With(logger, "ts", kitlog.DefaultTimestamp, "caller", kitlog.DefaultCaller)
	tagLogger := kitlog.With(logger, "service", "tag")
	tagDAL := NewDAL()
	tagSvc := NewService(tagLogger, tagDAL)
	tagA := Tag{
		TagName: "tagA",
		TagType: 1,
	}
	ctx := context.Background()
	resp, _ := tagSvc.CreateTag(ctx, &CreateTagRequest{
		&tagA,
	})
	tagSvc.RemoveTag(ctx, &RemoveTagRequest{
		TagID: resp.TagID,
	})

	tagResp, _ := tagSvc.LoadTag(ctx, &LoadTagRequest{
		TagNames: []string{"tagA"},
	})
	if tagResp.Err != ErrInvalidTag.Error() {
		t.Errorf("tag not removed")
	}
}

func TestAssignTagToPlaylist(t *testing.T) {
	logger := kitlog.NewLogfmtLogger(kitlog.NewSyncWriter(os.Stderr))
	logger = level.NewFilter(logger, level.AllowAll())
	logger = kitlog.With(logger, "ts", kitlog.DefaultTimestamp, "caller", kitlog.DefaultCaller)
	tagLogger := kitlog.With(logger, "service", "tag")
	tagDAL := NewDAL()
	tagSvc := NewService(tagLogger, tagDAL)
	tagA := Tag{
		TagName: "tagA",
		TagType: 1,
	}
	ctx := context.Background()
	resp, _ := tagSvc.CreateTag(ctx, &CreateTagRequest{
		&tagA,
	})

	tagSvc.AssignTagToPlaylist(ctx, &AssignTagToPlaylistRequest{
		TagID:      resp.TagID,
		PlaylistID: "p_1",
	})

	playlistIDSet, _ := tagSvc.LoadTagPlaylistID(ctx, "tagA")
	if !playlistIDSet.Has("p_1") {
		t.Error("tag is not assigned")
	}

	tagIDSet, _ := tagSvc.LoadPlaylistTag(ctx, "p_1")
	if !tagIDSet.Has(resp.TagID) {
		t.Error("playlist is not assigned")
	}
}

func TestUnAssignTagToPlaylist(t *testing.T) {
	logger := kitlog.NewLogfmtLogger(kitlog.NewSyncWriter(os.Stderr))
	logger = level.NewFilter(logger, level.AllowAll())
	logger = kitlog.With(logger, "ts", kitlog.DefaultTimestamp, "caller", kitlog.DefaultCaller)
	tagLogger := kitlog.With(logger, "service", "tag")
	tagDAL := NewDAL()
	tagSvc := NewService(tagLogger, tagDAL)
	tagA := Tag{
		TagName: "tagA",
		TagType: 1,
	}
	ctx := context.Background()
	resp, _ := tagSvc.CreateTag(ctx, &CreateTagRequest{
		&tagA,
	})

	tagSvc.AssignTagToPlaylist(ctx, &AssignTagToPlaylistRequest{
		TagID:      resp.TagID,
		PlaylistID: "p_1",
	})

	playlistIDSet, _ := tagSvc.LoadTagPlaylistID(ctx, "tagA")
	if !playlistIDSet.Has("p_1") {
		t.Error("tag is not assigned")
	}

	tagIDSet, _ := tagSvc.LoadPlaylistTag(ctx, "p_1")
	if !tagIDSet.Has(resp.TagID) {
		t.Error("playlist is not assigned")
	}

	tagSvc.UnAssignTagFromPlaylist(ctx, &UnAssignTagFromPlaylistRequest{
		TagID:      resp.TagID,
		PlaylistID: "p_1",
	})

	playlistIDSet, _ = tagSvc.LoadTagPlaylistID(ctx, "tagA")
	if playlistIDSet.Has("p_1") {
		t.Error("tag is not unassigned")
	}

	tagIDSet, _ = tagSvc.LoadPlaylistTag(ctx, "p_1")
	if tagIDSet.Has(resp.TagID) {
		t.Error("playlist is not unassigned")
	}
}
