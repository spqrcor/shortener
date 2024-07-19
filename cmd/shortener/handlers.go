package main

import (
	"io"
	"net/http"
	"net/url"
	"shortener/internal/app"
	"strings"
)

const shortStringLength = 6

func createShortHandler(store map[string]string, cfg Config) http.HandlerFunc {
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
				http.Error(res, err.Error(), http.StatusBadRequest)
			}
			err = req.Body.Close()
			if err != nil {
				http.Error(res, err.Error(), http.StatusBadRequest)
			}
		}

		_, err = url.ParseRequestURI(string(bodyBytes))
		if err == nil {
			res.WriteHeader(http.StatusCreated)
			res.Header().Set("Content-Type", "text/plain")
			genURL := app.GenerateShortURL(cfg.baseURL, shortStringLength)
			store[genURL] = string(bodyBytes)
			_, err = res.Write([]byte(genURL))
			if err != nil {
				http.Error(res, err.Error(), http.StatusBadRequest)
			}
		}
		res.WriteHeader(http.StatusBadRequest)
	}
}

func searchShortHandler(store map[string]string, cfg Config) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		redirectURL, ok := store[cfg.baseURL+req.URL.Path]
		if req.Method == http.MethodGet && ok {
			http.Redirect(res, req, redirectURL, http.StatusTemporaryRedirect)
		} else {
			res.WriteHeader(http.StatusBadRequest)
		}
	}
}
