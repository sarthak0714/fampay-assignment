package handlers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/sarthak0714/fampay-assignment/frontend/templates"
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

// Handles the landing page request and renders the landing template.
func (h *APIHandler) LandingHandler(c echo.Context) error {
	return templates.Landing().Render(c.Request().Context(), c.Response().Writer)
}

// Handles the video page request, retrieves videos from the store, and renders the video list.
func (h *APIHandler) VideoHandler(c echo.Context) error {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	if page < 1 {
		page = 1
	}

	size, _ := strconv.Atoi(c.QueryParam("size"))
	if size < 1 {
		size = 9
	}
	offset := (page - 1) * size

	videos, err := h.Store.GetVideos(size, offset)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch videos"})
	}

	if c.Request().Header.Get("HX-Request") == "true" {
		return templates.VideoList(videos, page).Render(c.Request().Context(), c.Response().Writer)
	}

	return templates.VideoPage(videos, page).Render(c.Request().Context(), c.Response().Writer)
}

// Handles the API request to fetch videos, returns RAW JSON.
func (h *APIHandler) GetVideosAPI(c echo.Context) error {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	if page < 1 {
		page = 1
	}

	size, _ := strconv.Atoi(c.QueryParam("size"))
	if size < 1 {
		size = 9
	}
	offset := (page - 1) * size

	videos, err := h.Store.GetVideos(size, offset)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch videos"})
	}

	return c.JSON(http.StatusOK, videos)
}
