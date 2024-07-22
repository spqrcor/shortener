package handlers

import (
	"io"
	"net/http"
	"shortener/internal/storage"
	"strings"
)

func CreateShortHandler() http.HandlerFunc {
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
				return
			}
			err = req.Body.Close()
			if err != nil {
				http.Error(res, err.Error(), http.StatusBadRequest)
				return
			}
		}

		genURL, err := storage.Add(string(bodyBytes))
		if err != nil {
			http.Error(res, err.Error(), http.StatusBadRequest)
			return
		}

		res.WriteHeader(http.StatusCreated)
		res.Header().Set("Content-Type", "text/plain")
		_, err = res.Write([]byte(genURL))
		if err != nil {
			http.Error(res, err.Error(), http.StatusBadRequest)
			return
		}
		res.WriteHeader(http.StatusBadRequest)
	}
}

func SearchShortHandler() http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		redirectURL, err := storage.Find(req.URL.Path)
		if err != nil {
			http.Error(res, err.Error(), http.StatusBadRequest)
			return
		}
		if req.Method == http.MethodGet {
			http.Redirect(res, req, redirectURL, http.StatusTemporaryRedirect)
		} else {
			res.WriteHeader(http.StatusBadRequest)
		}
	}
}
