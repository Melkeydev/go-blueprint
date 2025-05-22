package builder

import _ "embed"

//go:embed files/makefile.tmpl
var makeTemplate []byte

//go:embed files/justfile.tmpl
var justTemplate []byte

func MakeTemplate() []byte {
	return makeTemplate
}

func JustTemplate() []byte {
	return justTemplate
}
