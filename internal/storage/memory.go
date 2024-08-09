package storage

import (
	"context"
	"errors"
	"shortener/internal/app"
	"shortener/internal/config"
)

type MemoryStorage struct {
	Store map[string]string
}

func CreateMemoryStorage() {
	Source = MemoryStorage{
		Store: map[string]string{},
	}
}

func (m MemoryStorage) Add(ctx context.Context, inputURL string) (string, error) {
	genURL, err := app.CreateShortURL(inputURL)
	if err != nil {
		return "", err
	}
	m.Store[genURL] = inputURL
	return genURL, nil
}

func (m MemoryStorage) Find(ctx context.Context, key string) (string, error) {
	redirectURL, ok := m.Store[config.Cfg.BaseURL+key]
	if ok {
		return redirectURL, nil
	}
	return "", errors.New("ключ не найден")
}

func (m MemoryStorage) BatchAdd(ctx context.Context, inputURLs []BatchInputParams) ([]BatchOutputParams, error) {
	var output []BatchOutputParams
	for _, inputURL := range inputURLs {
		genURL, err := app.CreateShortURL(inputURL.URL)
		if err != nil {
			return nil, err
		}
		m.Store[genURL] = inputURL.URL
		output = append(output, BatchOutputParams{CorrelationID: inputURL.CorrelationID, ShortURL: genURL})
	}
	return output, nil
}
