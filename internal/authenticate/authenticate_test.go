package authenticate

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"net/http/httptest"
	"reflect"
	"shortener/internal/config"
	"shortener/internal/logger"
	"testing"
)

func BenchmarkAuthenticate_GetUserIDFromCookie(b *testing.B) {
	conf := config.NewConfig()
	loggerRes, _ := logger.NewLogger(zap.InfoLevel)
	a := &Authenticate{
		logger:    loggerRes,
		secretKey: conf.SecretKey,
		tokenExp:  conf.TokenExp,
	}
	b.Run("default", func(b *testing.B) {
		cookieValue := "30316566363932392d373133392d363036632d393466312d303031353564336462623865da1bc3dd8b597c82ec439df707f6fc8e988ef19082aac7950957830c1bb4aa2a"
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_, _ = a.GetUserIDFromCookie(cookieValue)
		}
	})
}

func TestAuthenticate_GetUserIDFromCookie(t *testing.T) {
	conf := config.NewConfig()
	loggerRes, _ := logger.NewLogger(zap.InfoLevel)

	a := &Authenticate{
		logger:    loggerRes,
		secretKey: conf.SecretKey,
		tokenExp:  conf.TokenExp,
	}
	token, _ := a.CreateToken(uuid.New())

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
		{
			"Success",
			token,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
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

func BenchmarkAuthenticate_createCookie(b *testing.B) {
	conf := config.NewConfig()
	loggerRes, _ := logger.NewLogger(zap.InfoLevel)
	a := &Authenticate{
		logger:    loggerRes,
		secretKey: conf.SecretKey,
		tokenExp:  conf.TokenExp,
	}
	b.Run("default", func(b *testing.B) {
		userID := uuid.New()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_, _ = a.createCookie(userID)
		}
	})
}

func TestAuthenticate_createToken(t *testing.T) {
	conf := config.NewConfig()
	loggerRes, _ := logger.NewLogger(zap.InfoLevel)

	tests := []struct {
		name string
		key  string
	}{
		{
			"success",
			conf.SecretKey,
		},
		{
			"false",
			"",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Authenticate{
				logger:    loggerRes,
				secretKey: conf.SecretKey,
				tokenExp:  conf.TokenExp,
			}
			_, err := a.CreateToken(uuid.New())
			assert.Nil(t, err)
		})
	}
}

func TestAuthenticate_SetCookie(t *testing.T) {
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
			err := a.SetCookie(httptest.NewRecorder(), uuid.New())
			assert.Nil(t, err)
		})
	}
}

func BenchmarkAuthenticate_createToken(b *testing.B) {
	conf := config.NewConfig()
	loggerRes, _ := logger.NewLogger(zap.InfoLevel)
	a := &Authenticate{
		logger:    loggerRes,
		secretKey: conf.SecretKey,
		tokenExp:  conf.TokenExp,
	}
	b.Run("default", func(b *testing.B) {
		userID := uuid.New()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_, _ = a.CreateToken(userID)
		}
	})
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
