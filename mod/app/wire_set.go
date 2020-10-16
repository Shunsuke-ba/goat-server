package app

import (
	"github.com/Shunsuke-ba/goat-server/mod/app/domain"
	"github.com/Shunsuke-ba/goat-server/mod/app/handler"
	"github.com/Shunsuke-ba/goat-server/mod/app/usecase"
	"github.com/google/wire"
	"github.com/labstack/echo/v4"
)

type AllRepository struct {
	Health echo.HandlerFunc
}

func NewAllRepository(
	health handler.ProvideHealth,
) AllRepository {
	return AllRepository{
		Health: health.Echoer(),
	}
}

var Set = wire.NewSet(
	domain.Set,
	usecase.Set,
	handler.Set,
	NewAllRepository,
)
