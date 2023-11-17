package dbdriver

import (
	_ "embed"
)

type MongoTemplate struct{}

//go:embed files/service/mongo.tmpl
var mongoServiceTemplate []byte

//go:embed files/env/mongo.tmpl
var mongoEnvTemplate []byte

func (m MongoTemplate) Service() []byte {
	return mongoServiceTemplate
}

func (m MongoTemplate) Env() []byte {
	return mongoEnvTemplate
}
