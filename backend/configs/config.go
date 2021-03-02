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
		Host string
		Port int
		Pwd  string
		Db   int
	}
}
