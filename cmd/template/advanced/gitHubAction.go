package advanced

import (
	_ "embed"
)

//go:embed files/workflow/github/github_action_goreliser.yml.tmpl
var gitHubActionBuildTemplate []byte

//go:embed files/workflow/github/github_action_gotest.yml.tmpl
var gitHubActionTestTemplate []byte

//go:embed files/workflow/github/github_action_reliser_config.yml.tmpl
var gitHubActionConfigTemplate []byte

func File_1() []byte {
	return gitHubActionBuildTemplate
}

func File_2() []byte {
	return gitHubActionTestTemplate
}

func File_3() []byte {
	return gitHubActionConfigTemplate
}
