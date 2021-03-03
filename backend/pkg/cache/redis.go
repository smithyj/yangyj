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
	prefix string
}

func (c *redisCache) buildKey(key string) string {
	if c.prefix == "" {
		return key
	}
	return fmt.Sprintf("%v:%v", c.prefix, key)
}

func (c *redisCache) Get(key string, obj interface{}) (err error) {
	result, err := c.client.Get(c.buildKey(key)).Result()
	if err == nil {
		err = json.Unmarshal([]byte(result), obj)
	}
	return
}

func (c *redisCache) Set(key string, value interface{}, expiration time.Duration) bool {
	v, err := json.Marshal(value)
	if err != nil {
		return false
	}
	if err := c.client.Set(c.buildKey(key), v, expiration).Err(); err != nil {
		return false
	}
	return true
}

func (c *redisCache) Del(keys ...string) bool {
	tmp := make([]string, 0, len(keys))
	for _, key := range keys {
		tmp = append(tmp, c.buildKey(key))
	}
	if _, err := c.client.Del(tmp...).Result(); err != nil {
		return false
	}
	return true
}

func newRedisCache() (cache *redisCache, err error) {
	redisCfg := config.Config.Redis.Cache
	cacheCfg := config.Config.Cache
	client, err := redis.New(&redisCfg)
	if err != nil {
		return
	}
	cache = &redisCache{
		client: client,
		prefix: cacheCfg.Prefix,
	}
	return
}
