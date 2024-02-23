package framework

import (
	_ "embed"

	"github.com/melkeydev/go-blueprint/cmd/template/advanced"
)

//go:embed files/routes/chi.go.tmpl
var chiRoutesTemplate []byte

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

func (c ChiTemplates) Routes() []byte {
	return chiRoutesTemplate
}

func (c ChiTemplates) TestHandler() []byte {
	return chiTestHandlerTemplate
}

func (c ChiTemplates) HtmxTemplImports() []byte {
	return advanced.StdLibHtmxTemplImportsTemplate()
}

func (c ChiTemplates) HtmxTemplRoutes() []byte {
	return advanced.ChiHtmxTemplRoutesTemplate()
}

func (c ChiTemplates) WebsocketImports() []byte {
	return advanced.StdLibWebsocketTemplImportsTemplate()
}
