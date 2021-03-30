package i18n

import (
	"fmt"
	"golang.org/x/text/language"
	"gopkg.in/yaml.v2"
	"io/fs"
	"strings"
	"yangyj/assets"
	"yangyj/pkg/config"
)

var locales map[string]map[string]string

type Option struct {
	ID string
	Default string
	Params []interface{}
	Lang string
}

func InitLang()  {
	var err error
	var dirs []fs.DirEntry
	baseDir := "lang"
	if dirs, err = assets.FS.ReadDir(baseDir); err != nil {
		panic("lang dir not exists")
	}
	locales = make(map[string]map[string]string)
	for _, dir := range dirs {
		if isDir := dir.IsDir(); !isDir {
			continue
		}
		name := dir.Name()

		dirname := fmt.Sprintf("%v/%v", baseDir, name)
		var files []fs.DirEntry
		if files, err = assets.FS.ReadDir(dirname); err != nil {
			continue
		}

		if _, ok := locales[name]; !ok {
			locales[name] = make(map[string]string)
		}

		for _, file := range files {
			yamlName := file.Name()
			filename := fmt.Sprintf("%v/%v", dirname, yamlName)
			var fileByte []byte
			if fileByte, err = assets.FS.ReadFile(filename); err != nil {
				panic(fmt.Sprintf("load lang file err, file path is: %v", filename))
			}
			var locale = make(map[string]string)
			if err = yaml.Unmarshal(fileByte, &locale); err != nil {
				panic(fmt.Sprintf("lang file invalid, file path is: %v", filename))
			}
			for key, val := range locale {
				locales[name][key] = val
			}
		}
	}
}

func Trans(option *Option) (message string) {
	var locale map[string]string
	ok := false
	message = option.Default
	lang := option.Lang
	if lang == "" {
		lang = config.Config.Lang
	}
	t, _, err := language.ParseAcceptLanguage(lang)
	// 解析语言失败
	if err != nil {
		if message == "" {
			message = option.ID
		}
		total := strings.Count(message, "%v")
		// 参数个数与占位标识个数一致才走格式化
		if total == len(option.Params) {
			message = fmt.Sprintf(message, option.Params...)
		}
		return
	}
	matched := false
	for _, tag := range t {
		if locale, ok = locales[tag.String()]; !ok {
			// 继续匹配下一个国际化语言
			continue
		}
		if message, ok = locale[option.ID]; !ok {
			// 继续匹配下一个国际化语言
			continue
		}
		// 匹配成功
		matched = true
		total := strings.Count(message, "%v")
		// 参数个数与占位标识个数一致才走格式化
		if total == len(option.Params) {
			message = fmt.Sprintf(message, option.Params...)
		}
		return
	}
	// 匹配失败
	if !matched {
		if message == "" {
			message = option.ID
		}
		total := strings.Count(message, "%v")
		// 参数个数与占位标识个数一致才走格式化
		if total == len(option.Params) {
			message = fmt.Sprintf(message, option.Params...)
		}
	}
	return
}