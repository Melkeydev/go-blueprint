// Package template provides utility functions that
// help with the templating of created files.
package template

import (
	_ "embed"
)

//go:embed files/main/main.go.tmpl
var mainTemplate []byte

//go:embed files/air.toml.tmpl
var airTomlTemplate []byte

//go:embed files/README.md.tmpl
var readmeTemplate []byte

//go:embed files/makefile.tmpl
var makeTemplate []byte

// MakeTemplate returns a byte slice that represents 
// the default Makefile template.
func MakeTemplate() []byte {
	return []byte(
		`
# Simple Makefile for a Go project

# Build the application
all: build

build:
	@echo "Building..."
	@go build -o main cmd/api/main.go

# Run the application
run:
	@go run cmd/api/main.go

# Test the application
test:
	@echo "Testing..."
	@go test ./...

# Clean the binary
clean:
	@echo "Cleaning..."
	@rm -f main

# Live Reload
watch:
	@if [ -x "$(GOPATH)/bin/air" ]; then \
	    "$(GOPATH)/bin/air"; \
		@echo "Watching...";\
	else \
	    read -p "air is not installed. Do you want to install it now? (y/n) " choice; \
	    if [ "$$choice" = "y" ]; then \
			go install github.com/cosmtrek/air@latest; \
	        "$(GOPATH)/bin/air"; \
				@echo "Watching...";\
	    else \
	        echo "You chose not to install air. Exiting..."; \
	        exit 1; \
	    fi; \
	fi

.PHONY: all build run test clean
		`)
}


func GitIgnoreTemplate() []byte {
	return []byte(
		`
# Binaries for programs and plugins
*.exe
*.exe~
*.dll
*.so
*.dylib

# Test binary, built with "go test -c"
*.test

# Output of the go coverage tool, specifically when used with LiteIDE
*.out

# Dependency directories (remove the comment below to include it)
# vendor/

# Go workspace file
go.work
tmp/

# IDE specific files
.vscode
.idea
		`)
}

func AirTomlTemplate() []byte {
	return airTomlTemplate
}


// ReadmeTemplate returns a byte slice that represents
// the default README.md file template.
func ReadmeTemplate() []byte {
	return []byte(
		`
# Project {{.ProjectName}}

One Paragraph of project description goes here

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.

## MakeFile

run all make commands with clean tests
` + "```bash" + `
make all build
` + "```" + `

build the application
` + "```bash" + `
make build
` + "```" + `

run the application
` + "```bash" + `
make run
` + "```" + `

live reload the application
` + "```bash" + `
make watch
` + "```" + `

run the test suite
` + "```bash" + `
make test
` + "```" + `

clean up binary from the last build
` + "```bash" + `
make clean
` + "```" + `
	`)

}
