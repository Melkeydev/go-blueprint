package advanced

import (
	_ "embed"
)

//go:embed files/docker/dockerfile.tmpl
var dockerfileTemplate []byte

//go:embed files/docker/docker_compose.yml.tmpl
var dockerComposeTemplate []byte

func Dockerfile() []byte {
	return dockerfileTemplate
}

func DockerCompose() []byte {
	return dockerComposeTemplate
}
