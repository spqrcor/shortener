package storage

import (
	"errors"
	"net/url"
	"shortener/internal/app"
	"shortener/internal/config"
)

var store = map[string]string{}

func Add(inputURL string) (string, error) {
	if inputURL == "" {
		return "", errors.New("входящее значение пустое")
	}

	_, err := url.ParseRequestURI(inputURL)
	if err != nil {
		return "", errors.New("неверный формат URL")
	}

	genURL := app.GenerateShortURL()

	store[genURL] = inputURL
	return genURL, nil
}

func Find(key string) (string, error) {
	redirectURL, ok := store[config.Cfg.BaseURL+key]
	if ok {
		return redirectURL, nil
	}
	return "", errors.New("ключ не найден")
}
