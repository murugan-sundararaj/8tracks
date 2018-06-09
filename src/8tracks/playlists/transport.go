package playlists

import (
	"8tracks/lib/econst"
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"strings"

	kitlog "github.com/go-kit/kit/log"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
)

// IDGenerator should return unique record identifiers, i.e. ULIDs.
type IDGenerator func() string

// PopulateRequestID populates ULID as request id into the context
func PopulateRequestID(idGen IDGenerator) kithttp.RequestFunc {
	return func(ctx context.Context, r *http.Request) context.Context {
		ctx = context.WithValue(ctx, econst.RequestID, idGen())
		return ctx
	}
}

// MakeHandler returns a handler for the auth service.
func MakeHandler(s Service, logger kitlog.Logger, idGen IDGenerator) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerErrorEncoder(encodeError),
		kithttp.ServerBefore(kithttp.PopulateRequestContext, PopulateRequestID(idGen)),
	}

	createPlaylistHandler := kithttp.NewServer(
		makeCreatePlaylistEndpoint(s),
		decodeCreatePlaylistRequest,
		encodeResponse,
		opts...,
	)

	loadPlaylistHandler := kithttp.NewServer(
		makeLoadPlaylistEndpoint(s),
		decodeLoadPlaylistRequest,
		encodeResponse,
		opts...,
	)

	upsertPlaylistHandler := kithttp.NewServer(
		makeUpsertPlaylistEndpoint(s),
		decodeUpsertPlaylistRequest,
		encodeResponse,
		opts...,
	)

	updatePlaylistNameHandler := kithttp.NewServer(
		makeUpdatePlaylistNameEndpoint(s),
		decodeUpdatePlaylistNameRequest,
		encodeResponse,
		opts...,
	)

	removePlaylistHandler := kithttp.NewServer(
		makeRemovePlaylistEndpoint(s),
		decodeRemovePlaylistRequest,
		encodeResponse,
		opts...,
	)

	addTrackHandler := kithttp.NewServer(
		makeAddTrackEndpoint(s),
		decodeAddTrackRequest,
		encodeResponse,
		opts...,
	)

	removeTrackHandler := kithttp.NewServer(
		makeRemoveTrackEndpoint(s),
		decodeRemoveTrackRequest,
		encodeResponse,
		opts...,
	)

	playsHandler := kithttp.NewServer(
		makePlaysEndpoint(s),
		decodePlaysRequest,
		encodeResponse,
		opts...,
	)

	likesHandler := kithttp.NewServer(
		makeLikesEndpoint(s),
		decodeLikesRequest,
		encodeResponse,
		opts...,
	)

	dislikesHandler := kithttp.NewServer(
		makeDislikesEndpoint(s),
		decodeDislikesRequest,
		encodeResponse,
		opts...,
	)

	r := mux.NewRouter()

	r.Handle("/playlists/playlist", createPlaylistHandler).Methods("POST")
	r.Handle("/playlists/playlist", loadPlaylistHandler).Methods("GET")
	r.Handle("/playlists/{id}", upsertPlaylistHandler).Methods("PUT")
	r.Handle("/playlists/{id}", updatePlaylistNameHandler).Methods("PATCH")
	r.Handle("/playlists/{id}", removePlaylistHandler).Methods("DELETE")
	r.Handle("/playlists/{playlistid}/tracks/{trackid}", addTrackHandler).Methods("PUT")
	r.Handle("/playlists/{playlistid}/tracks/{trackid}", removeTrackHandler).Methods("DELETE")
	r.Handle("/playlists/{id}/plays", playsHandler).Methods("POST")
	r.Handle("/playlists/{id}/likes", likesHandler).Methods("POST")
	r.Handle("/playlists/{id}/dislikes", dislikesHandler).Methods("POST")

	return r
}

// ErrDecode encapsulates the error occured during decoding
type ErrDecode struct {
	reason error
}

func (e ErrDecode) Error() string {
	return e.reason.Error()
}

func decodeCreatePlaylistRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req *CreatePlaylistRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, ErrDecode{errors.Wrap(err, "json decoding failed")}
	}
	return req, nil
}

func decodeLoadPlaylistRequest(_ context.Context, r *http.Request) (interface{}, error) {
	u, err := url.Parse(r.URL.String())
	if err != nil {
		return nil, ErrDecode{errors.Wrap(err, "url parsing failed")}
	}
	q := u.Query()
	playlistIDs := []string{}
	for _, ids := range q["ids"] {
		for _, id := range strings.Split(ids, ";") {
			playlistIDs = append(playlistIDs, id)
		}
	}

	playlistNames := []string{}
	for _, names := range q["names"] {
		for _, name := range strings.Split(names, ";") {
			playlistNames = append(playlistNames, name)
		}
	}
	return &LoadPlaylistRequest{
		PlaylistIDs:   playlistIDs,
		PlaylistNames: playlistNames,
	}, nil
}

func decodeUpsertPlaylistRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	playlistID, ok := vars["id"]
	if !ok {
		return nil, ErrDecode{errors.New("bad route")}
	}
	var req *UpsertPlaylistRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, ErrDecode{errors.Wrap(err, "json decoding failed")}
	}
	req.PlaylistID = playlistID
	return req, nil
}

func decodeUpdatePlaylistNameRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	playlistID, ok := vars["id"]
	if !ok {
		return nil, ErrDecode{errors.New("bad route")}
	}
	var req *UpdatePlaylistNameRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, ErrDecode{errors.Wrap(err, "json decoding failed")}
	}
	req.PlaylistID = playlistID
	return req, nil
}

func decodeRemovePlaylistRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	playlistID, ok := vars["id"]
	if !ok {
		return nil, ErrDecode{errors.New("bad route")}
	}
	return &RemovePlaylistRequest{PlaylistID: playlistID}, nil
}

func decodeAddTrackRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	playlistID, ok := vars["playlistid"]
	if !ok {
		return nil, ErrDecode{errors.New("bad route")}
	}
	trackID, ok := vars["trackid"]
	if !ok {
		return nil, ErrDecode{errors.New("bad route")}
	}
	var req *AddTrackRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, ErrDecode{errors.Wrap(err, "json decoding failed")}
	}
	req.PlaylistID = playlistID
	req.Track.ID = trackID
	return req, nil
}

func decodeRemoveTrackRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	playlistID, ok := vars["playlistid"]
	if !ok {
		return nil, ErrDecode{errors.New("bad route")}
	}
	trackID, ok := vars["trackid"]
	if !ok {
		return nil, ErrDecode{errors.New("bad route")}
	}
	return &RemoveTrackRequest{
		PlaylistID: playlistID,
		TrackID:    trackID,
	}, nil
}
func decodePlaysRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	playlistID, ok := vars["id"]
	if !ok {
		return nil, ErrDecode{errors.New("bad route")}
	}
	return &PlaysRequest{
		PlaylistID: playlistID,
	}, nil
}

func decodeLikesRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	playlistID, ok := vars["id"]
	if !ok {
		return nil, ErrDecode{errors.New("bad route")}
	}
	return &LikesRequest{
		PlaylistID: playlistID,
	}, nil
}

func decodeDislikesRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	playlistID, ok := vars["id"]
	if !ok {
		return nil, ErrDecode{errors.New("bad route")}
	}
	return &DislikesRequest{
		PlaylistID: playlistID,
	}, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(errorer); ok && e.error() != "" {
		w.WriteHeader(http.StatusBadRequest)
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

type errorer interface {
	error() string
}

// encode errors from business-logic
func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	switch err.(type) {
	case ErrDecode:
		w.WriteHeader(http.StatusBadRequest)
	case error:
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}
