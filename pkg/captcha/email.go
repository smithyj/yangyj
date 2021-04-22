package captcha

import (
	"fmt"
	"yangyj/pkg/email"
	"yangyj/pkg/helper"
)

type emailCaptcha struct {
	captcha *captcha
}

func (c *emailCaptcha) buildKey(key string) string {
	return fmt.Sprintf("email:%v", key)
}

// 创建邮箱验证码
//
// address: 邮箱地址
func (c *emailCaptcha) Create(address string) (err error) {
	code := helper.RandomStr(6, 4)
	if err = c.captcha.create(c.buildKey(address), code); err != nil {
		return
	}
	err = email.New().CaptchaCode(address, code)
	return
}

// 邮箱验证码校验
//
// address: 邮箱地址
// code: 验证码
func (c *emailCaptcha) Verify(address, code string) (ok bool) {
	return c.captcha.verify(c.buildKey(address), code)
}

func NewEmailCaptcha() *emailCaptcha {
	return &emailCaptcha{
		captcha: newCaptcha(),
	}
}
