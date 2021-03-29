package sys

import (
	"yangyj/pkg/cache"
	"yangyj/pkg/config"
	"yangyj/pkg/i18n"
	"yangyj/pkg/redis"
)

func init() {
	// 国际化
	i18n.InitLang()
	// 配置
	config.InitConfig()
	// Redis
	redis.InitRedis()
	// Cache
	cache.InitCache()
}