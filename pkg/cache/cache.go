package cache

import (
	"errors"
	"time"
	"yangyj/pkg/config"
)

type cache interface {
	Get(key string, obj interface{}) (err error)
	Set(key string, value interface{}, expiration time.Duration) (ok bool)
	Del(keys ...string) (ok bool)
}

var Cache cache

func init() {
	var err error
	var c *redisCache
	cfg := config.Config.Cache
	switch cfg.Kind {
	case "redis":
		if c, err = newRedisCache(); err != nil {
			panic(err)
		}
		Cache = c
	default:
		panic(errors.New("未知的缓存类型"))
	}
}
