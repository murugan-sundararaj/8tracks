package main

import (
	"8tracks/explore"
	"8tracks/playlists"
	"8tracks/tags"
	"flag"
	"math/rand"
	"net/http"
	"os"
	"time"

	kitlog "github.com/go-kit/kit/log"
	level "github.com/go-kit/kit/log/level"
	"github.com/oklog/ulid"
)

func main() {
	var (
		addr = flag.String("listen.address", ":8080", "server listen address")
	)
	flag.Parse()

	logger := kitlog.NewLogfmtLogger(kitlog.NewSyncWriter(os.Stderr))
	logger = level.NewFilter(logger, level.AllowAll())
	logger = kitlog.With(logger, "ts", kitlog.DefaultTimestamp, "caller", kitlog.DefaultCaller)

	// loggers
	playlistLogger := kitlog.With(logger, "service", "playlist")
	tagLogger := kitlog.With(logger, "service", "tag")
	exploreLogger := kitlog.With(logger, "service", "explore")

	// dal
	tagDAL := tags.NewDAL()
	playlistDAL := playlists.NewDAL()

	// services
	tagSvc := tags.NewService(tagLogger, tagDAL)
	tagSvc = tags.NewLoggingService(tagSvc, tagLogger)

	playlistSvc := playlists.NewService(tagSvc, playlistLogger, playlistDAL)
	playlistSvc = playlists.NewLoggingService(playlistSvc, playlistLogger)

	exploreSvc := explore.NewService(exploreLogger, tagSvc, playlistSvc)
	exploreSvc = explore.NewLoggingService(exploreSvc, exploreLogger)

	entropy := rand.New(rand.NewSource(time.Now().UnixNano()))
	requestIDGen := func() string { return ulid.MustNew(ulid.Now(), entropy).String() }

	mux := http.NewServeMux()

	mux.Handle("/tags/", tags.MakeHandler(tagSvc, tagLogger, requestIDGen))
	mux.Handle("/playlists/", playlists.MakeHandler(playlistSvc, playlistLogger, requestIDGen))
	mux.Handle("/explore/", explore.MakeHandler(exploreSvc, exploreLogger, requestIDGen))

	level.Info(logger).Log("transport", "http", "address", *addr, "msg", "listening")

	srv := &http.Server{
		Addr:           *addr,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   10 * time.Second,
		IdleTimeout:    120 * time.Second,
		MaxHeaderBytes: 4096,
		Handler:        &maxBytesHandler{h: mux, n: 10 * 4096},
	}

	level.Error(logger).Log("err", srv.ListenAndServe())
}

// safe guard against malicious clients
type maxBytesHandler struct {
	h http.Handler
	n int64
}

func (h *maxBytesHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")

	if r.Method == "OPTIONS" {
		return
	}
	r.Body = http.MaxBytesReader(w, r.Body, h.n)
	h.h.ServeHTTP(w, r)
}
