package docker

import (
	_ "embed"
)

type PostgresDockerTemplate struct{}

//go:embed files/docker-compose/postgres.tmpl
var postgresDockerTemplate []byte

func (m PostgresDockerTemplate) Docker() []byte {
	return postgresDockerTemplate
}