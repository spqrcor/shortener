package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"shortener/internal/config"
	"shortener/internal/db"
	"shortener/internal/storage"
	"strings"
)

type inputJSONData struct {
	URL string `json:"url,omitempty"`
}

type outputJSONData struct {
	Result string `json:"result,omitempty"`
}

type inputParams struct {
	Method      string
	ContentType string
}

func CreateShortHandler(res http.ResponseWriter, req *http.Request) {
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
	genURL, err := storage.Source.Add(req.Context(), string(bodyBytes))
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

func SearchShortHandler(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		res.WriteHeader(http.StatusBadRequest)
		return
	}
	redirectURL, err := storage.Source.Find(req.Context(), req.URL.Path)
	if errors.Is(err, storage.ErrShortIsRemoved) {
		http.Error(res, err.Error(), http.StatusGone)
		return
	} else if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(res, req, redirectURL, http.StatusTemporaryRedirect)
}

func CreateJSONShortHandler(res http.ResponseWriter, req *http.Request) {
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
	output.Result, err = storage.Source.Add(req.Context(), input.URL)
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

func CreateJSONBatchHandler(res http.ResponseWriter, req *http.Request) {
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

	output, err := storage.Source.BatchAdd(req.Context(), input)
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

func PingHandler(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet || config.Cfg.DatabaseDSN == "" {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err := db.Connect()
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
	res.WriteHeader(http.StatusOK)
}

func SearchByUserHandler(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet || config.Cfg.DatabaseDSN == "" {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	output, err := storage.Source.FindByUser(req.Context())
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

func RemoveShortHandler(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodDelete || config.Cfg.DatabaseDSN == "" {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	var input []string
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

	if err := storage.Source.Remove(req.Context(), input); err != nil {
		res.WriteHeader(http.StatusBadRequest)
		return
	}
	res.WriteHeader(http.StatusAccepted)
}

func isValidInputParams(req *http.Request, params inputParams) bool {
	if req.Method != params.Method {
		return false
	}
	if req.Header.Get(`Content-Encoding`) != `gzip` && !strings.Contains(req.Header.Get("Content-Type"), params.ContentType) {
		return false
	}
	return true
}
