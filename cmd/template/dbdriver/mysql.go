package dbdriver

import (
	_ "embed"
)

type MysqlTemplate struct{}

//go:embed files/service/mysql.tmpl
var mysqlServiceTemplate []byte

//go:embed files/env/mysql.tmpl
var mysqlEnvTemplate []byte

//go:embed files/tests/mysql.tmpl
var mysqlTestcontainersTemplate []byte

func (m MysqlTemplate) Service() []byte {
	return mysqlServiceTemplate
}

func (m MysqlTemplate) Env() []byte {
	return mysqlEnvTemplate
}

func (m MysqlTemplate) Tests() []byte {
	return mysqlTestcontainersTemplate
}
