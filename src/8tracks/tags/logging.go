package tags

import (
	"8tracks/lib/econst"
	"context"
	"time"

	kitlog "github.com/go-kit/kit/log"
	level "github.com/go-kit/kit/log/level"
	kithttp "github.com/go-kit/kit/transport/http"
)

type loggingService struct {
	Service
	logger kitlog.Logger
}

// NewLoggingService returns a new instance of a logging Service.
func NewLoggingService(s Service, logger kitlog.Logger) Service {
	return &loggingService{s, logger}
}

func (s *loggingService) CreateTag(ctx context.Context, r *CreateTagRequest) (resp *CreateTagResponse, err error) {
	defer func(begin time.Time) {
		level.Debug(s.logger).Log(
			"request_id", ctx.Value(econst.RequestID),
			"method", ctx.Value(kithttp.ContextKeyRequestMethod),
			"path", ctx.Value(kithttp.ContextKeyRequestPath),
			"func", "create_tag",
			"took", time.Since(begin),
			"err", err)
	}(time.Now())
	return s.Service.CreateTag(ctx, r)
}

func (s *loggingService) LoadTag(ctx context.Context, r *LoadTagRequest) (resp *LoadTagResponse, err error) {
	defer func(begin time.Time) {
		level.Debug(s.logger).Log(
			"request_id", ctx.Value(econst.RequestID),
			"method", ctx.Value(kithttp.ContextKeyRequestMethod),
			"path", ctx.Value(kithttp.ContextKeyRequestPath),
			"func", "load_tag",
			"took", time.Since(begin),
			"err", err)
	}(time.Now())
	return s.Service.LoadTag(ctx, r)
}

func (s *loggingService) UpsertTag(ctx context.Context, r *UpsertTagRequest) (resp *UpsertTagResponse, err error) {
	defer func(begin time.Time) {
		level.Debug(s.logger).Log(
			"request_id", ctx.Value(econst.RequestID),
			"method", ctx.Value(kithttp.ContextKeyRequestMethod),
			"path", ctx.Value(kithttp.ContextKeyRequestPath),
			"func", "upsert_tag",
			"tag_id", r.TagID,
			"took", time.Since(begin),
			"err", err)
	}(time.Now())
	return s.Service.UpsertTag(ctx, r)
}

func (s *loggingService) UpdateTag(ctx context.Context, r *UpdateTagRequest) (resp *UpdateTagResponse, err error) {
	defer func(begin time.Time) {
		level.Debug(s.logger).Log(
			"request_id", ctx.Value(econst.RequestID),
			"method", ctx.Value(kithttp.ContextKeyRequestMethod),
			"path", ctx.Value(kithttp.ContextKeyRequestPath),
			"func", "update_tag",
			"tag_id", r.TagID,
			"took", time.Since(begin),
			"err", err)
	}(time.Now())
	return s.Service.UpdateTag(ctx, r)
}

func (s *loggingService) RemoveTag(ctx context.Context, r *RemoveTagRequest) (resp *RemoveTagResponse, err error) {
	defer func(begin time.Time) {
		level.Debug(s.logger).Log(
			"request_id", ctx.Value(econst.RequestID),
			"method", ctx.Value(kithttp.ContextKeyRequestMethod),
			"path", ctx.Value(kithttp.ContextKeyRequestPath),
			"func", "remove_tag",
			"tag_id", r.TagID,
			"took", time.Since(begin),
			"err", err)
	}(time.Now())
	return s.Service.RemoveTag(ctx, r)
}

func (s *loggingService) LoadTagTypes(ctx context.Context, r *LoadTagTypesRequest) (resp *LoadTagTypesResponse, err error) {
	defer func(begin time.Time) {
		level.Debug(s.logger).Log(
			"request_id", ctx.Value(econst.RequestID),
			"method", ctx.Value(kithttp.ContextKeyRequestMethod),
			"path", ctx.Value(kithttp.ContextKeyRequestPath),
			"func", "load_tag",
			"took", time.Since(begin),
			"err", err)
	}(time.Now())
	return s.Service.LoadTagTypes(ctx, r)
}

func (s *loggingService) AssignTagToPlaylist(ctx context.Context, r *AssignTagToPlaylistRequest) (resp *AssignTagToPlaylistResponse, err error) {
	defer func(begin time.Time) {
		level.Debug(s.logger).Log(
			"request_id", ctx.Value(econst.RequestID),
			"method", ctx.Value(kithttp.ContextKeyRequestMethod),
			"path", ctx.Value(kithttp.ContextKeyRequestPath),
			"func", "assign_tag_to_playlist",
			"took", time.Since(begin),
			"tag_id", r.TagID,
			"err", err)
	}(time.Now())
	return s.Service.AssignTagToPlaylist(ctx, r)
}

func (s *loggingService) UnAssignTagFromPlaylist(ctx context.Context, r *UnAssignTagFromPlaylistRequest) (resp *UnAssignTagFromPlaylistResponse, err error) {
	defer func(begin time.Time) {
		level.Debug(s.logger).Log(
			"request_id", ctx.Value(econst.RequestID),
			"method", ctx.Value(kithttp.ContextKeyRequestMethod),
			"path", ctx.Value(kithttp.ContextKeyRequestPath),
			"func", "unassign_tag_from_playlist",
			"took", time.Since(begin),
			"tag_id", r.TagID,
			"err", err)
	}(time.Now())
	return s.Service.UnAssignTagFromPlaylist(ctx, r)
}
