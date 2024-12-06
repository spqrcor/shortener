package middlewares

import (
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

func Test_loggerMiddleware(t *testing.T) {
	loggerRes, _ := logger.NewLogger(zap.InfoLevel)
	conf := config.NewConfig()
	store := storage.CreateMemoryStorage(conf)

	r := chi.NewRouter()
	r.Use(LoggerMiddleware(loggerRes))

	r.Post(`/`, handlers.CreateShortHandler(store))
	srv := httptest.NewServer(r)
	defer srv.Close()

	t.Run("work_logger", func(t *testing.T) {
		r := httptest.NewRequest("POST", srv.URL+"/", nil)
		r.RequestURI = ""
		r.Header.Set("Content-Type", "text/html")
		resp, _ := http.DefaultClient.Do(r)
		require.Equal(t, http.StatusBadRequest, resp.StatusCode)
		defer func() {
			err := resp.Body.Close()
			require.NoError(t, err)
		}()
	})
}
