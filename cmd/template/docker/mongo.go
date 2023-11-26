package docker

import (
	_ "embed"
)

type MongoDockerTemplate struct{}

//go:embed files/docker-compose/mongo.tmpl
var mongoDockerTemplate []byte

func (m MongoDockerTemplate) Docker() []byte {
	return mongoDockerTemplate
}