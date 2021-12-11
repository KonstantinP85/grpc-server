package repository

import (
	"database/sql"
	"fmt"
	"github.com/Finnhub-Stock-API/finnhub-go"
	grpc_server "github.com/KonstantinP85/grpc-server"
)

type NewsDB struct {
	db *sql.DB
}

func NewNews(db *sql.DB) *NewsDB {
	return &NewsDB{db: db}
}

func (r *NewsDB) Upload(NewsList []finnhub.News) (int32, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var count int32
	count = 0
	for _, news := range NewsList {
		createQuery := fmt.Sprintf("INSERT INTO %s (category, datetime, headline, image, related, resource, summary, url) VALUES "+
			"(?, ?, ?, ?, ?, ?, ?, ?)", newsTable)
		stmt, err := r.db.Prepare(createQuery)
		if err != nil {
			return 0, err
		}
		if _, err := stmt.Exec(news.Category, news.Datetime, news.Headline, news.Image, news.Related, news.Source, news.Summary, news.Url); err != nil {
			tx.Rollback()
			return 0, err
		}
		count = count + 1
	}

	tx.Commit()

	return count, nil
}

func (r *NewsDB) GetNews(id int) (grpc_server.News, error) {
	var news grpc_server.News

	queryBook := fmt.Sprintf("SELECT * FROM %s WHERE id = ?", newsTable)
	row := r.db.QueryRow(queryBook, id)
	err := row.Scan(&news.Id, &news.Category, &news.Datetime, &news.Headline, &news.Image, &news.Related, &news.Resource, &news.Summary, &news.Url)

	return news, err
}
