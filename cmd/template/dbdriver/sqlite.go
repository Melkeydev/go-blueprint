package dbdriver

import (
	_ "embed"
)

type SqliteTemplate struct{}

//go:embed files/service/sqlite.tmpl
var sqliteServiceTemplate []byte

//go:embed files/env/sqlite.tmpl
var sqliteEnvTemplate []byte

//go:embed files/sqlc/sqlc.yaml.tmpl
var sqliteSqlcConfigTemplate []byte

//go:embed files/schema/users_sqlite.sql.tmpl
var sqliteUsersSchemaTemplate []byte

//go:embed files/schema/posts_sqlite.sql.tmpl
var sqlitePostsSchemaTemplate []byte

//go:embed files/query/users/sqlite.sql.tmpl
var sqliteUsersQueryTemplate []byte

//go:embed files/query/posts/sqlite.sql.tmpl
var sqlitePostsQueryTemplate []byte

func (m SqliteTemplate) Service() []byte {
	return sqliteServiceTemplate
}

func (m SqliteTemplate) Env() []byte {
	return sqliteEnvTemplate
}

func (m SqliteTemplate) Tests() []byte {
	return []byte{}
}

func (m SqliteTemplate) SqlcConfig() []byte {
	return sqliteSqlcConfigTemplate
}

func (m SqliteTemplate) SchemaExample() []byte {
	return sqliteUsersSchemaTemplate // This method is kept for backward compatibility
}

func (m SqliteTemplate) QueryExample() []byte {
	return sqliteUsersQueryTemplate // This method is kept for backward compatibility
}

func (m SqliteTemplate) UsersSchema() []byte {
	return sqliteUsersSchemaTemplate
}

func (m SqliteTemplate) PostsSchema() []byte {
	return sqlitePostsSchemaTemplate
}

func (m SqliteTemplate) UsersQuery() []byte {
	return sqliteUsersQueryTemplate
}

func (m SqliteTemplate) PostsQuery() []byte {
	return sqlitePostsQueryTemplate
}
