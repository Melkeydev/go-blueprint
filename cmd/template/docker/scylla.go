package docker

import (
	_ "embed"
)

type ScyllaDockerTemplate struct{}

//go:embed files/docker-compose/scylla.tmpl
var scyllaDockerTemplate []byte

func (r ScyllaDockerTemplate) Docker() []byte {
	return scyllaDockerTemplate
}
