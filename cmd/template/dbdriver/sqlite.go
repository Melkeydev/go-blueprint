package dbdriver

import (
	_ "embed"
)

type SqliteTemplate struct{}

//go:embed files/service/sqlite.tmpl
var sqliteServiceTemplate []byte

//go:embed files/env/sqlite.tmpl
var sqliteEnvTemplate []byte

//go:embed files/docker-compose/sqlite.tmpl
var sqliteDBTemplate []byte

func (m SqliteTemplate) Service() []byte {
	return sqliteServiceTemplate
}

func (m SqliteTemplate) Env() []byte {
	return sqliteEnvTemplate
}

func (m SqliteTemplate) DB() []byte {
	return sqliteDBTemplate
}