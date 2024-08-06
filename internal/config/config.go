package config

import (
	"flag"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

type Config struct {
	Addr              string        `env:"SERVER_ADDRESS"`
	BaseURL           string        `env:"BASE_URL"`
	ShortStringLength int           `env:"SHORT_STRING_LENGTH"`
	LogLevel          zapcore.Level `env:"LOG_LEVEL"`
	FileStoragePath   string        `env:"FILE_STORAGE_PATH"`
	DatabaseDSN       string        `env:"DATABASE_DSN"`
}

var Cfg = Config{
	Addr:              "localhost:8080",
	BaseURL:           "http://localhost:8080",
	ShortStringLength: 6,
	LogLevel:          zap.InfoLevel,
	FileStoragePath:   "",
	DatabaseDSN:       "",
}

func Init() {
	flag.StringVar(&Cfg.Addr, "a", Cfg.Addr, "address and port to run server")
	flag.StringVar(&Cfg.BaseURL, "b", Cfg.BaseURL, "base url")
	flag.StringVar(&Cfg.FileStoragePath, "f", Cfg.FileStoragePath, "file storage path")
	flag.StringVar(&Cfg.DatabaseDSN, "d", Cfg.DatabaseDSN, "database dsn")
	flag.Parse()

	serverAddressEnv, findAddress := os.LookupEnv("SERVER_ADDRESS")
	serverBaseURLEnv, findBaseURL := os.LookupEnv("BASE_URL")
	serverStoragePath, findStoragePath := os.LookupEnv("FILE_STORAGE_PATH")
	serverDatabaseDSN, findDatabaseDSN := os.LookupEnv("DATABASE_DSN")

	if findAddress {
		Cfg.Addr = serverAddressEnv
	}
	if findBaseURL {
		Cfg.BaseURL = serverBaseURLEnv
	}
	if findStoragePath {
		Cfg.FileStoragePath = serverStoragePath
	}
	if findDatabaseDSN {
		Cfg.DatabaseDSN = serverDatabaseDSN
	}
}
