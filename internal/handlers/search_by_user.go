package handlers

import (
	"encoding/json"
	"net/http"
	"shortener/internal/storage"
)

// SearchByUserHandler обработчик роута GET /api/user/urls
func SearchByUserHandler(s storage.Storage) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodGet {
			res.WriteHeader(http.StatusInternalServerError)
			return
		}

		output, err := s.FindByUser(req.Context())
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}
		resp, err := json.Marshal(output)
		if err != nil {
			http.Error(res, err.Error(), http.StatusBadRequest)
			return
		}

		res.Header().Set("Content-Type", "application/json; charset=utf-8")
		if len(output) > 0 {
			res.WriteHeader(http.StatusOK)
		} else {
			res.WriteHeader(http.StatusNoContent)
		}
		_, err = res.Write(resp)
		if err != nil {
			http.Error(res, err.Error(), http.StatusBadRequest)
			return
		}
	}
}
