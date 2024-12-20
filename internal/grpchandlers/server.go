// Package grpchandlers обработчики grpc запросов
package grpchandlers

import (
	"go.uber.org/zap"
	"shortener/internal/server/proto"
	"shortener/internal/services"
	"shortener/internal/storage"
)

// ShortenerServer структура grpc сервера
type ShortenerServer struct {
	proto.UnimplementedURLShortenerServiceServer
	Storage     storage.Storage
	Logger      *zap.Logger
	BatchRemove services.BatchRemover
}
