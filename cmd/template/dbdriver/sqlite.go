package dbdriver

import (
	_ "embed"
)

type SqliteTemplate struct{}

//go:embed files/service/mongo.tmpl
var sqliteServiceTemplate []byte

//go:embed files/env/mongo.tmpl
var sqliteEnvTemplate []byte

func (m SqliteTemplate) Service() []byte {
	return sqliteServiceTemplate
}

func (m SqliteTemplate) Env() []byte {
	return sqliteEnvTemplate
}
