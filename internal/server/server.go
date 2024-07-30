package server

import (
	"compress/gzip"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
	"shortener/internal/config"
	"shortener/internal/handlers"
	"shortener/internal/logger"
)

func Start() {
	r := chi.NewRouter()
	r.Use(middleware.Compress(5, "application/json", "text/html"))
	r.Use(getBodyMiddleware)

	r.Post("/", logger.RequestLogger(handlers.CreateShortHandler()))
	r.Post("/api/shorten", logger.RequestLogger(handlers.CreateJSONShortHandler()))
	r.Get("/{id}", logger.RequestLogger(handlers.SearchShortHandler()))
	r.HandleFunc(`/*`, func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(http.StatusBadRequest)
	})

	log.Fatal(http.ListenAndServe(config.Cfg.Addr, r))
}

func getBodyMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		if r.Header.Get(`Content-Encoding`) == `gzip` {
			gz, err := gzip.NewReader(r.Body)
			if err == nil {
				r.Body = gz
			}
			err = gz.Close()
			if err != nil {
				logger.Log.Error(err.Error())
			}
		}
		next.ServeHTTP(rw, r)
	})
}
