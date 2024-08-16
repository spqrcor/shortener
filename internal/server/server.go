package server

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
	"shortener/internal/config"
	"shortener/internal/handlers"
)

func Start() {
	r := chi.NewRouter()
	r.Use(authenticateMiddleware)
	r.Use(loggerMiddleware)
	r.Use(middleware.Compress(5, "application/json", "text/html"))
	r.Use(getBodyMiddleware)

	r.Post("/", handlers.CreateShortHandler)
	r.Post("/api/shorten", handlers.CreateJSONShortHandler)
	r.Post("/api/shorten/batch", handlers.CreateJSONBatchHandler)
	r.Get("/{id}", handlers.SearchShortHandler)
	r.Get("/ping", handlers.PingHandler)
	r.Get("/api/user/urls", handlers.SearchByUserHandler)
	r.HandleFunc(`/*`, func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(http.StatusBadRequest)
	})

	log.Fatal(http.ListenAndServe(config.Cfg.Addr, r))
}
