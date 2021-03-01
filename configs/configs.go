package configs

import "embed"

//go:embed *.yaml
var CONFIGS embed.FS
