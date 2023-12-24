package docker

import (
	_ "embed"
)

type SqlServerDockerTemplate struct{}

//go:embed files/docker-compose/sqlserver.tmpl
var sqlServerDockerTemplate []byte

func (m SqlServerDockerTemplate) Docker() []byte {
	return sqlServerDockerTemplate
}
