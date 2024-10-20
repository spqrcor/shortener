package logger

import (
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap/zapcore"
	"testing"
)

func TestNewLogger(t *testing.T) {
	tests := []struct {
		name     string
		logLevel zapcore.Level
	}{
		{
			name:     "success",
			logLevel: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logger, err := NewLogger(tt.logLevel)
			assert.Nil(t, err)
			assert.NotNil(t, logger)

		})
	}
}
