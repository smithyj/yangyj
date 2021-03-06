package configs

import "embed"

//go:embed *
var FS embed.FS

type Config struct {
	Mode    string
	Port    int
	Lang    string
	Db      DbConfig
	Redis   map[string]RedisConfig
	Cache   CacheConfig
	Captcha CaptchaConfig
	Email   []EmailConfig
	Sms     SmsConfig
}

type DbConfig struct {
	DSN string
}

type RedisConfig struct {
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
	Expired    int // 分钟
	Prefix     string
	SkipVerify bool `yaml:"skip_verify"`
}

type EmailConfig struct {
	Host     string
	Port     int
	Name     string
	Username string
	Password string
}

type SmsConfig struct {
	Kind     string
	Platform struct {
		Smsbao struct {
			Username string
			Password string
		}
		Aliyun struct {
			Appid     string
			Appsecret string
		}
	}
	Template struct {
		CaptchaCode map[string]SmsTemplateConfig `yaml:"captcha_code"`
	}
}

type SmsTemplateConfig struct {
	Zh string
	En string
}
