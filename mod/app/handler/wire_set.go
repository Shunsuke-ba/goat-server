package handler

import (
	"github.com/Shunsuke-ba/goat-server/mod/app/handler/basketball"
	"github.com/Shunsuke-ba/goat-server/mod/app/handler/soccer"
	"github.com/google/wire"
)

var Set = wire.NewSet(
	ProvideHealthHandler,
	basketball.Set,
	soccer.Set,
)
