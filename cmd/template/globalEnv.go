package template

import (
	_ "embed"
)

//go:embed backend/files/globalenv.tmpl
var globalEnvTemplate []byte

func GlobalEnvTemplate() []byte {
	return globalEnvTemplate
}
