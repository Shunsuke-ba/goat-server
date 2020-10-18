package sport_trader_domain

import (
	"time"
)

type GameSchedules struct {
	GeneratedAt string        `json:"ganerated_at"`
	Schema      string        `json:"schema"`
	SportEvents []SportEvents `json:"sport_events"`
}

type SportEvents struct {
	ID              string
	Scheduled       time.Time
	StartTimeTbd    bool
	Status          string
	TournamentRound TournamentRound
	Season          Season
	Tournament      Tournament
	Competitors     []Competitors
	Venue           Venue
}

// GlobalBasketballの中から日本バスケを取り出す(Bリーグ)
func (g GameSchedules) JapanValidation() (japanGames GameSchedules) {
	for _, game := range g.SportEvents {
		if game.Tournament.Category.CountryCode == "JPN" {
			japanGames.SportEvents = append(japanGames.SportEvents, game)
		}
	}
	return
}
