package advanced

import (
	_ "embed"
)

//go:embed files/redis/redis_interface.tmpl
var redisInterfaceTemplate []byte

//go:embed files/redis/redis_implementation.tmpl
var redisImplementationTemplate []byte

//go:embed files/redis/redis_config.tmpl
var redisConfigTemplate []byte

//go:embed files/redis/redis_env.tmpl
var redisEnvTemplate []byte

//go:embed files/redis/example_service.tmpl
var redisExampleServiceTemplate []byte

func RedisInterface() []byte {
	return redisInterfaceTemplate
}

func RedisImplementation() []byte {
	return redisImplementationTemplate
}

func RedisConfig() []byte {
	return redisConfigTemplate
}

func RedisEnv() []byte {
	return redisEnvTemplate
}

func RedisExampleService() []byte {
	return redisExampleServiceTemplate
}
