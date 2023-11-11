// Package steps provides utility for creating
// each step of the CLI
package steps

// A StepSchema contains the data that is used
// for an individual step of the CLI
type StepSchema struct {
	StepName string  // The name of a given step
	Options  []Item  // The slice of each option for a given step
	Headers  string  // The title displayed at the top of a given step
}

// Steps contains a slice of steps
type Steps struct {
	Steps map[string]StepSchema
}

// An Item contains the data for each option
// in a StepSchema.Options
type Item struct {
	Title, Desc string
}

// InitSteps initializes and returns the *Steps to be used in the CLI program
func InitSteps() *Steps {
	steps := &Steps{
		map[string]StepSchema{
			"framework": {
				StepName: "Go Project Framework",
				Options: []Item{
					{
						Title: "Standard library",
						Desc:  "The built-in Go standard library HTTP package",
					},
					{
						Title: "Chi",
						Desc:  "A lightweight, idiomatic and composable router for building Go HTTP services",
					},
					{
						Title: "Gin",
						Desc:  "Features a martini-like API with performance that is up to 40 times faster thanks to httprouter",
					},
					{
						Title: "Fiber",
						Desc:  "An Express inspired web framework built on top of Fasthttp",
					},
					{
						Title: "Gorilla/Mux",
						Desc:  "Package gorilla/mux implements a request router and dispatcher for matching incoming requests to their respective handler",
					},
					{
						Title: "HttpRouter",
						Desc:  "HttpRouter is a lightweight high performance HTTP request router for Go",
					},
					{Title: "Echo",
						Desc: "High performance, extensible, minimalist Go web framework",
					},
				},
				Headers: "What framework do you want to use in your Go project?",
			},
			"cicd": {
				StepName: "Go Project CICD Pipline",
				Options: []Item{
					{
						Title: "Jenkins",
						Desc:  "Jenkins pipline with Dockerfile for SSH Agnet",
					},
			
					{
						Title: "None",
						Desc:  "Choose this option if you don't want to use a CI/CD pipeline."},
				},
				Headers: "What CICD pipline do you want to use in your Go project?",
			},
		},
	}

	return steps
}
