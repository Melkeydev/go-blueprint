package template

import (
	_ "embed"
)

//go:embed files/github/github_action_goreliser.yml.tmpl
var gitHubActionBuildTemplate []byte

//go:embed files/github/github_action_gotest.yml.tmpl
var gitHubActionTestTemplate []byte

type GitHubActionTemplate struct{}

func (a GitHubActionTemplate) Action1() []byte {
	return gitHubActionBuildTemplate
}

func (a GitHubActionTemplate) Action2() []byte {
	return gitHubActionTestTemplate
}