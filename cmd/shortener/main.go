package main

import (
	"shortener/internal/config"
	"shortener/internal/logger"
	"shortener/internal/server"
	"shortener/internal/storage"
)

func main() {
	config.Init()
	logger.Init()
	storage.Init()

	server.Start()
}
