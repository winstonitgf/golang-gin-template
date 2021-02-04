package global

import (
	"golang-startup/structs"

	"github.com/casbin/casbin/v2"
	"github.com/gomodule/redigo/redis"
	"gorm.io/gorm"
)

var (
	EnvConfig      *structs.EnviromentModel
	Mysql          *gorm.DB
	Redis          *redis.Pool
	CasbinEnforcer *casbin.Enforcer
)
