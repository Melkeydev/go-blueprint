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
}

func InitSteps(options *Options) *Steps {
	steps := &Steps{
		[]StepSchema{
			{
				StepName: "Go Project Framework",
				Options: []Item{
					{
						Title: "Standard library",
						Desc:  "The built-in Go standard library HTTP package",
					},
					{Title: "Chi",
						Desc: "A lightweight, idiomatic and composable router for building Go HTTP services",
					},
					{Title: "Gin",
						Desc: "Features a martini-like API with performance that is up to 40 times faster thanks to httprouter",
					},
					{Title: "Fiber",
						Desc: "An Express inspired web framework built on top of Fasthttp",
					},
					{Title: "Gorilla/Mux",
						Desc: "Package gorilla/mux implements a request router and dispatcher for matching incoming requests to their respective handler",
					},
					{Title: "HttpRouter",
						Desc: "HttpRouter is a lightweight high performance HTTP request router for Go",
					},
				},
				Headers: "What framework do you want to use in your Go project?",
				Field:   &options.ProjectType,
			},
			{
				StepName: "Go Project Database",
				Options: []Item{
					{
						Title: "Mysql",
						Desc:  "The built-in Go standard library HTTP package",
					},
					{Title: "Postgres",
						Desc: "A lightweight, idiomatic and composable router for building Go HTTP services",
					},
					{Title: "Sqlite",
						Desc: "Features a martini-like API with performance that is up to 40 times faster thanks to httprouter",
					},
					{Title: "Mongodb",
						Desc: "An Express inspired web framework built on top of Fasthttp",
					},
				},
				Headers: "What database do you want to use in your Go project?",
				Field:   &options.ProjectType,
			},
		},
	}

	return steps
}
