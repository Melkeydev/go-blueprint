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

//go:embed files/sqlc/sqlc.yaml.tmpl
var mysqlSqlcConfigTemplate []byte

//go:embed files/schema/users_mysql.sql.tmpl
var mysqlUsersSchemaTemplate []byte

//go:embed files/schema/posts_mysql.sql.tmpl
var mysqlPostsSchemaTemplate []byte

//go:embed files/query/users/mysql.sql.tmpl
var mysqlUsersQueryTemplate []byte

//go:embed files/query/posts/mysql.sql.tmpl
var mysqlPostsQueryTemplate []byte

func (m MysqlTemplate) Service() []byte {
	return mysqlServiceTemplate
}

func (m MysqlTemplate) Env() []byte {
	return mysqlEnvTemplate
}

func (m MysqlTemplate) Tests() []byte {
	return mysqlTestcontainersTemplate
}

func (m MysqlTemplate) SqlcConfig() []byte {
	return mysqlSqlcConfigTemplate
}

func (m MysqlTemplate) SchemaExample() []byte {
	return mysqlUsersSchemaTemplate // This method is kept for backward compatibility
}

func (m MysqlTemplate) QueryExample() []byte {
	return mysqlUsersQueryTemplate // This method is kept for backward compatibility
}

func (m MysqlTemplate) UsersSchema() []byte {
	return mysqlUsersSchemaTemplate
}

func (m MysqlTemplate) PostsSchema() []byte {
	return mysqlPostsSchemaTemplate
}

func (m MysqlTemplate) UsersQuery() []byte {
	return mysqlUsersQueryTemplate
}

func (m MysqlTemplate) PostsQuery() []byte {
	return mysqlPostsQueryTemplate
}
