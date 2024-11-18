// Package steps provides utility for creating
// each step of the CLI
package steps

import "github.com/melkeydev/go-blueprint/cmd/flags"

// A StepSchema contains the data that is used
// for an individual step of the CLI
type StepSchema struct {
	StepName string // The name of a given step
	Options  []Item // The slice of each option for a given step
	Headers  string // The title displayed at the top of a given step
	Field    string
}

// Steps contains a slice of steps
type Steps struct {
	Steps map[string]StepSchema
}

// An Item contains the data for each option
// in a StepSchema.Options
type Item struct {
	Flag, Title, Desc string
}

// InitSteps initializes and returns the *Steps to be used in the CLI program
func InitSteps(projectType flags.Framework, databaseType flags.Database) *Steps {
	steps := &Steps{
		map[string]StepSchema{
			"framework": {
				StepName: "Go Project Framework",
				Options: []Item{
					{
						Title: "Standard-library",
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
					{
						Title: "Echo",
						Desc:  "High performance, extensible, minimalist Go web framework",
					},
				},
				Headers: "What framework do you want to use in your Go project?",
				Field:   projectType.String(),
			},
			"driver": {
				StepName: "Go Project Database Driver",
				Options: []Item{
					{
						Title: "Mysql",
						Desc:  "MySQL-Driver for Go's database/sql package",
					},
					{
						Title: "Postgres",
						Desc:  "Go postgres driver for Go's database/sql package"},
					{
						Title: "Sqlite",
						Desc:  "sqlite3 driver conforming to the built-in database/sql interface"},
					{
						Title: "Mongo",
						Desc:  "The MongoDB supported driver for Go."},
					{
						Title: "Redis",
						Desc:  "Redis driver for Go."},
					{
						Title: "Scylla",
						Desc:  "ScyllaDB Enhanced driver from GoCQL."},
					{
						Title: "None",
						Desc:  "Choose this option if you don't wish to install a specific database driver."},
				},
				Headers: "What database driver do you want to use in your Go project?",
				Field:   databaseType.String(),
			},
			"advanced": {
				StepName: "Advanced Features",
				Headers:  "Which advanced features do you want?",
				Options: []Item{
					{
						Flag:  "React",
						Title: "React",
						Desc:  "Use Vite to spin up a React project in TypeScript. This disables selecting HTMX/Templ",
					},
					{
						Flag:  "Htmx",
						Title: "HTMX/Templ",
						Desc:  "Add starter HTMX and Templ files. This disables selecting React",
					},
					{
						Flag:  "GitHubAction",
						Title: "Go Project Workflow",
						Desc:  "Workflow templates for testing, cross-compiling and releasing Go projects",
					},
					{
						Flag:  "Websocket",
						Title: "Websocket endpoint",
						Desc:  "Add a websocket endpoint",
					},
					{
						Flag:  "Tailwind",
						Title: "TailwindCSS",
						Desc:  "A utility-first CSS framework (selecting this will automatically add HTMX unless React is specified)",
					},
					{
						Flag:  "Docker",
						Title: "Docker",
						Desc:  "Dockerfile and docker-compose generic configuration for go project",
					},
				},
			},
			"git": {
				StepName: "Git Repository",
				Headers:  "Which git option would you like to select for your project?",
				Options: []Item{
					{
						Title: "Commit",
						Desc:  "Initialize a new git repository and commit all the changes",
					},
					{
						Title: "Stage",
						Desc:  "Initialize a new git repository but only stage the changes",
					},
					{
						Title: "Skip",
						Desc:  "Proceed without initializing a git repository",
					},
				},
			},
		},
	}

	return steps
}
