package program

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
	"os/exec"
	"path/filepath"

	tea "github.com/charmbracelet/bubbletea"
	tpl "github.com/melkeydev/go-blueprint/cmd/template"
	"github.com/spf13/cobra"
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

func executeCmd(name string, args []string, dir string) error {
	command := exec.Command(name, args...)
	command.Dir = dir
	var out bytes.Buffer
	command.Stdout = &out
	if err := command.Run(); err != nil {
		return err
	}
	return nil
}

func initGoMod(projectName string, appDir string) {
	if err := executeCmd("go",
		[]string{"mod", "init", projectName},
		appDir); err != nil {
		cobra.CheckErr(err)
	}
}

// things to do:
// create a Makefile
// create project structure

func (p *Project) CreateMainFile() error {
	// check if AbsolutePath exists
	if _, err := os.Stat(p.AbsolutePath); os.IsNotExist(err) {
		// create directory
		if err := os.Mkdir(p.AbsolutePath, 0754); err != nil {
			return err
		}
	}

	// First lets create a new director with the project name
	if _, err := os.Stat(fmt.Sprintf("%s/%s", p.AbsolutePath, p.ProjectName)); os.IsNotExist(err) {
		err := os.MkdirAll(fmt.Sprintf("%s/%s", p.AbsolutePath, p.ProjectName), 0751)
		if err != nil {
			fmt.Printf("Error creating root project directory %v\n", err)
		}
	}

	projectPath := fmt.Sprintf("%s/%s", p.AbsolutePath, p.ProjectName)

	// we need to create a go mod init
	initGoMod(p.ProjectName, projectPath)

	// create /cmd/api
	if _, err := os.Stat(fmt.Sprintf("%s/cmd/api", projectPath)); os.IsNotExist(err) {
		err := os.MkdirAll(fmt.Sprintf("%s/cmd/api", projectPath), 0751)
		if err != nil {
			fmt.Printf("Error creating directory %v\n", err)
		}
	}

	mainFile, err := os.Create(fmt.Sprintf("%s/cmd/api/main.go", projectPath))
	if err != nil {
		return err
	}

	defer mainFile.Close()

	// inject template
	mainTemplate := template.Must(template.New("main").Parse(string(tpl.MainTemplate())))
	err = mainTemplate.Execute(mainFile, p)
	if err != nil {
		return err
	}

	return nil

}

// We want to deprecate this approach
func (p *Project) CreateAPIProject() {
	appDir := filepath.Join(p.AbsolutePath, p.ProjectName)
	if _, err := os.Stat(appDir); os.IsNotExist(err) {
		if err := os.Mkdir(appDir, 0755); err != nil {
			fmt.Printf("Error creating project directory: %v\n", err)
			return
		}
	}

	scriptContent := `#!/bin/bash
		app_name="my-go-project"

		go mod init "$app_name"

		echo "Go project '$app_name' created and initialized with 'go mod init'."`

	tempScriptPath := filepath.Join(appDir, "temp_script.sh")
	if err := os.WriteFile(tempScriptPath, []byte(scriptContent), 0755); err != nil {
		fmt.Printf("Error creating temporary script file: %v\n", err)
		return
	}
	defer os.Remove(tempScriptPath)

	cmd := exec.Command("bash", tempScriptPath)

	cmd.Dir = appDir

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Printf("Error executing Bash script: %v\n", err)
		return
	}

	fmt.Println("Project structure created successfully in", appDir)
}
