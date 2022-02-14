package models

type Leaderboard struct {
	Email  string `json:"email"`
	Points int64  `json:"points"`
	Rank   int64  `json:"rank"`
}
