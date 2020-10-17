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
	Health      echo.HandlerFunc
	GameResults echo.HandlerFunc
}

func NewAllRepository(
	health handler.ProvideHealth,
	gameResults st.ProvideGameResults,
) AllRepository {
	return AllRepository{
		Health:      health.Echoer(),
		GameResults: gameResults.Echoer(),
	}
}

var Set = wire.NewSet(
	domain.Set,
	usecase.Set,
	handler.Set,
	NewAllRepository,
)
