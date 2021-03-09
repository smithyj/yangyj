package captcha

import (
	"fmt"
	"yangyj/backend/pkg/helper"
	"yangyj/backend/pkg/sms"
)

type phoneCaptcha struct {
	countryNo string
	op *operate
}

func (c *phoneCaptcha) buildKey(key string) string {
	return fmt.Sprintf("mobile:%v%v", c.countryNo, key)
}

// 创建手机验证码
//
// phone: 手机号码
func (c *phoneCaptcha) Create(phone string) (err error) {
	code := helper.RandomStr(4, 0)
	if err = c.op.create(c.buildKey(phone), code); err != nil {
		return
	}
	s, err := sms.New(c.countryNo)
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
	return c.op.verify(c.buildKey(phone), code)
}

func NewPhoneCaptcha(countryNo string) *phoneCaptcha {
	return &phoneCaptcha{
		countryNo: helper.FilterCountryNo(countryNo),
		op: newOperate(),
	}
}
