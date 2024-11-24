// Package config формирование и хранение конфига
package config

import (
	"encoding/json"
	"flag"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"os"
	"slices"
	"strings"
	"sync"
	"time"
)

// Config тип для хранение конфига
type Config struct {
	Addr              string        `env:"SERVER_ADDRESS" json:"server_address"`
	BaseURL           string        `env:"BASE_URL" json:"base_url"`
	ShortStringLength int           `env:"SHORT_STRING_LENGTH"`
	LogLevel          zapcore.Level `env:"LOG_LEVEL"`
	FileStoragePath   string        `env:"FILE_STORAGE_PATH" json:"file_storage_path"`
	DatabaseDSN       string        `env:"DATABASE_DSN" json:"database_dsn"`
	QueryTimeOut      time.Duration `env:"QUERY_TIME_OUT"`
	SecretKey         string        `env:"SECRET_KEY"`
	TokenExp          time.Duration `env:"TOKEN_EXPIRATION"`
	EnableTLS         bool          `env:"ENABLE_TLS" json:"enable_tls,omitempty"`
	ConfigPath        string        `env:"CONFIG"`
	TrustedSubnet     string        `env:"TRUSTED_SUBNET" json:"trusted_subnet"`
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
	ConfigPath:        "",
	TrustedSubnet:     "",
}

var once sync.Once
var boolVariants = []string{"t", "true", "1"}

// NewConfig получение конфига
func NewConfig() Config {
	once.Do(func() {
		var c, c1 string
		tempCfg := Config{}

		flag.StringVar(&c, "c", "", "config path")
		flag.StringVar(&c1, "config", "", "config path")
		flag.StringVar(&tempCfg.Addr, "a", "", "address and port to run server")
		flag.StringVar(&tempCfg.BaseURL, "b", "", "base url")
		flag.StringVar(&tempCfg.FileStoragePath, "f", "", "file storage path")
		flag.StringVar(&tempCfg.DatabaseDSN, "d", "", "database dsn")
		flag.StringVar(&tempCfg.TrustedSubnet, "t", "", "trusted subnet")
		flag.BoolVar(&tempCfg.EnableTLS, "s", false, "enable tls")
		flag.Parse()

		if c != "" {
			cfg.ConfigPath = c
		}
		if c1 != "" {
			cfg.ConfigPath = c1
		}
		serverConfig, findConfig := os.LookupEnv("CONFIG")
		if findConfig {
			cfg.ConfigPath = serverConfig
		}
		if cfg.ConfigPath != "" {
			raw, err := os.ReadFile(cfg.ConfigPath)
			if err != nil {
				log.Fatal("Error read config file")
				return
			}
			var cf Config
			if err = json.Unmarshal(raw, &cf); err != nil {
				log.Fatal("Error parse config file")
				return
			}
			err = nil
		}

		if tempCfg.Addr != "" {
			cfg.Addr = tempCfg.Addr
		}
		if tempCfg.BaseURL != "" {
			cfg.BaseURL = tempCfg.BaseURL
		}
		if tempCfg.FileStoragePath != "" {
			cfg.FileStoragePath = tempCfg.FileStoragePath
		}
		if tempCfg.DatabaseDSN != "" {
			cfg.DatabaseDSN = tempCfg.DatabaseDSN
		}
		if tempCfg.EnableTLS {
			cfg.EnableTLS = tempCfg.EnableTLS
		}
		if tempCfg.TrustedSubnet != "" {
			cfg.TrustedSubnet = tempCfg.TrustedSubnet
		}

		serverAddressEnv, findAddress := os.LookupEnv("SERVER_ADDRESS")
		serverBaseURLEnv, findBaseURL := os.LookupEnv("BASE_URL")
		serverStoragePath, findStoragePath := os.LookupEnv("FILE_STORAGE_PATH")
		serverDatabaseDSN, findDatabaseDSN := os.LookupEnv("DATABASE_DSN")
		serverEnableTLS, findEnableTLS := os.LookupEnv("ENABLE_TLS")
		serverTrustedSubnet, findTrustedSubnet := os.LookupEnv("TRUSTED_SUBNET")

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
		if findTrustedSubnet {
			cfg.TrustedSubnet = serverTrustedSubnet
		}
		if findEnableTLS && slices.IndexFunc(boolVariants, func(c string) bool { return c == strings.ToLower(serverEnableTLS) }) > -1 {
			cfg.EnableTLS = true
		}
	})
	return cfg
}
