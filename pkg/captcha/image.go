package captcha

import (
	"fmt"
	acaptcha "github.com/afocus/captcha"
	"image/color"
	"yangyj/assets"
)

type imageCaptcha struct {
	captcha *captcha
}

func (c *imageCaptcha) buildKey(key string) string {
	return fmt.Sprintf("image:%v", key)
}

// 创建图片验证码
//
// id: 验证码ID
//
// w: 宽度
//
// h: 高度
func (c *imageCaptcha) Create(id string, w, h int) (img *acaptcha.Image, err error) {
	var byteSlice []byte
	filename := "font/comic/comic.ttf"
	if byteSlice, err = assets.FS.ReadFile(filename); err != nil {
		return
	}
	draw := acaptcha.New()
	_ = draw.AddFontFromBytes(byteSlice)
	// 设置验证码大小
	draw.SetSize(w, h)
	// 设置干扰强度
	draw.SetDisturbance(acaptcha.HIGH)
	// 设置前景色 可以多个 随机替换文字颜色 默认黑色
	draw.SetFrontColor(
		color.RGBA{A: 255},
		color.RGBA{R: 241, G: 60, B: 60, A: 255},
		color.RGBA{R: 60, G: 81, B: 241, A: 255},
	)

	img, code := draw.Create(6, acaptcha.CLEAR)
	err = c.captcha.create(c.buildKey(id), code)
	return
}

// 图片验证码校验
//
// id: 图片验证码ID
// code: 验证码
func (c *imageCaptcha) Verify(id, code string) (ok bool) {
	return c.captcha.verify(c.buildKey(id), code)
}

func NewImageCaptcha() *imageCaptcha {
	return &imageCaptcha{
		captcha: newCaptcha(),
	}
}

