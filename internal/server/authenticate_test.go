package server

import (
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"net/http"
	"net/http/httptest"
	"shortener/internal/authenticate"
	"shortener/internal/config"
	"shortener/internal/handlers"
	"shortener/internal/logger"
	"shortener/internal/storage"
	"testing"
	"time"
)

func Test_authenticateMiddleware(t *testing.T) {
	loggerRes, _ := logger.NewLogger(zap.InfoLevel)
	conf := config.NewConfig()
	store := storage.CreateMemoryStorage(conf)
	authService := authenticate.NewAuthenticateService(
		authenticate.WithLogger(loggerRes),
		authenticate.WithSecretKey(conf.SecretKey),
		authenticate.WithTokenExp(conf.TokenExp),
	)
	r := chi.NewRouter()
	r.Use(authenticateMiddleware(loggerRes, authService))
	r.Post(`/`, handlers.CreateShortHandler(store))
	srv := httptest.NewServer(r)
	defer srv.Close()
	t.Run("400", func(t *testing.T) {
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

	r = chi.NewRouter()
	r.Use(authenticateMiddleware(loggerRes, authService))
	r.Post(`/`, handlers.CreateShortHandler(store))
	srv = httptest.NewServer(r)
	defer srv.Close()
	t.Run("500", func(t *testing.T) {
		r := httptest.NewRequest("POST", srv.URL+"/", nil)
		r.RequestURI = ""
		cookie := http.Cookie{Name: "Authorization", Value: "", Expires: time.Now().Add(conf.TokenExp), HttpOnly: true, Path: "/"}
		r.AddCookie(&cookie)
		r.Header.Set("Content-Type", "text/html")
		resp, _ := http.DefaultClient.Do(r)
		require.Equal(t, http.StatusInternalServerError, resp.StatusCode)
		defer func() {
			err := resp.Body.Close()
			require.NoError(t, err)
		}()
	})

	r = chi.NewRouter()
	r.Use(authenticateMiddleware(loggerRes, authService))
	r.Post(`/`, handlers.CreateShortHandler(store))
	srv = httptest.NewServer(r)
	defer srv.Close()
	t.Run("401", func(t *testing.T) {
		r := httptest.NewRequest("GET", srv.URL+"/api/user/urls", nil)
		r.RequestURI = ""
		r.Header.Set("Content-Type", "text/html")
		resp, _ := http.DefaultClient.Do(r)
		require.Equal(t, http.StatusUnauthorized, resp.StatusCode)
		defer func() {
			err := resp.Body.Close()
			require.NoError(t, err)
		}()
	})

	token, _ := authService.CreateToken(uuid.New())

	r = chi.NewRouter()
	r.Use(authenticateMiddleware(loggerRes, authService))
	r.Post(`/`, handlers.CreateShortHandler(store))
	srv = httptest.NewServer(r)
	defer srv.Close()
	t.Run("decode cookie", func(t *testing.T) {
		r := httptest.NewRequest("POST", srv.URL+"/", nil)
		r.RequestURI = ""
		cookie := http.Cookie{Name: "Authorization", Value: token, Expires: time.Now().Add(conf.TokenExp), HttpOnly: true, Path: "/"}
		r.AddCookie(&cookie)
		r.Header.Set("Content-Type", "text/html")
		resp, _ := http.DefaultClient.Do(r)
		require.Equal(t, http.StatusBadRequest, resp.StatusCode)
		defer func() {
			err := resp.Body.Close()
			require.NoError(t, err)
		}()
	})

}
