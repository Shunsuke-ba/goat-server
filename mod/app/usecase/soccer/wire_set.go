package soccer

import (
	"github.com/google/wire"
)

var Set = wire.NewSet(
	ProvideResultsCase,
	//ProvideGameShedulesCase,
	//ProvideHeadToHeadsCase,
)
