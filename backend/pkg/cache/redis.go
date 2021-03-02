package cache

import (
	"encoding/json"
	"fmt"
	"time"
	"yangyj/backend/pkg/config"
	"yangyj/backend/pkg/redis"
)

type redisCache struct {
	client *redis.Client
}

func (c *redisCache) Get(key string, obj interface{}) (err error) {
	key = fmt.Sprintf(PREFIX, key)
	result, err := c.client.Get(key).Result()
	if err == nil {
		err = json.Unmarshal([]byte(result), obj)
	}
	return
}

func (c *redisCache) Set(key string, value interface{}, expiration time.Duration) bool {
	var err error
	key = fmt.Sprintf(PREFIX, key)
	v, err := json.Marshal(value)
	if err != nil {
		return false
	}
	if err = c.client.Set(key, v, expiration).Err(); err != nil {
		return false
	}
	return true
}

func newRedisCache() (cache *redisCache, err error) {
	cfg := config.Config.Redis.Cache
	client, err := redis.New(
		cfg.Host,
		cfg.Port,
		cfg.Pwd,
		cfg.Db,
	)
	if err != nil {
		return
	}
	cache = &redisCache{
		client: client,
	}
	return
}
