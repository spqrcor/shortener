package storage

import (
	"shortener/internal/config"
)

type Storage interface {
	Add(inputURL string) (string, error)
	Find(key string) (string, error)
	BatchAdd(inputURLs []BatchInputParams) ([]BatchOutputParams, error)
}

type BatchInputParams struct {
	CorrelationID string `json:"correlation_id"`
	URL           string `json:"original_url"`
}

type BatchOutputParams struct {
	CorrelationID string `json:"correlation_id"`
	ShortURL      string `json:"short_url"`
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
