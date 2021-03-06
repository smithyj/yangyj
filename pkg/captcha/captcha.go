package captcha

import (
	"fmt"
	"strings"
	"time"
	"yangyj/pkg/config"
	"yangyj/pkg/redis"
)

type captcha struct {
	expired    int
	prefix     string
	skipVerify bool
}

func (v *captcha) buildKey(key string) string {
	if v.prefix == "" {
		return key
	}
	return fmt.Sprintf("%v:%v", v.prefix, key)
}

func (v *captcha) create(id, code string) (err error) {
	expiration := 10 * 60 * time.Second
	err = redis.Redis.Set(v.buildKey(id), code, expiration).Err()
	return
}

func (v *captcha) del(id string) (num int64, ok bool) {
	var err error
	if num, err = redis.Redis.Del(v.buildKey(id)).Result(); err != nil {
		return
	}
	ok = true
	return
}

func (v *captcha) verify(id, code string) (ok bool) {
	var err error
	var value string
	if v.skipVerify {
		return true
	}
	if value, err = redis.Redis.Get(v.buildKey(id)).Result(); err != nil {
		return
	}
	if strings.ToLower(value) == strings.ToLower(code) {
		// 验证成功，删除验证码
		_, ok = v.del(id)
		return
	}
	return
}

func newCaptcha() *captcha {
	cfg := config.Config.Captcha
	return &captcha{
		expired:    cfg.Expired,
		prefix:     cfg.Prefix,
		skipVerify: cfg.SkipVerify,
	}
}
