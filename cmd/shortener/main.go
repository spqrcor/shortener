package main

import (
	"io"
	"net/http"
	"net/url"
	"shortener/internal/app"
)

const baseUrl = "http://localhost:8080"
const shortStringLength = 6

var store = map[string]string{}

func createShortHandler(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	var bodyBytes []byte
	var err error
	if req.Body != nil {
		bodyBytes, err = io.ReadAll(req.Body)
		if err != nil {
			panic(err)
		}
		err = req.Body.Close()
		if err != nil {
			panic(err)
		}
	}

	_, err = url.ParseRequestURI(string(bodyBytes))
	if err == nil {
		res.WriteHeader(http.StatusCreated)
		res.Header().Set("Content-Type", "text/plain")
		genUrl := app.GenerateShortUrl(baseUrl, shortStringLength)
		store[genUrl] = string(bodyBytes)
		_, err = res.Write([]byte(genUrl))
		if err != nil {
			panic(err)
		}
	}
	res.WriteHeader(http.StatusBadRequest)
}

func searchShortHandler(res http.ResponseWriter, req *http.Request) {
	redirectUrl, ok := store[baseUrl+req.URL.Path]
	if req.Method == http.MethodGet && ok {
		res.Header().Set("Location", redirectUrl)
		res.WriteHeader(http.StatusTemporaryRedirect)
	} else {
		res.WriteHeader(http.StatusBadRequest)
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc(`/`, createShortHandler)
	mux.HandleFunc(`/{id}`, searchShortHandler)

	err := http.ListenAndServe(`:8080`, mux)
	if err != nil {
		panic(err)
	}
}
