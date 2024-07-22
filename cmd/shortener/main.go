package main

import (
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"shortener/internal/config"
	"shortener/internal/handlers"
)

func main() {
	config.ParseFlags()
	r := chi.NewRouter()
	r.Post("/", handlers.CreateShortHandler())
	r.Get("/{id}", handlers.SearchShortHandler())
	r.HandleFunc(`/*`, func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(http.StatusBadRequest)
	})

	err := http.ListenAndServe(config.Cfg.Addr, r)
	if err != nil {
		log.Fatal(err)
	}
}
