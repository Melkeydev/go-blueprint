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

//go:embed files/sqlc/sqlc.yaml.tmpl
var postgresSqlcConfigTemplate []byte

//go:embed files/schema/users_postgres.sql.tmpl
var postgresUsersSchemaTemplate []byte

//go:embed files/schema/posts_postgres.sql.tmpl
var postgresPostsSchemaTemplate []byte

//go:embed files/query/users/postgres.sql.tmpl
var postgresUsersQueryTemplate []byte

//go:embed files/query/posts/postgres.sql.tmpl
var postgresPostsQueryTemplate []byte

func (m PostgresTemplate) Service() []byte {
	return postgresServiceTemplate
}

func (m PostgresTemplate) Env() []byte {
	return postgresEnvTemplate
}

func (m PostgresTemplate) Tests() []byte {
	return postgresTestcontainersTemplate
}

func (m PostgresTemplate) SqlcConfig() []byte {
	return postgresSqlcConfigTemplate
}

func (m PostgresTemplate) SchemaExample() []byte {
	return postgresUsersSchemaTemplate // This method is kept for backward compatibility
}

func (m PostgresTemplate) QueryExample() []byte {
	return postgresUsersQueryTemplate // This method is kept for backward compatibility
}

func (m PostgresTemplate) UsersSchema() []byte {
	return postgresUsersSchemaTemplate
}

func (m PostgresTemplate) PostsSchema() []byte {
	return postgresPostsSchemaTemplate
}

func (m PostgresTemplate) UsersQuery() []byte {
	return postgresUsersQueryTemplate
}

func (m PostgresTemplate) PostsQuery() []byte {
	return postgresPostsQueryTemplate
}
