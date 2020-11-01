package basketball

import (
	"github.com/google/wire"
)

var Set = wire.NewSet(
	ProvideResultsCase,
	ProvideShedulesCase,
	ProvideMatchesCase,
)
