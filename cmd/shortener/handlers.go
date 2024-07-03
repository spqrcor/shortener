package main

import (
	"io"
	"net/http"
	"net/url"
	"shortener/internal/app"
	"strings"
)

const baseURL = "http://localhost:8080"
const shortStringLength = 6

func createShortHandler(store map[string]string) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodPost || !strings.Contains(req.Header.Get("Content-Type"), "text/plain") {
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
			genURL := app.GenerateShortURL(baseURL, shortStringLength)
			store[genURL] = string(bodyBytes)
			_, err = res.Write([]byte(genURL))
			if err != nil {
				panic(err)
			}
		}
		res.WriteHeader(http.StatusBadRequest)
	}
}

func searchShortHandler(store map[string]string) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		redirectURL, ok := store[baseURL+req.URL.Path]
		if req.Method == http.MethodGet && ok {
			http.Redirect(res, req, redirectURL, http.StatusTemporaryRedirect)
		} else {
			res.WriteHeader(http.StatusBadRequest)
		}
	}
}
