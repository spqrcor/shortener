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

func TestCreateShortHandler(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	storage := mocks.NewMockStorage(mockCtrl)

	tests := []struct {
		name        string
		method      string
		contentType string
		body        []byte
		statusCode  int
	}{
		{
			name:        "method error",
			method:      "GET",
			contentType: "text/plain",
			body:        []byte(`<num>3333</num>`),
			statusCode:  http.StatusBadRequest,
		},
		{
			name:        "content type error",
			method:      "POST",
			contentType: "application/json",
			body:        []byte(`<num>3333</num>`),
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

			req := httptest.NewRequest(tt.method, "/api/user/orders", body)
			req.Header.Add("Content-Type", tt.contentType)
			CreateShortHandler(storage)(rw, req)

			resp := rw.Result()
			assert.Equal(t, tt.statusCode, resp.StatusCode, "Error http status code")
			_ = resp.Body.Close()
		})
	}
}
