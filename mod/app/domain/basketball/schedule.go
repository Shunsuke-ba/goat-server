package basketball

import (
	"time"
)

type BasketballSchedules struct {
	Schema      string        `json:"schema"`
	GeneratedAt time.Time     `json:"generated_at"`
	SportEvents []SportEvents `json:"sport_events"`
}

type SportEvents struct {
	ID              string          `json:"ganerated_at"`
	Scheduled       time.Time       `json:"scheduled"`
	StartTimeTbd    bool            `json:"start_time_tbd"`
	Status          string          `json:"status"`
	TournamentRound TournamentRound `json:"tournament_round"`
	Season          Season          `json:"season"`
	Tournament      Tournament      `json:"tournament"`
	Competitors     []Competitors   `json:"competitors"`
	Venue           Venue           `json:"venue"`
}

// GlobalBasketballの中から日本バスケを取り出す(Bリーグ)
func (g BasketballSchedules) JapanValidation() (japanGames BasketballSchedules) {
	for _, game := range g.SportEvents {
		if game.Tournament.Category.CountryCode == "JPN" {
			japanGames.SportEvents = append(japanGames.SportEvents, game)
		}
	}
	return
}
