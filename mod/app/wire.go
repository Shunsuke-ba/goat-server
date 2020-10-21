// +build wireinject

package main

import (
	"context"

	"github.com/Shunsuke-ba/goat-server/mod/app/domain"
	"github.com/Shunsuke-ba/goat-server/mod/app/handler"
	"github.com/Shunsuke-ba/goat-server/mod/app/usecase"
	"github.com/google/wire"
)

func initApplication(ctx context.Context) AllRepository {
	panic(wire.Build(
		domain.Set,
		usecase.Set,
		handler.Set,
		NewAllRepository,
	))
}
