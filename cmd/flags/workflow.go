package flags

import (
	"fmt"
	"strings"
)

type Workflow string

// These are all the current workflows supported. If you want to add one, you
// can simply copy and past a line here. Do not forget to also add it into the
// AllowedWorkflows slice too!
const (
	GitHubAction    Workflow = "githubaction"
	none     		Workflow = "none"
)

var AllowedWorkflows = []string{string(GitHubAction), string(none)}

func (f Workflow) String() string {
	return string(f)
}

func (f *Workflow) Type() string {
	return "Workflow"
}

func (f *Workflow) Set(value string) error {
	// Contains isn't available in 1.20 yet
	// if AllowedWorkflows.Contains(value) {
	for _, workflow := range AllowedWorkflows {
		if workflow == value {
			*f = Workflow(value)
			return nil
		}
	}

	return fmt.Errorf("Workflow to use. Allowed values: %s", strings.Join(AllowedWorkflows, ", "))
}