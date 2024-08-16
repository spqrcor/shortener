package storage

import (
	"context"
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

	fileData := map[string]string{}
	if len(data) != 0 {
		if err := json.Unmarshal([]byte(data), &fileData); err != nil {
			logger.Log.Fatal(err.Error())
		}
	}
	Source = FileStorage{
		Store: fileData,
	}
}

func (f FileStorage) Add(ctx context.Context, inputURL string) (string, error) {
	err := app.ValidateURL(inputURL)
	if err != nil {
		return "", err
	}
	genURL := app.GenerateShortURL()
	f.Store[genURL] = inputURL
	updateFileStorage(f.Store)
	return genURL, nil
}

func (f FileStorage) Find(ctx context.Context, key string) (string, error) {
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

func (f FileStorage) BatchAdd(ctx context.Context, inputURLs []BatchInputParams) ([]BatchOutputParams, error) {
	var output []BatchOutputParams
	for _, inputURL := range inputURLs {
		err := app.ValidateURL(inputURL.URL)
		if err != nil {
			return nil, err
		}
		genURL := app.GenerateShortURL()
		f.Store[genURL] = inputURL.URL
		output = append(output, BatchOutputParams{CorrelationID: inputURL.CorrelationID, ShortURL: genURL})
	}
	updateFileStorage(f.Store)
	return output, nil
}

func (f FileStorage) FindByUser(ctx context.Context) ([]FindByUserOutputParams, error) {
	var output []FindByUserOutputParams
	return output, nil
}
