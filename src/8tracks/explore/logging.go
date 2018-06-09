package explore

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

func (s *loggingService) Explore(ctx context.Context, r *ExploreRequest) (resp *ExploreResponse, err error) {
	defer func(begin time.Time) {
		level.Debug(s.logger).Log(
			"request_id", ctx.Value(econst.RequestID),
			"method", ctx.Value(kithttp.ContextKeyRequestMethod),
			"path", ctx.Value(kithttp.ContextKeyRequestPath),
			"func", "explore",
			"took", time.Since(begin),
			"err", err)
	}(time.Now())
	return s.Service.Explore(ctx, r)
}
