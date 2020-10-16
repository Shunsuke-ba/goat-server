package main

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	c := context.Background()
	repo := initApplication(c)

	e.GET("/", repo.Health)
	e.GET("/game", repo.GameList)

	e.Logger.Fatal(e.Start(":8080"))
}

func healthCheck(c echo.Context) error {
	c.JSON(http.StatusOK, "OK")
	return nil
}
