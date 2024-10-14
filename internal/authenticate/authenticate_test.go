package authenticate

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"reflect"
	"shortener/internal/config"
	"shortener/internal/logger"
	"testing"
)

func TestAuthenticate_GetUserIDFromCookie(t *testing.T) {
	conf := config.NewConfig()
	loggerRes, _ := logger.NewLogger(zap.InfoLevel)

	tests := []struct {
		name   string
		token  string
		result bool
	}{
		{
			"Error",
			"_____",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Authenticate{
				logger:    loggerRes,
				secretKey: conf.SecretKey,
				tokenExp:  conf.TokenExp,
			}
			_, err := a.GetUserIDFromCookie(tt.token)
			assert.Equal(t, err == nil, tt.result)
		})
	}
}

func TestAuthenticate_createCookie(t *testing.T) {
	conf := config.NewConfig()
	loggerRes, _ := logger.NewLogger(zap.InfoLevel)

	tests := []struct {
		name string
	}{
		{
			"success",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Authenticate{
				logger:    loggerRes,
				secretKey: conf.SecretKey,
				tokenExp:  conf.TokenExp,
			}
			_, err := a.createCookie(uuid.New())
			assert.Nil(t, err)
		})
	}
}

func TestAuthenticate_createToken(t *testing.T) {
	conf := config.NewConfig()
	loggerRes, _ := logger.NewLogger(zap.InfoLevel)

	tests := []struct {
		name string
	}{
		{
			"success",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Authenticate{
				logger:    loggerRes,
				secretKey: conf.SecretKey,
				tokenExp:  conf.TokenExp,
			}
			_, err := a.createToken(uuid.New())
			assert.Nil(t, err)
		})
	}
}

func TestNewAuthenticateService(t *testing.T) {
	conf := config.NewConfig()
	loggerRes, _ := logger.NewLogger(zap.InfoLevel)
	store := NewAuthenticateService(
		WithLogger(loggerRes),
		WithSecretKey(conf.SecretKey),
		WithTokenExp(conf.TokenExp),
	)
	assert.Equal(t, reflect.TypeOf(store).String() == "*authenticate.Authenticate", true)
}
