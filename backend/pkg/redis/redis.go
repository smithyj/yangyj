package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
	"yangyj/backend/pkg/config"
)

const (
	PREFIX = "yangyj:%s"
)

var Redis *Client

type Client struct {
	ctx    context.Context
	client *redis.Client
}

func (c *Client) Get(key string) *redis.StringCmd {
	key = fmt.Sprintf(PREFIX, key)
	return c.client.Get(c.ctx, key)
}

func (c *Client) Set(key string, value interface{}, expiration time.Duration) *redis.StatusCmd {
	key = fmt.Sprintf(PREFIX, key)
	return c.client.Set(c.ctx, key, value, expiration)
}

func (c *Client) Del(keys ...string) *redis.IntCmd {
	tmp := make([]string, 0, len(keys))
	for _, v := range keys {
		tmp = append(tmp, fmt.Sprintf(PREFIX, v))
	}
	return c.client.Del(c.ctx, tmp...)
}

func init() {
	cfg := config.Config.Redis.Default
	client, err := New(
		cfg.Host,
		cfg.Port,
		cfg.Pwd,
		cfg.Db,
	)
	if err != nil {
		panic(err)
	}
	Redis = client
}

func New(Host string, Port int, Pwd string, Db int) (client *Client, err error) {
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%v", Host, Port),
		Password: Pwd,
		DB:       Db,
	})
	if err = rdb.Ping(ctx).Err(); err != nil {
		return
	}
	client = &Client{
		ctx:    ctx,
		client: rdb,
	}
	return
}
