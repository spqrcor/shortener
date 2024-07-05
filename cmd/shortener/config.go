package main

import (
	"flag"
	"github.com/caarlos0/env/v6"
	"log"
)

var flagRunAddr string
var flagBaseURL string

type Config struct {
	addr    string `env:"SERVER_ADDRESS"`
	baseURL string `env:"BASE_URL"`
}

func parseFlags() {
	var cfg Config
	err := env.Parse(&cfg)
	if err != nil {
		log.Fatal(err)
	}

	if cfg.addr != "" {
		flagRunAddr = cfg.addr
	} else {
		flag.StringVar(&flagRunAddr, "a", ":8080", "address and port to run server")
	}

	if cfg.baseURL != "" {
		flagBaseURL = cfg.baseURL
	} else {
		flag.StringVar(&flagBaseURL, "b", "http://localhost:8080", "base url")
	}
	flag.Parse()
}
