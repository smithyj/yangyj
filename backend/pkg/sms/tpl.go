package sms

const (
	// aliyun 模板
	// ALIYUN_CAPTCHACODE_TPL_CONTENT = "【YANGYJ】验证码 ${code}，10分钟内有效，请妥善保管"
	ALIYUN_CAPTCHACODE_TPL = "YANGYJ,SMS_190095340"
	ALIYUN_CAPTCHACODE_EN_TPL = "YANGYJ,SMS_212471167"

	// smsbao 模板
	SMSBAO_CAPTCHACODE_TPL    = "【YANGYJ】您的验证码：%v，%v分钟内有效，请勿泄露给他人！"
	SMSBAO_CAPTCHACODE_EN_TPL = "【YANGYJ】Your code is：%v，valid for %v minutes，do not disclose it to others！"
)
