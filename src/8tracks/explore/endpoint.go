package explore

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

func makeExploreEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, r interface{}) (interface{}, error) {
		req := r.(*ExploreRequest)
		res, err := s.Explore(ctx, req)
		return res, err
	}
}
