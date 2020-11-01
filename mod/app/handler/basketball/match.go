package basketball

import (
	"net/http"

	"github.com/labstack/echo/v4"

	usecase "github.com/Shunsuke-ba/goat-server/mod/app/usecase/basketball"
)

type (
	ProvideMatches echo.HandlerFunc

	MatchesParam struct {
		Team1ID string `json:"team1_id"`
		Team2ID string `json:"team2_id"`
	}
)

func (p ProvideMatches) Echoer() echo.HandlerFunc {
	return echo.HandlerFunc(p)
}

func ProvideMatchesHandler(uc usecase.Matches) ProvideMatches {
	return func(c echo.Context) error {
		param := MatchesParam{}
		if err := c.Bind(&param); err != nil {
			return err
		}
		matches, err := uc(c.Request().Context(), param.Team1ID, param.Team2ID)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, matches)
	}
}
