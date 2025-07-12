package dbdriver

import (
	_ "embed"
)

type ScyllaTemplate struct{}

//go:embed files/service/scylla.tmpl
var scyllaServiceTemplate []byte

//go:embed files/env/scylla.tmpl
var scyllaEnvTemplate []byte

//go:embed files/tests/scylla.tmpl
var scyllaTestcontainersTemplate []byte

func (r ScyllaTemplate) Service() []byte {
	return scyllaServiceTemplate
}

func (r ScyllaTemplate) Env() []byte {
	return scyllaEnvTemplate
}

func (r ScyllaTemplate) Tests() []byte {
	return scyllaTestcontainersTemplate
}

func (r ScyllaTemplate) SqlcConfig() []byte {
	return []byte{} // ScyllaDB doesn't use sqlc
}

func (r ScyllaTemplate) SchemaExample() []byte {
	return []byte{} // ScyllaDB doesn't use sqlc
}

func (r ScyllaTemplate) QueryExample() []byte {
	return []byte{} // ScyllaDB doesn't use sqlc
}

func (r ScyllaTemplate) UsersSchema() []byte {
	return []byte{} // ScyllaDB doesn't use sqlc
}

func (r ScyllaTemplate) PostsSchema() []byte {
	return []byte{} // ScyllaDB doesn't use sqlc
}

func (r ScyllaTemplate) UsersQuery() []byte {
	return []byte{} // ScyllaDB doesn't use sqlc
}

func (r ScyllaTemplate) PostsQuery() []byte {
	return []byte{} // ScyllaDB doesn't use sqlc
}
