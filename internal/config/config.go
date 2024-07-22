package config

import (
	"flag"
	"github.com/caarlos0/env/v6"
	"log"
)

type Config struct {
	Addr              string `env:"SERVER_ADDRESS"`
	BaseURL           string `env:"BASE_URL"`
	ShortStringLength int    `env:"SHORT_STRING_LENGTH"`
}

var Cfg = Config{
	Addr:              ":8080",
	BaseURL:           "http://localhost:8080",
	ShortStringLength: 6,
}

func ParseFlags() {
	err := env.Parse(&Cfg)
	if err != nil {
		log.Fatal(err)
	}

	if Cfg.Addr == "" {
		flag.StringVar(&Cfg.Addr, "a", Cfg.Addr, "address and port to run server")
	}

	if Cfg.BaseURL == "" {
		flag.StringVar(&Cfg.BaseURL, "b", Cfg.BaseURL, "base url")
	}
	flag.Parse()
}
