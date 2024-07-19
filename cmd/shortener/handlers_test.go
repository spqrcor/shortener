package main

import (
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var testStore = map[string]string{"http://localhost:8080/FaKeLiNk": "https://ya.ru"}
var testCfg = Config{addr: ":8080", baseURL: "http://localhost:8080"}

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
			h := createShortHandler(testStore, testCfg)
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
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := httptest.NewRequest(tt.method, tt.target, nil)
			w := httptest.NewRecorder()
			h := searchShortHandler(testStore, testCfg)
			h(w, request)
			result := w.Result()

			err := result.Body.Close()
			if err != nil {
				assert.Equal(t, tt.want.code, result.StatusCode)
			}
		})
	}
}
