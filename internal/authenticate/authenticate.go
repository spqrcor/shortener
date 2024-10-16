// Package authenticate аутентификация
package authenticate

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"net/http"
	"strings"
	"time"
)

// claims тип по работе с токеном
type claims struct {
	jwt.RegisteredClaims
	UserID uuid.UUID
}

// ContextKey тип для хранения UserID в контексте
type ContextKey string

// ContextUserID для хранения UserID в контексте
var ContextUserID ContextKey = "UserID"

// Auth интерфейс аутентификации
type Auth interface {
	GetUserIDFromCookie(tokenString string) (uuid.UUID, error)
	SetCookie(rw http.ResponseWriter, UserID uuid.UUID)
}

// Authenticate аутентификация
type Authenticate struct {
	logger    *zap.Logger
	secretKey string
	tokenExp  time.Duration
}

// NewAuthenticateService создание Authenticate, opts - набор параметров
func NewAuthenticateService(opts ...func(*Authenticate)) *Authenticate {
	auth := &Authenticate{}
	for _, opt := range opts {
		opt(auth)
	}
	return auth
}

// WithLogger добавление logger
func WithLogger(logger *zap.Logger) func(*Authenticate) {
	return func(a *Authenticate) {
		a.logger = logger
	}
}

// WithSecretKey добавление secretKey
func WithSecretKey(secretKey string) func(*Authenticate) {
	return func(a *Authenticate) {
		a.secretKey = secretKey
	}
}

// WithTokenExp добавление tokenExp
func WithTokenExp(tokenExp time.Duration) func(*Authenticate) {
	return func(a *Authenticate) {
		a.tokenExp = tokenExp
	}
}

// createCookie создание cookie, UserID - guid пользователя
func (a *Authenticate) createCookie(UserID uuid.UUID) (http.Cookie, error) {
	token, err := a.createToken(UserID)
	if err != nil {
		return http.Cookie{}, err
	}
	return http.Cookie{Name: "Authorization", Value: token, Expires: time.Now().Add(a.tokenExp), HttpOnly: true, Path: "/"}, nil
}

// createToken создание токена, UserID - guid пользователя
func (a *Authenticate) createToken(UserID uuid.UUID) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(a.tokenExp)),
		},
		UserID: UserID,
	})

	tokenString, err := token.SignedString([]byte(a.secretKey))
	if err != nil {
		return "", err
	}
	return "Bearer " + tokenString, nil
}

// GetUserIDFromCookie получение UserID из токена, tokenString - токен
func (a *Authenticate) GetUserIDFromCookie(tokenString string) (uuid.UUID, error) {
	claims := &claims{}
	token, err := jwt.ParseWithClaims(strings.TrimPrefix(tokenString, "Bearer "), claims,
		func(t *jwt.Token) (interface{}, error) {
			return []byte(a.secretKey), nil
		})
	if err != nil {
		return uuid.Nil, err
	}

	if !token.Valid {
		return uuid.Nil, fmt.Errorf("invalid token")
	}
	return claims.UserID, nil
}

// SetCookie установка cookie, rw - http.ResponseWriter, UserID - guid пользователя
func (a *Authenticate) SetCookie(rw http.ResponseWriter, UserID uuid.UUID) {
	cookie, err := a.createCookie(UserID)
	if err != nil {
		a.logger.Error(err.Error())
	} else {
		http.SetCookie(rw, &cookie)
	}
}
