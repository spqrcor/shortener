package handlers

import (
	"encoding/json"
	"net/http"
	"shortener/internal/storage"
)

// InternalStatHandler обработчик роута GET /api/internal/stats
func InternalStatHandler(s storage.Storage) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodGet {
			res.WriteHeader(http.StatusInternalServerError)
			return
		}

		stat, err := s.Stat(req.Context())
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}
		res.Header().Set("Content-Type", "application/json; charset=utf-8")
		res.WriteHeader(http.StatusOK)
		enc := json.NewEncoder(res)
		if err := enc.Encode(stat); err != nil {
			res.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}
