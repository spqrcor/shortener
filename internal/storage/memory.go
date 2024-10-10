package storage

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"shortener/internal/app"
	"shortener/internal/config"
)

type MemoryStorage struct {
	config config.Config
	Store  map[string]string
}

func CreateMemoryStorage(config config.Config) Storage {
	return MemoryStorage{
		config: config,
		Store:  map[string]string{},
	}
}

func (m MemoryStorage) Add(ctx context.Context, inputURL string) (string, error) {
	err := app.ValidateURL(inputURL)
	if err != nil {
		return "", err
	}
	genURL := app.GenerateShortURL(m.config.ShortStringLength, m.config.BaseURL)
	m.Store[genURL] = inputURL
	return genURL, nil
}

func (m MemoryStorage) Find(ctx context.Context, key string) (string, error) {
	redirectURL, ok := m.Store[m.config.BaseURL+key]
	if ok {
		return redirectURL, nil
	}
	return "", errors.New("ключ не найден")
}

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

func (m MemoryStorage) FindByUser(ctx context.Context) ([]FindByUserOutputParams, error) {
	var output []FindByUserOutputParams
	return output, nil
}

func (m MemoryStorage) Remove(ctx context.Context, UserID uuid.UUID, shorts []string) error {
	return nil
}
