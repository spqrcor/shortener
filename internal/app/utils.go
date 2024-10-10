package app

import (
	"fmt"
	"math/rand"
	"net/url"
	"time"
)

var ErrURLFormat = fmt.Errorf("url format error")
var ErrURLEmpty = fmt.Errorf("url empty error")

func GenerateShortURL(stringLength int, baseURL string) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	random := rand.New(rand.NewSource(time.Now().UnixNano()))

	buf := make([]byte, stringLength)
	for i := range buf {
		buf[i] = charset[random.Intn(len(charset))]
	}
	return baseURL + "/" + string(buf)
}

func ValidateURL(inputURL string) error {
	if inputURL == "" {
		return ErrURLEmpty
	}

	_, err := url.ParseRequestURI(inputURL)
	if err != nil {
		return ErrURLFormat
	}
	return nil
}
