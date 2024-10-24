// Package server http server
package server

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go.uber.org/zap"
	"net/http"
	"shortener/internal/authenticate"
	"shortener/internal/config"
	"shortener/internal/handlers"
	"shortener/internal/services"
	"shortener/internal/storage"
)

// HTTPServer тип http server
type HTTPServer struct {
	config      config.Config
	logger      *zap.Logger
	storage     storage.Storage
	auth        authenticate.Auth
	batchRemove services.BatchRemover
}

// NewServer создание HTTPServer, opts набор параметров
func NewServer(opts ...func(*HTTPServer)) *HTTPServer {
	server := &HTTPServer{}
	for _, opt := range opts {
		opt(server)
	}
	return server
}

// WithLogger добавление logger
func WithLogger(logger *zap.Logger) func(*HTTPServer) {
	return func(h *HTTPServer) {
		h.logger = logger
	}
}

// WithConfig добавление config
func WithConfig(config config.Config) func(*HTTPServer) {
	return func(h *HTTPServer) {
		h.config = config
	}
}

// WithStorage добавление storage
func WithStorage(storage storage.Storage) func(*HTTPServer) {
	return func(h *HTTPServer) {
		h.storage = storage
	}
}

// WithAuthenticate добавление auth
func WithAuthenticate(auth authenticate.Auth) func(*HTTPServer) {
	return func(h *HTTPServer) {
		h.auth = auth
	}
}

// WithBatchRemove добавление batchRemove
func WithBatchRemove(batchRemove services.BatchRemover) func(*HTTPServer) {
	return func(h *HTTPServer) {
		h.batchRemove = batchRemove
	}
}

// Start старт сервера
func (s *HTTPServer) Start() error {
	r := chi.NewRouter()
	r.Use(loggerMiddleware(s.logger))
	r.Use(middleware.Compress(5, "application/json", "text/html"))
	r.Use(getBodyMiddleware(s.logger))
	r.Use(authenticateMiddleware(s.logger, s.auth))
	r.Mount("/debug", middleware.Profiler())

	r.Post("/", handlers.CreateShortHandler(s.storage))
	r.Post("/api/shorten", handlers.CreateJSONShortHandler(s.storage))
	r.Post("/api/shorten/batch", handlers.CreateJSONBatchHandler(s.storage))
	r.Get("/{id}", handlers.SearchShortHandler(s.storage))
	r.Get("/ping", handlers.PingHandler())
	r.Get("/api/user/urls", handlers.SearchByUserHandler(s.storage))
	r.Delete("/api/user/urls", handlers.RemoveShortHandler(s.batchRemove))

	r.HandleFunc(`/*`, func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(http.StatusBadRequest)
	})
	return http.ListenAndServe(s.config.Addr, r)
}
