package framework

import (
	_ "embed"

	"github.com/melkeydev/go-blueprint/cmd/template/advanced"
)

//go:embed files/routes/bone.go.tmpl
var boneRoutesTemplate []byte

//go:embed files/dbRoutes/bone.go.tmpl
var boneDBRoutesTemplate []byte

// BoneTemplates contains the methods used for building
// an app that uses [github.com/go-zoo/bone]
type BoneTemplates struct{}

func (e BoneTemplates) Main() []byte {
	return mainTemplate
}
func (e BoneTemplates) Server() []byte {
	return standardServerTemplate
}

func (e BoneTemplates) ServerWithDB() []byte {
	return standardDBServerTemplate
}

func (e BoneTemplates) Routes() []byte {
	return boneRoutesTemplate
}

func (e BoneTemplates) RoutesWithDB() []byte {
	return boneDBRoutesTemplate
}
func (e BoneTemplates) TestHandler() []byte {
	return standardTestHandlerTemplate
}

func (e BoneTemplates) HtmxTemplImports() []byte {
	return advanced.StdLibHtmxTemplImportsTemplate()
}

func (e BoneTemplates) HtmxTemplRoutes() []byte {
	return advanced.StdLibHtmxTemplRoutesTemplate()
}
