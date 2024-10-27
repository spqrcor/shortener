package services

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"reflect"
	"shortener/internal/config"
	"shortener/internal/logger"
	"shortener/internal/storage"
	"testing"
)

func TestBatchRemove_DeleteShortURL(t *testing.T) {
	conf := config.NewConfig()
	loggerRes, _ := logger.NewLogger(zap.InfoLevel)

	store := storage.CreateMemoryStorage(conf)
	obj := NewBatchRemoveService(loggerRes, store)
	obj.DeleteShortURL(uuid.New(), []string{"fake"})
}

func TestNewBatchRemoveService(t *testing.T) {
	conf := config.NewConfig()
	loggerRes, _ := logger.NewLogger(zap.InfoLevel)

	store := storage.CreateMemoryStorage(conf)
	obj := NewBatchRemoveService(loggerRes, store)
	assert.Equal(t, reflect.TypeOf(obj).String() == "*services.BatchRemove", true)
}
