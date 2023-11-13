package dbdriver

import (
	_ "embed"
)

type MysqlTemplate struct{}

//go:embed files/service/mongo.tmpl
var mysqlServiceTemplate []byte

//go:embed files/env/mongo.tmpl
var mysqlEnvTemplate []byte

func (m MysqlTemplate) Service() []byte {
	return mysqlServiceTemplate
}

func (m MysqlTemplate) Env() []byte {
	return mysqlEnvTemplate
}
