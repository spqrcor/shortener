// Package config формирование и хранение конфига
package config

import (
	"flag"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"sync"
	"time"
)

// Config тип для хранение конфига
type Config struct {
	Addr              string        `env:"SERVER_ADDRESS"`
	BaseURL           string        `env:"BASE_URL"`
	ShortStringLength int           `env:"SHORT_STRING_LENGTH"`
	LogLevel          zapcore.Level `env:"LOG_LEVEL"`
	FileStoragePath   string        `env:"FILE_STORAGE_PATH"`
	DatabaseDSN       string        `env:"DATABASE_DSN"`
	QueryTimeOut      time.Duration `env:"QUERY_TIME_OUT"`
	SecretKey         string        `env:"SECRET_KEY"`
	TokenExp          time.Duration `env:"TOKEN_EXPIRATION"`
}

// cfg переменная конфига
var cfg = Config{
	Addr:              "localhost:8080",
	BaseURL:           "http://localhost:8080",
	ShortStringLength: 6,
	LogLevel:          zap.InfoLevel,
	FileStoragePath:   "",
	DatabaseDSN:       "",
	QueryTimeOut:      3,
	SecretKey:         "KLJ-fo3Fksd3fl!=",
	TokenExp:          time.Hour * 3,
}

var once sync.Once

// NewConfig получение конфига
func NewConfig() Config {
	once.Do(func() {
		flag.StringVar(&cfg.Addr, "a", cfg.Addr, "address and port to run server")
		flag.StringVar(&cfg.BaseURL, "b", cfg.BaseURL, "base url")
		flag.StringVar(&cfg.FileStoragePath, "f", cfg.FileStoragePath, "file storage path")
		flag.StringVar(&cfg.DatabaseDSN, "d", cfg.DatabaseDSN, "database dsn")
		flag.Parse()

		serverAddressEnv, findAddress := os.LookupEnv("SERVER_ADDRESS")
		serverBaseURLEnv, findBaseURL := os.LookupEnv("BASE_URL")
		serverStoragePath, findStoragePath := os.LookupEnv("FILE_STORAGE_PATH")
		serverDatabaseDSN, findDatabaseDSN := os.LookupEnv("DATABASE_DSN")

		if findAddress {
			cfg.Addr = serverAddressEnv
		}
		if findBaseURL {
			cfg.BaseURL = serverBaseURLEnv
		}
		if findStoragePath {
			cfg.FileStoragePath = serverStoragePath
		}
		if findDatabaseDSN {
			cfg.DatabaseDSN = serverDatabaseDSN
		}
	})
	return cfg
}
