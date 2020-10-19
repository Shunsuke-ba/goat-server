package usecase

import (
	st "github.com/Shunsuke-ba/goat-server/mod/app/usecase/sport_trader_case"
	"github.com/google/wire"
)

var Set = wire.NewSet(
	ProvideHealthCase,
	st.ProvideGameResultsCase,
	st.ProvideGameShedulesCase,
	st.ProvideHeadToHeadsCase,
)
