// Package template provides utility functions that
// help with the templating of created files.
package framework

import (
	_ "embed"
)

//go:embed files/main/main.go.tmpl
var mainTemplate []byte

//go:embed files/air.toml.tmpl
var airTomlTemplate []byte

//go:embed files/htmx/hello.templ.tmpl
var helloTemplTemplate []byte

//go:embed files/htmx/base.templ.tmpl
var baseTemplTemplate []byte

//go:embed files/htmx/htmx.min.js.tmpl
var htmxMinJsTemplate []byte

//go:embed files/htmx/efs.go.tmpl
var efsTemplate []byte

//go:embed files/htmx/hello.go.tmpl
var helloGoTemplate []byte

//go:embed files/README.md.tmpl
var readmeTemplate []byte

//go:embed files/makefile.tmpl
var makeTemplate []byte

//go:embed files/gitignore.tmpl
var gitIgnoreTemplate []byte

// MakeTemplate returns a byte slice that represents
// the default Makefile template.
func MakeTemplate() []byte {
	return makeTemplate
}

func GitIgnoreTemplate() []byte {
	return gitIgnoreTemplate
}

func AirTomlTemplate() []byte {
	return airTomlTemplate
}

func HelloTemplTemplate() []byte {
	return helloTemplTemplate
}

func BaseTemplTemplate() []byte {
	return baseTemplTemplate
}

func HtmxJSTemplate() []byte {
	return htmxMinJsTemplate
}

func EfsTemplate() []byte {
	return efsTemplate
}

func HelloGoTemplate() []byte {
	return helloGoTemplate
}

// ReadmeTemplate returns a byte slice that represents
// the default README.md file template.
func ReadmeTemplate() []byte {
	return readmeTemplate
}
