package app

import (
	"math/rand"
	"time"
)

func GenerateShortUrl(baseUrl string, length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	seed := rand.NewSource(time.Now().UnixNano())
	random := rand.New(seed)

	buf := make([]byte, length)
	for i := range buf {
		buf[i] = charset[random.Intn(len(charset))]
	}
	return baseUrl + "/" + string(buf)
}
