package config

import (
	"flag"
	"gopkg.in/yaml.v2"
	"yangyj/configs"
)

type config configs.Config

var Config *config

func InitConfig() {
	var err error
	var buf []byte
	var filename string
	flag.StringVar(&filename, "f", "config.yaml", "程序配置文件")
	flag.Parse()
	if buf, err = configs.FS.ReadFile(filename); err != nil {
		panic(err)
	}
	if err = yaml.Unmarshal(buf, &Config); err != nil {
		panic(err)
	}
}
