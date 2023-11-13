package framework

import (
	_ "embed"
)

//go:embed files/routes/fiber.go.tmpl
var fiberRoutesTemplate []byte

//go:embed files/dbRoutes/fiber.go.tmpl
var fiberDBRoutesTemplate []byte

//go:embed files/server/fiber.go.tmpl
var fiberServerTemplate []byte

//go:embed files/dbServer/fiber.go.tmpl
var fiberDBServerTemplate []byte

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

func (f FiberTemplates) ServerWithDB() []byte {
	return fiberDBServerTemplate
}

func (f FiberTemplates) Routes() []byte {
	return fiberRoutesTemplate
}

func (f FiberTemplates) RoutesWithDB() []byte {
	return fiberDBRoutesTemplate
}
func (f FiberTemplates) TestHandler() []byte {
    return fiberTestHandlerTemplate
}
