package captcha

import (
	"fmt"
	"github.com/afocus/captcha"
	"image/color"
	"yangyj/backend/assets"
)

type imageCode struct {
	op *operate
}

func (c *imageCode) buildKey(key string) string {
	return fmt.Sprintf("image:%v", key)
}

func (c *imageCode) Create(id string, w, h int) (img *captcha.Image, code string, err error) {
	filename := "font/comic/comic.ttf"
	bytes, err := assets.FS.ReadFile(filename)
	draw := captcha.New()
	if err == nil {
		_ = draw.AddFontFromBytes(bytes)
	}
	// 设置验证码大小
	draw.SetSize(w, h)
	// 设置干扰强度
	draw.SetDisturbance(captcha.HIGH)
	// 设置前景色 可以多个 随机替换文字颜色 默认黑色
	draw.SetFrontColor(
		color.RGBA{A: 255},
		color.RGBA{R: 241, G: 60, B: 60, A: 255},
		color.RGBA{R: 60, G: 81, B: 241, A: 255},
	)

	img, code = draw.Create(6, captcha.CLEAR)
	err = c.op.create(c.buildKey(id), code)
	return
}

func (c *imageCode) Verify(id, code string) (result bool) {
	return c.op.verify(c.buildKey(id), code)
}

func NewImageCode() *imageCode {
	return &imageCode{
		op: newOperate(),
	}
}

