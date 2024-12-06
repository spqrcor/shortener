package grpchandlers

import (
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	"shortener/internal/server/proto"
)

// Stats Статистика
func (s *ShortenerServer) Stats(ctx context.Context, req *emptypb.Empty) (*proto.StatsResponse, error) {
	stat, err := s.Storage.Stat(ctx)
	if err != nil {
		return nil, err
	}
	return &proto.StatsResponse{
		UrlsAmount:  uint64(stat.Urls),
		UsersAmount: uint32(stat.Users),
	}, nil
}
