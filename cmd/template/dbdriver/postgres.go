package dbdriver

import (
	_ "embed"
)

type PostgresTemplate struct{}

//go:embed files/service/postgres.tmpl
var postgresServiceTemplate []byte

//go:embed files/env/postgres.tmpl
var postgresEnvTemplate []byte

//go:embed files/tests/postgres.tmpl
var postgresTestcontainersTemplate []byte

func (m PostgresTemplate) Service() []byte {
	return postgresServiceTemplate
}

func (m PostgresTemplate) Env() []byte {
	return postgresEnvTemplate
}

func (m PostgresTemplate) Tests() []byte {
	return postgresTestcontainersTemplate
}
