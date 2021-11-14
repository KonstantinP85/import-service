package service

import (
	"context"
	"github.com/Finnhub-Stock-API/finnhub-go"
	import_service "github.com/KonstantinP85/import-service"
	"github.com/KonstantinP85/import-service/pkg/repository"
	"os"
)

type NewsService struct {
	repo repository.News
}

func NewNewsService(repo repository.News) *NewsService {
	return &NewsService{repo: repo}
}

func (s *NewsService) Upload() error {

	cfg := finnhub.NewConfiguration()
	cfg.AddDefaultHeader("X-Finnhub-Token", os.Getenv("FINNHUB_TOKEN"))
	finnhubClient := finnhub.NewAPIClient(cfg).DefaultApi

	generalNews, _, err := finnhubClient.GeneralNews(context.Background(), "general", nil)
	if err != nil {
		return err
	}

	err = s.repo.Upload(generalNews)
	if err != nil {
		return err
	}

	return nil
}

func (s *NewsService) GetNews(id int) (import_service.News, error) {
	return s.repo.GetNews(id)
}


