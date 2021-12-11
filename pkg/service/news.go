package service

import (
	"context"
	"github.com/Finnhub-Stock-API/finnhub-go"
	grpc_server "github.com/KonstantinP85/grpc-server"
	"github.com/KonstantinP85/grpc-server/pkg/repository"
	"os"
)

type NewsService struct {
	repo repository.News
}

func NewNewsService(repo repository.News) *NewsService {
	return &NewsService{repo: repo}
}

func (s *NewsService) Upload() (int32, error) {

	cfg := finnhub.NewConfiguration()
	cfg.AddDefaultHeader("X-Finnhub-Token", os.Getenv("FINNHUB_TOKEN"))
	finnhubClient := finnhub.NewAPIClient(cfg).DefaultApi

	generalNews, _, err := finnhubClient.GeneralNews(context.Background(), "general", nil)
	if err != nil {
		return 0, err
	}

	var count int32
	count, err = s.repo.Upload(generalNews)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (s *NewsService) GetNews(id int) (grpc_server.News, error) {
	return s.repo.GetNews(id)
}
