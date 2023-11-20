package framework

import (
	_ "embed"
)

//go:embed files/routes/gin.go.tmpl
var ginRoutesTemplate []byte

//go:embed files/dbRoutes/gin.go.tmpl
var ginDBRoutesTemplate []byte
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

func (g GinTemplates) ServerWithDB() []byte {
	return standardDBServerTemplate
}

func (g GinTemplates) Routes() []byte {
	return ginRoutesTemplate
}

func (g GinTemplates) RoutesWithDB() []byte {
	return ginDBRoutesTemplate
}
func (g GinTemplates) TestHandler() []byte {
    return ginTestHandlerTemplate
}
