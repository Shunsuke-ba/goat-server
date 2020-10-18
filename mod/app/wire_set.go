package app

import (
	"github.com/Shunsuke-ba/goat-server/mod/app/domain"
	"github.com/Shunsuke-ba/goat-server/mod/app/handler"
	st "github.com/Shunsuke-ba/goat-server/mod/app/handler/sport_trader_handler"
	"github.com/Shunsuke-ba/goat-server/mod/app/usecase"
	"github.com/google/wire"
	"github.com/labstack/echo/v4"
)

type AllRepository struct {
	Health        echo.HandlerFunc
	GameResults   echo.HandlerFunc
	GameSchedules echo.HandlerFunc
}

func NewAllRepository(
	health handler.ProvideHealth,
	gameResults st.ProvideGameResults,
	gameSchedules st.ProvideGameSchedules,
) AllRepository {
	return AllRepository{
		Health:        health.Echoer(),
		GameResults:   gameResults.Echoer(),
		GameSchedules: gameSchedules.Echoer(),
	}
}

var Set = wire.NewSet(
	domain.Set,
	usecase.Set,
	handler.Set,
	NewAllRepository,
)
