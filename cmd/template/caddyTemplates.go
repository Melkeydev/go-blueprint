package template

import (
	_ "embed"
)

//go:embed files/main/caddy_main.go.tmpl
var caddyMainTemplate []byte

//go:embed files/plugin/caddy_plugin.go.tmpl
var caddyPluginTemplate []byte

type CaddyTemplates struct{}

func (c CaddyTemplates) Main() []byte {
	return caddyMainTemplate
}

func (c CaddyTemplates) Server() []byte {
	return nil
}

func (c CaddyTemplates) Routes() []byte {
	return nil
}

func (c CaddyTemplates) Plugin() []byte {
	return caddyPluginTemplate
}
