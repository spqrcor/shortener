package handlers

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"shortener/internal/mocks"
	"shortener/internal/storage"
	"testing"
)

func TestSearchShortHandler(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	m := mocks.NewMockStorage(mockCtrl)
	m.EXPECT().Find(context.Background(), "/success").Return("xxx", nil).AnyTimes()
	m.EXPECT().Find(context.Background(), "/delete").Return("", storage.ErrShortIsRemoved).AnyTimes()
	m.EXPECT().Find(context.Background(), "/notfound").Return("", storage.ErrKeyNotFound).AnyTimes()

	tests := []struct {
		name       string
		method     string
		url        string
		statusCode int
	}{
		{
			name:       "method error",
			method:     http.MethodPost,
			url:        "/xxxxxx",
			statusCode: http.StatusBadRequest,
		},
		{
			name:       "success",
			method:     http.MethodGet,
			url:        "/success",
			statusCode: http.StatusTemporaryRedirect,
		},
		{
			name:       "is delete",
			method:     http.MethodGet,
			url:        "/delete",
			statusCode: http.StatusGone,
		},
		{
			name:       "key not found",
			method:     http.MethodGet,
			url:        "/notfound",
			statusCode: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rw := httptest.NewRecorder()
			req := httptest.NewRequest(tt.method, tt.url, nil)
			SearchShortHandler(m)(rw, req)

			resp := rw.Result()
			assert.Equal(t, tt.statusCode, resp.StatusCode, "Error http status code")
			_ = resp.Body.Close()
		})
	}

}
