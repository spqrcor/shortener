package app

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerateShortURL(t *testing.T) {
	tests := []struct {
		name         string
		stringLength int
		baseURL      string
		resultLength int
	}{
		{
			name:         "success",
			stringLength: 6,
			baseURL:      "http://localhost",
			resultLength: 23,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.NotNil(t, GenerateShortURL(tt.stringLength, tt.baseURL))
			assert.Equal(t, tt.resultLength, len(GenerateShortURL(tt.stringLength, tt.baseURL)))
		})
	}
}

func TestValidateURL(t *testing.T) {
	tests := []struct {
		name   string
		url    string
		result error
	}{
		{
			name:   "success",
			url:    "https://ya.ru",
			result: nil,
		},
		{
			name:   "empty error",
			url:    "",
			result: ErrURLEmpty,
		},
		{
			name:   "format error",
			url:    "1https://ya.ru",
			result: ErrURLFormat,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, ValidateURL(tt.url), tt.result)
		})
	}
}
