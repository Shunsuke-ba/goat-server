package sport_trader_handler

import (
	"net/http"

	"github.com/labstack/echo/v4"

	usecase "github.com/Shunsuke-ba/goat-server/mod/app/usecase/sport_trader_case"
)

type (
	ProvideGameSchedules echo.HandlerFunc
)

func (p ProvideGameSchedules) Echoer() echo.HandlerFunc {
	return echo.HandlerFunc(p)
}

func ProvideGameSchedulesHandler(uc usecase.GameSchedules) ProvideGameSchedules {
	return func(c echo.Context) error {
		gameSchedules, err := uc(c.Request().Context())
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, gameSchedules)

	}
}
