package template

// defines the files to be created

func MainTemplate() []byte {
	return []byte(`
package main

import "fmt"

func main() {
	fmt.Println("Hello World")
}
`)
}

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
