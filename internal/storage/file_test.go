package storage

import (
	"context"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"os"
	"reflect"
	"shortener/internal/config"
	"shortener/internal/logger"
	"testing"
)

func TestFileStorage_FindByUser(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		{
			"Success",
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := FileStorage{
				Store:  map[string]string{"http://localhost:8080/fakeurl": "http://ya.ru"},
				config: config.Config{BaseURL: "http://localhost:8080"},
			}
			_, err := m.FindByUser(context.Background())
			assert.Equal(t, tt.want, err == nil)
		})
	}
}

func TestFileStorage_Remove(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		{
			"Success",
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := FileStorage{
				Store:  map[string]string{"http://localhost:8080/fakeurl": "http://ya.ru"},
				config: config.Config{BaseURL: "http://localhost:8080"},
			}
			err := m.Remove(context.Background(), uuid.New(), []string{"xxx"})
			assert.Equal(t, tt.want, err == nil)
		})
	}
}

func TestCreateFileStorage(t *testing.T) {
	conf := config.NewConfig()
	loggerRes, _ := logger.NewLogger(zap.InfoLevel)

	if conf.FileStoragePath == "" {
		conf.FileStoragePath = os.TempDir() + "/fake_short.txt"
	}
	store := CreateFileStorage(conf, loggerRes)
	assert.Equal(t, reflect.TypeOf(store).String() == "storage.FileStorage", true)
}

func TestFileStorage_Add(t *testing.T) {
	conf := config.NewConfig()
	loggerRes, _ := logger.NewLogger(zap.InfoLevel)

	tests := []struct {
		name     string
		inputURL string
		want     bool
	}{
		{
			"Error add",
			"1http://ya.ru",
			false,
		},
		{
			"Current add",
			"http://ya.ru",
			true,
		},
	}
	if conf.FileStoragePath == "" {
		conf.FileStoragePath = os.TempDir() + "/fake_short.txt"
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := CreateFileStorage(conf, loggerRes)
			_, err := m.Add(context.Background(), tt.inputURL)
			assert.Equal(t, tt.want, err == nil)
		})
	}
}

func TestFileStorage_BatchAdd(t *testing.T) {
	conf := config.NewConfig()
	loggerRes, _ := logger.NewLogger(zap.InfoLevel)
	tests := []struct {
		name        string
		inputParams []BatchInputParams
		want        bool
	}{
		{
			"Error add",
			[]BatchInputParams{
				{
					CorrelationID: "b9253cb9-03e9-4850-a3cb-16e84e9f8a37",
					URL:           "1http://lenta.ru",
				},
			},
			false,
		},
		{
			"Current add",
			[]BatchInputParams{
				{
					CorrelationID: "b9253cb9-03e9-4850-a3cb-16e84e9f8a37",
					URL:           "http://lenta.ru",
				},
			},
			true,
		},
	}
	if conf.FileStoragePath == "" {
		conf.FileStoragePath = os.TempDir() + "/fake_short.txt"
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := CreateFileStorage(conf, loggerRes)
			_, err := m.BatchAdd(context.Background(), tt.inputParams)
			assert.Equal(t, tt.want, err == nil)
		})
	}
}

func TestFileStorage_Find(t *testing.T) {
	tests := []struct {
		name     string
		inputURI string
		want     bool
	}{
		{
			"Error find",
			"/xxx",
			false,
		},
		{
			"Current find",
			"/fakeurl",
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := FileStorage{
				Store:  map[string]string{"http://localhost:8080/fakeurl": "http://ya.ru"},
				config: config.Config{BaseURL: "http://localhost:8080"},
			}
			_, err := m.Find(context.Background(), tt.inputURI)
			assert.Equal(t, tt.want, err == nil)
		})
	}
}

func TestFileStorage_CreateFileStorage(t *testing.T) {
	conf := config.NewConfig()
	loggerRes, _ := logger.NewLogger(zap.InfoLevel)

	if conf.FileStoragePath == "" {
		conf.FileStoragePath = os.TempDir() + "/fake_short.txt"
	}

	store := CreateFileStorage(conf, loggerRes)
	assert.Equal(t, reflect.TypeOf(store).String() == "storage.FileStorage", true)
}

func TestFileStorage_ShutDown(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			"Success",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := FileStorage{
				Store:  map[string]string{"http://localhost:8080/fakeurl": "http://ya.ru"},
				config: config.Config{BaseURL: "http://localhost:8080"},
			}
			err := m.ShutDown()
			assert.Nil(t, err)
		})
	}
}

func TestFileStorage_Stat(t *testing.T) {
	tests := []struct {
		name      string
		urlsCount int
	}{
		{
			"success",
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := FileStorage{
				Store:  map[string]string{"http://localhost:8080/fakeurl": "http://ya.ru"},
				config: config.Config{BaseURL: "http://localhost:8080"},
			}
			stat, _ := m.Stat(context.Background())
			assert.Equal(t, tt.urlsCount, stat.Urls)
		})
	}
}
