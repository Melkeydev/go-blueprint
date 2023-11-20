package framework

import (
	_ "embed"
)

//go:embed files/routes/echo.go.tmpl
var echoRoutesTemplate []byte

//go:embed files/dbRoutes/echo.go.tmpl
var echoDBRoutesTemplate []byte
//go:embed files/tests/echo-test.go.tmpl
var echoTestHandlerTemplate []byte

// EchoTemplates contains the methods used for building
// an app that uses [github.com/labstack/echo]
type EchoTemplates struct{}

func (e EchoTemplates) Main() []byte {
	return mainTemplate
}
func (e EchoTemplates) Server() []byte {
	return standardServerTemplate
}

func (e EchoTemplates) ServerWithDB() []byte {
	return standardDBServerTemplate
}

func (e EchoTemplates) Routes() []byte {
	return echoRoutesTemplate
}

func (e EchoTemplates) RoutesWithDB() []byte {
	return echoDBRoutesTemplate
}
func (e EchoTemplates) TestHandler() []byte {
    return echoTestHandlerTemplate
}
