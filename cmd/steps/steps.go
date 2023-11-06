package steps

import textinput "github.com/melkeydev/go-blueprint/cmd/ui/textinput"

type StepSchema struct {
	StepName string
	Options  []Item
	Headers  string
	Field    *string
}

type Steps struct {
	Steps []StepSchema
}

type Item struct {
	Title, Desc string
}

type Options struct {
	ProjectName *textinput.Output
	ProjectType string
	DBDriver    string
}

func InitFrameworkSteps(options *Options) *Steps {
	steps := &Steps{
		[]StepSchema{
			{
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
					{
						Title: "Echo",
						Desc:  "High performance, extensible, minimalist Go web framework",
					},
				},
				Headers: "What framework do you want to use in your Go project?",
				Field:   &options.ProjectType,
			},
		},
	}
	return steps
}

func InitDBDriverSteps(options *Options) *Steps {
	steps := &Steps{
		[]StepSchema{
			{
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
						Title: "None",
						Desc:  "Choose this option if you don't wish to install a specific database driver."},
				},
				Headers: "What database driver do you want to use in your Go project?",
				Field:   &options.DBDriver,
			},
		},
	}

	return steps
}
