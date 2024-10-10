package handlers

import (
	"errors"
	"net/http"
	"shortener/internal/storage"
)

func SearchShortHandler(s storage.Storage) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodGet {
			res.WriteHeader(http.StatusBadRequest)
			return
		}
		redirectURL, err := s.Find(req.Context(), req.URL.Path)
		if errors.Is(err, storage.ErrShortIsRemoved) {
			http.Error(res, err.Error(), http.StatusGone)
			return
		} else if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(res, req, redirectURL, http.StatusTemporaryRedirect)
	}
}
