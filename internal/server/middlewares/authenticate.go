package middlewares

import (
	"context"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"net"
	"net/http"
	"shortener/internal/authenticate"
)

// AuthenticateMiddleware middlewares для аутентификации, logger - логгер, auth - сервис аутентификации, trustedSubnet - доверенная подсеть
func AuthenticateMiddleware(logger *zap.Logger, auth authenticate.Auth, trustedSubnet string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
			if r.RequestURI == "/api/internal/stats" && trustedSubnet != "" {
				_, ipnet, _ := net.ParseCIDR(trustedSubnet)
				ipB := net.ParseIP(r.Header.Get(`X-Real-IP`))
				if !ipnet.Contains(ipB) {
					http.Error(rw, "403", http.StatusForbidden)
					return
				}
			}

			UserID := uuid.New()
			cookie, err := r.Cookie("Authorization")
			if err != nil {
				if r.RequestURI == "/api/user/urls" && r.Method == http.MethodGet {
					http.Error(rw, err.Error(), http.StatusUnauthorized)
					return
				} else {
					if err = auth.SetCookie(rw, UserID); err != nil {
						logger.Error(err.Error())
					}
				}
			} else {
				decodeUserID, err := auth.GetUserIDFromCookie(cookie.Value)
				if err != nil {
					logger.Error(err.Error())
					http.Error(rw, err.Error(), http.StatusInternalServerError)
					return
				} else {
					UserID = decodeUserID
				}
			}
			ctx := context.WithValue(r.Context(), authenticate.ContextUserID, UserID)
			next.ServeHTTP(rw, r.WithContext(ctx))
		})
	}
}
