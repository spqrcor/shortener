package grpchandlers

import (
	"context"
	"shortener/internal/server/proto"
	"shortener/internal/storage"
)

// ShortenBatch групповое добавление
func (s *ShortenerServer) ShortenBatch(ctx context.Context, req *proto.ShortenBatchRequest) (*proto.ShortenBatchResponse, error) {
	var input []storage.BatchInputParams
	for _, row := range req.Urls {
		input = append(input, storage.BatchInputParams{CorrelationID: row.CorrelationId, URL: row.OriginalUrl})
	}
	URLs, err := s.Storage.BatchAdd(ctx, input)
	if err != nil {
		return nil, err
	}

	respURLs := make([]*proto.ShortenBatchResponse_URL, len(URLs))
	for i, url := range URLs {
		respURLs[i] = &proto.ShortenBatchResponse_URL{
			CorrelationId: url.CorrelationID,
			ShortenUrl:    url.ShortURL,
		}
	}
	response := &proto.ShortenBatchResponse{
		Urls: respURLs,
	}
	return response, nil
}
