package app

import (
	"math/rand"
	"time"
)

func GenerateShortURL(baseURL string, length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	random := rand.New(rand.NewSource(time.Now().UnixNano()))

	buf := make([]byte, length)
	for i := range buf {
		buf[i] = charset[random.Intn(len(charset))]
	}
	return baseURL + "/" + string(buf)
}
