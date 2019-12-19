package env

var Enviroment *EnviromentSetting

type EnviromentSetting struct {
	Server   Server   `mapstructure:"server"`
	Database Database `mapstructure:"database"`
	Mode     string   `mapstructure:"mode"`
	Redis    Redis    `mapstructure:"redis"`
	Cors     Cors     `mapstructure:"cors"`
	Jwt      Jwt      `mapstructure:"jwt"`
	Ag       Ag       `mapstructure:"ag"`
}

type Ag struct {
	Actype string `mapstructure:"actype"`
}

type Server struct {
	Port string `mapstructure:"port"`
}

type Database struct {
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	Host         string `mapstructure:"host"`
	Port         string `mapstructure:"port"`
	Name         string `mapstructure:"name"`
	Parameter    string `mapstructure:"parameter"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	MaxLifetime  int    `mapstructure:"max_lifetime"`
}

type Redis struct {
	Url      string `mapstructure:"url"`
	Password string `mapstructure:"password"`
	Database int    `mapstructure:"database"`
}

type Cors struct {
	Allow CorsAllow `mapstructure:"allow"`
}

type CorsAllow struct {
	Origins []string `mapstructure:"origins"`
	Headers []string `mapstructure:"headers"`
}

type Jwt struct {
	Secret  string `mapstructure:"secret"`
	Expired int    `mapstructure:"expired"`
	Refresh int    `mapstructure:"refresh"`
}
