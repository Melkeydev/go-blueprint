package backend

import (
	_ "embed"

	"github.com/melkeydev/go-blueprint/cmd/template/advanced"
	"github.com/melkeydev/go-blueprint/cmd/template/frontend"
)

//go:embed files/routes/standard_library.go.tmpl
var standardRoutesTemplate []byte

//go:embed files/server/standard_library.go.tmpl
var standardServerTemplate []byte

//go:embed files/tests/default-test.go.tmpl
var standardTestHandlerTemplate []byte

// StandardLibTemplate contains the methods used for building
// an app that uses [net/http]
type StandardLibTemplate struct{}

func (s StandardLibTemplate) Main() []byte {
	return mainTemplate
}

func (s StandardLibTemplate) Server() []byte {
	return standardServerTemplate
}

func (s StandardLibTemplate) Routes() []byte {
	return standardRoutesTemplate
}

func (s StandardLibTemplate) TestHandler() []byte {
	return standardTestHandlerTemplate
}

func (s StandardLibTemplate) HtmxTemplImports() []byte {
	return frontend.StdLibHtmxTemplImportsTemplate()
}

func (s StandardLibTemplate) HtmxTemplRoutes() []byte {
	return frontend.StdLibHtmxTemplRoutesTemplate()
}

func (s StandardLibTemplate) WebsocketImports() []byte {
	return advanced.StdLibWebsocketTemplImportsTemplate()
}
