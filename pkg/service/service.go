package service

import (
	grpc_server "github.com/KonstantinP85/grpc-server"
	"github.com/KonstantinP85/grpc-server/pkg/repository"
)

type News interface {
	Upload() (int32, error)
	GetNews(id int) (grpc_server.News, error)
}

type Service struct {
	News
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		News: NewNewsService(repo.News),
	}
}
