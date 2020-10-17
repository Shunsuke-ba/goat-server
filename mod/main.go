package main

import (
	"context"

	"github.com/Shunsuke-ba/goat-server/mod/util"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	util.CheckEnv(e)

	c := context.Background()
	repo := initApplication(c)

	e.GET("/", repo.Health)
	e.GET("/gamelist", repo.GameResults)

	e.Logger.Fatal(e.Start(":8080"))
}
