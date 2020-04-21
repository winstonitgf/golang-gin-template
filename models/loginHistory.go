package models

type LoginHistory struct {
	Id        uint   `json:"id" form:"id"`
	UserId    uint   `json:"user_id" form:"user_id"`
	Ip        string `json:"ip" form:"ip"`
	CreatedAt uint   `json:"created_at"`
}
