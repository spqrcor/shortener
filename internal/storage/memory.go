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
	err := app.ValidateURL(inputURL)
	if err != nil {
		return "", err
	}
	genURL := app.GenerateShortURL()
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
		err := app.ValidateURL(inputURL.URL)
		if err != nil {
			return nil, err
		}
		genURL := app.GenerateShortURL()
		m.Store[genURL] = inputURL.URL
		output = append(output, BatchOutputParams{CorrelationID: inputURL.CorrelationID, ShortURL: genURL})
	}
	return output, nil
}

func (m MemoryStorage) FindByUser(ctx context.Context) ([]FindByUserOutputParams, error) {
	var output []FindByUserOutputParams
	return output, nil
}
