package main

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

var store = map[string]string{}

func main() {
	r := chi.NewRouter()
	r.Post("/", createShortHandler(store))
	r.Get("/{id}", searchShortHandler(store))
	r.HandleFunc(`/*`, func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(http.StatusBadRequest)
	})
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		panic(err)
	}
}
