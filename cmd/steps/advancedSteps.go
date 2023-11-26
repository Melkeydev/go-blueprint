package steps

import "github.com/melkeydev/go-blueprint/cmd/flags"

// InitSteps initializes and returns the *Steps to be used in the CLI program
func InitAdvancedSteps(projectType flags.Framework, databaseType flags.Database) *Steps {
	advancedSteps := &Steps{
		map[string]StepSchema{
			"htmxTempl": {
				StepName: "HTMX/Templ",
				Headers:  "Add starter HTMX and Templ files?",
			},
		},
	}

	return advancedSteps
}
