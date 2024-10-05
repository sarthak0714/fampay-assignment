package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/sarthak0714/fampay-assignment/internal/config"
	"github.com/sarthak0714/fampay-assignment/internal/handlers"
	"github.com/sarthak0714/fampay-assignment/internal/services"
	"github.com/sarthak0714/fampay-assignment/internal/store"
	"github.com/sarthak0714/fampay-assignment/pkg/utils"
)

func main() {

	// load config
	cfg := config.Load()
	// connect to db
	store, err := store.NewPostgresStore(cfg.DbPath)
	if err != nil {
		log.Fatal(err)
	}
	// init db if not exists or migrate
	if err := store.Init(); err != nil {
		log.Fatal(err)
	}
	// create new youtube api service
	ytAPIService, err := services.NewService(cfg.YouTubeAPIKeys, store)
	if err != nil {
		log.Fatal(err)
	}
	// create new handler
	h := handlers.NewAPIHandler(ytAPIService, store)

	e := echo.New()
	e.HideBanner = true
	e.Use(utils.CustomLogger())

	e.Static("/static", "web/static")

	e.GET("/", h.LandingHandler)
	e.GET("/video", h.VideoHandler)
	e.GET("/api/video", h.GetVideosAPI)

	// background worker
	go h.YouTubeVideoService.FetchVideosWorker(cfg.SearchQuery, cfg.FetchInterval, store)

	e.Logger.Fatal(e.Start(":8080"))
}
