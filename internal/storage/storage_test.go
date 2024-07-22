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
			genUrl, _ := Add(tt.inputURL)
			if tt.wantErr {
				assert.Empty(t, genUrl)
			} else {
				assert.NotEmpty(t, genUrl)
			}
		})
	}
}

func TestFind(t *testing.T) {
	type args struct {
		key string
	}
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
				genUrl, _ := Add("https://ya.ru")
				result, _ := url.ParseRequestURI(genUrl)
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
