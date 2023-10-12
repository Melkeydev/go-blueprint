#!/bin/bash

app_name="my-go-project"

go mod init "$app_name"

touch Makefile

mkdir -p cmd/api
echo 'package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
' > cmd/api/main.go

echo "Go project '$app_name' created and initialized with 'go mod init'."