package testcontainers

import (
	_ "embed"
)

type RedisTestcontainersTemplate struct{}

//go:embed files/redis.tmpl
var redisTestcontainersTemplate []byte

func (r RedisTestcontainersTemplate) Testcontainers() []byte {
	return redisTestcontainersTemplate
}
