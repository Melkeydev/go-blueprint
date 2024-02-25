// Package program provides the
// main functionality of Blueprint
package program

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"os"
	"path/filepath"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/melkeydev/go-blueprint/cmd/flags"
	tpl "github.com/melkeydev/go-blueprint/cmd/template"
	"github.com/melkeydev/go-blueprint/cmd/template/advanced"
	"github.com/melkeydev/go-blueprint/cmd/template/dbdriver"
	"github.com/melkeydev/go-blueprint/cmd/template/docker"
	"github.com/melkeydev/go-blueprint/cmd/template/framework"
	"github.com/melkeydev/go-blueprint/cmd/utils"
	"github.com/spf13/cobra"
)

// A Project contains the data for the project folder
// being created, and methods that help with that process
type Project struct {
	ProjectName       string
	Exit              bool
	AbsolutePath      string
	ProjectType       flags.Framework
	DBDriver          flags.Database
	Docker            flags.Database
	FrameworkMap      map[flags.Framework]Framework
	DBDriverMap       map[flags.Database]Driver
	DockerMap         map[flags.Database]Docker
	AdvancedOptions   map[string]bool
	AdvancedTemplates AdvancedTemplates
}

type AdvancedTemplates struct {
	TemplateRoutes  template.HTML
	TemplateImports template.HTML
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

type Docker struct {
	packageName []string
	templater   DockerTemplater
}

// A Templater has the methods that help build the files
// in the Project folder, and is specific to a Framework
type Templater interface {
	Main() []byte
	Server() []byte
	Routes() []byte
	TestHandler() []byte
	HtmxTemplRoutes() []byte
	HtmxTemplImports() []byte
	WebsocketImports() []byte
}

type DBDriverTemplater interface {
	Service() []byte
	Env() []byte
}

type DockerTemplater interface {
	Docker() []byte
}

type WorkflowTemplater interface {
	Releaser() []byte
	Test() []byte
	ReleaserConfig() []byte
}

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
	redisDriver    = []string{"github.com/redis/go-redis/v9"}
	mongoDriver    = []string{"go.mongodb.org/mongo-driver"}

	godotenvPackage = []string{"github.com/joho/godotenv"}
	templPackage    = []string{"github.com/a-h/templ"}
)

const (
	root                 = "/"
	cmdApiPath           = "cmd/api"
	cmdWebPath           = "cmd/web"
	internalServerPath   = "internal/server"
	internalDatabasePath = "internal/database"
	gitHubActionPath     = ".github/workflows"
	testHandlerPath      = "tests"
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
	p.FrameworkMap[flags.Chi] = Framework{
		packageName: chiPackage,
		templater:   framework.ChiTemplates{},
	}

	p.FrameworkMap[flags.StandardLibrary] = Framework{
		packageName: []string{},
		templater:   framework.StandardLibTemplate{},
	}

	p.FrameworkMap[flags.Gin] = Framework{
		packageName: ginPackage,
		templater:   framework.GinTemplates{},
	}

	p.FrameworkMap[flags.Fiber] = Framework{
		packageName: fiberPackage,
		templater:   framework.FiberTemplates{},
	}

	p.FrameworkMap[flags.GorillaMux] = Framework{
		packageName: gorillaPackage,
		templater:   framework.GorillaTemplates{},
	}

	p.FrameworkMap[flags.HttpRouter] = Framework{
		packageName: routerPackage,
		templater:   framework.RouterTemplates{},
	}

	p.FrameworkMap[flags.Echo] = Framework{
		packageName: echoPackage,
		templater:   framework.EchoTemplates{},
	}
}

func (p *Project) createDBDriverMap() {
	p.DBDriverMap[flags.MySql] = Driver{
		packageName: mysqlDriver,
		templater:   dbdriver.MysqlTemplate{},
	}
	p.DBDriverMap[flags.Postgres] = Driver{
		packageName: postgresDriver,
		templater:   dbdriver.PostgresTemplate{},
	}
	p.DBDriverMap[flags.Sqlite] = Driver{
		packageName: sqliteDriver,
		templater:   dbdriver.SqliteTemplate{},
	}
	p.DBDriverMap[flags.Mongo] = Driver{
		packageName: mongoDriver,
		templater:   dbdriver.MongoTemplate{},
	}
	p.DBDriverMap[flags.Redis] = Driver{
		packageName: redisDriver,
		templater:   dbdriver.RedisTemplate{},
	}
}

func (p *Project) createDockerMap() {
	p.DockerMap = make(map[flags.Database]Docker)

	p.DockerMap[flags.MySql] = Docker{
		packageName: []string{},
		templater:   docker.MysqlDockerTemplate{},
	}
	p.DockerMap[flags.Postgres] = Docker{
		packageName: []string{},
		templater:   docker.PostgresDockerTemplate{},
	}
	p.DockerMap[flags.Mongo] = Docker{
		packageName: []string{},
		templater:   docker.MongoDockerTemplate{},
	}
	p.DockerMap[flags.Redis] = Docker{
		packageName: []string{},
		templater:   docker.RedisDockerTemplate{},
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

	nameSet, err := utils.CheckGitConfig("user.name")
	if err != nil {
		cobra.CheckErr(err)
	}
	if !nameSet {
		fmt.Println("user.name is not set in git config.")
		fmt.Println("Please set up git config before trying again.")
		panic("\nGIT CONFIG ISSUE: user.name is not set in git config.\n")
	}

	// Check if user.email is set.
	emailSet, err := utils.CheckGitConfig("user.email")
	if err != nil {
		cobra.CheckErr(err)
	}
	if !emailSet {
		fmt.Println("user.email is not set in git config.")
		fmt.Println("Please set up git config before trying again.")
		panic("\nGIT CONFIG ISSUE: user.email is not set in git config.\n")
	}

	p.ProjectName = strings.TrimSpace(p.ProjectName)

	// Create a new directory with the project name
	projectPath := filepath.Join(p.AbsolutePath, p.ProjectName)
	if _, err := os.Stat(projectPath); os.IsNotExist(err) {
		err := os.MkdirAll(projectPath, 0751)
		if err != nil {
			log.Printf("Error creating root project directory %v\n", err)
			return err
		}
	}

	// Create the map for our program
	p.createFrameworkMap()

	// Create go.mod
	err = utils.InitGoMod(p.ProjectName, projectPath)
	if err != nil {
		log.Printf("Could not initialize go.mod in new project %v\n", err)
		cobra.CheckErr(err)
	}

	// Install the correct package for the selected framework
	if p.ProjectType != flags.StandardLibrary {
		err = utils.GoGetPackage(projectPath, p.FrameworkMap[p.ProjectType].packageName)
		if err != nil {
			log.Printf("Could not install go dependency for the chosen framework %v\n", err)
			cobra.CheckErr(err)
		}
	}

	// Install the correct package for the selected driver
	if p.DBDriver != "none" {
		p.createDBDriverMap()
		err = utils.GoGetPackage(projectPath, p.DBDriverMap[p.DBDriver].packageName)
		if err != nil {
			log.Printf("Could not install go dependency for chosen driver %v\n", err)
			cobra.CheckErr(err)
		}

		err = p.CreatePath(internalDatabasePath, projectPath)
		if err != nil {
			log.Printf("Error creating path: %s", internalDatabasePath)
			cobra.CheckErr(err)
			return err
		}

		err = p.CreateFileWithInjection(internalDatabasePath, projectPath, "database.go", "database")
		if err != nil {
			log.Printf("Error injecting database.go file: %v", err)
			cobra.CheckErr(err)
			return err
		}
	}

	// Create correct docker compose for the selected driver
	if p.DBDriver != "none" {

		if p.DBDriver != "sqlite" {
			p.createDockerMap()
			p.Docker = p.DBDriver

			err = p.CreateFileWithInjection(root, projectPath, "docker-compose.yml", "db-docker")
			if err != nil {
				log.Printf("Error injecting docker-compose.yml file: %v", err)
				cobra.CheckErr(err)
				return err
			}
		} else {
			fmt.Println(" We are unable to create docker-compose.yml file for an SQLite database")
		}
	}

	// Install the godotenv package
	err = utils.GoGetPackage(projectPath, godotenvPackage)
	if err != nil {
		log.Printf("Could not install go dependency %v\n", err)
		cobra.CheckErr(err)
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

	err = p.CreatePath(testHandlerPath, projectPath)
	if err != nil {
		log.Printf("Error creating path: %s", projectPath)
		cobra.CheckErr(err)
		return err
	}
	// inject testhandler template
	err = p.CreateFileWithInjection(testHandlerPath, projectPath, "handler_test.go", "tests")
	if err != nil {
		cobra.CheckErr(err)
		return err
	}

	makeFile, err := os.Create(filepath.Join(projectPath, "Makefile"))
	if err != nil {
		cobra.CheckErr(err)
		return err
	}

	defer makeFile.Close()

	if p.DBDriver == "sqlite" || p.DBDriver == "none" {
		// inject makefile template
		makeFileTemplate := template.Must(template.New("makefile").Parse(string(framework.NonDbMakeFileTemplate())))
		err = makeFileTemplate.Execute(makeFile, p)
		if err != nil {
			return err
		}
	} else {
		// inject makefile template for database excluding sqlite
		makeFileTemplate := template.Must(template.New("makefile").Parse(string(framework.MakeTemplate())))
		err = makeFileTemplate.Execute(makeFile, p)
		if err != nil {
			return err
		}
	}

	readmeFile, err := os.Create(filepath.Join(projectPath, "README.md"))
	if err != nil {
		cobra.CheckErr(err)
		return err
	}
	defer readmeFile.Close()

	// inject readme template
	readmeFileTemplate := template.Must(template.New("readme").Parse(string(framework.ReadmeTemplate())))
	err = readmeFileTemplate.Execute(readmeFile, p)
	if err != nil {
		return err
	}

	err = p.CreatePath(internalServerPath, projectPath)
	if err != nil {
		log.Printf("Error creating path: %s", internalServerPath)
		cobra.CheckErr(err)
		return err
	}

	if p.AdvancedOptions[string(flags.Htmx)] {
		// create folders and hello world file
		err = p.CreatePath(cmdWebPath, projectPath)
		if err != nil {
			cobra.CheckErr(err)
			return err
		}
		helloTemplFile, err := os.Create(fmt.Sprintf("%s/%s/hello.templ", projectPath, cmdWebPath))
		if err != nil {
			cobra.CheckErr(err)
		}
		defer helloTemplFile.Close()

		//inject hello.templ template
		helloTemplTemplate := template.Must(template.New("hellotempl").Parse((string(advanced.HelloTemplTemplate()))))
		err = helloTemplTemplate.Execute(helloTemplFile, p)
		if err != nil {
			return err
		}

		baseTemplFile, err := os.Create(fmt.Sprintf("%s/%s/base.templ", projectPath, cmdWebPath))
		if err != nil {
			cobra.CheckErr(err)
		}
		defer baseTemplFile.Close()

		baseTemplTemplate := template.Must(template.New("basetempl").Parse((string(advanced.BaseTemplTemplate()))))
		err = baseTemplTemplate.Execute(baseTemplFile, p)
		if err != nil {
			return err
		}

		err = os.Mkdir(fmt.Sprintf("%s/%s/js", projectPath, cmdWebPath), 0755)
		if err != nil {
			cobra.CheckErr(err)
		}

		htmxMinJsFile, err := os.Create(fmt.Sprintf("%s/%s/js/htmx.min.js", projectPath, cmdWebPath))
		if err != nil {
			cobra.CheckErr(err)
		}
		defer htmxMinJsFile.Close()

		htmxMinJsTemplate := advanced.HtmxJSTemplate()
		err = os.WriteFile(fmt.Sprintf("%s/%s/js/htmx.min.js", projectPath, cmdWebPath), htmxMinJsTemplate, 0644)
		if err != nil {
			return err
		}

		efsFile, err := os.Create(fmt.Sprintf("%s/%s/efs.go", projectPath, cmdWebPath))
		if err != nil {
			cobra.CheckErr(err)
		}
		defer efsFile.Close()

		efsTemplate := template.Must(template.New("efs").Parse((string(advanced.EfsTemplate()))))
		err = efsTemplate.Execute(efsFile, p)
		if err != nil {
			return err
		}
		err = utils.GoGetPackage(projectPath, templPackage)
		if err != nil {
			log.Printf("Could not install go dependency %v\n", err)
			cobra.CheckErr(err)
		}

		helloGoFile, err := os.Create(fmt.Sprintf("%s/%s/hello.go", projectPath, cmdWebPath))
		if err != nil {
			cobra.CheckErr(err)
		}
		defer efsFile.Close()

		if p.ProjectType == "fiber" {
			helloGoTemplate := template.Must(template.New("efs").Parse((string(advanced.HelloFiberGoTemplate()))))
			err = helloGoTemplate.Execute(helloGoFile, p)
			if err != nil {
				return err
			}
			err = utils.GoGetPackage(projectPath, []string{"github.com/gofiber/fiber/v2/middleware/adaptor"})
			if err != nil {
				log.Printf("Could not install go dependency %v\n", err)
				cobra.CheckErr(err)
			}
			if err != nil {
				log.Printf("Could not install go dependency %v\n", err)
				cobra.CheckErr(err)
			}
		} else {
			helloGoTemplate := template.Must(template.New("efs").Parse((string(advanced.HelloGoTemplate()))))
			err = helloGoTemplate.Execute(helloGoFile, p)
			if err != nil {
				return err
			}
		}

		p.CreateHtmxTemplates()
	}

	// Create .github/workflows folder and inject release.yml and go-test.yml
	if p.AdvancedOptions[string(flags.GoProjectWorkflow)] {
		err = p.CreatePath(gitHubActionPath, projectPath)
		if err != nil {
			log.Printf("Error creating path: %s", gitHubActionPath)
			cobra.CheckErr(err)
			return err
		}

		err = p.CreateFileWithInjection(gitHubActionPath, projectPath, "release.yml", "releaser")
		if err != nil {
			log.Printf("Error injecting release.yml file: %v", err)
			cobra.CheckErr(err)
			return err
		}

		err = p.CreateFileWithInjection(gitHubActionPath, projectPath, "go-test.yml", "go-test")
		if err != nil {
			log.Printf("Error injecting go-test.yml file: %v", err)
			cobra.CheckErr(err)
			return err
		}

		err = p.CreateFileWithInjection(root, projectPath, ".goreleaser.yml", "releaser-config")
		if err != nil {
			log.Printf("Error injecting .goreleaser.yml file: %v", err)
			cobra.CheckErr(err)
			return err
		}
	}

	// if the websocket option is checked, a websocket dependency needs to
	// be added to the routes depending on the framework choosen.
	// Only fiber uses a different websocket library, the other frameworks
	// all work with the same one
	if p.AdvancedOptions["Websocket"] {
		p.CreateWebsocketImports(projectPath)
	}

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

	err = p.CreateFileWithInjection(root, projectPath, ".env", "env")
	if err != nil {
		log.Printf("Error injecting .env file: %v", err)
		cobra.CheckErr(err)
		return err
	}

	// Initialize git repo
	err = utils.ExecuteCmd("git", []string{"init"}, projectPath)
	if err != nil {
		log.Printf("Error initializing git repo: %v", err)
		cobra.CheckErr(err)
		return err
	}
	// Create gitignore
	gitignoreFile, err := os.Create(filepath.Join(projectPath, ".gitignore"))
	if err != nil {
		cobra.CheckErr(err)
		return err
	}
	defer gitignoreFile.Close()

	// inject gitignore template
	gitignoreTemplate := template.Must(template.New(".gitignore").Parse(string(framework.GitIgnoreTemplate())))
	err = gitignoreTemplate.Execute(gitignoreFile, p)
	if err != nil {
		return err
	}

	// Create .air.toml file
	airTomlFile, err := os.Create(filepath.Join(projectPath, ".air.toml"))
	if err != nil {
		cobra.CheckErr(err)
		return err
	}

	defer airTomlFile.Close()

	// inject air.toml template
	airTomlTemplate := template.Must(template.New("airtoml").Parse(string(framework.AirTomlTemplate())))
	err = airTomlTemplate.Execute(airTomlFile, p)
	if err != nil {
		return err
	}

	err = utils.GoTidy(projectPath)
	if err != nil {
		log.Printf("Could not go tidy in new project %v\n", err)
		cobra.CheckErr(err)
	}

	err = utils.GoFmt(projectPath)
	if err != nil {
		log.Printf("Could not gofmt in new project %v\n", err)
		cobra.CheckErr(err)
		return err
	}
	// Git add files
	err = utils.ExecuteCmd("git", []string{"add", "."}, projectPath)
	if err != nil {
		log.Printf("Error adding files to git repo: %v", err)
		cobra.CheckErr(err)
		return err
	}
	// Git commit files
	err = utils.ExecuteCmd("git", []string{"commit", "-m", "Initial commit"}, projectPath)
	if err != nil {
		log.Printf("Error committing files to git repo: %v", err)
		cobra.CheckErr(err)
		return err
	}
	return nil
}

// CreatePath creates the given directory in the projectPath
func (p *Project) CreatePath(pathToCreate string, projectPath string) error {
	path := filepath.Join(projectPath, pathToCreate)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.MkdirAll(path, 0751)
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
	createdFile, err := os.Create(filepath.Join(projectPath, pathToCreate, fileName))
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
		routeFileBytes := p.FrameworkMap[p.ProjectType].templater.Routes()
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
	case "tests":
		createdTemplate := template.Must(template.New(fileName).Parse(string(p.FrameworkMap[p.ProjectType].templater.TestHandler())))
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

func (p *Project) CreateHtmxTemplates() {
	routesPlaceHolder := ""
	importsPlaceHolder := ""
	if p.AdvancedOptions[string(flags.Htmx)] {
		routesPlaceHolder += string(p.FrameworkMap[p.ProjectType].templater.HtmxTemplRoutes())
		importsPlaceHolder += string(p.FrameworkMap[p.ProjectType].templater.HtmxTemplImports())
	}

	routeTmpl, err := template.New("routes").Parse(routesPlaceHolder)
	if err != nil {
		log.Fatal(err)
	}
	importTmpl, err := template.New("imports").Parse(importsPlaceHolder)
	if err != nil {
		log.Fatal(err)
	}
	var routeBuffer bytes.Buffer
	var importBuffer bytes.Buffer
	err = routeTmpl.Execute(&routeBuffer, p)
	if err != nil {
		log.Fatal(err)
	}
	err = importTmpl.Execute(&importBuffer, p)
	if err != nil {
		log.Fatal(err)
	}
	p.AdvancedTemplates.TemplateRoutes = template.HTML(routeBuffer.String())
	p.AdvancedTemplates.TemplateImports = template.HTML(importBuffer.String())
}

func (p *Project) CreateWebsocketImports(appDir string) {
	websocketDependency := []string{"nhooyr.io/websocket"}
	if p.ProjectType == flags.Fiber {
		websocketDependency = []string{"github.com/gofiber/contrib/websocket"}
	}

	// Websockets require a different package depending on what framework is
	// choosen. The application calls go mod tidy at the end so we don't
	// have to here
	err := utils.GoGetPackage(appDir, websocketDependency)
	if err != nil {
		log.Fatal(err)
	}

	importsPlaceHolder := ""
	if p.AdvancedOptions["Websocket"] {
		importsPlaceHolder += string(p.FrameworkMap[p.ProjectType].templater.WebsocketImports())
	}

	importTmpl, err := template.New("imports").Parse(importsPlaceHolder)
	if err != nil {
		log.Fatal(err)
	}
	var importBuffer bytes.Buffer
	err = importTmpl.Execute(&importBuffer, p)
	if err != nil {
		log.Fatal(err)
	}
	newImports := strings.Join([]string{string(p.AdvancedTemplates.TemplateImports), importBuffer.String()}, "\n")
	p.AdvancedTemplates.TemplateImports = template.HTML(newImports)
}
