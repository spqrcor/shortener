package grpchandlers

import (
	"context"
	"shortener/internal/server/proto"
)

// GetOriginalURL поиск по url
func (s *ShortenerServer) GetOriginalURL(ctx context.Context, req *proto.GetOriginalURLRequest) (*proto.GetAnOriginalURLResponse, error) {
	redirectURL, err := s.Storage.Find(ctx, req.ShortUrl)
	if err != nil {
		return nil, err
	}
	res := &proto.GetAnOriginalURLResponse{
		Url: redirectURL,
	}
	return res, nil
}
