package structs

type EnviromentModel struct {
	Server   serverModel
	Mode     string
	Redis    redisModel
	Cors     corsModel
	Jwt      jwt
	Database database
}

type serverModel struct {
	Port string
}

type redisModel struct {
	Url      string
	Password string
	Database int
	Idle     int
	Active   int
	Protocol string
}

type corsModel struct {
	Allow           corsAllowModel
	DefaultAllowUrl string
}

type corsAllowModel struct {
	Headers []string
}

type jwt struct {
	Secret string
}

type database struct {
	Client      string
	MaxIdle     uint
	MaxLifeTime string
	MaxOpenConn uint
	User        string
	Password    string
	Host        string
	Db          string
	Params      string
	Port        string
}
