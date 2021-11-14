package repository

import (
	"database/sql"
	"fmt"
	"github.com/Finnhub-Stock-API/finnhub-go"
	import_service "github.com/KonstantinP85/import-service"
)

type NewsDB struct {
	db *sql.DB
}

func NewNews(db *sql.DB) *NewsDB {
	return &NewsDB{db: db}
}

func (r *NewsDB) Upload(NewsList []finnhub.News) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	for _, news := range NewsList {
		createQuery := fmt.Sprintf("INSERT INTO %s (category, datetime, headline, source_id, image, related, resource, summary, url) VALUES " +
			"(?, ?, ?, ?, ?, ?, ?, ?, ?)", newsTable)
		stmt, err := r.db.Prepare(createQuery)
		if err != nil {
			return err
		}
		if _, err := stmt.Exec(news.Category, news.Datetime, news.Headline, news.Id, news.Image, news.Related, news.Source, news.Summary, news.Url);
		err != nil {
			tx.Rollback()
			return err
		}
	}

	tx.Commit()

	return nil
}

func (r *NewsDB) GetNews(id int) (import_service.News, error) {
	var news import_service.News

	queryBook := fmt.Sprintf("SELECT * FROM %s WHERE id = ?", newsTable)
	row := r.db.QueryRow(queryBook, id)
	err := row.Scan(&news.Id, &news.Category, &news.Datetime, &news.Headline, &news.SourceId, &news.Image, &news.Related, &news.Resource, &news.Summary, &news.Url)

	return news, err
}
