// Package template provides utility functions that
// help with the templating of created files.
package template

// MakeHTTPRoutes returns a byte slice that represents 
// the default cmd/api/main.go file.
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
