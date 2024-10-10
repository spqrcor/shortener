package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"shortener/internal/storage"
)

func CreateJSONBatchHandler(s storage.Storage) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		if !isValidInputParams(req, inputParams{Method: http.MethodPost, ContentType: "application/json"}) {
			res.WriteHeader(http.StatusBadRequest)
			return
		}

		var input []storage.BatchInputParams
		var buf bytes.Buffer

		_, err := buf.ReadFrom(req.Body)
		if err != nil {
			res.WriteHeader(http.StatusBadRequest)
			return
		}
		if err = json.Unmarshal(buf.Bytes(), &input); err != nil {
			res.WriteHeader(http.StatusBadRequest)
			return
		}
		if len(input) == 0 {
			res.WriteHeader(http.StatusBadRequest)
			return
		}

		output, err := s.BatchAdd(req.Context(), input)
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
		res.WriteHeader(http.StatusCreated)
		_, err = res.Write(resp)
		if err != nil {
			http.Error(res, err.Error(), http.StatusBadRequest)
			return
		}
	}
}
