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
			logger.Error(status.Errorf(codes.Unauthenticated, "missing metadata").Error())
			return nil, status.Errorf(codes.Unauthenticated, "missing metadata")
		}

		if info.FullMethod == "/grpc_server.URLShortenerService/Stats" && trustedSubnet != "" {
			values := md.Get("X-Real-IP")
			if len(values) == 0 {
				logger.Error(status.Errorf(codes.Unauthenticated, "missing IP address").Error())
				return nil, status.Errorf(codes.Unauthenticated, "missing IP address")
			}
			_, ipnet, _ := net.ParseCIDR(trustedSubnet)
			ipB := net.ParseIP(values[0])
			if !ipnet.Contains(ipB) {
				logger.Error(status.Errorf(codes.PermissionDenied, "access denied").Error())
				return nil, status.Errorf(codes.PermissionDenied, "access denied")
			}
		}

		values := md.Get("token")
		if len(values) > 0 {
			decodeUserID, err := auth.GetUserIDFromCookie(values[0])
			if err != nil {
				logger.Error(status.Errorf(codes.Unauthenticated, "invalid token").Error())
				return nil, status.Errorf(codes.Unauthenticated, "invalid token")
			}
			return handler(context.WithValue(ctx, authenticate.ContextUserID, decodeUserID), req)
		}
		return handler(ctx, req)
	}
}
