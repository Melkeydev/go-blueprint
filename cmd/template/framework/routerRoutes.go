package framework

import (
	_ "embed"
)

//go:embed files/routes/http_router.go.tmpl
var httpRouterRoutesTemplate []byte

//go:embed files/dbRoutes/http_router.go.tmpl
var httpDBRouterRoutesTemplate []byte
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

func (r RouterTemplates) ServerWithDB() []byte {
	return standardDBServerTemplate
}

func (r RouterTemplates) Routes() []byte {
	return httpRouterRoutesTemplate
}

func (r RouterTemplates) RoutesWithDB() []byte {
	return httpDBRouterRoutesTemplate
}
func (r RouterTemplates) TestHandler() []byte {
    return httpRouterTestHandlerTemplate
}
