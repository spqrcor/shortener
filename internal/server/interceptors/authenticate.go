package interceptors

import (
	"context"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"net"
	"shortener/internal/authenticate"
)

const (
	statMethod = "/grpc_server.URLShortenerService/Stats"
	ParamIP    = "X-Real-IP"
)

// ErrMissingMetadata ошибка получения метадаты
var ErrMissingMetadata = status.Errorf(codes.Unauthenticated, "missing metadata")

// ErrMissingIPAddress пропущен IP адрес
var ErrMissingIPAddress = status.Errorf(codes.Unauthenticated, "missing IP address")

// ErrAccessDenied доступ запрещен
var ErrAccessDenied = status.Errorf(codes.PermissionDenied, "access denied")

// ErrInvalidToken неверный токен
var ErrInvalidToken = status.Errorf(codes.Unauthenticated, "invalid token")

// AuthenticateInterceptor для аутентификации, logger - логгер, auth - сервис аутентификации, trustedSubnet - доверенная подсеть
func AuthenticateInterceptor(logger *zap.Logger, auth authenticate.Auth, trustedSubnet string) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			logger.Error(ErrMissingMetadata.Error())
			return nil, ErrMissingMetadata
		}

		if info.FullMethod == statMethod && trustedSubnet != "" {
			values := md.Get(ParamIP)
			if len(values) == 0 {
				logger.Error(ErrMissingIPAddress.Error())
				return nil, ErrMissingIPAddress
			}
			_, ipnet, _ := net.ParseCIDR(trustedSubnet)
			ipB := net.ParseIP(values[0])
			if !ipnet.Contains(ipB) {
				logger.Error(ErrAccessDenied.Error())
				return nil, ErrAccessDenied
			}
		}

		values := md.Get("token")
		if len(values) > 0 {
			decodeUserID, err := auth.GetUserIDFromCookie(values[0])
			if err != nil {
				logger.Error(ErrInvalidToken.Error())
				return nil, ErrInvalidToken
			}
			return handler(context.WithValue(ctx, authenticate.ContextUserID, decodeUserID), req)
		}
		return handler(ctx, req)
	}
}
