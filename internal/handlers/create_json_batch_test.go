package handlers

import (
	"bytes"
	"context"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"shortener/internal/mocks"
	"shortener/internal/storage"
	"testing"
)

func TestCreateJSONBatchHandler(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	m := mocks.NewMockStorage(mockCtrl)
	m.EXPECT().BatchAdd(context.Background(), []storage.BatchInputParams{
		{
			CorrelationID: "b9253cb9-03e9-4850-a3cb-16e84e9f8a37",
			URL:           "https://ya.ru",
		},
	}).Return([]storage.BatchOutputParams{
		{
			CorrelationID: "b9253cb9-03e9-4850-a3cb-16e84e9f8a37",
			ShortURL:      "http://localhost/xxxxxx",
		},
	}, nil).AnyTimes()

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
			contentType: "application/json",
			body:        []byte(`{"url":"https://ya.ru"}`),
			statusCode:  http.StatusBadRequest,
		},
		{
			name:        "content type error",
			method:      http.MethodPost,
			contentType: "text/plain",
			body:        []byte(`{"url":"https://ya.ru"}`),
			statusCode:  http.StatusBadRequest,
		},
		{
			name:        "success",
			method:      http.MethodPost,
			contentType: "application/json",
			body:        []byte(`[{"correlation_id": "b9253cb9-03e9-4850-a3cb-16e84e9f8a37", "original_url": "https://ya.ru"}]`),
			statusCode:  http.StatusCreated,
		},
		{
			name:        "empty input",
			method:      http.MethodPost,
			contentType: "application/json",
			body:        []byte(`[]`),
			statusCode:  http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rw := httptest.NewRecorder()

			body := bytes.NewBuffer([]byte{})
			if tt.method == http.MethodPost {
				body = bytes.NewBuffer(tt.body)
			}

			req := httptest.NewRequest(tt.method, "/api/shorten/batch", body)
			req.Header.Add("Content-Type", tt.contentType)
			CreateJSONBatchHandler(m)(rw, req)

			resp := rw.Result()
			assert.Equal(t, tt.statusCode, resp.StatusCode, "Error http status code")
			_ = resp.Body.Close()
		})
	}
}
