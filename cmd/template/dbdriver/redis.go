package dbdriver

import (
	_ "embed"
)

type RedisTemplate struct{}

//go:embed files/service/redis.tmpl
var redisServiceTemplate []byte

//go:embed files/env/redis.tmpl
var redisEnvTemplate []byte

func (r RedisTemplate) Service() []byte {
	return redisServiceTemplate
}

func (r RedisTemplate) Env() []byte {
	return redisEnvTemplate
}
