// Package program provides the
// main functionality of Blueprint
package program

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	tpl "github.com/melkeydev/go-blueprint/cmd/template"
	"github.com/melkeydev/go-blueprint/cmd/template/dbdriver"
	"github.com/melkeydev/go-blueprint/cmd/template/framework"
	"github.com/melkeydev/go-blueprint/cmd/utils"
	"github.com/spf13/cobra"
)

// A Project contains the data for the project folder
// being created, and methods that help with that process
type Project struct {
	ProjectName  string
	Exit         bool
	AbsolutePath string
	ProjectType  string
	DBDriver     string
	FrameworkMap map[string]Framework
	DBDriverMap  map[string]Driver
}

// A Framework contains the name and templater for a
// given Framework
type Framework struct {
	packageName []string
	templater   Templater
}

type Driver struct {
	packageName []string
	templater   DBDriverTemplater
}

// A Templater has the methods that help build the files
// in the Project folder, and is specific to a Framework
type Templater interface {
	Main() []byte
	Server() []byte
	Routes() []byte
	RoutesWithDB() []byte
	ServerWithDB() []byte
	TestHandler() []byte
}

type DBDriverTemplater interface {
	Service() []byte
	Env() []byte
}

// Move these
var (
	chiPackage     = []string{"github.com/go-chi/chi/v5"}
	gorillaPackage = []string{"github.com/gorilla/mux"}
	routerPackage  = []string{"github.com/julienschmidt/httprouter"}
	ginPackage     = []string{"github.com/gin-gonic/gin"}
	fiberPackage   = []string{"github.com/gofiber/fiber/v2"}
	echoPackage    = []string{"github.com/labstack/echo/v4", "github.com/labstack/echo/v4/middleware"}

	mysqlDriver    = []string{"github.com/go-sql-driver/mysql"}
	postgresDriver = []string{"github.com/lib/pq"}
	sqliteDriver   = []string{"github.com/mattn/go-sqlite3"}
	mongoDriver    = []string{"go.mongodb.org/mongo-driver"}

	godotenvPackage = []string{"github.com/joho/godotenv"}
)

const (
	root                 = "/"
	cmdApiPath           = "cmd/api"
	internalServerPath   = "internal/server"
	internalDatabasePath = "internal/database"
	testHandlerPath      = "tests"
	standardLib          = "standard = library"
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

// Move this
// createFrameWorkMap adds the current supported
// Frameworks into a Project's FrameworkMap
func (p *Project) createFrameworkMap() {
	p.FrameworkMap["chi"] = Framework{
		packageName: chiPackage,
		templater:   framework.ChiTemplates{},
	}

	p.FrameworkMap["standard library"] = Framework{
		packageName: []string{},
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

	p.FrameworkMap["echo"] = Framework{
		packageName: echoPackage,
		templater:   framework.EchoTemplates{},
	}
}

func (p *Project) createDBDriverMap() {
	p.DBDriverMap["mysql"] = Driver{
		packageName: mysqlDriver,
		templater:   dbdriver.MysqlTemplate{},
	}
	p.DBDriverMap["postgres"] = Driver{
		packageName: postgresDriver,
		templater:   dbdriver.PostgresTemplate{},
	}
	p.DBDriverMap["sqlite"] = Driver{
		packageName: sqliteDriver,
		templater:   dbdriver.SqliteTemplate{},
	}
	p.DBDriverMap["mongo"] = Driver{
		packageName: mongoDriver,
		templater:   dbdriver.MongoTemplate{},
	}
}

// CreateMainFile creates the project folders and files,
// and writes to them depending on the selected options
func (p *Project) CreateMainFile() error {
	projectPath := fmt.Sprintf("%s/%s", p.AbsolutePath, p.ProjectName)

	// check if AbsolutePath exists
	if err := utils.CreateDirectoryIfNotExist(p.AbsolutePath); err != nil {
		cobra.CheckErr(err)
		return err
	}

	p.ProjectName = strings.TrimSpace(p.ProjectName)

	// Create a new directory with the project name
	if err := utils.CreateDirectoryIfNotExist(fmt.Sprintf("%s/%s", p.AbsolutePath, p.ProjectName)); err != nil {
		cobra.CheckErr(err)
	}

	// Create the framework map for our program
	p.createFrameworkMap()

	if err := initializeProject(projectPath, p); err != nil {
		cobra.CheckErr(err)
	}

	if err := setUpFramework(projectPath, p); err != nil {
		cobra.CheckErr(err)
	}

	var err error
	// stop here

	// Install the correct package for the selected driver
	if p.DBDriver != "none" {
		p.createDBDriverMap()
		err = utils.GoGetPackage(projectPath, p.DBDriverMap[p.DBDriver].packageName)
		if err != nil {
			log.Printf("Could not install go dependency for chosen driver %v\n", err)
			cobra.CheckErr(err)
		}

		err = utils.CreateDirectoryIfNotExist(fmt.Sprintf("%s/%s", projectPath, internalDatabasePath))
		if err != nil {
			log.Printf("Error creating path: %s", internalDatabasePath)
			cobra.CheckErr(err)
			return err
		}

		err = p.CreateFileWithInjection(internalDatabasePath, projectPath, "database.go", "database")
		if err != nil {
			log.Printf("Error injecting server.go file: %v", err)
			cobra.CheckErr(err)
			return err
		}
	}

	// One
	err = utils.CreateDirectoryIfNotExist(fmt.Sprintf("%s/%s", projectPath, cmdApiPath))
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

	// Two
	err = utils.CreateDirectoryIfNotExist(fmt.Sprintf("%s/%s", projectPath, testHandlerPath))
	if err != nil {
		log.Printf("Error creating path: %s", projectPath)
		cobra.CheckErr(err)
		return err
	}

	err = p.CreateFileWithInjection(testHandlerPath, projectPath, "handler_test.go", "tests")
	if err != nil {
		cobra.CheckErr(err)
		return err
	}

	// Three
	err = utils.CreateDirectoryIfNotExist(fmt.Sprintf("%s/%s", projectPath, internalServerPath))
	if err != nil {
		log.Printf("Error creating path: %s", internalServerPath)
		cobra.CheckErr(err)
		return err
	}

	if p.DBDriver != "none" {
		err = p.CreateFileWithInjection(internalServerPath, projectPath, "routes.go", "routesWithDB")
		if err != nil {
			log.Printf("Error injecting routes.go file: %v", err)
			cobra.CheckErr(err)
			return err
		}
		err = p.CreateFileWithInjection(internalServerPath, projectPath, "server.go", "serverWithDB")
		if err != nil {
			log.Printf("Error injecting server.go file: %v", err)
			cobra.CheckErr(err)
			return err
		}
	} else {
		err = p.CreateFileWithInjection(internalServerPath, projectPath, "routes.go", "routes")
		if err != nil {
			log.Printf("Error injecting routes.go file: %v", err)
			cobra.CheckErr(err)
			return err
		}
		err = p.CreateFileWithInjection(internalServerPath, projectPath, "server.go", "server")
		if err != nil {
			log.Printf("Error injecting server.go file: %v", err)
			cobra.CheckErr(err)
			return err
		}
	}

	if err := createCommonFiles(projectPath, p); err != nil {
		cobra.CheckErr(err)
	}

	// These can be saved for the end
	err = utils.GoFmt(projectPath)
	if err != nil {
		log.Printf("Could not gofmt in new project %v\n", err)
		cobra.CheckErr(err)
		return err
	}

	err = utils.GoTidy(projectPath)
	if err != nil {
		log.Printf("Could not go tidy in new project %v\n", err)
		cobra.CheckErr(err)
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
	case "serverWithDB":
		createdTemplate := template.Must(template.New(fileName).Parse(string(p.FrameworkMap[p.ProjectType].templater.ServerWithDB())))
		err = createdTemplate.Execute(createdFile, p)
	case "routes":
		createdTemplate := template.Must(template.New(fileName).Parse(string(p.FrameworkMap[p.ProjectType].templater.Routes())))
		err = createdTemplate.Execute(createdFile, p)
	case "routesWithDB":
		createdTemplate := template.Must(template.New(fileName).Parse(string(p.FrameworkMap[p.ProjectType].templater.RoutesWithDB())))
		err = createdTemplate.Execute(createdFile, p)
	case "database":
		createdTemplate := template.Must(template.New(fileName).Parse(string(p.DBDriverMap[p.DBDriver].templater.Service())))
		err = createdTemplate.Execute(createdFile, p)
	case "tests":
		createdTemplate := template.Must(template.New(fileName).Parse(string(p.FrameworkMap[p.ProjectType].templater.TestHandler())))
		err = createdTemplate.Execute(createdFile, p)
	}

	if err != nil {
		return err
	}

	return nil
}

// initializeProject initializes go.mod and other initial setups
func initializeProject(projectPath string, p *Project) error {
	// Initialize go.mod
	if err := utils.InitGoMod(p.ProjectName, projectPath); err != nil {
		log.Printf("Could not initialize go.mod in new project %v\n", err)
		return err
	}

	// Initialize git repository
	if err := utils.ExecuteCmd("git", []string{"init"}, projectPath); err != nil {
		log.Printf("Error initializing git repo: %v", err)
		return err
	}

	return nil
}

// createCommonFiles creates files like Makefile, README, etc., and .env file.
func createCommonFiles(projectPath string, p *Project) error {
	commonFiles := map[string][]byte{
		"Makefile":   framework.MakeTemplate(),
		"README.md":  framework.ReadmeTemplate(),
		".gitignore": framework.GitIgnoreTemplate(),
		".air.toml":  framework.AirTomlTemplate(),
	}

	for fileName, templateContent := range commonFiles {
		if err := utils.CreateFileFromTemplate(fmt.Sprintf("%s/%s", projectPath, fileName), templateContent, p); err != nil {
			return err
		}
	}

	// Create .env file
	return createEnvFile(projectPath, p)
}

// createEnvFile creates the .env file based on the project settings.
func createEnvFile(projectPath string, p *Project) error {
	envFilePath := fmt.Sprintf("%s/.env", projectPath)
	createdFile, err := os.Create(envFilePath)
	if err != nil {
		return err
	}
	defer createdFile.Close()

	var createdTemplate *template.Template
	if p.DBDriver != "none" {
		envBytes := [][]byte{
			tpl.GlobalEnvTemplate(),
			p.DBDriverMap[p.DBDriver].templater.Env(),
		}
		createdTemplate = template.Must(template.New(".env").Parse(string(bytes.Join(envBytes, []byte("\n")))))
	} else {
		createdTemplate = template.Must(template.New(".env").Parse(string(tpl.GlobalEnvTemplate())))
	}

	return createdTemplate.Execute(createdFile, p)
}

func setUpFramework(projectPath string, p *Project) error {
	if p.ProjectType != standardLib {
		err := utils.GoGetPackage(projectPath, p.FrameworkMap[p.ProjectType].packageName)
		if err != nil {
			log.Printf("Could not install go dependency for the chosen framework %v\n", err)
			return err
		}
	}

	// Install the godotenv package
	err := utils.GoGetPackage(projectPath, godotenvPackage)
	if err != nil {
		log.Printf("Could not install go dependency %v\n", err)
		return err
	}

	return nil
}
