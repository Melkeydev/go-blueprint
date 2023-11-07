// Package template provides utility functions that
// help with the templating of created files.
package framework

// MakeHTTPRoutes returns a byte slice that represents
// the default cmd/api/main.go file template.
func MainTemplate() []byte {
	return []byte(`package main

import (
	"{{.ProjectName}}/internal/server"
)

func main() {

	server := server.NewServer()

	err := server.ListenAndServe()
	if err != nil {
		panic("cannot start server")
	}
}
`)
}

// MakeHTTPRoutes returns a byte slice that represents
// the default Makefile.
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

.PHONY: all build run test clean
		`)
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
