package storage

import (
	"encoding/json"
	"errors"
	"net/url"
	"os"
	"shortener/internal/app"
	"shortener/internal/config"
	"shortener/internal/logger"
)

var store = map[string]string{}

func Init() {
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
	if err := json.Unmarshal([]byte(data), &store); err != nil {
		logger.Log.Fatal(err.Error())
	}
}

func updateFileStorage() {
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
	updateFileStorage()
	return genURL, nil
}

func Find(key string) (string, error) {
	redirectURL, ok := store[config.Cfg.BaseURL+key]
	if ok {
		return redirectURL, nil
	}
	return "", errors.New("ключ не найден")
}
