package server

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
	"shortener/internal/config"
	"shortener/internal/handlers"
	"shortener/internal/logger"
)

func Start() {
	r := chi.NewRouter()
	r.Use(middleware.Compress(5, "application/json", "text/html"))

	r.Post("/", logger.RequestLogger(handlers.CreateShortHandler()))
	r.Post("/api/shorten", logger.RequestLogger(handlers.CreateJSONShortHandler()))
	r.Get("/{id}", logger.RequestLogger(handlers.SearchShortHandler()))
	r.HandleFunc(`/*`, func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(http.StatusBadRequest)
	})

	err := http.ListenAndServe(config.Cfg.Addr, r)
	if err != nil {
		logger.Log.Fatal(err.Error())
	}
}