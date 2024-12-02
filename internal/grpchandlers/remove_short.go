package grpchandlers

import (
	"context"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/emptypb"
	"shortener/internal/authenticate"
	"shortener/internal/server/proto"
)

// DeleteURLs удаление
func (s *ShortenerServer) DeleteURLs(ctx context.Context, req *proto.DeleteURLsRequest) (*emptypb.Empty, error) {
	UserID, ok := ctx.Value(authenticate.ContextUserID).(uuid.UUID)
	if ok {
		go s.BatchRemove.DeleteShortURL(UserID, req.URLs)
	}
	return &emptypb.Empty{}, nil
}
