package program

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"text/template"

	tpl "github.com/melkeydev/go-blueprint/cmd/template"
	"github.com/melkeydev/go-blueprint/cmd/template/advanced"
)

// CreatePath creates the given directory in the projectPath
func (p *Project) CreatePath(pathToCreate string, projectPath string) error {
	path := filepath.Join(projectPath, pathToCreate)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.MkdirAll(path, 0o751)
		if err != nil {
			log.Printf("Error creating directory %v\n", err)
			return err
		}
	}

	return nil
}

// CheckOs checks Operation system and generates MakeFile and `go build` command
// Based on Project.Unixbase
func (p *Project) CheckOS() {
	p.OSCheck = make(map[string]bool)

	if runtime.GOOS != "windows" {
		p.OSCheck["UnixBased"] = true
	}
	if runtime.GOOS == "linux" {
		p.OSCheck["linux"] = true
	}
	if runtime.GOOS == "darwin" {
		p.OSCheck["darwin"] = true
	}
}

func checkNpmInstalled() error {
	cmd := exec.Command("npm", "--version")
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("npm is not installed: %w", err)
	}
	return nil
}

// CreateFileWithInjection creates the given file at the
// project path, and injects the appropriate template
func (p *Project) CreateFileWithInjection(pathToCreate string, projectPath string, fileName string, methodName string) error {
	createdFile, err := os.Create(filepath.Join(projectPath, pathToCreate, fileName))
	if err != nil {
		return err
	}

	defer createdFile.Close()

	switch methodName {
	case "main":
		createdTemplate := template.Must(template.New(fileName).Parse(string(p.BackendMap[p.ProjectType].templater.Main())))
		err = createdTemplate.Execute(createdFile, p)
	case "server":
		createdTemplate := template.Must(template.New(fileName).Parse(string(p.BackendMap[p.ProjectType].templater.Server())))
		err = createdTemplate.Execute(createdFile, p)
	case "routes":
		routeFileBytes := p.BackendMap[p.ProjectType].templater.Routes()
		createdTemplate := template.Must(template.New(fileName).Parse(string(routeFileBytes)))
		err = createdTemplate.Execute(createdFile, p)
	case "releaser":
		createdTemplate := template.Must(template.New(fileName).Parse(string(advanced.Releaser())))
		err = createdTemplate.Execute(createdFile, p)
	case "go-test":
		createdTemplate := template.Must(template.New(fileName).Parse(string(advanced.Test())))
		err = createdTemplate.Execute(createdFile, p)
	case "releaser-config":
		createdTemplate := template.Must(template.New(fileName).Parse(string(advanced.ReleaserConfig())))
		err = createdTemplate.Execute(createdFile, p)
	case "database":
		createdTemplate := template.Must(template.New(fileName).Parse(string(p.DBDriverMap[p.DBDriver].templater.Service())))
		err = createdTemplate.Execute(createdFile, p)
	case "db-docker":
		createdTemplate := template.Must(template.New(fileName).Parse(string(p.DockerMap[p.Docker].templater.Docker())))
		err = createdTemplate.Execute(createdFile, p)
	case "integration-tests":
		createdTemplate := template.Must(template.New(fileName).Parse(string(p.DBDriverMap[p.DBDriver].templater.Tests())))
		err = createdTemplate.Execute(createdFile, p)
	case "tests":
		createdTemplate := template.Must(template.New(fileName).Parse(string(p.BackendMap[p.ProjectType].templater.TestHandler())))
		err = createdTemplate.Execute(createdFile, p)
	case "env":
		if p.DBDriver != "none" {

			envBytes := [][]byte{
				tpl.GlobalEnvTemplate(),
				p.DBDriverMap[p.DBDriver].templater.Env(),
			}
			createdTemplate := template.Must(template.New(fileName).Parse(string(bytes.Join(envBytes, []byte("\n")))))
			err = createdTemplate.Execute(createdFile, p)

		} else {
			createdTemplate := template.Must(template.New(fileName).Parse(string(tpl.GlobalEnvTemplate())))
			err = createdTemplate.Execute(createdFile, p)
		}
	}

	if err != nil {
		return err
	}

	return nil
}
