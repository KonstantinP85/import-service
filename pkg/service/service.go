package service

import (
	import_service "github.com/KonstantinP85/import-service"
	"github.com/KonstantinP85/import-service/pkg/repository"
)

type News interface {
	Upload() error
	GetNews(id int) (import_service.News, error)
}

type Service struct {
	News
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		News: NewNewsService(repo.News),
	}
}