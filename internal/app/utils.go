package app

import (
	"math/rand"
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
