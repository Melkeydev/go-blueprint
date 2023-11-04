package program

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/melkeydev/go-blueprint/cmd/template/DBDriver"
	"github.com/melkeydev/go-blueprint/cmd/template/framework"
	"github.com/melkeydev/go-blueprint/cmd/utils"
	"github.com/spf13/cobra"
	"html/template"
	"log"
	"os"
)

type Project struct {
	ProjectName  string
	Exit         bool
	AbsolutePath string
	ProjectType  string
	DBDriver     string
	FrameworkMap map[string]Framework
	DBDriverMap  map[string]Driver
}

type Framework struct {
	packageName string
	templater   Templater
}

type Driver struct {
	packageName string
	templater   DBDriverTemplater
}

type Templater interface {
	Main() []byte
	Server() []byte
	Routes() []byte
}

type DBDriverTemplater interface {
	Service() []byte
}

const (
	chiPackage     = "github.com/go-chi/chi/v5"
	gorillaPackage = "github.com/gorilla/mux"
	routerPackage  = "github.com/julienschmidt/httprouter"
	ginPackage     = "github.com/gin-gonic/gin"
	fiberPackage   = "github.com/gofiber/fiber/v2"

	mysqlDriver    = "github.com/go-sql-driver/mysql"
	postgresDriver = "github.com/lib/pq"
	sqliteDriver   = "github.com/mattn/go-sqlite3"
	mongoDriver    = "go.mongodb.org/mongo-driver"

	cmdApiPath         = "cmd/api"
	internalServerPath = "internal/server"
	service            = "services"
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
		templater:   framework.ChiTemplates{},
	}

	p.FrameworkMap["standard library"] = Framework{
		packageName: "",
		templater:   framework.StandardLibTemplate{},
	}

	p.FrameworkMap["gin"] = Framework{
		packageName: ginPackage,
		templater:   framework.GinTemplates{},
	}

	p.FrameworkMap["fiber"] = Framework{
		packageName: fiberPackage,
		templater:   framework.FiberTemplates{},
	}

	p.FrameworkMap["gorilla/mux"] = Framework{
		packageName: gorillaPackage,
		templater:   framework.GorillaTemplates{},
	}

	p.FrameworkMap["httprouter"] = Framework{
		packageName: routerPackage,
		templater:   framework.RouterTemplates{},
	}
}

func (p *Project) createDBDriverMap() {
	p.DBDriverMap["mysql"] = Driver{
		packageName: mysqlDriver,
		templater:   DBDriver.MysqlTemplate{},
	}
	p.DBDriverMap["postgres"] = Driver{
		packageName: postgresDriver,
		templater:   DBDriver.PostgresTemplate{},
	}
	p.DBDriverMap["sqlite"] = Driver{
		packageName: sqliteDriver,
		templater:   DBDriver.SqliteTemplate{},
	}
	p.DBDriverMap["mongo"] = Driver{
		packageName: mongoDriver,
		templater:   DBDriver.MysqlTemplate{},
	}
}

func (p *Project) CreateMainFile() error {
	// check if AbsolutePath exists
	if _, err := os.Stat(p.AbsolutePath); os.IsNotExist(err) {
		// create directory
		if err := os.Mkdir(p.AbsolutePath, 0754); err != nil {
			log.Printf("Could not create directory: %v", err)
			return err
		}
	}

	// First lets create a new director with the project name
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

	// Create go mod
	err := utils.InitGoMod(p.ProjectName, projectPath)
	if err != nil {
		log.Printf("Could not init go mod in new project %v\n", err)
		cobra.CheckErr(err)
	}

	// We need to install the correct package
	if p.ProjectType != "standard library" {
		err = utils.GoGetPackage(projectPath, p.FrameworkMap[p.ProjectType].packageName)
		if err != nil {
			log.Printf("Could not install go dependency for chosen framework %v\n", err)
			cobra.CheckErr(err)
		}
	}

	if p.DBDriver != "None" {
		// Create the map for our program
		p.createDBDriverMap()

		err = utils.GoGetPackage(projectPath, p.DBDriverMap[p.DBDriver].packageName)
		if err != nil {
			log.Printf("Could not install go dependency for chosen driver %v\n", err)
			cobra.CheckErr(err)
		}

		err = p.CreatePath(service, projectPath)
		if err != nil {
			log.Printf("Error creating path: %s", service)
			cobra.CheckErr(err)
			return err
		}

		err = p.CreateFileWithInjection(service, projectPath, "service.go", "services")
		if err != nil {
			log.Printf("Error injecting server.go file: %v", err)
			cobra.CheckErr(err)
			return err
		}
	}

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

	makeFile, err := os.Create(fmt.Sprintf("%s/Makefile", projectPath))
	if err != nil {
		cobra.CheckErr(err)
		return err
	}

	defer makeFile.Close()

	// inject makefile template
	makeFileTemplate := template.Must(template.New("makefile").Parse(string(framework.MakeTemplate())))
	err = makeFileTemplate.Execute(makeFile, p)
	if err != nil {
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

	return nil
}

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

	case "services":
		createdTemplate := template.Must(template.New(fileName).Parse(string(p.DBDriverMap[p.DBDriver].templater.Service())))
		err = createdTemplate.Execute(createdFile, p)
	}

	if err != nil {
		return err
	}

	return nil
}
