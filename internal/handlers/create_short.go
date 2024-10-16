package handlers

import (
	"errors"
	"io"
	"net/http"
	"shortener/internal/storage"
)

// CreateShortHandler обработчик роута /
func CreateShortHandler(s storage.Storage) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		if !isValidInputParams(req, inputParams{Method: http.MethodPost, ContentType: "text/plain"}) {
			res.WriteHeader(http.StatusBadRequest)
			return
		}

		var bodyBytes []byte
		bodyBytes, err := io.ReadAll(req.Body)
		if err != nil {
			http.Error(res, err.Error(), http.StatusBadRequest)
			return
		}
		if err = req.Body.Close(); err != nil {
			http.Error(res, err.Error(), http.StatusBadRequest)
			return
		}

		successStatus := http.StatusCreated
		genURL, err := s.Add(req.Context(), string(bodyBytes))
		if errors.Is(err, storage.ErrURLExists) {
			successStatus = http.StatusConflict
		} else if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}

		res.Header().Set("Content-Type", "text/plain")
		res.WriteHeader(successStatus)
		_, err = res.Write([]byte(genURL))
		if err != nil {
			http.Error(res, err.Error(), http.StatusBadRequest)
			return
		}
	}
}
