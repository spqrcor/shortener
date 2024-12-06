package handlers

import (
	"context"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"shortener/internal/mocks"
	"shortener/internal/storage"
	"testing"
)

func TestInternalStatHandler(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	m := mocks.NewMockStorage(mockCtrl)
	m.EXPECT().Stat(context.Background()).Return(storage.Stat{Urls: 1, Users: 1}, nil).MaxTimes(1)
	m.EXPECT().Stat(context.Background()).Return(storage.Stat{}, fmt.Errorf("500")).MinTimes(1)

	tests := []struct {
		name       string
		method     string
		statusCode int
	}{
		{
			name:       "method error",
			method:     http.MethodPost,
			statusCode: http.StatusInternalServerError,
		},
		{
			name:       "success",
			method:     http.MethodGet,
			statusCode: http.StatusOK,
		},
		{
			name:       "method error",
			method:     http.MethodGet,
			statusCode: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rw := httptest.NewRecorder()
			req := httptest.NewRequest(tt.method, "/api/internal/stats", nil)
			InternalStatHandler(m)(rw, req)

			resp := rw.Result()
			assert.Equal(t, tt.statusCode, resp.StatusCode, "Error http status code")
			_ = resp.Body.Close()
		})
	}
}
