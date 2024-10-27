package server

import (
	"bytes"
	"compress/gzip"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"net/http"
	"net/http/httptest"
	"shortener/internal/config"
	"shortener/internal/handlers"
	"shortener/internal/logger"
	"shortener/internal/storage"
	"testing"
)

func Test_getBodyMiddleware(t *testing.T) {
	loggerRes, _ := logger.NewLogger(zap.InfoLevel)
	conf := config.NewConfig()
	store := storage.CreateMemoryStorage(conf)

	r := chi.NewRouter()
	r.Use(getBodyMiddleware(loggerRes))

	r.Post(`/`, handlers.CreateShortHandler(store))
	srv := httptest.NewServer(r)
	defer srv.Close()

	t.Run("work_without_gzip", func(t *testing.T) {
		r := httptest.NewRequest("POST", srv.URL+"/", nil)
		r.RequestURI = ""
		r.Header.Set("Content-Type", "text/html")
		resp, err := http.DefaultClient.Do(r)
		if err != nil {
		}
		require.Equal(t, http.StatusBadRequest, resp.StatusCode)
		defer func() {
			err := resp.Body.Close()
			require.NoError(t, err)
		}()
	})

	t.Run("work_with_gzip", func(t *testing.T) {
		buf := bytes.NewBuffer(nil)
		zb := gzip.NewWriter(buf)
		_, err := zb.Write([]byte{})
		require.NoError(t, err)
		err = zb.Close()
		require.NoError(t, err)

		r := httptest.NewRequest("POST", srv.URL+"/", buf)
		r.RequestURI = ""
		r.Header.Set("Content-Encoding", "gzip")
		r.Header.Set("Content-Type", "text/html")
		resp, err := http.DefaultClient.Do(r)
		if err != nil {
		}
		require.Equal(t, http.StatusInternalServerError, resp.StatusCode)
		defer func() {
			err := resp.Body.Close()
			require.NoError(t, err)
		}()
	})
}
