package grpchandlers

import (
	"context"
	"shortener/internal/server/proto"
)

// Shorten добавление
func (s *ShortenerServer) Shorten(ctx context.Context, req *proto.ShortenRequest) (*proto.ShortenResponse, error) {
	genURL, err := s.Storage.Add(ctx, req.OriginalUrl)
	if err != nil {
		return nil, err
	}
	return &proto.ShortenResponse{Shorten: genURL}, nil
}
