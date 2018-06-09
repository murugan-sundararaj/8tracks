package tags

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

	createTagHandler := kithttp.NewServer(
		makeCreateTagEndpoint(s),
		decodeCreateTagRequest,
		encodeResponse,
		opts...,
	)

	loadTagHandler := kithttp.NewServer(
		makeLoadTagEndpoint(s),
		decodeLoadTagRequest,
		encodeResponse,
		opts...,
	)

	upsertTagHandler := kithttp.NewServer(
		makeUpsertTagEndpoint(s),
		decodeUpsertTagRequest,
		encodeResponse,
		opts...,
	)

	updateTagHandler := kithttp.NewServer(
		makeUpdateTagEndpoint(s),
		decodeUpdateTagRequest,
		encodeResponse,
		opts...,
	)

	removeTagHandler := kithttp.NewServer(
		makeRemoveTagEndpoint(s),
		decodeRemoveTagRequest,
		encodeResponse,
		opts...,
	)

	loadTagTypesHandler := kithttp.NewServer(
		makeLoadTagTypesEndpoint(s),
		decodeLoadTagTypesRequest,
		encodeResponse,
		opts...,
	)

	assignTagToPlaylistHandler := kithttp.NewServer(
		makeAssignTagToPlaylistEndpoint(s),
		decodeAssignTagToPlaylistRequest,
		encodeResponse,
		opts...,
	)

	unAssignTagFromPlaylistHandler := kithttp.NewServer(
		makeUnAssignTagFromPlaylistEndpoint(s),
		decodeUnAssignTagFromPlaylistRequest,
		encodeResponse,
		opts...,
	)

	r := mux.NewRouter()
	r.Handle("/tags/tag", createTagHandler).Methods("POST")
	r.Handle("/tags/tag", loadTagHandler).Methods("GET")
	r.Handle("/tags/{id}", upsertTagHandler).Methods("PUT")
	r.Handle("/tags/{id}", updateTagHandler).Methods("PATCH")
	r.Handle("/tags/{id}", removeTagHandler).Methods("DELETE")
	r.Handle("/tags/types", loadTagTypesHandler).Methods("GET")

	r.Handle("/tags/{tagid}/playlists/{playlistid}", assignTagToPlaylistHandler).Methods("PUT")
	r.Handle("/tags/{tagid}/playlists/{playlistid}", unAssignTagFromPlaylistHandler).Methods("DELETE")
	return r
}

// ErrDecode encapsulates the error occured during decoding
type ErrDecode struct {
	reason error
}

func (e ErrDecode) Error() string {
	return e.reason.Error()
}

func decodeCreateTagRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req *CreateTagRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, ErrDecode{errors.Wrap(err, "json decoding failed")}
	}
	return req, nil
}

func decodeLoadTagRequest(_ context.Context, r *http.Request) (interface{}, error) {
	u, err := url.Parse(r.URL.String())
	if err != nil {
		return nil, ErrDecode{errors.Wrap(err, "url parsing failed")}
	}
	q := u.Query()

	tagIDs := []string{}
	for _, ids := range q["ids"] {
		for _, id := range strings.Split(ids, ";") {
			tagIDs = append(tagIDs, string(id))
		}
	}

	tagNames := []string{}
	for _, names := range q["names"] {
		for _, name := range strings.Split(names, ";") {
			tagNames = append(tagNames, name)
		}
	}

	return &LoadTagRequest{
		TagIDs:   tagIDs,
		TagNames: tagNames,
	}, nil
}

func decodeUpsertTagRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	tagID, ok := vars["id"]
	if !ok {
		return nil, ErrDecode{errors.New("bad route")}
	}
	var req *UpsertTagRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, ErrDecode{errors.Wrap(err, "json decoding failed")}
	}
	req.TagID = tagID
	return req, nil
}

func decodeUpdateTagRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	tagID, ok := vars["id"]
	if !ok {
		return nil, ErrDecode{errors.New("bad route")}
	}
	var req *UpdateTagRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, ErrDecode{errors.Wrap(err, "json decoding failed")}
	}
	req.TagID = tagID
	return req, nil
}

func decodeRemoveTagRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	tagID, ok := vars["id"]
	if !ok {
		return nil, ErrDecode{errors.New("bad route")}
	}
	return &RemoveTagRequest{TagID: tagID}, nil
}

func decodeLoadTagTypesRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return &LoadTagTypesRequest{}, nil
}

func decodeAssignTagToPlaylistRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	tagID, ok := vars["tagid"]
	if !ok {
		return nil, ErrDecode{errors.New("bad route")}
	}
	playlistID, ok := vars["playlistid"]
	if !ok {
		return nil, ErrDecode{errors.New("bad route")}
	}
	return &AssignTagToPlaylistRequest{
		TagID:      tagID,
		PlaylistID: string(playlistID),
	}, nil
}

func decodeUnAssignTagFromPlaylistRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	tagID, ok := vars["tagid"]
	if !ok {
		return nil, ErrDecode{errors.New("bad route")}
	}
	playlistID, ok := vars["playlistid"]
	if !ok {
		return nil, ErrDecode{errors.New("bad route")}
	}
	return &UnAssignTagFromPlaylistRequest{
		TagID:      tagID,
		PlaylistID: string(playlistID),
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
