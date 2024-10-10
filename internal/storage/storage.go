package storage

import (
	"context"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"shortener/internal/config"
)

type Storage interface {
	Add(ctx context.Context, inputURL string) (string, error)
	Find(ctx context.Context, key string) (string, error)
	BatchAdd(ctx context.Context, inputURLs []BatchInputParams) ([]BatchOutputParams, error)
	FindByUser(ctx context.Context) ([]FindByUserOutputParams, error)
	Remove(ctx context.Context, UserID uuid.UUID, shorts []string) error
}

type BatchInputParams struct {
	CorrelationID string `json:"correlation_id"`
	URL           string `json:"original_url"`
}

type BatchOutputParams struct {
	CorrelationID string `json:"correlation_id"`
	ShortURL      string `json:"short_url"`
}

type FindByUserOutputParams struct {
	ShortURL    string `json:"short_url"`
	OriginalURL string `json:"original_url"`
}

func NewStorage(config config.Config, logger *zap.Logger) Storage {
	if config.DatabaseDSN != "" {
		return CreateDBStorage(config, logger)
	} else if config.FileStoragePath != "" {
		return CreateFileStorage(config, logger)
	} else {
		return CreateMemoryStorage(config)
	}
}
