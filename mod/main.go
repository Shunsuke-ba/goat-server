package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	e.GET("/", HelloWorld)

	e.Logger.Fatal(e.Start(":8080"))
}

func HelloWorld(c echo.Context) error {
	fmt.Println("Hello World")
	return nil
}
