package soccer

import (
	"github.com/google/wire"
)

var Set = wire.NewSet(
	ProvideResultsHandler,
	//ProvideGameSchedulesHandler,
	//ProvideHeadToHeadsHandler,
)
