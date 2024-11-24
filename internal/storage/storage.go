// Package storage
package storage

import (
	"context"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"shortener/internal/config"
	"shortener/internal/db"
)

// Storage интерфейс хранилища
type Storage interface {
	Add(ctx context.Context, inputURL string) (string, error)
	Find(ctx context.Context, key string) (string, error)
	BatchAdd(ctx context.Context, inputURLs []BatchInputParams) ([]BatchOutputParams, error)
	FindByUser(ctx context.Context) ([]FindByUserOutputParams, error)
	Remove(ctx context.Context, UserID uuid.UUID, shorts []string) error
	ShutDown() error
	Stat(ctx context.Context) (Stat, error)
}

// Stat тип возвращаемый методом InternalStatHandler
type Stat struct {
	Urls  int `json:"urls"`
	Users int `json:"users"`
}

// BatchInputParams тип для входящих данных роута /api/shorten/batch
type BatchInputParams struct {
	CorrelationID string `json:"correlation_id"`
	URL           string `json:"original_url"`
}

// BatchOutputParams тип для ответа роута /api/shorten/batch
type BatchOutputParams struct {
	CorrelationID string `json:"correlation_id"`
	ShortURL      string `json:"short_url"`
}

// FindByUserOutputParams тип возвращаемый методом FindByUser
type FindByUserOutputParams struct {
	ShortURL    string `json:"short_url"`
	OriginalURL string `json:"original_url"`
}

// NewStorage создание хранилища, config конфиг, logger - логгер
func NewStorage(config config.Config, logger *zap.Logger) Storage {
	if config.DatabaseDSN != "" {
		res, err := db.Connect(config.DatabaseDSN)
		if err != nil {
			logger.Fatal(err.Error())
		}
		if err := db.Migrate(res); err != nil {
			logger.Fatal(err.Error())
		}
		return CreateDBStorage(config, logger, res)
	} else if config.FileStoragePath != "" {
		return CreateFileStorage(config, logger)
	} else {
		return CreateMemoryStorage(config)
	}
}
