package dbdriver

import (
	_ "embed"
)

type MysqlTemplate struct{}

//go:embed files/service/mysql.tmpl
var mysqlServiceTemplate []byte

//go:embed files/env/mysql.tmpl
var mysqlEnvTemplate []byte

//go:embed files/docker-compose/mysql.tmpl
var mysqlDBTemplate []byte

func (m MysqlTemplate) Service() []byte {
	return mysqlServiceTemplate
}

func (m MysqlTemplate) Env() []byte {
	return mysqlEnvTemplate
}

func (m MysqlTemplate) DB() []byte {
	return mysqlDBTemplate
}