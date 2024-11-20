package dbdriver

import (
	_ "embed"
)

type ScyllaTemplate struct{}

//go:embed files/service/scylla.tmpl
var scyllaServiceTemplate []byte

//go:embed files/env/scylla.tmpl
var scyllaEnvTemplate []byte

//go:embed files/tests/scylla.tmpl
var scyllaTestcontainersTemplate []byte

func (r ScyllaTemplate) Service() []byte {
	return scyllaServiceTemplate
}

func (r ScyllaTemplate) Env() []byte {
	return scyllaEnvTemplate
}

func (r ScyllaTemplate) Tests() []byte {
	return scyllaTestcontainersTemplate
}
