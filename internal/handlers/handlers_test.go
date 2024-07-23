package handlers

import (
	"bytes"
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
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := httptest.NewRequest(tt.method, "/", strings.NewReader(tt.body))
			request.Header.Add("Content-Type", tt.contentType)
			w := httptest.NewRecorder()
			h := CreateShortHandler()
			h(w, request)
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
	genURL, _ := storage.Add("https://ya.ru")
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.name == "GET current" {
				tt.target = genURL
			}
			request := httptest.NewRequest(tt.method, tt.target, nil)
			w := httptest.NewRecorder()
			h := SearchShortHandler()
			h(w, request)
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
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := httptest.NewRequest(tt.method, "/", bytes.NewReader(tt.body))
			request.Header.Add("Content-Type", tt.contentType)
			w := httptest.NewRecorder()
			h := CreateJSONShortHandler()
			h(w, request)
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
