package storage

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"shortener/internal/app"
	"shortener/internal/config"
)

// MemoryStorage тип memory хранилища
type MemoryStorage struct {
	config config.Config
	Store  map[string]string
}

// CreateMemoryStorage создание memory хранилища, config - конфиг
func CreateMemoryStorage(config config.Config) Storage {
	return MemoryStorage{
		config: config,
		Store:  map[string]string{},
	}
}

// Add добавление, ctx - контекст, inputURL - входящий url
func (m MemoryStorage) Add(ctx context.Context, inputURL string) (string, error) {
	err := app.ValidateURL(inputURL)
	if err != nil {
		return "", err
	}
	genURL := app.GenerateShortURL(m.config.ShortStringLength, m.config.BaseURL)
	m.Store[genURL] = inputURL
	return genURL, nil
}

// Find поиск, ctx - контекст, key - шорткей
func (m MemoryStorage) Find(ctx context.Context, key string) (string, error) {
	redirectURL, ok := m.Store[m.config.BaseURL+key]
	if ok {
		return redirectURL, nil
	}
	return "", errors.New("ключ не найден")
}

// BatchAdd групповое добавление, ctx - контекст, inputURLs массив данных
func (m MemoryStorage) BatchAdd(ctx context.Context, inputURLs []BatchInputParams) ([]BatchOutputParams, error) {
	var output []BatchOutputParams
	for _, inputURL := range inputURLs {
		err := app.ValidateURL(inputURL.URL)
		if err != nil {
			return nil, err
		}
		genURL := app.GenerateShortURL(m.config.ShortStringLength, m.config.BaseURL)
		m.Store[genURL] = inputURL.URL
		output = append(output, BatchOutputParams{CorrelationID: inputURL.CorrelationID, ShortURL: genURL})
	}
	return output, nil
}

// FindByUser поиск по пользователю, ctx - контекст
func (m MemoryStorage) FindByUser(ctx context.Context) ([]FindByUserOutputParams, error) {
	var output []FindByUserOutputParams
	return output, nil
}

// Remove удаление, ctx - контекст, UserID - guid пользователя, shorts - массив шорткеев
func (m MemoryStorage) Remove(ctx context.Context, UserID uuid.UUID, shorts []string) error {
	return nil
}

// ShutDown завершение работы с хранилищем
func (m MemoryStorage) ShutDown() error {
	return nil
}
