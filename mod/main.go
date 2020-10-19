package main

import (
	"context"

	"github.com/Shunsuke-ba/goat-server/mod/util"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// echo
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	// env
	util.CheckEnv(e)

	// wire
	c := context.Background()
	repo := initApplication(c)

	// routing
	e.GET("/", repo.Health)
	e.GET("/game/results", repo.GameResults)
	e.GET("/game/schedules", repo.GameSchedules)

	// server start
	e.Logger.Fatal(e.Start(":8080"))
}
