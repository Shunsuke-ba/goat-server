package usecase

import (
	"github.com/Shunsuke-ba/goat-server/mod/app/usecase/basketball"
	"github.com/Shunsuke-ba/goat-server/mod/app/usecase/soccer"
	"github.com/google/wire"
)

var Set = wire.NewSet(
	ProvideHealthCase,
	basketball.Set,
	soccer.Set,
)
