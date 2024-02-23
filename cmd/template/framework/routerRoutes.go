package framework

import (
	_ "embed"

	"github.com/melkeydev/go-blueprint/cmd/template/advanced"
)

//go:embed files/routes/http_router.go.tmpl
var httpRouterRoutesTemplate []byte

//go:embed files/tests/default-test.go.tmpl
var httpRouterTestHandlerTemplate []byte

// RouterTemplates contains the methods used for building
// an app that uses [github.com/julienschmidt/httprouter]
type RouterTemplates struct{}

func (r RouterTemplates) Main() []byte {
	return mainTemplate
}
func (r RouterTemplates) Server() []byte {
	return standardServerTemplate
}

func (r RouterTemplates) Routes() []byte {
	return httpRouterRoutesTemplate
}

func (r RouterTemplates) TestHandler() []byte {
	return httpRouterTestHandlerTemplate
}

func (r RouterTemplates) HtmxTemplImports() []byte {
	return advanced.StdLibHtmxTemplImportsTemplate()
}

func (r RouterTemplates) HtmxTemplRoutes() []byte {
	return advanced.HttpRouterHtmxTemplRoutesTemplate()
}

func (r RouterTemplates) WebsocketImports() []byte {
	return advanced.StdLibWebsocketTemplImportsTemplate()
}
