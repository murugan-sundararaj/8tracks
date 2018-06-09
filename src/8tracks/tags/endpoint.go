package tags

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

func makeCreateTagEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, r interface{}) (interface{}, error) {
		req := r.(*CreateTagRequest)
		res, err := s.CreateTag(ctx, req)
		return res, err
	}
}

func makeLoadTagEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, r interface{}) (interface{}, error) {
		req := r.(*LoadTagRequest)
		res, err := s.LoadTag(ctx, req)
		return res, err
	}
}

func makeUpsertTagEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, r interface{}) (interface{}, error) {
		req := r.(*UpsertTagRequest)
		res, err := s.UpsertTag(ctx, req)
		return res, err
	}
}

func makeUpdateTagEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, r interface{}) (interface{}, error) {
		req := r.(*UpdateTagRequest)
		res, err := s.UpdateTag(ctx, req)
		return res, err
	}
}

func makeRemoveTagEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, r interface{}) (interface{}, error) {
		req := r.(*RemoveTagRequest)
		res, err := s.RemoveTag(ctx, req)
		return res, err
	}
}

func makeLoadTagTypesEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, r interface{}) (interface{}, error) {
		req := r.(*LoadTagTypesRequest)
		res, err := s.LoadTagTypes(ctx, req)
		return res, err
	}
}

func makeAssignTagToPlaylistEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, r interface{}) (interface{}, error) {
		req := r.(*AssignTagToPlaylistRequest)
		res, err := s.AssignTagToPlaylist(ctx, req)
		return res, err
	}
}

func makeUnAssignTagFromPlaylistEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, r interface{}) (interface{}, error) {
		req := r.(*UnAssignTagFromPlaylistRequest)
		res, err := s.UnAssignTagFromPlaylist(ctx, req)
		return res, err
	}
}
