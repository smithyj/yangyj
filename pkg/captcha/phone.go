package captcha

import (
	"fmt"
	"yangyj/pkg/helper"
	"yangyj/pkg/sms"
)

type phoneCaptcha struct {
	countryCode string
	captcha *captcha
}

func (c *phoneCaptcha) buildKey(key string) string {
	return fmt.Sprintf("mobile:%v%v", c.countryCode, key)
}

// 创建手机验证码
//
// phone: 手机号码
func (c *phoneCaptcha) Create(phone string) (err error) {
	code := helper.RandomStr(4, 0)
	if err = c.captcha.create(c.buildKey(phone), code); err != nil {
		return
	}
	s, err := sms.New(c.countryCode)
	if err != nil {
		return
	}
	return s.CaptchaCode(phone, code)
}

// 手机验证码校验
//
// phone: 手机号码
// code: 验证码
func (c *phoneCaptcha) Verify(phone, code string) (ok bool) {
	return c.captcha.verify(c.buildKey(phone), code)
}

func NewPhoneCaptcha(countryCode string) *phoneCaptcha {
	return &phoneCaptcha{
		countryCode: helper.FilterCountryCode(countryCode),
		captcha:   newCaptcha(),
	}
}

