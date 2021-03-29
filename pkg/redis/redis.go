package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
	"yangyj/configs"
	"yangyj/pkg/config"
)

var Redis *Client

type Client struct {
	ctx    context.Context
	client *redis.Client
	prefix string
}

func (c *Client) buildKey(key string) string {
	if c.prefix == "" {
		return key
	}
	return fmt.Sprintf("%v:%v", c.prefix, key)
}

func (c *Client) Get(key string) *redis.StringCmd {
	return c.client.Get(c.ctx, c.buildKey(key))
}

func (c *Client) Set(key string, value interface{}, expiration time.Duration) *redis.StatusCmd {
	return c.client.Set(c.ctx, c.buildKey(key), value, expiration)
}

func (c *Client) Del(keys ...string) *redis.IntCmd {
	tmp := make([]string, 0, len(keys))
	for _, key := range keys {
		tmp = append(tmp, c.buildKey(key))
	}
	return c.client.Del(c.ctx, tmp...)
}

func InitRedis() {
	cfg := config.Config.Redis["default"]
	client, err := New(&cfg)
	if err != nil {
		panic(err)
	}
	Redis = client
}

func New(cfg *configs.RedisConfig) (client *Client, err error) {
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%v", cfg.Host, cfg.Port),
		Password: cfg.Pwd,
		DB:       cfg.Db,
	})
	if err = rdb.Ping(ctx).Err(); err != nil {
		return
	}
	client = &Client{
		ctx:    ctx,
		client: rdb,
		prefix: cfg.Prefix,
	}
	return
}
