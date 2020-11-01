package basketball

import (
	"github.com/google/wire"
)

var Set = wire.NewSet(
	ProvideResultsHandler,
	ProvideSchedulesHandler,
	ProvideMatchesHandler,
)
