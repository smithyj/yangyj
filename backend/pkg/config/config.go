package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
	"yangyj/backend/configs"
)

type config configs.Config

var Config *config

func init() {
	var err error
	var bytes []byte
	env := os.Getenv("ENV")
	filename := "config"
	ext := "yaml"
	if env != "" {
		filename = fmt.Sprintf("%s_%s", filename, env)
	}
	filename = fmt.Sprintf("%s.%s", filename, ext)
	if bytes, err = configs.FS.ReadFile(filename); err != nil {
		panic(err)
	}
	if err = yaml.Unmarshal(bytes, &Config); err != nil {
		panic(err)
	}
}
