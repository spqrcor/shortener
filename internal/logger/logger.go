package logger

import (
	"go.uber.org/zap"
	"log"
	"shortener/internal/config"
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
