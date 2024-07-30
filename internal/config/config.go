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
}

var Cfg = Config{
	Addr:              "localhost:8080",
	BaseURL:           "http://localhost:8080",
	ShortStringLength: 6,
	LogLevel:          zap.InfoLevel,
	FileStoragePath:   os.TempDir() + "/storage.json",
}

func Init() {
	flag.StringVar(&Cfg.Addr, "a", Cfg.Addr, "address and port to run server")
	flag.StringVar(&Cfg.BaseURL, "b", Cfg.BaseURL, "base url")
	flag.StringVar(&Cfg.BaseURL, "f", Cfg.BaseURL, "file storage path")
	flag.Parse()

	serverAddressEnv, findAddress := os.LookupEnv("SERVER_ADDRESS")
	serverBaseURLEnv, findBaseURL := os.LookupEnv("BASE_URL")
	serverStoragePath, findStoragePath := os.LookupEnv("FILE_STORAGE_PATH")

	if findAddress {
		Cfg.Addr = serverAddressEnv
	}
	if findBaseURL {
		Cfg.BaseURL = serverBaseURLEnv
	}
	if findStoragePath {
		Cfg.FileStoragePath = serverStoragePath
	}
}
