package main

import (
	"log"
	"shortener/internal/authenticate"
	"shortener/internal/config"
	"shortener/internal/logger"
	"shortener/internal/server"
	"shortener/internal/services"
	"shortener/internal/storage"
)

func main() {
	conf := config.NewConfig()
	loggerRes, err := logger.NewLogger(conf.LogLevel)
	if err != nil {
		log.Fatal(err)
	}

	storageService := storage.NewStorage(conf, loggerRes)
	authService := authenticate.NewAuthenticateService(
		authenticate.WithLogger(loggerRes),
		authenticate.WithSecretKey(conf.SecretKey),
		authenticate.WithTokenExp(conf.TokenExp),
	)
	batchRemoveService := services.NewBatchRemoveService(loggerRes, storageService)

	appServer := server.NewServer(
		server.WithConfig(conf),
		server.WithLogger(loggerRes),
		server.WithStorage(storageService),
		server.WithAuthenticate(authService),
		server.WithBatchRemove(batchRemoveService),
	)

	if err := appServer.Start(); err != nil {
		loggerRes.Error(err.Error())
	}
}
