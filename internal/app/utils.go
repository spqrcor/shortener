package app

import (
	"errors"
	"math/rand"
	"net/url"
	"time"
)

func GenerateShortURL(stringLength int, baseUrl string) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	random := rand.New(rand.NewSource(time.Now().UnixNano()))

	buf := make([]byte, stringLength)
	for i := range buf {
		buf[i] = charset[random.Intn(len(charset))]
	}
	return baseUrl + "/" + string(buf)
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
