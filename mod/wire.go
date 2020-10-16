// +build wireinject

package main

import (
	"context"

	"github.com/Shunsuke-ba/goat-server/mod/app"
	"github.com/google/wire"
)

func initApplication(ctx context.Context) app.AllRepository {
	panic(wire.Build(
		app.Set,
	))
}
