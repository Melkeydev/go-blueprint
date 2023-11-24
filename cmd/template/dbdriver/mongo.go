package dbdriver

import (
	_ "embed"
)

type MongoTemplate struct{}

//go:embed files/service/mongo.tmpl
var mongoServiceTemplate []byte

//go:embed files/env/mongo.tmpl
var mongoEnvTemplate []byte

//go:embed files/env/mongo.example.tmpl
var mongoEnvExampleTemplate []byte

func (m MongoTemplate) Service() []byte {
	return mongoServiceTemplate
}

func (m MongoTemplate) Env() []byte {
	return mongoEnvTemplate
}

func (m MongoTemplate) EnvExample() []byte {
	return mongoEnvExampleTemplate
}