package docker

import (
	_ "embed"
)

type RedisDockerTemplate struct{}

//go:embed files/docker-compose/redis.tmpl
var redisDockerTemplate []byte

func (r RedisDockerTemplate) Docker() []byte {
	return redisDockerTemplate
}
