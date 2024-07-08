package main

import (
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

func main() {
	store := map[string]string{}
	cfg := parseFlags()
	r := chi.NewRouter()
	r.Post("/", createShortHandler(store, cfg))
	r.Get("/{id}", searchShortHandler(store, cfg))
	r.HandleFunc(`/*`, func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(http.StatusBadRequest)
	})

	err := http.ListenAndServe(cfg.addr, r)
	if err != nil {
		log.Fatal(err)
	}
}
