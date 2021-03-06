package configs

import "embed"

//go:embed *.yaml
var FS embed.FS

type Config struct {
	Mode    string
	Port    int
	Db      DbConfig
	Redis   RedisConfig
	Cache   CacheConfig
	Captcha CaptchaConfig
	Email   []EmailConfig
	Sms     SmsConfig
}

type DbConfig struct {
	DSN string
}

type RedisConfig struct {
	Default RedisItemConfig
	Cache   RedisItemConfig
}

type RedisItemConfig struct {
	Host   string
	Port   int
	Pwd    string
	Db     int
	Prefix string
}

type CacheConfig struct {
	Kind   string
	Prefix string
}

type CaptchaConfig struct {
	Expired int // 分钟
	Prefix  string
}

type EmailConfig struct {
	Host     string
	Port     int
	Username string
	Password string
}

type SmsConfig struct {
	Kind string
	Smsbao  struct {
		Username string
		Password string
	}
	Aliyun struct {
		Appid     string
		Appsecret string
	}
}
