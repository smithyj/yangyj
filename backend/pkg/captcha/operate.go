package captcha

import (
	"fmt"
	"strings"
	"time"
	"yangyj/backend/pkg/config"
	"yangyj/backend/pkg/redis"
)

type operate struct {
	expired int
	prefix  string
}

func (v *operate) buildKey(key string) string {
	if v.prefix == "" {
		return key
	}
	return fmt.Sprintf("%v:%v", v.prefix, key)
}

func (v *operate) create(id, code string) (err error) {
	expiration := 10 * 60 * time.Second
	err = redis.Redis.Set(v.buildKey(id), code, expiration).Err()
	return
}

func (v *operate) del(id string) (num int64, result bool) {
	num, err := redis.Redis.Del(v.buildKey(id)).Result()
	if err != nil {
		return num, false
	}
	return num, true
}

func (v *operate) verify(id, code string) bool {
	value, err := redis.Redis.Get(v.buildKey(id)).Result()
	if err != nil {
		return false
	}
	if strings.ToLower(value) == strings.ToLower(code) {
		// 验证成功，删除验证码
		if _, ok := v.del(id); !ok {
			return false
		}
		return true
	}
	return false
}

func newOperate() *operate {
	cfg := config.Config.Captcha
	return &operate{
		expired: cfg.Expired,
		prefix:  cfg.Prefix,
	}
}
