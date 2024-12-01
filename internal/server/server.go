// Package server http server
package server

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
	"net/http"
	"os"
	"os/signal"
	"shortener/internal/authenticate"
	"shortener/internal/config"
	"shortener/internal/grpcHandlers"
	"shortener/internal/handlers"
	"shortener/internal/server/proto"
	"shortener/internal/services"
	"shortener/internal/storage"
	"syscall"
)

// AppServer тип server
type AppServer struct {
	config      config.Config
	logger      *zap.Logger
	storage     storage.Storage
	auth        authenticate.Auth
	batchRemove services.BatchRemover
}

// NewServer создание HTTPServer, opts набор параметров
func NewServer(opts ...func(*AppServer)) *AppServer {
	server := &AppServer{}
	for _, opt := range opts {
		opt(server)
	}
	return server
}

// WithLogger добавление logger
func WithLogger(logger *zap.Logger) func(*AppServer) {
	return func(a *AppServer) {
		a.logger = logger
	}
}

// WithConfig добавление config
func WithConfig(config config.Config) func(*AppServer) {
	return func(a *AppServer) {
		a.config = config
	}
}

// WithStorage добавление storage
func WithStorage(storage storage.Storage) func(*AppServer) {
	return func(a *AppServer) {
		a.storage = storage
	}
}

// WithAuthenticate добавление auth
func WithAuthenticate(auth authenticate.Auth) func(*AppServer) {
	return func(a *AppServer) {
		a.auth = auth
	}
}

// WithBatchRemove добавление batchRemove
func WithBatchRemove(batchRemove services.BatchRemover) func(*AppServer) {
	return func(a *AppServer) {
		a.batchRemove = batchRemove
	}
}

// NewHTTPServer создание http сервера
func (a *AppServer) NewHTTPServer() *http.Server {
	r := chi.NewRouter()
	r.Use(loggerMiddleware(a.logger))
	r.Use(middleware.Compress(5, "application/json", "text/html"))
	r.Use(getBodyMiddleware(a.logger))
	r.Use(authenticateMiddleware(a.logger, a.auth, a.config.TrustedSubnet))
	r.Mount("/debug", middleware.Profiler())

	r.Post("/", handlers.CreateShortHandler(a.storage))
	r.Post("/api/shorten", handlers.CreateJSONShortHandler(a.storage))
	r.Post("/api/shorten/batch", handlers.CreateJSONBatchHandler(a.storage))
	r.Get("/{id}", handlers.SearchShortHandler(a.storage))
	r.Get("/ping", handlers.PingHandler())
	r.Get("/api/user/urls", handlers.SearchByUserHandler(a.storage))
	r.Get("/api/internal/stats", handlers.InternalStatHandler(a.storage))
	r.Delete("/api/user/urls", handlers.RemoveShortHandler(a.batchRemove))

	r.HandleFunc(`/*`, func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(http.StatusBadRequest)
	})

	return &http.Server{
		Handler: r,
		Addr:    a.config.Addr,
	}
}

// Start запуск приложения
func (a *AppServer) Start() {
	httpServer := a.NewHTTPServer()
	grpcServer := grpc.NewServer()
	proto.RegisterURLShortenerServiceServer(grpcServer, &grpcHandlers.ShortenerServer{
		Storage:     a.storage,
		Logger:      a.logger,
		BatchRemove: a.batchRemove,
	})

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)

	go func() {
		<-stop
		if err := httpServer.Shutdown(context.Background()); err != nil {
			a.logger.Error(err.Error())
		}
		if err := a.storage.ShutDown(); err != nil {
			a.logger.Error(err.Error())
		}
		grpcServer.Stop()

	}()

	if a.config.EnableTLS {
		if err := initCertificate(); err != nil {
			a.logger.Error(err.Error())
		}
		if err := httpServer.ListenAndServeTLS(certCfg.certPath, certCfg.keyPath); err != nil {
			a.logger.Error(err.Error())
		}
	} else {
		if err := httpServer.ListenAndServe(); err != nil {
			a.logger.Error(err.Error())
		}
	}

	listen, err := net.Listen("tcp", a.config.GRPCAddr)
	if err != nil {
		a.logger.Error(err.Error())
	}

	err = grpcServer.Serve(listen)
	if err != nil {
		a.logger.Error(err.Error())
	}

	a.logger.Info("graceful shutdown")
}
