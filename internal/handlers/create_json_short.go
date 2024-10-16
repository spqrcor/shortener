package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"shortener/internal/storage"
)

// CreateJSONShortHandler обработчик роута /api/shorten
func CreateJSONShortHandler(s storage.Storage) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		if !isValidInputParams(req, inputParams{Method: http.MethodPost, ContentType: "application/json"}) {
			res.WriteHeader(http.StatusBadRequest)
			return
		}

		var input inputJSONData
		var buf bytes.Buffer
		body := req.Body

		_, err := buf.ReadFrom(body)
		if err != nil {
			res.WriteHeader(http.StatusBadRequest)
			return
		}
		if err = json.Unmarshal(buf.Bytes(), &input); err != nil {
			res.WriteHeader(http.StatusBadRequest)
			return
		}

		var output outputJSONData
		successStatus := http.StatusCreated
		output.Result, err = s.Add(req.Context(), input.URL)
		if errors.Is(err, storage.ErrURLExists) {
			successStatus = http.StatusConflict
		} else if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}

		resp, err := json.Marshal(output)
		if err != nil {
			http.Error(res, err.Error(), http.StatusBadRequest)
			return
		}

		res.Header().Set("Content-Type", "application/json; charset=utf-8")
		res.WriteHeader(successStatus)
		_, err = res.Write(resp)
		if err != nil {
			http.Error(res, err.Error(), http.StatusBadRequest)
			return
		}
	}
}
