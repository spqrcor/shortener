package grpchandlers

import (
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	"shortener/internal/config"
	"shortener/internal/db"
)

// PingDB Ping
func (s *ShortenerServer) PingDB(ctx context.Context, req *emptypb.Empty) (*emptypb.Empty, error) {
	conf := config.NewConfig()
	if conf.DatabaseDSN != "" {
		_, err := db.Connect(conf.DatabaseDSN)
		if err != nil {
			return &emptypb.Empty{}, err
		}
	}
	return &emptypb.Empty{}, nil
}
