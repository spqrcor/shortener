package services

import (
	"context"
	"github.com/google/uuid"
	"shortener/internal/logger"
	"shortener/internal/storage"
	"sync"
)

const batchSize = 25

func DeleteShortURL(UserID uuid.UUID, shorts []string) {
	urlChan := make(chan string, len(shorts))
	var wg sync.WaitGroup

	go func() {
		for _, shortURL := range shorts {
			urlChan <- shortURL
		}
		close(urlChan)
	}()

	go func() {
		defer wg.Wait()
		var buffer []string
		for shortURL := range urlChan {
			buffer = append(buffer, shortURL)
			if len(buffer) >= batchSize {
				wg.Add(1)
				go func(urls []string) {
					defer wg.Done()
					err := storage.Source.Remove(context.Background(), UserID, urls)
					if err != nil {
						logger.Log.Error("Remove error " + err.Error())
					}
				}(buffer)
				buffer = nil
			}
		}
		if len(buffer) > 0 {
			wg.Add(1)
			go func(urls []string) {
				defer wg.Done()
				err := storage.Source.Remove(context.Background(), UserID, urls)
				if err != nil {
					logger.Log.Error("Remove error " + err.Error())
				}
			}(buffer)
		}
	}()
}
