package captcha

import (
	"fmt"
	"github.com/afocus/captcha"
	"image/color"
	"time"
	"yangyj/backend/assets"
	"yangyj/backend/pkg/redis"
)

const (
	PREFIX = "captcha:%s"
)

type Captcha struct {
	id string
}

func (c *Captcha) Image(w, h int) (img *captcha.Image, code string, err error) {
	filename := "font/comic/comic.ttf"
	bytes, err := assets.FS.ReadFile(filename)
	cc := captcha.New()
	if err == nil {
		_ = cc.AddFontFromBytes(bytes)
	}
	// 设置验证码大小
	cc.SetSize(w, h)
	// 设置干扰强度
	cc.SetDisturbance(captcha.HIGH)
	// 设置前景色 可以多个 随机替换文字颜色 默认黑色
	cc.SetFrontColor(
		color.RGBA{A: 255},
		color.RGBA{R: 241, G: 60, B: 60, A: 255},
		color.RGBA{R: 60, G: 81, B: 241, A: 255},
	)

	img, code = cc.Create(6, captcha.ALL)
	err = c.Create(code)
	return
}

func (c *Captcha) Create(code string) (err error) {
	err = redis.Redis.Set(fmt.Sprintf(PREFIX, c.id), code, 10*60*time.Second).Err()
	return
}

func (c *Captcha) Verify(id, code string) bool {
	value, err := redis.Redis.Get(fmt.Sprintf(PREFIX, id)).Result()
	if err != nil {
		return false
	}
	if value == code {
		return true
	}
	return false
}

func New(id string) *Captcha {
	return &Captcha{
		id: id,
	}
}
