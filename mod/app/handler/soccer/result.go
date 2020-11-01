package soccer

import (
	"net/http"

	"github.com/labstack/echo/v4"

	usecase "github.com/Shunsuke-ba/goat-server/mod/app/usecase/soccer"
)

type (
	ProvideResults echo.HandlerFunc
)

func (p ProvideResults) Echoer() echo.HandlerFunc {
	return echo.HandlerFunc(p)
}

func ProvideResultsHandler(uc usecase.Results) ProvideResults {
	return func(c echo.Context) error {
		results, err := uc(c.Request().Context())
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, results)
	}
}
