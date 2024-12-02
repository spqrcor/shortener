package grpchandlers

import (
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	"shortener/internal/server/proto"
)

// UserURLs поиск по пользователю
func (s *ShortenerServer) UserURLs(ctx context.Context, req *emptypb.Empty) (*proto.UsersURLsResponse, error) {
	output, err := s.Storage.FindByUser(ctx)
	if err != nil {
		return nil, err
	}

	response := &proto.UsersURLsResponse{
		Urls: make([]*proto.UsersURLsResponse_URL, len(output)),
	}
	for i, u := range output {
		response.Urls[i].Original = u.OriginalURL
		response.Urls[i].Short = u.ShortURL
	}
	return response, nil
}
