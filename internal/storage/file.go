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

// FileStorage тип file хранилища
type FileStorage struct {
	config config.Config
	logger *zap.Logger
	Store  map[string]string
}

// CreateFileStorage создание file хранилища, config - конфиг, logger - логгер
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

// Add добавление, ctx - контекст, inputURL - входящий url
func (f FileStorage) Add(ctx context.Context, inputURL string) (string, error) {
	if err := app.ValidateURL(inputURL); err != nil {
		return "", err
	}
	genURL := app.GenerateShortURL(f.config.ShortStringLength, f.config.BaseURL)
	f.Store[genURL] = inputURL
	f.updateFileStorage(f.Store)
	return genURL, nil
}

// Find поиск, ctx - контекст, key - шорткей
func (f FileStorage) Find(ctx context.Context, key string) (string, error) {
	redirectURL, ok := f.Store[f.config.BaseURL+key]
	if ok {
		return redirectURL, nil
	}
	return "", errors.New("ключ не найден")
}

// updateFileStorage обновление файла
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

// BatchAdd групповое добавление, ctx - контекст, inputURLs массив данных
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

// FindByUser поиск по пользователю, ctx - контекст
func (f FileStorage) FindByUser(ctx context.Context) ([]FindByUserOutputParams, error) {
	var output []FindByUserOutputParams
	return output, nil
}

// Remove удаление, ctx - контекст, UserID - guid пользователя, shorts - массив шорткеев
func (f FileStorage) Remove(ctx context.Context, UserID uuid.UUID, shorts []string) error {
	return nil
}

// ShutDown завершение работы с хранилищем
func (f FileStorage) ShutDown() error {
	return nil
}
