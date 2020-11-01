package domain

import (
	"github.com/Shunsuke-ba/goat-server/mod/app/domain/basketball"
	"github.com/Shunsuke-ba/goat-server/mod/app/domain/soccer"
	"github.com/google/wire"
)

var Set = wire.NewSet(
	// Health
	wire.Struct(new(Health)),

	// Basketball
	basketball.Set,
	soccer.Set,
)
