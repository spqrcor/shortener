// Package services сервисы
package services

import (
	"context"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"shortener/internal/storage"
	"sync"
)

// batchSize размер пачки для группового удаление
const batchSize = 25

// BatchRemover интерфейс сервиса группового удаления
type BatchRemover interface {
	DeleteShortURL(UserID uuid.UUID, shorts []string)
}

// BatchRemove тип сервиса группового удаления
type BatchRemove struct {
	logger  *zap.Logger
	storage storage.Storage
}

// NewBatchRemoveService создание сервиса группового удаления, logger - логгер, storage - сервис хранилища
func NewBatchRemoveService(logger *zap.Logger, storage storage.Storage) *BatchRemove {
	return &BatchRemove{
		logger:  logger,
		storage: storage,
	}
}

// DeleteShortURL удаление записей, UserID - guid пользователя, shorts массив записей
func (b *BatchRemove) DeleteShortURL(UserID uuid.UUID, shorts []string) {
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
					err := b.storage.Remove(context.Background(), UserID, urls)
					if err != nil {
						b.logger.Error("Remove error " + err.Error())
					}
				}(buffer)
				buffer = nil
			}
		}
		if len(buffer) > 0 {
			wg.Add(1)
			go func(urls []string) {
				defer wg.Done()
				err := b.storage.Remove(context.Background(), UserID, urls)
				if err != nil {
					b.logger.Error("Remove error " + err.Error())
				}
			}(buffer)
		}
	}()
}
