package logger

import (
	"github.com/go-chi/chi/v5/middleware"
	"go.uber.org/zap"
	"log"
	"net/http"
	"shortener/internal/config"
	"time"
)

var Log *zap.Logger = zap.NewNop()

func Init() {
	cfg := zap.NewProductionConfig()
	cfg.Level = zap.NewAtomicLevelAt(config.Cfg.LogLevel)
	zl, err := cfg.Build()
	if err != nil {
		log.Fatal(err)
	}
	Log = zl
}

func RequestLogger(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)
		h(ww, r)
		duration := time.Since(start)

		Log.Info("HTTP request",
			zap.String("method", r.Method),
			zap.String("path", r.URL.Path),
			zap.Int("status", ww.Status()),
			zap.Int("content-length", ww.BytesWritten()),
			zap.String("duration", duration.String()),
		)
	}
}
