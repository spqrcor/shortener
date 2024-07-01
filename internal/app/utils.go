package app

import (
	"math/rand"
	"time"
)

func GenerateShortURL(baseURL string, length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	seed := rand.NewSource(time.Now().UnixNano())
	random := rand.New(seed)

	buf := make([]byte, length)
	for i := range buf {
		buf[i] = charset[random.Intn(len(charset))]
	}
	return baseURL + "/" + string(buf)
}
