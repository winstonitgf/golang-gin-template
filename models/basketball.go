package models

type BasketballTeam struct {
	Score int    `json:"score" redis:"score"`
	Name  string `json:"name" redis:"name"`
}
