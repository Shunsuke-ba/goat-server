package basketball

import "time"

type BasketballMatches struct {
	Schema       string       `json:"schema"`
	GeneratedAt  time.Time    `json:"generated_at"`
	Teams        []Teams      `json:"teams"`
	LastMeetings LastMeetings `json:"last_meetings"`
}

type Teams struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Country      string `json:"country"`
	CountryCode  string `json:"country_code"`
	Abbreviation string `json:"abbreviation"`
}

type LastMeetings struct {
	Results []Results `json:"results"`
}
