package structs

type EnviromentModel struct {
	Database database
	Server   server
	Redis    redisModel
	Cors     corsModel
}

type server struct {
	Port string
	Mode string
}

type corsModel struct {
	Allow           corsAllowModel
	DefaultAllowUrl []string
}

type corsAllowModel struct {
	Headers []string
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
	LogEnable   int
}

type redisModel struct {
	Url      string
	Password string
	Database int
	Idle     int
	Active   int
	Protocol string
	Expire   int
	Unit     int
}
