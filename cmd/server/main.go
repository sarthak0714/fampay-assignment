package main

import (
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/sarthak0714/fampay-assignment/internal/config"
	"github.com/sarthak0714/fampay-assignment/internal/handlers"
	"github.com/sarthak0714/fampay-assignment/internal/services"
	"github.com/sarthak0714/fampay-assignment/internal/store"
	"github.com/sarthak0714/fampay-assignment/pkg/utils"
)

func main() {

	cfg := config.Load()
	store, err := store.NewPostgresStore(cfg.DbPath)
	if err != nil {
		log.Fatal(err)
	}
	if err := store.Init(); err != nil {
		log.Fatal(err)
	}
	ytAPIService, err := services.NewService(cfg.YouTubeAPIKeys, store)
	if err != nil {
		log.Fatal(err)
	}
	h := handlers.NewAPIHandler(ytAPIService, store)

	e := echo.New()
	e.HideBanner = true
	e.Use(utils.CustomLogger())

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"message":   "works",
			"timestamp": time.Now().Format("2006-01-02 15:04:05"),
		})
	})
	e.GET("/video", h.GetVideos)

	go h.YouTubeVideoService.FetchVideosWorker(cfg.SearchQuery, cfg.FetchInterval, store)

	e.Logger.Fatal(e.Start(":8080"))
}
