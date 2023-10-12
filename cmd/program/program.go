package program

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	tea "github.com/charmbracelet/bubbletea"
)

type Project struct {
	ProjectName  string
	Exit         bool
	AbsolutePath string
}

func (p *Project) ExitCLI(tprogram *tea.Program) {
	if p.Exit {
		// logo render here
		tprogram.ReleaseTerminal()
		os.Exit(1)
	}
}

func (p *Project) CreateAPIProject() {
	appDir := filepath.Join(p.AbsolutePath, p.ProjectName)
	if _, err := os.Stat(appDir); os.IsNotExist(err) {
		if err := os.Mkdir(appDir, 0755); err != nil {
			fmt.Printf("Error creating project directory: %v\n", err)
			return
		}
	}

	// Define the Bash script content
	scriptContent := `#!/bin/bash

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

echo "Go project '$app_name' created and initialized with 'go mod init'."`

	// Create a temporary Bash script file
	tempScriptPath := filepath.Join(appDir, "temp_script.sh")
	if err := os.WriteFile(tempScriptPath, []byte(scriptContent), 0755); err != nil {
		fmt.Printf("Error creating temporary script file: %v\n", err)
		return
	}
	defer os.Remove(tempScriptPath) // Clean up the temporary script file

	// Create the command to execute the Bash script
	cmd := exec.Command("bash", tempScriptPath)

	// Set the working directory to appDir
	cmd.Dir = appDir

	// Redirect standard output and error to the current process's standard output and error
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Run the command
	if err := cmd.Run(); err != nil {
		fmt.Printf("Error executing Bash script: %v\n", err)
		return
	}

	fmt.Println("Project structure created successfully in", appDir)
}
