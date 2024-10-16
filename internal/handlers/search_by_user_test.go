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

func TestSearchByUserHandler(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	m := mocks.NewMockStorage(mockCtrl)
	m.EXPECT().FindByUser(context.Background()).Return([]storage.FindByUserOutputParams{}, nil).MaxTimes(1)
	m.EXPECT().FindByUser(context.Background()).Return([]storage.FindByUserOutputParams{
		{
			ShortURL:    "http://localhost/xxxxx",
			OriginalURL: "https://ya.ru",
		},
	}, nil).MinTimes(1)

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
			name:       "no content",
			method:     http.MethodGet,
			statusCode: http.StatusNoContent,
		},
		{
			name:       "content exists",
			method:     http.MethodGet,
			statusCode: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rw := httptest.NewRecorder()
			req := httptest.NewRequest(tt.method, "/api/user/urls", nil)
			SearchByUserHandler(m)(rw, req)

			resp := rw.Result()
			assert.Equal(t, tt.statusCode, resp.StatusCode, "Error http status code")
			_ = resp.Body.Close()
		})
	}
}
