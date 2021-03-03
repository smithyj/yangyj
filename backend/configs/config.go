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
		Default redisConfig
		Cache   redisConfig
	}
	Cache struct {
		Type string
		// 过期时间，分钟
		Expired int64
	}
}

type redisConfig struct {
	Host string
	Port int
	Pwd  string
	Db   int
}
