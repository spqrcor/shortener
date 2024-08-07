package handlers

import (
	"bytes"
	"encoding/json"
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
	body := req.Body

	bodyBytes, err := io.ReadAll(body)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
	if err = req.Body.Close(); err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	genURL, err := storage.Source.Add(string(bodyBytes))
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	res.Header().Set("Content-Type", "text/plain")
	res.WriteHeader(http.StatusCreated)
	_, err = res.Write([]byte(genURL))
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
}

func SearchShortHandler(res http.ResponseWriter, req *http.Request) {
	redirectURL, err := storage.Source.Find(req.URL.Path)
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
	output.Result, err = storage.Source.Add(input.URL)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
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

func CreateJSONBatchHandler(res http.ResponseWriter, req *http.Request) {
	if !isValidInputParams(req, inputParams{Method: http.MethodPost, ContentType: "application/json"}) {
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	var input []storage.BatchInputParams
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
	if len(input) == 0 {
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	output, err := storage.Source.BatchAdd(input)
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

func isValidInputParams(req *http.Request, params inputParams) bool {
	if req.Method != params.Method {
		return false
	}
	if req.Header.Get(`Content-Encoding`) != `gzip` && !strings.Contains(req.Header.Get("Content-Type"), params.ContentType) {
		return false
	}
	return true
}
