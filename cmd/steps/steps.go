package steps

import textinput "github.com/melkeydev/go-blueprint/cmd/ui/textinput"

type StepSchema struct {
	StepName string
	Options  []string
	Headers  string
	Field    *string
}

type Steps struct {
	Steps []StepSchema
}

type Options struct {
	ProjectName *textinput.Output
	ProjectType string
}

func InitSteps(options *Options) *Steps {
	steps := &Steps{
		[]StepSchema{
			{
				StepName: "Go Project Framework",
				Options:  []string{"standard lib", "chi", "gin", "fiber", "gorilla/mux", "httpRouter"},
				Headers:  "What framework do you want to use in your Go project?",
				Field:    &options.ProjectType,
			},
		},
	}

	return steps
}
