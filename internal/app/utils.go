package app

import (
	"errors"
	"math/rand"
	"net/url"
	"shortener/internal/config"
	"time"
)

func GenerateShortURL() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	random := rand.New(rand.NewSource(time.Now().UnixNano()))

	buf := make([]byte, config.Cfg.ShortStringLength)
	for i := range buf {
		buf[i] = charset[random.Intn(len(charset))]
	}
	return config.Cfg.BaseURL + "/" + string(buf)
}

func ValidateURL(inputURL string) error {
	if inputURL == "" {
		return errors.New("входящее значение пустое")
	}

	_, err := url.ParseRequestURI(inputURL)
	if err != nil {
		return errors.New("неверный формат URL")
	}
	return nil
}
