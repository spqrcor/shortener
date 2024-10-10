package storage

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"os"
	"shortener/internal/app"
	"shortener/internal/config"
)

type FileStorage struct {
	config config.Config
	logger *zap.Logger
	Store  map[string]string
}

func CreateFileStorage(config config.Config, logger *zap.Logger) Storage {
	data, err := os.ReadFile(config.FileStoragePath)
	if err != nil {
		file, err := os.OpenFile(config.FileStoragePath, os.O_CREATE, 0666)
		if err != nil {
			logger.Fatal(err.Error())
		}
		if err := file.Close(); err != nil {
			logger.Fatal(err.Error())
		}
	}

	fileData := map[string]string{}
	if len(data) != 0 {
		if err := json.Unmarshal([]byte(data), &fileData); err != nil {
			logger.Fatal(err.Error())
		}
	}
	return FileStorage{
		config: config,
		logger: logger,
		Store:  fileData,
	}
}

func (f FileStorage) Add(ctx context.Context, inputURL string) (string, error) {
	if err := app.ValidateURL(inputURL); err != nil {
		return "", err
	}
	genURL := app.GenerateShortURL(f.config.ShortStringLength, f.config.BaseURL)
	f.Store[genURL] = inputURL
	f.updateFileStorage(f.Store)
	return genURL, nil
}

func (f FileStorage) Find(ctx context.Context, key string) (string, error) {
	redirectURL, ok := f.Store[f.config.BaseURL+key]
	if ok {
		return redirectURL, nil
	}
	return "", errors.New("ключ не найден")
}

func (f FileStorage) updateFileStorage(store map[string]string) {
	x, err := json.Marshal(store)
	if err != nil {
		f.logger.Fatal(err.Error())
	}
	if err := os.Truncate(f.config.FileStoragePath, 0); err != nil {
		f.logger.Fatal(err.Error())
	}
	if err := os.WriteFile(f.config.FileStoragePath, x, 0666); err != nil {
		f.logger.Fatal(err.Error())
	}
}

func (f FileStorage) BatchAdd(ctx context.Context, inputURLs []BatchInputParams) ([]BatchOutputParams, error) {
	var output []BatchOutputParams
	for _, inputURL := range inputURLs {
		err := app.ValidateURL(inputURL.URL)
		if err != nil {
			return nil, err
		}
		genURL := app.GenerateShortURL(f.config.ShortStringLength, f.config.BaseURL)
		f.Store[genURL] = inputURL.URL
		output = append(output, BatchOutputParams{CorrelationID: inputURL.CorrelationID, ShortURL: genURL})
	}
	f.updateFileStorage(f.Store)
	return output, nil
}

func (f FileStorage) FindByUser(ctx context.Context) ([]FindByUserOutputParams, error) {
	var output []FindByUserOutputParams
	return output, nil
}

func (f FileStorage) Remove(ctx context.Context, UserID uuid.UUID, shorts []string) error {
	return nil
}
