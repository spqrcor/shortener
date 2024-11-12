package server

import (
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"reflect"
	"shortener/internal/authenticate"
	"shortener/internal/config"
	"shortener/internal/logger"
	"shortener/internal/services"
	"shortener/internal/storage"
	"syscall"
	"testing"
	"time"
)

func TestNewServer(t *testing.T) {
	conf := config.NewConfig()
	loggerRes, _ := logger.NewLogger(zap.InfoLevel)

	store := storage.CreateMemoryStorage(conf)

	auth := authenticate.NewAuthenticateService(
		authenticate.WithLogger(loggerRes),
		authenticate.WithSecretKey(conf.SecretKey),
		authenticate.WithTokenExp(conf.TokenExp),
	)

	batchRemove := services.NewBatchRemoveService(loggerRes, store)

	server := NewServer(
		WithLogger(loggerRes),
		WithConfig(conf),
		WithAuthenticate(auth),
		WithStorage(store),
		WithBatchRemove(batchRemove),
	)
	assert.Equal(t, reflect.TypeOf(server).String() == "*server.HTTPServer", true)
}

func TestHTTPServer_Start(t *testing.T) {
	conf := config.NewConfig()
	loggerRes, _ := logger.NewLogger(zap.InfoLevel)

	store := storage.CreateMemoryStorage(conf)

	auth := authenticate.NewAuthenticateService(
		authenticate.WithLogger(loggerRes),
		authenticate.WithSecretKey(conf.SecretKey),
		authenticate.WithTokenExp(conf.TokenExp),
	)

	batchRemove := services.NewBatchRemoveService(loggerRes, store)

	server := NewServer(
		WithLogger(loggerRes),
		WithConfig(conf),
		WithAuthenticate(auth),
		WithStorage(store),
		WithBatchRemove(batchRemove),
	)
	go func() {
		time.Sleep(1 * time.Second)
		_ = syscall.Kill(syscall.Getpid(), syscall.SIGINT)
	}()
	server.Start()
}

func TestHTTPServer_StartHttps(t *testing.T) {
	conf := config.NewConfig()
	conf.EnableTLS = true
	loggerRes, _ := logger.NewLogger(zap.InfoLevel)

	store := storage.CreateMemoryStorage(conf)

	auth := authenticate.NewAuthenticateService(
		authenticate.WithLogger(loggerRes),
		authenticate.WithSecretKey(conf.SecretKey),
		authenticate.WithTokenExp(conf.TokenExp),
	)

	batchRemove := services.NewBatchRemoveService(loggerRes, store)

	server := NewServer(
		WithLogger(loggerRes),
		WithConfig(conf),
		WithAuthenticate(auth),
		WithStorage(store),
		WithBatchRemove(batchRemove),
	)
	go func() {
		time.Sleep(1 * time.Second)
		_ = syscall.Kill(syscall.Getpid(), syscall.SIGINT)
	}()
	server.Start()
}
