// Package config формирование и хранение конфига
package config

import (
	"flag"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"slices"
	"strings"
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
	EnableTLS         bool          `env:"ENABLE_TLS"`
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
	EnableTLS:         false,
}

var once sync.Once
var boolVariants = []string{"t", "true", "1"}

// NewConfig получение конфига
func NewConfig() Config {
	once.Do(func() {
		flag.StringVar(&cfg.Addr, "a", cfg.Addr, "address and port to run server")
		flag.StringVar(&cfg.BaseURL, "b", cfg.BaseURL, "base url")
		flag.StringVar(&cfg.FileStoragePath, "f", cfg.FileStoragePath, "file storage path")
		flag.StringVar(&cfg.DatabaseDSN, "d", cfg.DatabaseDSN, "database dsn")
		flag.BoolVar(&cfg.EnableTLS, "s", cfg.EnableTLS, "enable tls")
		flag.Parse()

		serverAddressEnv, findAddress := os.LookupEnv("SERVER_ADDRESS")
		serverBaseURLEnv, findBaseURL := os.LookupEnv("BASE_URL")
		serverStoragePath, findStoragePath := os.LookupEnv("FILE_STORAGE_PATH")
		serverDatabaseDSN, findDatabaseDSN := os.LookupEnv("DATABASE_DSN")
		serverEnableTLS, findEnableTLS := os.LookupEnv("ENABLE_TLS")

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
		if findEnableTLS && slices.IndexFunc(boolVariants, func(c string) bool { return c == strings.ToLower(serverEnableTLS) }) > -1 {
			cfg.EnableTLS = true
		}
	})
	return cfg
}
