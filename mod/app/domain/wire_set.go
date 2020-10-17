package domain

import (
	"github.com/google/wire"

	st "github.com/Shunsuke-ba/goat-server/mod/app/domain/sport_trader_domain"
)

var Set = wire.NewSet(
	wire.Struct(new(Health)),
	wire.Struct(new(st.GameResults)),
	wire.Struct(new(st.Results)),
	wire.Struct(new(st.SportEvent)),
	wire.Struct(new(st.TournamentRound)),
	wire.Struct(new(st.Season)),
	wire.Struct(new(st.Tournament)),
	wire.Struct(new(st.Sport)),
	wire.Struct(new(st.Category)),
	wire.Struct(new(st.Competitors)),
	wire.Struct(new(st.Venue)),
	wire.Struct(new(st.SportEventStatus)),
	wire.Struct(new(st.PeriodScores)),
)
