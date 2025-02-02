package backend

import (
	_ "embed"

	"github.com/melkeydev/go-blueprint/cmd/template/advanced"
	"github.com/melkeydev/go-blueprint/cmd/template/frontend"
)

//go:embed files/routes/gorilla.go.tmpl
var gorillaRoutesTemplate []byte

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

func (g GorillaTemplates) Routes() []byte {
	return gorillaRoutesTemplate
}

func (g GorillaTemplates) TestHandler() []byte {
	return gorillaTestHandlerTemplate
}

func (g GorillaTemplates) HtmxTemplImports() []byte {
	return frontend.StdLibHtmxTemplImportsTemplate()
}

func (g GorillaTemplates) HtmxTemplRoutes() []byte {
	return frontend.GorillaHtmxTemplRoutesTemplate()
}

func (g GorillaTemplates) WebsocketImports() []byte {
	return advanced.StdLibWebsocketTemplImportsTemplate()
}
