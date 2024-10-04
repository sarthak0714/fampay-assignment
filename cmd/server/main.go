package main

import (
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/sarthak0714/fampay-assignment/internal/store"
)

func main() {

	_, err := store.NewPostgresStore("test.db")
	if err != nil {
		log.Fatal(err)
	}

	e := echo.New()
	e.HideBanner = true

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"message":   "works",
			"timestamp": time.Now().Format("2006-01-02 15:04:05"),
		})
	})

	e.Logger.Fatal(e.Start(":8080"))
}
