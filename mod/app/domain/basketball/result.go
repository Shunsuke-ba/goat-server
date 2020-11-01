package basketball

import "time"

type (
	BasketballResults struct {
		Schema      string    `json:"schema"`
		GeneratedAt time.Time `json:"generated_at"`
		Results     []Results `json:"results"`
	}

	Results struct {
		SportEvent       SportEvent       `json:"sport_event"`
		SportEventStatus SportEventStatus `json:"sport_event_status"`
	}

	SportEvent struct {
		ID              string
		Schedule        time.Time
		StartTimeTbd    bool
		TournamentRound TournamentRound
		Season          Season
		Tournament      Tournament
		Competitors     []Competitors
		Venue           Venue
	}

	TournamentRound struct {
		Type                string `json:"type"`
		Name                string `json:"quarterfinal"`
		CupRoundMatchNumber int    `json:"cup_round_match_number"`
		CupRoundMatches     int    `json:"cup_round_matches"`
	}

	Season struct {
		ID           string `json:"id"`
		Name         string `json:"name"`
		StartDate    string `json:"start_date"`
		EndDate      string `json:"end_date"`
		Year         string `json:"year"`
		TournamentID string `json:"tournament_id"`
		Tournament   string `json:"tournament"`
	}

	Tournament struct {
		ID       string `json:"id"`
		Name     string `json:"name"`
		Sport    Sport
		Category Category
	}

	Sport struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	}

	Category struct {
		ID          string `json:"id"`
		Name        string `json:"name"`
		CountryCode string `json:"country_code"`
	}

	Competitors struct {
		ID           string `json:"id"`
		Name         string `json:"name"`
		Country      string `json:"country"`
		CountryCode  string `json:"country_code"`
		Abbreviation string `json:"abbreviation"`
		Qualifier    string `json:"qualifier"`
	}

	Venue struct {
		ID   string `json:"id"`
		Name string `json:"name"`
		//Capacity string `json:"capacity"`
		CityName       string `json:"city_name"`
		CountryName    string `json:"country_name"`
		MapCoordinates string `json:"map_coordinates"`
		CountryCode    string `json:"country_code"`
	}

	SportEventStatus struct {
		Status       string         `json:"status"`
		MatchStatus  string         `json:"match_status"`
		HomeScore    int            `json:"home_score"`
		AwayScore    int            `json:"away_score"`
		WinnerID     string         `json:"winner_id"`
		PeriodScores []PeriodScores `json:"period_scores"`
	}

	PeriodScores struct {
		HomeScore int    `json:"home_score"`
		AwayScore int    `json:"away_score"`
		Type      string `json:"type"`
		Number    int    `json:"number"`
	}
)

// GlobalBasketballの中から日本バスケを取り出す(Bリーグ)
func (g BasketballResults) JapanValidation() (japanGames BasketballResults) {
	for _, game := range g.Results {
		if game.SportEvent.Tournament.Category.CountryCode == "JPN" {
			japanGames.Results = append(japanGames.Results, game)
		}
	}
	return
}
