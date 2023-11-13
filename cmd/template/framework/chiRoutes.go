package framework

import (
	_ "embed"
)

//go:embed files/routes/chi.go.tmpl
var chiRoutesTemplate []byte

//go:embed files/dbRoutes/chi.go.tmpl
var chiDBRoutesTemplate []byte
//go:embed files/tests/default-test.go.tmpl
var chiTestHandlerTemplate []byte

// ChiTemplates contains the methods used for building
// an app that uses [github.com/go-chi/chi]
type ChiTemplates struct{}

func (c ChiTemplates) Main() []byte {
	return mainTemplate
}

func (c ChiTemplates) Server() []byte {
	return standardServerTemplate
}

func (c ChiTemplates) ServerWithDB() []byte {
	return standardDBServerTemplate
}

func (c ChiTemplates) Routes() []byte {
	return chiRoutesTemplate
}

func (c ChiTemplates) RoutesWithDB() []byte {
	return chiDBRoutesTemplate
}

func (c ChiTemplates) TestHandler() []byte {
    return chiTestHandlerTemplate
}

