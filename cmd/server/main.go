package main

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func main() {
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
