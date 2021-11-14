package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) upload(c * gin.Context) {

	err := h.services.News.Upload()
	if err != nil {
		return
	}
}

func (h *Handler) getNews(c *gin.Context)  {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"result": "Invalid id param",
		})
		return
	}
	news, err := h.services.News.GetNews(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"result": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, news)
}