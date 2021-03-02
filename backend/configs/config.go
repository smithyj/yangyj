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
		// 过期时间，分钟
		Expired int64
	}
}

type RedisConfig struct {
	Host string
	Port int
	Pwd  string
	Db   int
}
