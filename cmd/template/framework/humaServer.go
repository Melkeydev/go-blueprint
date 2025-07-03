package framework

import (
	_ "embed"

	"github.com/melkeydev/go-blueprint/cmd/template/advanced"
)

//go:embed files/huma/huma_main.go.tmpl
var humaMainTemplate []byte

//go:embed files/huma/huma_server.go.tmpl
var humaServerTemplate []byte

//go:embed files/huma/huma_routes.go.tmpl
var humaRoutesTemplate []byte

//go:embed files/huma/huma_test.go.tmpl
var humaTestHandlerTemplate []byte

// HumaTemplates contains the methods used for building
// an app that uses [github.com/danielgtaylor/huma]
type HumaTemplates struct{}

func (h HumaTemplates) Main() []byte {
	return humaMainTemplate
}

func (h HumaTemplates) Server() []byte {
	return humaServerTemplate
}

func (h HumaTemplates) Routes() []byte {
	return humaRoutesTemplate
}

func (h HumaTemplates) TestHandler() []byte {
	return humaTestHandlerTemplate
}

// Note: Huma might have different ways of handling HTMX and WebSockets.
// These methods are included for consistency with FiberTemplates but may
// need adjustment or removal based on Huma's specific features.
func (h HumaTemplates) HtmxTemplImports() []byte {
	// Placeholder: Return Fiber's HTMX imports for now.
	// This will need to be adapted if Huma has its own way of handling HTMX.
	return advanced.FiberHtmxTemplImportsTemplate()
}

func (h HumaTemplates) HtmxTemplRoutes() []byte {
	// Placeholder: Return Fiber's HTMX routes for now.
	// This will need to be adapted if Huma has its own way of handling HTMX.
	return advanced.FiberHtmxTemplRoutesTemplate()
}

func (h HumaTemplates) WebsocketImports() []byte {
	// Placeholder: Return Fiber's WebSocket imports for now.
	// This will need to be adapted if Huma has its own way of handling WebSockets.
	return advanced.FiberWebsocketTemplImportsTemplate()
}
