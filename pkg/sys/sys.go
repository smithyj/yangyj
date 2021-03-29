package sys

import (
	"yangyj/pkg/cache"
	"yangyj/pkg/config"
	"yangyj/pkg/redis"
)

func init() {
	// 配置
	config.InitConfig()
	// Redis
	redis.InitRedis()
	// Cache
	cache.InitCache()
}