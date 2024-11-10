// Package config формирование и хранение конфига
package config

import (
	"encoding/json"
	"flag"
	"fmt"
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
}

var once sync.Once
var boolVariants = []string{"t", "true", "1"}

// NewConfig получение конфига
func NewConfig() Config {
	once.Do(func() {
		var c, c1 string
		flag.StringVar(&c, "c", "", "config path")
		flag.StringVar(&c1, "config", "", "config path")
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
		fmt.Print("Config:", cfg)
	})
	return cfg
}

func configFromJSON() {
	var c, c1 string
	flag.StringVar(&c, "c", "", "config path")
	flag.StringVar(&c1, "config", "", "config path")
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
}
