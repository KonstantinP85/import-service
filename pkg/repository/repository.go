package repository

import (
	"database/sql"
	"github.com/Finnhub-Stock-API/finnhub-go"
	import_service "github.com/KonstantinP85/import-service"
)

type News interface {
	Upload([]finnhub.News) error
	GetNews(id int) (import_service.News, error)
}

type Repository struct {
	News
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		News: NewNews(db),
	}
}
