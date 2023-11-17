package workflow

import (
	_ "embed"
)

//go:embed files/github/github_action_goreliser.yml.tmpl
var gitHubActionBuildTemplate []byte

//go:embed files/github/github_action_gotest.yml.tmpl
var gitHubActionTestTemplate []byte

// GitHubActionTemplates contains the methods used for building
type GitHubActionTemplate struct{}

func (a GitHubActionTemplate) File_1() []byte {
	return gitHubActionBuildTemplate
}

func (a GitHubActionTemplate) File_2() []byte {
	return gitHubActionTestTemplate
}