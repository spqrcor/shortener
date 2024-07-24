package handlers

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"io"
	"net/http"
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

func CreateShortHandler() http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		if !isValidInputParams(req, inputParams{Method: http.MethodPost, ContentType: "text/plain"}) {
			res.WriteHeader(http.StatusBadRequest)
			return
		}

		var bodyBytes []byte
		var err error
		body, err := getBody(req)
		if err != nil {
			res.WriteHeader(http.StatusBadRequest)
			return
		}

		bodyBytes, err = io.ReadAll(body)
		if err != nil {
			http.Error(res, err.Error(), http.StatusBadRequest)
			return
		}
		err = req.Body.Close()
		if err != nil {
			http.Error(res, err.Error(), http.StatusBadRequest)
			return
		}

		genURL, err := storage.Add(string(bodyBytes))
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

func CreateJSONShortHandler() http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		if !isValidInputParams(req, inputParams{Method: http.MethodPost, ContentType: "application/json"}) {
			res.WriteHeader(http.StatusBadRequest)
			return
		}

		var input inputJSONData
		var buf bytes.Buffer
		body, err := getBody(req)
		if err != nil {
			res.WriteHeader(http.StatusBadRequest)
			return
		}

		_, err = buf.ReadFrom(body)
		if err != nil {
			res.WriteHeader(http.StatusBadRequest)
			return
		}
		if err = json.Unmarshal(buf.Bytes(), &input); err != nil {
			res.WriteHeader(http.StatusBadRequest)
			return
		}

		var output outputJSONData
		output.Result, err = storage.Add(input.URL)
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

func getBody(req *http.Request) (io.Reader, error) {
	var reader io.Reader

	if req.Header.Get(`Content-Encoding`) == `gzip` {
		gz, err := gzip.NewReader(req.Body)
		if err != nil {
			return nil, err
		}
		reader = gz
		err = gz.Close()
		if err != nil {
			return nil, err
		}
	} else {
		reader = req.Body
	}
	return reader, nil
}
