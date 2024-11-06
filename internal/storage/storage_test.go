package storage

import (
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"os"
	"reflect"
	"shortener/internal/config"
	"shortener/internal/logger"
	"testing"
)

func TestNewStorage(t *testing.T) {
	conf := config.NewConfig()
	loggerRes, _ := logger.NewLogger(zap.InfoLevel)

	tests := []struct {
		name string
		want string
	}{
		{
			"MemoryStorage",
			"storage.MemoryStorage",
		},
		{
			"FileStorage",
			"storage.FileStorage",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.name == "FileStorage" {
				conf.FileStoragePath = os.TempDir() + "/fake_short.txt"
			}
			obj := NewStorage(conf, loggerRes)
			assert.Equal(t, reflect.TypeOf(obj).String() == tt.want, true)
		})
	}
}
