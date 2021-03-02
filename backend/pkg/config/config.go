package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
	"yangyj/backend/configs"
)

type Setting configs.Config

var Config *Setting

func init() {
	var err error
	env := os.Getenv("ENV")
	filename := "config"
	ext := "yaml"
	if env != "" {
		filename = fmt.Sprintf("%s_%s", filename, env)
	}
	filename = fmt.Sprintf("%s.%s", filename, ext)
	bytes, err := configs.FS.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	if err = yaml.Unmarshal(bytes, &Config); err != nil {
		panic(err)
	}
}
