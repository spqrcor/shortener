package db

import (
	"github.com/stretchr/testify/assert"
	"shortener/internal/config"
	"testing"
)

func TestConnect(t *testing.T) {
	conf := config.NewConfig()
	if conf.DatabaseDSN == "" {
		t.Skip("Skipping testing...")
	}

	tests := []struct {
		name string
		dsn  string
		want bool
	}{
		{
			"Error",
			"",
			false,
		},
		{
			"Success",
			conf.DatabaseDSN,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := Connect(tt.dsn)
			assert.Equal(t, tt.want, err == nil)
		})
	}
}
