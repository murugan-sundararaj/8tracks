package explore

import (
	"8tracks/lib/econst"
	"context"
	"encoding/json"
	"net/http"
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

	exploreHandler := kithttp.NewServer(
		makeExploreEndpoint(s),
		decodeExploreRequest,
		encodeResponse,
		opts...,
	)

	r := mux.NewRouter()
	r.Handle("/explore/{tags}", exploreHandler).Methods("GET")
	return r
}

// ErrDecode encapsulates the error occured during decoding
type ErrDecode struct {
	reason error
}

func (e ErrDecode) Error() string {
	return e.reason.Error()
}

func decodeExploreRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	tags, ok := vars["tags"]
	if !ok {
		return nil, ErrDecode{errors.New("bad route")}
	}
	tagNames := strings.Split(tags, "+")
	return &ExploreRequest{
		TagNames: tagNames,
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
