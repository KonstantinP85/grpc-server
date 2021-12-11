package repository

import (
	"database/sql"
	"github.com/Finnhub-Stock-API/finnhub-go"
	grpc_server "github.com/KonstantinP85/grpc-server"
)

type News interface {
	Upload([]finnhub.News) (int32, error)
	GetNews(id int) (grpc_server.News, error)
}

type Repository struct {
	News
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		News: NewNews(db),
	}
}
