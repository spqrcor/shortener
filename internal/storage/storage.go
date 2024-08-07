package storage

import (
	"shortener/internal/config"
)

type Storage interface {
	Add(inputURL string) (string, error)
	Find(key string) (string, error)
	BatchAdd(inputURLs []BatchParams) error
}

type BatchParams struct {
	ShortURL string `json:"correlation_id"`
	URL      string `json:"original_url"`
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
