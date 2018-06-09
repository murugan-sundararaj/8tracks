package playlists

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

func makeCreatePlaylistEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, r interface{}) (interface{}, error) {
		req := r.(*CreatePlaylistRequest)
		res, err := s.CreatePlaylist(ctx, req)
		return res, err
	}
}

func makeLoadPlaylistEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, r interface{}) (interface{}, error) {
		req := r.(*LoadPlaylistRequest)
		res, err := s.LoadPlaylist(ctx, req)
		return res, err
	}
}

func makeUpsertPlaylistEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, r interface{}) (interface{}, error) {
		req := r.(*UpsertPlaylistRequest)
		res, err := s.UpsertPlaylist(ctx, req)
		return res, err
	}
}

func makeUpdatePlaylistNameEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, r interface{}) (interface{}, error) {
		req := r.(*UpdatePlaylistNameRequest)
		res, err := s.UpdatePlaylistName(ctx, req)
		return res, err
	}
}

func makeRemovePlaylistEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, r interface{}) (interface{}, error) {
		req := r.(*RemovePlaylistRequest)
		res, err := s.RemovePlaylist(ctx, req)
		return res, err
	}
}

func makeAddTrackEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, r interface{}) (interface{}, error) {
		req := r.(*AddTrackRequest)
		res, err := s.AddTrack(ctx, req)
		return res, err
	}
}

func makeRemoveTrackEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, r interface{}) (interface{}, error) {
		req := r.(*RemoveTrackRequest)
		res, err := s.RemoveTrack(ctx, req)
		return res, err
	}
}

func makePlaysEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, r interface{}) (interface{}, error) {
		req := r.(*PlaysRequest)
		res, err := s.Plays(ctx, req)
		return res, err
	}
}

func makeLikesEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, r interface{}) (interface{}, error) {
		req := r.(*LikesRequest)
		res, err := s.Likes(ctx, req)
		return res, err
	}
}

func makeDislikesEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, r interface{}) (interface{}, error) {
		req := r.(*DislikesRequest)
		res, err := s.Dislikes(ctx, req)
		return res, err
	}
}
