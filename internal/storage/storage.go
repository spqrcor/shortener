package storage

import "shortener/internal/config"

type Storage interface {
	Add(inputURL string) (string, error)
	Find(key string) (string, error)
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
