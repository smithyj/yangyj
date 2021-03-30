package e

import (
	"yangyj/pkg/config"
	"yangyj/pkg/i18n"
)

// 默认获取
func Msg(code int) string {
	return I18NMsg(config.Config.Lang, code)
}

// 国际化获取
func I18NMsg(lang string, code int) string {
	id, ok := codeMap[code]
	if !ok {
		id = codeMap[ERROR]
	}
	return i18n.Trans(&i18n.Option{
		ID: id,
		Lang: lang,
	})
}