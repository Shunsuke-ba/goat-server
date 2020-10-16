package handler

import (
	"net/http"

	"github.com/Shunsuke-ba/goat-server/mod/app/usecase"
	"github.com/labstack/echo/v4"
)

type (
	ProvideHealth echo.HandlerFunc
)

func (p ProvideHealth) Echoer() echo.HandlerFunc {
	return echo.HandlerFunc(p)
}

func ProvideHealthHandler(uc usecase.Health) ProvideHealth {
	return func(c echo.Context) error {
		message := uc(c.Request().Context())
		return c.JSON(http.StatusOK, message)
	}
}
