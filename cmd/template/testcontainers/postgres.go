package testcontainers

import (
	_ "embed"
)

type PostgresTestcontainersTemplate struct{}

//go:embed files/postgres.tmpl
var postgresTestcontainersTemplate []byte

func (m PostgresTestcontainersTemplate) Testcontainers() []byte {
	return postgresTestcontainersTemplate
}
