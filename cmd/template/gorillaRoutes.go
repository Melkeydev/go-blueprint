package template

import (
	_ "embed"
)

//go:embed files/routes/gorilla.go.tmpl
var gorillaRoutesTemplate []byte

// GorillaTemplates contains the methods used for building
// an app that uses [github.com/gorilla/mux]
type GorillaTemplates struct{}

func (g GorillaTemplates) Main() []byte {
	return mainTemplate
}
func (g GorillaTemplates) Server() []byte {
	return standardServerTemplate
}
func (g GorillaTemplates) Routes() []byte {
	return gorillaRoutesTemplate
}

func (g GorillaTemplates) Plugin() []byte {
	return nil
}
