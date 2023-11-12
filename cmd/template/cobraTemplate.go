package template

import _ "embed"

//go:embed files/cobra/cobra_cmd_root.go.tmpl
var cobraCMDRootTemplate []byte

//go:embed files/cobra/cobra_main.go.tmpl
var cobraMainTemplate []byte

// CobraTemplates contains the methods used for building
// an app that uses [github.com/spf13/cobra] for building a CLI app.
type CobraTemplates struct{}

func (c CobraTemplates) Main() []byte {
	return cobraMainTemplate
}

func (c CobraTemplates) Server() []byte {
	return []byte("")
}

func (c CobraTemplates) Routes() []byte {
	return []byte("")
}

// MakeCobraCMDRoot returns a byte slice that represents
// the cmd/root.go file when using Cobra.
func MakeCobraCMDRoot() []byte {
	return cobraCMDRootTemplate
}
