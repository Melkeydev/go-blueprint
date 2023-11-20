package framework

import (
	_ "embed"
)

//go:embed files/routes/gorilla.go.tmpl
var gorillaRoutesTemplate []byte

//go:embed files/dbRoutes/gorilla.go.tmpl
var gorillaDBRoutesTemplate []byte
//go:embed files/tests/default-test.go.tmpl
var gorillaTestHandlerTemplate []byte

// GorillaTemplates contains the methods used for building
// an app that uses [github.com/gorilla/mux]
type GorillaTemplates struct{}

func (g GorillaTemplates) Main() []byte {
	return mainTemplate
}

func (g GorillaTemplates) Server() []byte {
	return standardServerTemplate
}

func (g GorillaTemplates) ServerWithDB() []byte {
	return standardDBServerTemplate
}

func (g GorillaTemplates) Routes() []byte {
	return gorillaRoutesTemplate
}

func (g GorillaTemplates) RoutesWithDB() []byte {
	return gorillaDBRoutesTemplate
}
func (g GorillaTemplates) TestHandler() []byte {
    return gorillaTestHandlerTemplate
}
