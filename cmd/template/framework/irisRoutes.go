package framework

import (
	_ "embed"
)

//go:embed files/routes/iris.go.tmpl
var irisRoutesTemplate []byte

//go:embed files/dbRoutes/iris.go.tmpl
var irisDBRoutesTemplate []byte

//go:embed files/server/iris.go.tmpl
var irisServerTemplate []byte

//go:embed files/dbServer/iris.go.tmpl
var irisDBServerTemplate []byte

//go:embed files/main/iris_main.go.tmpl
var irisMainTemplate []byte

//go:embed files/tests/iris-test.go.tmpl
var irisTestHandlerTemplate []byte

// irisTemplates contains the methods used for building
// an app that uses [github.com/labstack/iris]
type IrisTemplates struct{}

func (i IrisTemplates) Main() []byte {
	return irisMainTemplate
}

func (i IrisTemplates) Server() []byte {
	return irisServerTemplate
}

func (i IrisTemplates) ServerWithDB() []byte {
	return irisDBServerTemplate
}

func (i IrisTemplates) Routes() []byte {
	return irisRoutesTemplate
}

func (e IrisTemplates) RoutesWithDB() []byte {
	return irisDBRoutesTemplate
}

func (e IrisTemplates) TestHandler() []byte {
	return irisTestHandlerTemplate
}
