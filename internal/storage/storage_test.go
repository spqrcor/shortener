package storage

import (
	"github.com/stretchr/testify/assert"
	"net/url"
	"testing"
)

func TestAdd(t *testing.T) {
	tests := []struct {
		name     string
		inputURL string
		wantErr  bool
	}{
		{
			"Empty inputURL",
			"",
			true,
		},
		{
			"Invalid format inputURL",
			"1https://ya.ru",
			true,
		},
		{
			"Current inputURL",
			"https://ya.ru",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			genURL, _ := Add(tt.inputURL)
			if tt.wantErr {
				assert.Empty(t, genURL)
			} else {
				assert.NotEmpty(t, genURL)
			}
		})
	}
}

func TestFind(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			"Invalid Key",
			true,
		},
		{
			"Current Key",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			key := "INVALID_KEY"

			if !tt.wantErr {
				genURL, _ := Add("https://ya.ru")
				result, _ := url.ParseRequestURI(genURL)
				key = result.Path
			}

			redirectURL, _ := Find(key)
			if tt.wantErr {
				assert.Empty(t, redirectURL)
			} else {
				assert.NotEmpty(t, redirectURL)
			}
		})
	}
}
