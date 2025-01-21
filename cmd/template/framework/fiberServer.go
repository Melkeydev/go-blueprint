package framework

import (
	_ "embed"

	"github.com/melkeydev/go-blueprint/cmd/template/advanced"
	"github.com/melkeydev/go-blueprint/cmd/template/frontend"
)

//go:embed files/routes/fiber.go.tmpl
var fiberRoutesTemplate []byte

//go:embed files/server/fiber.go.tmpl
var fiberServerTemplate []byte

//go:embed files/main/fiber_main.go.tmpl
var fiberMainTemplate []byte

//go:embed files/tests/fiber-test.go.tmpl
var fiberTestHandlerTemplate []byte

// FiberTemplates contains the methods used for building
// an app that uses [github.com/gofiber/fiber]
type FiberTemplates struct{}

func (f FiberTemplates) Main() []byte {
	return fiberMainTemplate
}
func (f FiberTemplates) Server() []byte {
	return fiberServerTemplate
}

func (f FiberTemplates) Routes() []byte {
	return fiberRoutesTemplate
}

func (f FiberTemplates) TestHandler() []byte {
	return fiberTestHandlerTemplate
}

func (f FiberTemplates) HtmxTemplImports() []byte {
	return frontend.FiberHtmxTemplImportsTemplate()
}

func (f FiberTemplates) HtmxTemplRoutes() []byte {
	return frontend.FiberHtmxTemplRoutesTemplate()
}

func (f FiberTemplates) WebsocketImports() []byte {
	return advanced.FiberWebsocketTemplImportsTemplate()
}
