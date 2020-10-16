package domain

import "github.com/google/wire"

var Set = wire.NewSet(
	wire.Struct(new(HealthRepository)),
)
