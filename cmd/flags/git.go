package flags

import (
	"fmt"
	"strings"
)

type Git string

const (
	Commit = "commit"
	Stage  = "stage"
	NoGit  = "none"
)

var AllowedGitsOptions = []string{string(Commit), string(Stage), string(None)}

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
