package handler

import (
	st "github.com/Shunsuke-ba/goat-server/mod/app/handler/sport_trader_handler"
	"github.com/google/wire"
)

var Set = wire.NewSet(
	ProvideHealthHandler,
	st.ProvideGameResultsHandler,
	st.ProvideGameSchedulesHandler,
	st.ProvideHeadToHeadsHandler,
)
