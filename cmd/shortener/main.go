package main

import (
	"net/http"
)

var store = map[string]string{}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc(`/`, createShortHandler(store))
	mux.HandleFunc(`/{id}`, searchShortHandler(store))

	err := http.ListenAndServe(`:8080`, mux)
	if err != nil {
		panic(err)
	}
}
