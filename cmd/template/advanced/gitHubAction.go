package advanced

import (
	_ "embed"
)

//go:embed files/workflow/github/github_action_goreleaser.yml.tmpl
var gitHubActionBuildTemplate []byte

//go:embed files/workflow/github/github_action_gotest.yml.tmpl
var gitHubActionTestTemplate []byte

//go:embed files/workflow/github/github_action_releaser_config.yml.tmpl
var gitHubActionConfigTemplate []byte

func Releaser() []byte {
	return gitHubActionBuildTemplate
}

func Test() []byte {
	return gitHubActionTestTemplate
}

func ReleaserConfig() []byte {
	return gitHubActionConfigTemplate
}
