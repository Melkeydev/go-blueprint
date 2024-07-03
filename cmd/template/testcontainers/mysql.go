package testcontainers

import (
	_ "embed"
)

type MysqlTestcontainersTemplate struct{}

//go:embed files/mysql.tmpl
var mysqlTestcontainersTemplate []byte

func (m MysqlTestcontainersTemplate) Testcontainers() []byte {
	return mysqlTestcontainersTemplate
}
