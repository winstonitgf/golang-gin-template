package app

import (
	"bao-bet365-api/database"
	"bao-bet365-api/redis"
)

type AppInit struct{}

func (AppInit) Init() {
	// 開啟DB連線
	database.Open()

	// 開啟Redis連線
	redis.Init()
}
