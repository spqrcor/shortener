package main

import (
	"shortener/internal/config"
	"shortener/internal/logger"
	"shortener/internal/server"
)

func main() {
	config.Init()
	logger.Init()

	server.Start()
}
