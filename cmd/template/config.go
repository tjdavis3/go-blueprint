package template

import (
	_ "embed"
)

//go:embed framework/files/config.tmpl
var configTemplate []byte

func ConfigTemplate() []byte {
	return configTemplate
}
