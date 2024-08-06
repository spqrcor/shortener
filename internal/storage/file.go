package storage

import (
	"encoding/json"
	"errors"
	"os"
	"shortener/internal/app"
	"shortener/internal/config"
	"shortener/internal/logger"
)

type FileStorage struct {
	Store map[string]string
}

func CreateFileStorage() {
	data, err := os.ReadFile(config.Cfg.FileStoragePath)
	if err != nil {
		file, err := os.OpenFile(config.Cfg.FileStoragePath, os.O_CREATE, 0666)
		if err != nil {
			logger.Log.Fatal(err.Error())
		}
		if err := file.Close(); err != nil {
			logger.Log.Fatal(err.Error())
		}
	}

	if len(data) == 0 {
		return
	}
	fileData := map[string]string{}
	if err := json.Unmarshal([]byte(data), &fileData); err != nil {
		logger.Log.Fatal(err.Error())
	}
	Source = FileStorage{
		Store: fileData,
	}
}

func (f FileStorage) Add(inputURL string) (string, error) {
	genURL, err := app.CreateShortURL(inputURL)
	if err != nil {
		return "", err
	}
	f.Store[genURL] = inputURL
	updateFileStorage(f.Store)
	return genURL, nil
}

func (f FileStorage) Find(key string) (string, error) {
	redirectURL, ok := f.Store[config.Cfg.BaseURL+key]
	if ok {
		return redirectURL, nil
	}
	return "", errors.New("ключ не найден")
}

func updateFileStorage(store map[string]string) {
	x, err := json.Marshal(store)
	if err != nil {
		logger.Log.Fatal(err.Error())
	}
	if err := os.Truncate(config.Cfg.FileStoragePath, 0); err != nil {
		logger.Log.Fatal(err.Error())
	}
	if err := os.WriteFile(config.Cfg.FileStoragePath, x, 0666); err != nil {
		logger.Log.Fatal(err.Error())
	}
}
