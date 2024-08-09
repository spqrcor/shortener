package server

import (
	"compress/gzip"
	"github.com/go-chi/chi/v5/middleware"
	"go.uber.org/zap"
	"net/http"
	"shortener/internal/logger"
	"time"
)

func getBodyMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		if r.Header.Get(`Content-Encoding`) == `gzip` {
			gz, err := gzip.NewReader(r.Body)
			if err != nil {
				logger.Log.Error(err.Error())
			} else {
				r.Body = gz
			}
			if err = gz.Close(); err != nil {
				logger.Log.Error(err.Error())
			}
		}
		next.ServeHTTP(rw, r)
	})
}

func loggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		start := time.Now()
		ww := middleware.NewWrapResponseWriter(rw, r.ProtoMajor)
		next.ServeHTTP(ww, r)
		duration := time.Since(start)

		logger.Log.Info("HTTP request",
			zap.String("method", r.Method),
			zap.String("path", r.URL.Path),
			zap.Int("status", ww.Status()),
			zap.Int("content-length", ww.BytesWritten()),
			zap.String("duration", duration.String()),
		)
	})
}
