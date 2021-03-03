package cache

import (
	"errors"
	"time"
	"yangyj/backend/pkg/config"
)

type cache interface {
	Get(key string, obj interface{}) (err error)
	Set(key string, value interface{}, expiration time.Duration) (result bool)
	Del(keys ...string) (result bool)
}

var Cache cache

func init() {
	cfg := config.Config.Cache
	switch cfg.Type {
	case "redis":
		c, err := newRedisCache()
		if err != nil {
			panic(err)
		}
		Cache = c
	default:
		panic(errors.New("未知的缓存类型"))
	}
}
