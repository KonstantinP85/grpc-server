package news

import (
	"context"
	"github.com/KonstantinP85/grpc-server/pkg/api"
	"github.com/KonstantinP85/grpc-server/pkg/service"
	"strconv"
)

type Server struct {
	services *service.Service
}

func NewServer(services *service.Service) *Server {
	return &Server{services: services}
}

func (s *Server) GetNews(ctx context.Context, r *api.GetNewsRequest) (*api.GetNewsResponse, error) {

	data, err := s.services.News.GetNews(int(r.Id))
	if err != nil {
		return nil, err
	}

	news := &api.News{
		Id:       strconv.Itoa(data.Id),
		Category: data.Category,
		Datetime: int64(data.Datetime),
		Headline: data.Headline,
		Image:    data.Image,
		Related:  data.Related,
		Resource: data.Resource,
		Summary:  data.Summary,
		Url:      data.Url,
	}

	return &api.GetNewsResponse{News: news}, nil
}

func (s *Server) UploadNews(ctx context.Context, r *api.UploadNewsRequest) (*api.UploadNewsResponse, error) {

	count, err := s.services.News.Upload()
	if err != nil {
		return &api.UploadNewsResponse{Count: 0}, err
	}

	return &api.UploadNewsResponse{Count: count}, nil
}
