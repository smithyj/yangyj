package sms

import (
	"crypto/md5"
	"errors"
	"fmt"
	"net/url"
	"yangyj/backend/pkg/config"
	"yangyj/backend/pkg/helper"
)

var smsbaoErr = map[string]string{
	"30": "错误密码",
	"40": "账号不存在",
	"41": "余额不足",
	"43": "IP地址限制",
	"50": "内容含有敏感词",
	"51": "手机号码不正确",
}

type smsbao struct {
	countryNo string
}

func (s *smsbao) send(phone, content string) error {
	cfg := config.Config.Sms.Smsbao
	baseUrl := "https://api.smsbao.com/sms"
	// 国际手机号码
	if s.countryNo != "86" {
		baseUrl = "https://api.smsbao.com/wsms"
		phone = url.QueryEscape(fmt.Sprintf("+%v%v", s.countryNo, phone))
	}
	content = url.QueryEscape(content)
	pwd := fmt.Sprintf("%x", md5.Sum([]byte(cfg.Password)))
	reqUrl := fmt.Sprintf("%v?u=%v&p=%v&m=%v&c=%v", baseUrl, cfg.Username, pwd, phone, content)

	result, err := helper.Get(reqUrl)
	if err != nil {
		return errors.New("发送失败，短信平台故障")
	}
	if result == "" {
		return errors.New("请求短信接口失败")
	}
	if result != "0" {
		message, ok := smsbaoErr[result]
		if !ok {
			message = "未知的错误码"
		}
		return errors.New(fmt.Sprintf("发送失败，来自短信平台提示：%v", message))
	}

	return nil
}

func (s *smsbao) CaptchaCode(phone, code string) (err error) {
	expired := config.Config.Captcha.Expired
	content := fmt.Sprintf(SMSBAO_CAPTCHACODE_TPL, code, expired)
	if s.countryNo != "86" {
		content = fmt.Sprintf(SMSBAO_CAPTCHACODE_EN_TPL, code, expired)
	}
	return s.send(phone, content)
}