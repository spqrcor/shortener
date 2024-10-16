package handlers

import (
	"bytes"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"shortener/internal/mocks"
	"testing"
)

func TestRemoveShortHandler(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	m := mocks.NewMockBatchRemover(mockCtrl)

	tests := []struct {
		name        string
		method      string
		contentType string
		body        []byte
		statusCode  int
	}{
		{
			name:        "method error",
			method:      http.MethodPost,
			contentType: "application/json",
			body:        []byte(`{"url":"https://ya.ru"}`),
			statusCode:  http.StatusInternalServerError,
		},
		{
			name:        "empty list error",
			method:      http.MethodDelete,
			contentType: "application/json",
			body:        []byte(`[]`),
			statusCode:  http.StatusBadRequest,
		},
		{
			name:        "success",
			method:      http.MethodDelete,
			contentType: "application/json",
			body:        []byte(`["xxxxxx"]`),
			statusCode:  http.StatusAccepted,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rw := httptest.NewRecorder()
			req := httptest.NewRequest(tt.method, "/api/user/urls", bytes.NewBuffer(tt.body))
			req.Header.Add("Content-Type", tt.contentType)
			RemoveShortHandler(m)(rw, req)

			resp := rw.Result()
			assert.Equal(t, tt.statusCode, resp.StatusCode, "Error http status code")
			_ = resp.Body.Close()
		})
	}
}
