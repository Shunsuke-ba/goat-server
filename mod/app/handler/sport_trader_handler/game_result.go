package sport_trader_handler

import (
	"net/http"

	"github.com/labstack/echo/v4"

	usecase "github.com/Shunsuke-ba/goat-server/mod/app/usecase/sport_trader_case"
)

type (
	ProvideGameResults echo.HandlerFunc
)

func (p ProvideGameResults) Echoer() echo.HandlerFunc {
	return echo.HandlerFunc(p)
}

func ProvideGameResultsHandler(uc usecase.GameResults) ProvideGameResults {
	return func(c echo.Context) error {
		gameResults, err := uc(c.Request().Context())
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}

		return c.JSON(http.StatusOK, gameResults)
	}
}
