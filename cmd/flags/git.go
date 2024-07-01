package flags

import (
	"fmt"
	"strings"
)

type Git string

const (
	GitCommit Git = "commit"
	GitStage  Git = "stage"
	NoGit     Git = "none"
)

var AllowedGitsOptions = []string{string(GitCommit), string(GitStage), string(NoGit)}

func (f Git) String() string {
	return string(f)
}

func (f *Git) Type() string {
	return "Git"
}

func (f *Git) Set(value string) error {
	for _, gitOption := range AllowedGitsOptions {
		if gitOption == value {
			*f = Git(value)
			return nil
		}
	}

	return fmt.Errorf("Git to use. Allowed values: %s", strings.Join(AllowedGitsOptions, ", "))
}
