package handlers

import (
	"bytes"
	"context"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"shortener/internal/app"
	"shortener/internal/mocks"
	"shortener/internal/storage"
	"testing"
)

func TestCreateShortHandler(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	m := mocks.NewMockStorage(mockCtrl)
	m.EXPECT().Add(context.Background(), "https://ya.ru").Return("-", nil).MaxTimes(1)
	m.EXPECT().Add(context.Background(), "https://ya.ru").Return("", storage.ErrURLExists).MinTimes(1)
	m.EXPECT().Add(context.Background(), "1https://ya.ru").Return("", app.ErrURLFormat).AnyTimes()

	tests := []struct {
		name        string
		method      string
		contentType string
		body        []byte
		statusCode  int
	}{
		{
			name:        "method error",
			method:      http.MethodGet,
			contentType: "text/plain",
			body:        []byte(`<num>3333</num>`),
			statusCode:  http.StatusBadRequest,
		},
		{
			name:        "content type error",
			method:      http.MethodPost,
			contentType: "application/json",
			body:        []byte(`<num>3333</num>`),
			statusCode:  http.StatusBadRequest,
		},
		{
			name:        "success",
			method:      http.MethodPost,
			contentType: "text/plain",
			body:        []byte(`https://ya.ru`),
			statusCode:  http.StatusCreated,
		},
		{
			name:        "conflict",
			method:      http.MethodPost,
			contentType: "text/plain",
			body:        []byte(`https://ya.ru`),
			statusCode:  http.StatusConflict,
		},
		{
			name:        "invalid url error",
			method:      http.MethodPost,
			contentType: "text/plain",
			body:        []byte(`1https://ya.ru`),
			statusCode:  http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rw := httptest.NewRecorder()

			body := bytes.NewBuffer([]byte{})
			if tt.method == http.MethodPost {
				body = bytes.NewBuffer(tt.body)
			}

			req := httptest.NewRequest(tt.method, "/api/user/orders", body)
			req.Header.Add("Content-Type", tt.contentType)
			CreateShortHandler(m)(rw, req)

			resp := rw.Result()
			assert.Equal(t, tt.statusCode, resp.StatusCode, "Error http status code")
			_ = resp.Body.Close()
		})
	}
}
