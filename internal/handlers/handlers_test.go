package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"shortener/internal/storage"
	"strings"
	"testing"
)

func Test_createShortHandler(t *testing.T) {
	type want struct {
		code int
	}
	tests := []struct {
		name        string
		method      string
		body        string
		contentType string
		want        want
	}{
		{
			"NOT POST",
			http.MethodGet,
			"",
			"text/plain",
			want{
				code: http.StatusBadRequest,
			},
		},
		{
			"POST current",
			http.MethodPost,
			"https://ya.ru",
			"text/plain",
			want{
				code: http.StatusCreated,
			},
		},
		{
			"POST error contentType",
			http.MethodPost,
			"https://ya.ru",
			"text/html",
			want{
				code: http.StatusBadRequest,
			},
		},
		{
			"POST error link",
			http.MethodPost,
			"4http://ya.ru",
			"text/plain",
			want{
				code: http.StatusBadRequest,
			},
		},
	}
	storage.CreateMemoryStorage()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := httptest.NewRequest(tt.method, "/", strings.NewReader(tt.body))
			request.Header.Add("Content-Type", tt.contentType)
			w := httptest.NewRecorder()
			CreateShortHandler(w, request)
			result := w.Result()

			assert.Equal(t, tt.want.code, result.StatusCode)
			if result.StatusCode == http.StatusCreated {
				bodyBytes, _ := io.ReadAll(result.Body)
				err := result.Body.Close()
				if err == nil {
					assert.NotEmpty(t, string(bodyBytes))
				}
			}
		})
	}
}

func Test_searchShortHandler(t *testing.T) {
	type want struct {
		code int
	}
	tests := []struct {
		name   string
		method string
		target string
		want   want
	}{
		{
			"NOT GET",
			http.MethodPost,
			"/test",
			want{
				code: http.StatusBadRequest,
			},
		},
		{
			name:   "GET current",
			method: http.MethodGet,
			target: "/FaKeLiNk",
			want: want{
				code: http.StatusTemporaryRedirect,
			},
		},
		{
			name:   "GET error",
			method: http.MethodGet,
			target: "/XxXxXxX",
			want: want{
				code: http.StatusBadRequest,
			},
		},
	}
	storage.CreateMemoryStorage()
	genURL, _ := storage.Source.Add(context.Background(), "https://ya.ru")
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.name == "GET current" {
				tt.target = genURL
			}
			request := httptest.NewRequest(tt.method, tt.target, nil)
			w := httptest.NewRecorder()
			SearchShortHandler(w, request)
			result := w.Result()

			_ = result.Body.Close()
			assert.Equal(t, tt.want.code, result.StatusCode)
		})
	}
}

func Test_createJsonShortHandler(t *testing.T) {
	type want struct {
		code int
	}
	tests := []struct {
		name        string
		method      string
		body        []byte
		contentType string
		want        want
	}{
		{
			"NOT POST",
			http.MethodGet,
			[]byte(``),
			"text/plain",
			want{
				code: http.StatusBadRequest,
			},
		},
		{
			"POST error contentType",
			http.MethodPost,
			[]byte(`{"url":"https://ya.ru"}`),
			"text/plain",
			want{
				code: http.StatusBadRequest,
			},
		},
		{
			"POST error json param",
			http.MethodPost,
			[]byte(`{"url1":"https://ya.ru"}`),
			"application/json",
			want{
				code: http.StatusBadRequest,
			},
		},
		{
			"POST error json value",
			http.MethodPost,
			[]byte(`{"url":"1https://ya.ru"}`),
			"application/json",
			want{
				code: http.StatusBadRequest,
			},
		},
		{
			"POST current",
			http.MethodPost,
			[]byte(`{"url":"https://ya.ru"}`),
			"application/json",
			want{
				code: http.StatusCreated,
			},
		},
	}
	storage.CreateMemoryStorage()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := httptest.NewRequest(tt.method, "/", bytes.NewReader(tt.body))
			request.Header.Add("Content-Type", tt.contentType)
			w := httptest.NewRecorder()
			CreateJSONShortHandler(w, request)
			result := w.Result()

			assert.Equal(t, tt.want.code, result.StatusCode)
			if result.StatusCode == http.StatusCreated {
				var output outputJSONData
				decoder := json.NewDecoder(result.Body)
				err := decoder.Decode(&output)
				if err == nil {
					assert.NotEmpty(t, output.Result)
				}
			}
			_ = result.Body.Close()
		})
	}
}

func Test_isValidInputParams(t *testing.T) {
	tests := []struct {
		name                   string
		requestMethod          string
		currentMethod          string
		requestContentType     string
		currentContentType     string
		requestContentEncoding string
		want                   bool
	}{
		{
			"NOT POST",
			http.MethodGet,
			http.MethodPost,
			"text/plain",
			"text/plain",
			"",
			false,
		}, {
			"NOT error contentType",
			http.MethodPost,
			http.MethodPost,
			"text/xml",
			"text/plain",
			"",
			false,
		},
		{
			"POST current",
			http.MethodPost,
			http.MethodPost,
			"text/plain",
			"text/plain",
			"",
			true,
		},
		{
			"POST current with encoding",
			http.MethodPost,
			http.MethodPost,
			"text/plain",
			"text/plain",
			"gzip",
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := httptest.NewRequest(tt.requestMethod, "/", strings.NewReader(""))
			request.Header.Add("Content-Type", tt.requestContentType)
			if tt.requestContentEncoding != "" {
				request.Header.Add("Content-Encoding", tt.requestContentEncoding)
			}
			assert.Equal(t, tt.want, isValidInputParams(request, inputParams{Method: tt.currentMethod, ContentType: tt.currentContentType}))
		})
	}
}

func TestPingHandler(t *testing.T) {
	tests := []struct {
		name           string
		requestMethod  string
		responceStatus int
	}{
		{
			"NOT current",
			http.MethodGet,
			http.StatusOK,
		},
		{
			"GET current",
			http.MethodPost,
			http.StatusInternalServerError,
		},
	}
	storage.CreateMemoryStorage()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := httptest.NewRequest(tt.requestMethod, "/ping", strings.NewReader(""))
			w := httptest.NewRecorder()
			PingHandler(w, request)
			result := w.Result()
			assert.Equal(t, tt.responceStatus, result.StatusCode)
		})
	}
}

func TestCreateJSONBatchHandler(t *testing.T) {
	type want struct {
		code int
	}
	tests := []struct {
		name        string
		method      string
		body        []byte
		contentType string
		want        want
	}{
		{
			"NOT POST",
			http.MethodGet,
			[]byte(``),
			"text/plain",
			want{
				code: http.StatusBadRequest,
			},
		},
		{
			"POST error contentType",
			http.MethodPost,
			[]byte(`{"url":"https://ya.ru"}`),
			"text/plain",
			want{
				code: http.StatusBadRequest,
			},
		},
		{
			"POST empty rows",
			http.MethodPost,
			[]byte(`{}`),
			"application/json",
			want{
				code: http.StatusBadRequest,
			},
		},
		{
			"POST current",
			http.MethodPost,
			[]byte(`[{"correlation_id": "b9253cb9-03e9-4850-a3cb-16e84e9f8a37", "original_url": "http://lenta.ru"}]`),
			"application/json",
			want{
				code: http.StatusCreated,
			},
		},
	}
	storage.CreateMemoryStorage()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := httptest.NewRequest(tt.method, "/api/shorten/batch", bytes.NewReader(tt.body))
			request.Header.Add("Content-Type", tt.contentType)
			w := httptest.NewRecorder()
			CreateJSONBatchHandler(w, request)
			result := w.Result()

			assert.Equal(t, tt.want.code, result.StatusCode)
			if result.StatusCode == http.StatusCreated {
				var output outputJSONData
				decoder := json.NewDecoder(result.Body)
				err := decoder.Decode(&output)
				if err == nil {
					assert.NotEmpty(t, output.Result)
				}
			}
			_ = result.Body.Close()
		})
	}
}
