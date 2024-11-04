// go run -ldflags "-X main.buildVersion=v1.0.1 -X 'main.buildDate=$(date +'%Y/%m/%d %H:%M:%S')' -X main.buildCommit=hello world" main.go
package main

import (
	"fmt"
	"log"
	"shortener/internal/authenticate"
	"shortener/internal/config"
	"shortener/internal/logger"
	"shortener/internal/server"
	"shortener/internal/services"
	"shortener/internal/storage"
)

var (
	buildVersion string = "N/A"
	buildDate    string = "N/A"
	buildCommit  string = "N/A"
)

func main() {
	fmt.Printf("Build version:%s\nBuild date:%s\nBuild commit:%s\n", buildVersion, buildDate, buildCommit)
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
