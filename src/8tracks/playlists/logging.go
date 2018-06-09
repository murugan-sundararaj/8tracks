package playlists

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

func (s *loggingService) CreatePlaylist(ctx context.Context, r *CreatePlaylistRequest) (resp *CreatePlaylistResponse, err error) {
	defer func(begin time.Time) {
		level.Debug(s.logger).Log(
			"request_id", ctx.Value(econst.RequestID),
			"method", ctx.Value(kithttp.ContextKeyRequestMethod),
			"path", ctx.Value(kithttp.ContextKeyRequestPath),
			"func", "create_playlist",
			"took", time.Since(begin),
			"err", err)
	}(time.Now())
	return s.Service.CreatePlaylist(ctx, r)
}

func (s *loggingService) LoadPlaylist(ctx context.Context, r *LoadPlaylistRequest) (resp *LoadPlaylistResponse, err error) {
	defer func(begin time.Time) {
		level.Debug(s.logger).Log(
			"request_id", ctx.Value(econst.RequestID),
			"method", ctx.Value(kithttp.ContextKeyRequestMethod),
			"path", ctx.Value(kithttp.ContextKeyRequestPath),
			"func", "load_playlist",
			"took", time.Since(begin),
			"err", err)
	}(time.Now())
	return s.Service.LoadPlaylist(ctx, r)
}

func (s *loggingService) UpsertPlaylist(ctx context.Context, r *UpsertPlaylistRequest) (resp *UpsertPlaylistResponse, err error) {
	defer func(begin time.Time) {
		level.Debug(s.logger).Log(
			"request_id", ctx.Value(econst.RequestID),
			"method", ctx.Value(kithttp.ContextKeyRequestMethod),
			"path", ctx.Value(kithttp.ContextKeyRequestPath),
			"func", "upsert_playlist",
			"took", time.Since(begin),
			"err", err)
	}(time.Now())
	return s.Service.UpsertPlaylist(ctx, r)
}

func (s *loggingService) UpdatePlaylistName(ctx context.Context, r *UpdatePlaylistNameRequest) (resp *UpdatePlaylistNameResponse, err error) {
	defer func(begin time.Time) {
		level.Debug(s.logger).Log(
			"request_id", ctx.Value(econst.RequestID),
			"method", ctx.Value(kithttp.ContextKeyRequestMethod),
			"path", ctx.Value(kithttp.ContextKeyRequestPath),
			"func", "update_playlist",
			"took", time.Since(begin),
			"err", err)
	}(time.Now())
	return s.Service.UpdatePlaylistName(ctx, r)
}

func (s *loggingService) RemovePlaylist(ctx context.Context, r *RemovePlaylistRequest) (resp *RemovePlaylistResponse, err error) {
	defer func(begin time.Time) {
		level.Debug(s.logger).Log(
			"request_id", ctx.Value(econst.RequestID),
			"method", ctx.Value(kithttp.ContextKeyRequestMethod),
			"path", ctx.Value(kithttp.ContextKeyRequestPath),
			"func", "remove_playlist",
			"took", time.Since(begin),
			"err", err)
	}(time.Now())
	return s.Service.RemovePlaylist(ctx, r)
}

func (s *loggingService) AddTrack(ctx context.Context, r *AddTrackRequest) (resp *AddTrackResponse, err error) {
	defer func(begin time.Time) {
		level.Debug(s.logger).Log(
			"request_id", ctx.Value(econst.RequestID),
			"method", ctx.Value(kithttp.ContextKeyRequestMethod),
			"path", ctx.Value(kithttp.ContextKeyRequestPath),
			"func", "add_track",
			"took", time.Since(begin),
			"err", err)
	}(time.Now())
	return s.Service.AddTrack(ctx, r)
}

func (s *loggingService) RemoveTrack(ctx context.Context, r *RemoveTrackRequest) (resp *RemoveTrackResponse, err error) {
	defer func(begin time.Time) {
		level.Debug(s.logger).Log(
			"request_id", ctx.Value(econst.RequestID),
			"method", ctx.Value(kithttp.ContextKeyRequestMethod),
			"path", ctx.Value(kithttp.ContextKeyRequestPath),
			"func", "remove_track",
			"took", time.Since(begin),
			"err", err)
	}(time.Now())
	return s.Service.RemoveTrack(ctx, r)
}

func (s *loggingService) Plays(ctx context.Context, r *PlaysRequest) (resp *PlaysResponse, err error) {
	defer func(begin time.Time) {
		level.Debug(s.logger).Log(
			"request_id", ctx.Value(econst.RequestID),
			"method", ctx.Value(kithttp.ContextKeyRequestMethod),
			"path", ctx.Value(kithttp.ContextKeyRequestPath),
			"func", "plays",
			"took", time.Since(begin),
			"err", err)
	}(time.Now())
	return s.Service.Plays(ctx, r)
}

func (s *loggingService) Likes(ctx context.Context, r *LikesRequest) (resp *LikesResponse, err error) {
	defer func(begin time.Time) {
		level.Debug(s.logger).Log(
			"request_id", ctx.Value(econst.RequestID),
			"method", ctx.Value(kithttp.ContextKeyRequestMethod),
			"path", ctx.Value(kithttp.ContextKeyRequestPath),
			"func", "likes",
			"took", time.Since(begin),
			"err", err)
	}(time.Now())
	return s.Service.Likes(ctx, r)
}

func (s *loggingService) Dislikes(ctx context.Context, r *DislikesRequest) (resp *DislikesResponse, err error) {
	defer func(begin time.Time) {
		level.Debug(s.logger).Log(
			"request_id", ctx.Value(econst.RequestID),
			"method", ctx.Value(kithttp.ContextKeyRequestMethod),
			"path", ctx.Value(kithttp.ContextKeyRequestPath),
			"func", "dislikes",
			"took", time.Since(begin),
			"err", err)
	}(time.Now())
	return s.Service.Dislikes(ctx, r)
}
