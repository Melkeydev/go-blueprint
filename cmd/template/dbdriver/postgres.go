package dbdriver

import (
	_ "embed"
)

type PostgresTemplate struct{}

//go:embed files/service/mongo.tmpl
var postgresServiceTemplate []byte

//go:embed files/env/mongo.tmpl
var postgresEnvTemplate []byte

func (m PostgresTemplate) Service() []byte {
	return postgresServiceTemplate
}

func (m PostgresTemplate) Env() []byte {
	return postgresEnvTemplate
}
