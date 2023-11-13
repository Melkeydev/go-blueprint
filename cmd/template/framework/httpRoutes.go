package framework

import (
	_ "embed"
)

//go:embed files/routes/standard_library.go.tmpl
var standardRoutesTemplate []byte

//go:embed files/dbRoutes/standard_library.go.tmpl
var standardDBRoutesTemplate []byte

//go:embed files/server/standard_library.go.tmpl
var standardServerTemplate []byte

//go:embed files/dbServer/standard_library.go.tmpl
var standardDBServerTemplate []byte

// StandardLibTemplate contains the methods used for building
// an app that uses [net/http]
type StandardLibTemplate struct{}

func (s StandardLibTemplate) Main() []byte {
	return mainTemplate
}

func (s StandardLibTemplate) Server() []byte {
	return standardServerTemplate
}

func (s StandardLibTemplate) ServerWithDB() []byte {
	return standardRoutesTemplate
}

func (s StandardLibTemplate) Routes() []byte {
	return standardRoutesTemplate
}

func (s StandardLibTemplate) RoutesWithDB() []byte {
	return standardDBRoutesTemplate
}
