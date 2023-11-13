package template

import (
	_ "embed"
)

//go:embed framework/files/globalenv.tmpl
var globalEnvTemplate []byte

func GlobalEnvTemplate() []byte {
	return globalEnvTemplate
}
