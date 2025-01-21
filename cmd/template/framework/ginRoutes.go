package framework

import (
	_ "embed"

	"github.com/melkeydev/go-blueprint/cmd/template/advanced"
	"github.com/melkeydev/go-blueprint/cmd/template/frontend"
)

//go:embed files/routes/gin.go.tmpl
var ginRoutesTemplate []byte

//go:embed files/tests/gin-test.go.tmpl
var ginTestHandlerTemplate []byte

// GinTemplates contains the methods used for building
// an app that uses [github.com/gin-gonic/gin]
type GinTemplates struct{}

func (g GinTemplates) Main() []byte {
	return mainTemplate
}

func (g GinTemplates) Server() []byte {
	return standardServerTemplate
}

func (g GinTemplates) Routes() []byte {
	return ginRoutesTemplate
}

func (g GinTemplates) TestHandler() []byte {
	return ginTestHandlerTemplate
}

func (g GinTemplates) HtmxTemplImports() []byte {
	return frontend.GinHtmxTemplImportsTemplate()
}

func (g GinTemplates) HtmxTemplRoutes() []byte {
	return frontend.GinHtmxTemplRoutesTemplate()
}

func (g GinTemplates) WebsocketImports() []byte {
	return advanced.StdLibWebsocketTemplImportsTemplate()
}
