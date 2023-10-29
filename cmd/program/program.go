package program

import (
	"fmt"
	"html/template"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	tpl "github.com/melkeydev/go-blueprint/cmd/template"
	"github.com/melkeydev/go-blueprint/cmd/utils"
	"github.com/spf13/cobra"
)

type Project struct {
	ProjectName  string
	Exit         bool
	AbsolutePath string
	ProjectType  string
	FrameworkMap map[string]Framework
}

type Framework struct {
	packageName string
	templater   Templater
}

type Templater interface {
	Main() []byte
	Server() []byte
	Routes() []byte
}

const (
	chiPackage     = "github.com/go-chi/chi/v5"
	gorillaPackage = "github.com/gorilla/mux"
	routerPackage  = "github.com/julienschmidt/httprouter"
	ginPackage     = "github.com/gin-gonic/gin"
	fiberPackage   = "github.com/gofiber/fiber/v2"
)

func (p *Project) ExitCLI(tprogram *tea.Program) {
	if p.Exit {
		// logo render here
		tprogram.ReleaseTerminal()
		os.Exit(1)
	}
}

func (p *Project) createFrameworkMap() {

	p.FrameworkMap["chi"] = Framework{
		packageName: chiPackage,
		templater:   tpl.ChiTemplates{},
	}

	p.FrameworkMap["standard lib"] = Framework{
		packageName: "",
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

	p.FrameworkMap["httpRouter"] = Framework{
		packageName: routerPackage,
		templater:   tpl.RouterTemplates{},
	}
}

// We can clean this up after
// seperate it
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

	// i hate my life
	p.createFrameworkMap()

	// we need to create a go mod init
	utils.InitGoMod(p.ProjectName, projectPath)

	// we need to install the correct package
	if p.ProjectType != "standard lib" {
		utils.GoGetPackage(projectPath, p.FrameworkMap[p.ProjectType].packageName)
	}

	err := p.CreatePath("cmd/api", projectPath)
	if err != nil {
		cobra.CheckErr(err)
		return err
	}

	err = p.CreateFileWithInjection("cmd/api", projectPath, "main.go", "main")
	if err != nil {
		cobra.CheckErr(err)
		return err
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

	err = p.CreatePath("internal/server", projectPath)
	if err != nil {
		cobra.CheckErr(err)
		return err
	}

	err = p.CreateFileWithInjection("internal/server", projectPath, "server.go", "server")
	if err != nil {
		cobra.CheckErr(err)
		return err
	}

	err = p.CreateFileWithInjection("internal/server", projectPath, "routes.go", "routes")
	if err != nil {
		cobra.CheckErr(err)
		return err
	}

	return nil
}

// cmd/api
func (p *Project) CreatePath(pathToCreate string, projectPath string) error {
	if _, err := os.Stat(fmt.Sprintf("%s/%s", projectPath, pathToCreate)); os.IsNotExist(err) {
		err := os.MkdirAll(fmt.Sprintf("%s/%s", projectPath, pathToCreate), 0751)
		if err != nil {
			fmt.Printf("Error creating directory %v\n", err)
			return err
		}
	}

	return nil
}

// cmd/api
func (p *Project) CreateFileWithInjection(pathToCreate string, projectPath string, fileName string, methodName string) error {
	createdFile, err := os.Create(fmt.Sprintf("%s/%s/%s", projectPath, pathToCreate, fileName))
	if err != nil {
		return err
	}

	defer createdFile.Close()

	// inject template
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
	}

	if err != nil {
		return err
	}

	return nil
}
