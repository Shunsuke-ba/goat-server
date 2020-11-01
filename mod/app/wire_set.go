package app

import (
	"github.com/Shunsuke-ba/goat-server/mod/app/domain"
	"github.com/Shunsuke-ba/goat-server/mod/app/handler"
	"github.com/Shunsuke-ba/goat-server/mod/app/handler/basketball"
	"github.com/Shunsuke-ba/goat-server/mod/app/handler/soccer"
	"github.com/Shunsuke-ba/goat-server/mod/app/usecase"
	"github.com/google/wire"
	"github.com/labstack/echo/v4"
)

type AllRepository struct {
	Health              echo.HandlerFunc
	BasketballResults   echo.HandlerFunc
	BasketballSchedules echo.HandlerFunc
	BasketballMatches   echo.HandlerFunc
	SoccerResults       echo.HandlerFunc
}

func NewAllRepository(
	health handler.ProvideHealth,
	basketballResults basketball.ProvideResults,
	basketballSchedules basketball.ProvideSchedules,
	basketballMatches basketball.ProvideMatches,
	soccerResults soccer.ProvideResults,
) AllRepository {
	return AllRepository{
		Health:              health.Echoer(),
		BasketballResults:   basketballResults.Echoer(),
		BasketballSchedules: basketballSchedules.Echoer(),
		BasketballMatches:   basketballMatches.Echoer(),
		SoccerResults:       soccerResults.Echoer(),
	}
}

var Set = wire.NewSet(
	domain.Set,
	usecase.Set,
	handler.Set,
	NewAllRepository,
)
