package interceptors

import (
	"context"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

// LoggerInterceptor для логирования, logger - логгер
func LoggerInterceptor(logger *zap.Logger) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		logger.Info("GRPC request",
			zap.String("method", info.FullMethod),
		)
		return handler(ctx, req)
	}
}
