package storage

import (
	"context"
	"github.com/stretchr/testify/assert"
	"reflect"
	"shortener/internal/config"
	"testing"
)

func TestCreateMemoryStorage(t *testing.T) {
	conf := config.NewConfig()

	store := CreateMemoryStorage(conf)
	assert.Equal(t, reflect.TypeOf(store).String() == "storage.MemoryStorage", true)
}

func TestMemoryStorage_Add(t *testing.T) {
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
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := MemoryStorage{
				Store: map[string]string{},
			}
			_, err := m.Add(context.Background(), tt.inputURL)
			assert.Equal(t, tt.want, err == nil)
		})
	}
}

func TestMemoryStorage_BatchAdd(t *testing.T) {
	tests := []struct {
		name        string
		inputParams []BatchInputParams
		want        bool
	}{
		{
			"Error add",
			[]BatchInputParams{
				BatchInputParams{
					CorrelationID: "b9253cb9-03e9-4850-a3cb-16e84e9f8a37",
					URL:           "1http://lenta.ru",
				},
			},
			false,
		},
		{
			"Current add",
			[]BatchInputParams{
				BatchInputParams{
					CorrelationID: "b9253cb9-03e9-4850-a3cb-16e84e9f8a37",
					URL:           "http://lenta.ru",
				},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := MemoryStorage{
				Store: map[string]string{},
			}
			_, err := m.BatchAdd(context.Background(), tt.inputParams)
			assert.Equal(t, tt.want, err == nil)
		})
	}
}

func TestMemoryStorage_Find(t *testing.T) {
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
			m := MemoryStorage{
				Store:  map[string]string{"http://localhost:8080/fakeurl": "http://ya.ru"},
				config: config.Config{BaseURL: "http://localhost:8080"},
			}
			_, err := m.Find(context.Background(), tt.inputURI)
			assert.Equal(t, tt.want, err == nil)
		})
	}
}
