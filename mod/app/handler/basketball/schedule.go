package basketball

import (
	"net/http"

	"github.com/labstack/echo/v4"

	usecase "github.com/Shunsuke-ba/goat-server/mod/app/usecase/basketball"
)

type (
	ProvideSchedules echo.HandlerFunc
)

func (p ProvideSchedules) Echoer() echo.HandlerFunc {
	return echo.HandlerFunc(p)
}

func ProvideSchedulesHandler(uc usecase.Schedules) ProvideSchedules {
	return func(c echo.Context) error {
		schedules, err := uc(c.Request().Context())
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, schedules)

	}
}
