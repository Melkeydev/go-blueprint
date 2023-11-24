package dbdriver

import (
	_ "embed"
)

type SqliteTemplate struct{}

//go:embed files/service/sqlite.tmpl
var sqliteServiceTemplate []byte

//go:embed files/env/sqlite.tmpl
var sqliteEnvTemplate []byte

//go:embed files/env/sqlite.example.tmpl
var sqliteEnvExampleTemplate []byte

func (m SqliteTemplate) Service() []byte {
	return sqliteServiceTemplate
}

func (m SqliteTemplate) Env() []byte {
	return sqliteEnvTemplate
}

func (m SqliteTemplate) EnvExample() []byte {
	return sqliteEnvExampleTemplate
}