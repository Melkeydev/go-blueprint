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

// ReadmeTemplate returns a byte slice that represents
// the default README.md file template.
func ReadmeTemplate() []byte {
	return readmeTemplate
}
