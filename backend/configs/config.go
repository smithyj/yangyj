package configs

import "embed"

//go:embed *.yaml
var FS embed.FS

type Config struct {
	Mode string
	Port int
	Db   struct {
		DSN string
	}
	Redis struct {
		Default RedisConfig
		Cache   RedisConfig
	}
	Cache struct {
		Type string
		Prefix  string
	}
	Captcha struct {
		Expired int // 分钟
		Prefix string
	} `yaml:"captcha"`
}

type RedisConfig struct {
	Host   string
	Port   int
	Pwd    string
	Db     int
	Prefix string
}
