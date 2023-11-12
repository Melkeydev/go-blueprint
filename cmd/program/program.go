// Package program provides the
// main functionality of Blueprint
package program

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	tpl "github.com/melkeydev/go-blueprint/cmd/template"
	"github.com/melkeydev/go-blueprint/cmd/utils"
	"github.com/spf13/cobra"
	"html/template"
	"log"
	"os"
	"strings"
)

// A Project contains the data for the project folder
// being created, and methods that help with that process
type Project struct {
	ProjectName  string
	Exit         bool
	AbsolutePath string
	ProjectType  string
	FrameworkMap map[string]Framework
}

// A Framework contains the name and templater for a
// given Framework
type Framework struct {
	packageName []string
	templater   Templater
}

// A Templater has the methods that help build the files
// in the Project folder, and is specific to a Framework
type Templater interface {
	Main() []byte
	Server() []byte
	Routes() []byte
}

var (
	chiPackage     = []string{"github.com/go-chi/chi/v5"}
	gorillaPackage = []string{"github.com/gorilla/mux"}
	routerPackage  = []string{"github.com/julienschmidt/httprouter"}
	ginPackage     = []string{"github.com/gin-gonic/gin"}
	fiberPackage   = []string{"github.com/gofiber/fiber/v2"}
	echoPackage    = []string{"github.com/labstack/echo/v4", "github.com/labstack/echo/v4/middleware"}
	cobraPackage   = []string{"github.com/spf13/cobra"}

	cmdApiPath         = "cmd/api"
	cmdRootPath        = "cmd"
	internalServerPath = "internal/server"
)

// ExitCLI checks if the Project has been exited, and closes
// out of the CLI if it has
func (p *Project) ExitCLI(tprogram *tea.Program) {
	if p.Exit {
		// logo render here
		if err := tprogram.ReleaseTerminal(); err != nil {
			log.Fatal(err)
		}
		os.Exit(1)
	}
}

// createFrameWorkMap adds the current supported
// Frameworks into a Project's FrameworkMap
func (p *Project) createFrameworkMap() {
	p.FrameworkMap["chi"] = Framework{
		packageName: chiPackage,
		templater:   tpl.ChiTemplates{},
	}

	p.FrameworkMap["standard library"] = Framework{
		packageName: []string{},
		templater:   tpl.StandardLibTemplate{},
	}

	p.FrameworkMap["gin"] = Framework{
		packageName: ginPackage,
		templater:   tpl.GinTemplates{},
	}

	p.FrameworkMap["fiber"] = Framework{
		packageName: fiberPackage,
		templater:   tpl.FiberTemplates{},
	}

	p.FrameworkMap["gorilla/mux"] = Framework{
		packageName: gorillaPackage,
		templater:   tpl.GorillaTemplates{},
	}

	p.FrameworkMap["httprouter"] = Framework{
		packageName: routerPackage,
		templater:   tpl.RouterTemplates{},
	}

	p.FrameworkMap["echo"] = Framework{
		packageName: echoPackage,
		templater:   tpl.EchoTemplates{},
	}

	p.FrameworkMap["cobra"] = Framework{
		packageName: cobraPackage,
		templater:   tpl.CobraTemplates{},
	}
}

// CreateMainFile creates the project folders and files,
// and writes to them depending on the selected options
func (p *Project) CreateMainFile() error {
	// check if AbsolutePath exists
	if _, err := os.Stat(p.AbsolutePath); os.IsNotExist(err) {
		// create directory
		if err := os.Mkdir(p.AbsolutePath, 0754); err != nil {
			log.Printf("Could not create directory: %v", err)
			return err
		}
	}

	p.ProjectName = strings.TrimSpace(p.ProjectName)

	// Create a new directory with the project name
	if _, err := os.Stat(fmt.Sprintf("%s/%s", p.AbsolutePath, p.ProjectName)); os.IsNotExist(err) {
		err := os.MkdirAll(fmt.Sprintf("%s/%s", p.AbsolutePath, p.ProjectName), 0751)
		if err != nil {
			log.Printf("Error creating root project directory %v\n", err)
			return err
		}
	}

	projectPath := fmt.Sprintf("%s/%s", p.AbsolutePath, p.ProjectName)

	// Create the map for our program
	p.createFrameworkMap()

	// Create go.mod
	err := utils.InitGoMod(p.ProjectName, projectPath)
	if err != nil {
		log.Printf("Could not initialize go.mod in new project %v\n", err)
		cobra.CheckErr(err)
	}

	// Install the correct package for the selected framework
	if p.ProjectType != "standard library" {
		err = utils.GoGetPackage(projectPath, p.FrameworkMap[p.ProjectType].packageName)
		if err != nil {
			log.Printf("Could not install go dependency for the chosen framework %v\n", err)
			cobra.CheckErr(err)
		}
	}

	makeFile, err := os.Create(fmt.Sprintf("%s/Makefile", projectPath))
	if err != nil {
		cobra.CheckErr(err)
		return err
	}

	defer makeFile.Close()

	// inject makefile template
	makeFileTemplate := template.Must(template.New("makefile").Parse(string(tpl.MakeTemplate())))
	err = makeFileTemplate.Execute(makeFile, p)
	if err != nil {
		return err
	}

	readmeFile, err := os.Create(fmt.Sprintf("%s/README.md", projectPath))
	if err != nil {
		cobra.CheckErr(err)
		return err
	}

	defer readmeFile.Close()

	// inject readme template
	readmeFileTemplate := template.Must(template.New("readme").Parse(string(tpl.ReadmeTemplate())))
	err = readmeFileTemplate.Execute(readmeFile, p)
	if err != nil {
		return err
	}

	// We only want a server.go and routes.go file if the user is not using cobra
	// We also want to create a cmd/api/main.go file if the user is not using cobra
	if p.ProjectType != "cobra" {
		err = p.CreatePath(cmdApiPath, projectPath)
		if err != nil {
			log.Printf("Error creating path: %s", projectPath)
			cobra.CheckErr(err)
			return err
		}

		err = p.CreateFileWithInjection(cmdApiPath, projectPath, "main.go", "main")
		if err != nil {
			cobra.CheckErr(err)
			return err
		}

		err = p.CreatePath(internalServerPath, projectPath)
		if err != nil {
			log.Printf("Error creating path: %s", internalServerPath)
			cobra.CheckErr(err)
			return err
		}

		err = p.CreateFileWithInjection(internalServerPath, projectPath, "server.go", "server")
		if err != nil {
			log.Printf("Error injecting server.go file: %v", err)
			cobra.CheckErr(err)
			return err
		}

		err = p.CreateFileWithInjection(internalServerPath, projectPath, "routes.go", "routes")
		if err != nil {
			log.Printf("Error injecting routes.go file: %v", err)
			cobra.CheckErr(err)
			return err
		}
	} else {
		// we want to create a cmd/root.go file if the user is using cobra
		err = p.CreatePath(cmdRootPath, projectPath)
		if err != nil {
			log.Printf("Error creating path: %s", projectPath)
			cobra.CheckErr(err)
			return err
		}

		err = p.CreateFileWithInjection(cmdRootPath, projectPath, "root.go", "root")
		if err != nil {
			log.Printf("Error injecting root.go file: %v", err)
			cobra.CheckErr(err)
			return err
		}

		err = p.CreateFileWithInjection("/", projectPath, "main.go", "main")
		if err != nil {
			log.Printf("Error injecting main.go file: %v", err)
			cobra.CheckErr(err)
			return err
		}
	}

	// Initialize git repo
	err = utils.ExecuteCmd("git", []string{"init"}, projectPath)
	if err != nil {
		log.Printf("Error initializing git repo: %v", err)
		cobra.CheckErr(err)
		return err
	}
	// Create gitignore
	gitignoreFile, err := os.Create(fmt.Sprintf("%s/.gitignore", projectPath))
	if err != nil {
		cobra.CheckErr(err)
		return err
	}
	defer gitignoreFile.Close()

	// inject gitignore template
	gitignoreTemplate := template.Must(template.New(".gitignore").Parse(string(tpl.GitIgnoreTemplate())))
	err = gitignoreTemplate.Execute(gitignoreFile, p)
	if err != nil {
		return err
	}

	// Create .air.toml file
	airTomlFile, err := os.Create(fmt.Sprintf("%s/.air.toml", projectPath))
	if err != nil {
		cobra.CheckErr(err)
		return err
	}

	defer airTomlFile.Close()

	// inject air.toml template
	airTomlTemplate := template.Must(template.New("airtoml").Parse(string(tpl.AirTomlTemplate())))
	err = airTomlTemplate.Execute(airTomlFile, p)
	if err != nil {
		return err
	}

	err = utils.GoFmt(projectPath)
	if err != nil {
		log.Printf("Could not gofmt in new project %v\n", err)
		cobra.CheckErr(err)
		return err
	}
	return nil
}

// CreatePath creates the given directory in the projectPath
func (p *Project) CreatePath(pathToCreate string, projectPath string) error {
	if _, err := os.Stat(fmt.Sprintf("%s/%s", projectPath, pathToCreate)); os.IsNotExist(err) {
		err := os.MkdirAll(fmt.Sprintf("%s/%s", projectPath, pathToCreate), 0751)
		if err != nil {
			log.Printf("Error creating directory %v\n", err)
			return err
		}
	}

	return nil
}

// CreateFileWithInjection creates the given file at the
// project path, and injects the appropriate template
func (p *Project) CreateFileWithInjection(pathToCreate string, projectPath string, fileName string, methodName string) error {
	createdFile, err := os.Create(fmt.Sprintf("%s/%s/%s", projectPath, pathToCreate, fileName))
	if err != nil {
		return err
	}

	defer createdFile.Close()

	switch methodName {
	case "main":
		createdTemplate := template.Must(template.New(fileName).Parse(string(p.FrameworkMap[p.ProjectType].templater.Main())))
		err = createdTemplate.Execute(createdFile, p)
	case "server":
		createdTemplate := template.Must(template.New(fileName).Parse(string(p.FrameworkMap[p.ProjectType].templater.Server())))
		err = createdTemplate.Execute(createdFile, p)
	case "routes":
		createdTemplate := template.Must(template.New(fileName).Parse(string(p.FrameworkMap[p.ProjectType].templater.Routes())))
		err = createdTemplate.Execute(createdFile, p)
	case "root":
		// Doing this directly because Cobra is the only framework that uses this template and the others shouldn't need to implement it
		createdTemplate := template.Must(template.New(fileName).Parse(string(tpl.MakeCobraCMDRoot())))
		err = createdTemplate.Execute(createdFile, p)
	}

	if err != nil {
		return err
	}

	return nil
}
