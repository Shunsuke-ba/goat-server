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
	e.GET("/health", repo.Health)

	game := e.Group("/game")
	{
		japan := game.Group("/japan")
		{
			basketball := japan.Group("/basketball")
			{
				basketball.GET("/results", repo.BasketballResults)
				basketball.GET("/schedules", repo.BasketballSchedules)
				basketball.GET("/headtohead", repo.BasketballMatches)
			}
			soccer := japan.Group("/soccer")
			{
				soccer.GET("/results", repo.SoccerResults)
			}
		}
	}

	// server start
	e.Logger.Fatal(e.Start(":8080"))
}
