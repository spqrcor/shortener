package config

import (
	"flag"
	"os"
)

type Config struct {
	Addr              string `env:"SERVER_ADDRESS"`
	BaseURL           string `env:"BASE_URL"`
	ShortStringLength int    `env:"SHORT_STRING_LENGTH"`
}

var Cfg = Config{
	Addr:              "localhost:8080",
	BaseURL:           "http://localhost:8080",
	ShortStringLength: 6,
}

func ParseFlags() {
	flag.StringVar(&Cfg.Addr, "a", "localhost:8080", "server address")
	flag.StringVar(&Cfg.BaseURL, "b", "http://localhost:8080", "base result server address")
	flag.Parse()

	serverAddressEnv, findAddress := os.LookupEnv("SERVER_ADDRESS")
	serverBaseURLEnv, findBaseURL := os.LookupEnv("BASE_URL")

	if findAddress {
		Cfg.Addr = serverAddressEnv
	}
	if findBaseURL {
		Cfg.BaseURL = serverBaseURLEnv
	}
}
