package handler

import (
	"github.com/KonstantinP85/import-service/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api/v1")
	{
		news := api.Group("/news")
		{
			news.GET("/upload", h.upload)
			news.GET("/:id", h.getNews)
		}
	}

	return router
}