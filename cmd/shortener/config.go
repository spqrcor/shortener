package main

import (
	"flag"
	"github.com/caarlos0/env/v6"
	"log"
)

type Config struct {
	addr    string `env:"SERVER_ADDRESS"`
	baseURL string `env:"BASE_URL"`
}

func parseFlags() Config {
	var cfg Config
	err := env.Parse(&cfg)
	if err != nil {
		log.Fatal(err)
	}

	if cfg.addr == "" {
		flag.StringVar(&cfg.addr, "a", ":8080", "address and port to run server")
	}

	if cfg.baseURL == "" {
		flag.StringVar(&cfg.baseURL, "b", "http://localhost:8080", "base url")
	}
	flag.Parse()
	return cfg
}
