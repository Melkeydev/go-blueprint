package dbdriver

import (
	_ "embed"
)

type MongoTemplate struct{}

//go:embed files/service/mongo.tmpl
var mongoServiceTemplate []byte

//go:embed files/env/mongo.tmpl
var mongoEnvTemplate []byte

//go:embed files/tests/mongo.tmpl
var mongoTestcontainersTemplate []byte

func (m MongoTemplate) Service() []byte {
	return mongoServiceTemplate
}

func (m MongoTemplate) Env() []byte {
	return mongoEnvTemplate
}

func (m MongoTemplate) Tests() []byte {
	return mongoTestcontainersTemplate
}
