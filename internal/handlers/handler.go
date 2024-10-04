package handlers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/sarthak0714/fampay-assignment/internal/services"
	"github.com/sarthak0714/fampay-assignment/internal/store"
)

type APIHandler struct {
	YouTubeVideoService services.YouTubeVideoService
	Store               store.Store
}

func NewAPIHandler(ytService services.YouTubeVideoService, store store.Store) *APIHandler {
	return &APIHandler{
		YouTubeVideoService: ytService,
		Store:               store,
	}
}

func (h *APIHandler) GetVideos(c echo.Context) error {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	if page < 1 {
		page = 1
	}

	size, _ := strconv.Atoi(c.QueryParam("page_size"))
	if size < 1 {
		size = 20
	}
	offset := (page - 1) * size

	videos, err := h.Store.GetVideos(size, offset)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch videos"})
	}

	return c.JSON(http.StatusOK, videos)
}
