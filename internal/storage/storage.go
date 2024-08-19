package storage

import (
	"context"
	"shortener/internal/config"
)

type Storage interface {
	Add(ctx context.Context, inputURL string) (string, error)
	Find(ctx context.Context, key string) (string, error)
	BatchAdd(ctx context.Context, inputURLs []BatchInputParams) ([]BatchOutputParams, error)
	FindByUser(ctx context.Context) ([]FindByUserOutputParams, error)
	Remove(ctx context.Context, shorts []string) error
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

var Source Storage

func Init() {
	if config.Cfg.DatabaseDSN != "" {
		CreateDBStorage()
	} else if config.Cfg.FileStoragePath != "" {
		CreateFileStorage()
	} else {
		CreateMemoryStorage()
	}
}
