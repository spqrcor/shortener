package server

import (
	"github.com/go-chi/chi/v5"
	"net/http"
	"shortener/internal/config"
	"shortener/internal/handlers"
	"shortener/internal/logger"
)

func Start() {
	r := chi.NewRouter()
	r.Post("/", logger.RequestLogger(handlers.CreateShortHandler()))
	r.Get("/{id}", logger.RequestLogger(handlers.SearchShortHandler()))
	r.HandleFunc(`/*`, func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(http.StatusBadRequest)
	})

	err := http.ListenAndServe(config.Cfg.Addr, r)
	if err != nil {
		logger.Log.Fatal(err.Error())
	}
}
