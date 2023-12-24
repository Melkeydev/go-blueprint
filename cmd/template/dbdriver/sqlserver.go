package dbdriver

import (
	_ "embed"
)

type SqlServerTemplate struct{}

//go:embed files/service/sqlserver.tmpl
var sqlServerServiceTemplate []byte

//go:embed files/env/sqlserver.tmpl
var sqlServerEnvTemplate []byte

//go:embed files/env/sqlserver.example.tmpl
var sqlServerEnvExampleTemplate []byte

func (m SqlServerTemplate) Service() []byte {
	return sqlServerServiceTemplate
}

func (m SqlServerTemplate) Env() []byte {
	return sqlServerEnvTemplate
}

func (m SqlServerTemplate) EnvExample() []byte {
	return sqlServerEnvExampleTemplate
}
