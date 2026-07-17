package framework

import (
	_ "embed"

	"github.com/melkeydev/go-blueprint/cmd/template/advanced"
)

//go:embed files/routes/huma.go.tmpl
var humaRoutesTemplate []byte

//go:embed files/tests/huma-test.go.tmpl
var humaTestHandlerTemplate []byte

// HumaTemplates contains the methods used for building
// an app that uses [github.com/danielgtaylor/huma] on top of [github.com/go-chi/chi]
type HumaTemplates struct{}

func (h HumaTemplates) Main() []byte {
	return mainTemplate
}

func (h HumaTemplates) Server() []byte {
	return standardServerTemplate
}

func (h HumaTemplates) Routes() []byte {
	return humaRoutesTemplate
}

func (h HumaTemplates) TestHandler() []byte {
	return humaTestHandlerTemplate
}

func (h HumaTemplates) HtmxTemplImports() []byte {
	return advanced.StdLibHtmxTemplImportsTemplate()
}

func (h HumaTemplates) HtmxTemplRoutes() []byte {
	return advanced.ChiHtmxTemplRoutesTemplate()
}

func (h HumaTemplates) WebsocketImports() []byte {
	return advanced.StdLibWebsocketTemplImportsTemplate()
}
