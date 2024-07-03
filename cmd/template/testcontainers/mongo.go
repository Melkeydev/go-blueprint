package testcontainers

import (
	_ "embed"
)

type MongoTestcontainersTemplate struct{}

//go:embed files/mongo.tmpl
var mongoTestcontainersTemplate []byte

func (m MongoTestcontainersTemplate) Testcontainers() []byte {
	return mongoTestcontainersTemplate
}
