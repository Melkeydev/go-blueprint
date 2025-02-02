// Package program provides the
// main functionality of Blueprint
package program

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/melkeydev/go-blueprint/cmd/flags"
	"github.com/melkeydev/go-blueprint/cmd/template/advanced"
	"github.com/melkeydev/go-blueprint/cmd/template/backend"
	"github.com/melkeydev/go-blueprint/cmd/template/dbdriver"
	"github.com/melkeydev/go-blueprint/cmd/template/docker"
	"github.com/melkeydev/go-blueprint/cmd/template/frontend"
	"github.com/melkeydev/go-blueprint/cmd/utils"
)

// A Project contains the data for the project folder
// being created, and methods that help with that process
type Project struct {
	ProjectName         string
	Exit                bool
	AbsolutePath        string
	BackendFramework    flags.BackendFramework
	DBDriver            flags.Database
	Docker              flags.Database
	FrontendFramework   flags.FrontendFramework
	BackendFrameworkMap map[flags.BackendFramework]BackendFramework
	DBDriverMap         map[flags.Database]Driver
	DockerMap           map[flags.Database]Docker
	FrontendTemplates   FrontendTemplates
	FrontendOptions     map[string]bool
	AdvancedTemplates   AdvancedTemplates
	AdvancedOptions     map[string]bool
	GitOptions          flags.Git
	OSCheck             map[string]bool
}

type FrontendTemplates struct {
	TemplateRoutes  string
	TemplateImports string
}

type AdvancedTemplates struct {
	TemplateRoutes  string
	TemplateImports string
}

// A Backend Framework contains the name and templater for a
// given Backend Framework
type BackendFramework struct {
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
// in the Project folder, and is specific to a BackendFramework
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
	Tests() []byte
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
	ginPackage     = []string{"github.com/gin-gonic/gin"}
	fiberPackage   = []string{"github.com/gofiber/fiber/v2"}
	echoPackage    = []string{"github.com/labstack/echo/v4", "github.com/labstack/echo/v4/middleware"}

	mysqlDriver    = []string{"github.com/go-sql-driver/mysql"}
	postgresDriver = []string{"github.com/lib/pq"}
	sqliteDriver   = []string{"github.com/mattn/go-sqlite3"}
	redisDriver    = []string{"github.com/redis/go-redis/v9"}
	mongoDriver    = []string{"go.mongodb.org/mongo-driver"}
	gocqlDriver    = []string{"github.com/gocql/gocql"}
	scyllaDriver   = "github.com/scylladb/gocql@v1.14.4" // Replacement for GoCQL

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
// BackendFrameworks into a Project's BackendFrameworkMap
func (p *Project) createBackendFrameworkMap() {
	p.BackendFrameworkMap[flags.Chi] = BackendFramework{
		packageName: chiPackage,
		templater:   backend.ChiTemplates{},
	}

	p.BackendFrameworkMap[flags.StandardLibrary] = BackendFramework{
		packageName: []string{},
		templater:   backend.StandardLibTemplate{},
	}

	p.BackendFrameworkMap[flags.Gin] = BackendFramework{
		packageName: ginPackage,
		templater:   backend.GinTemplates{},
	}

	p.BackendFrameworkMap[flags.Fiber] = BackendFramework{
		packageName: fiberPackage,
		templater:   backend.FiberTemplates{},
	}

	p.BackendFrameworkMap[flags.GorillaMux] = BackendFramework{
		packageName: gorillaPackage,
		templater:   backend.GorillaTemplates{},
	}

	p.BackendFrameworkMap[flags.Echo] = BackendFramework{
		packageName: echoPackage,
		templater:   backend.EchoTemplates{},
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

	p.DBDriverMap[flags.Scylla] = Driver{
		packageName: gocqlDriver,
		templater:   dbdriver.ScyllaTemplate{},
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
	p.DockerMap[flags.Scylla] = Docker{
		packageName: []string{},
		templater:   docker.ScyllaDockerTemplate{},
	}
}

// CreateMainFile creates the project folders and files,
// and writes to them depending on the selected options
func (p *Project) CreateMainFile() error {

	var err error

	// check if AbsolutePath exists
	if _, err := os.Stat(p.AbsolutePath); os.IsNotExist(err) {
		// create directory
		if err := os.Mkdir(p.AbsolutePath, 0o754); err != nil {
			log.Printf("Could not create directory: %v", err)
			return err
		}
	}

	p.ProjectName = strings.TrimSpace(p.ProjectName)

	// Create a new directory with the project name
	projectPath := filepath.Join(p.AbsolutePath, utils.GetRootDir(p.ProjectName))
	if _, err := os.Stat(projectPath); os.IsNotExist(err) {
		err := os.MkdirAll(projectPath, 0o751)
		if err != nil {
			log.Printf("Error creating root project directory %v\n", err)
			return err
		}
	}

	// Define Operating system
	p.CheckOS()

	// Create the map for our program
	p.createBackendFrameworkMap()

	// Create go.mod
	err = utils.InitGoMod(p.ProjectName, projectPath)
	if err != nil {
		log.Printf("Could not initialize go.mod in new project %v\n", err)
		return err
	}

	// Install the correct package for the selected backend
	if p.BackendFramework != flags.StandardLibrary {
		err = utils.GoGetPackage(projectPath, p.BackendFrameworkMap[p.BackendFramework].packageName)
		if err != nil {
			log.Printf("Could not install go dependency for the chosen backend %v\n", err)
			return err
		}
	}

	// Install the correct package for the selected driver
	if p.DBDriver != "none" {
		p.createDBDriverMap()
		err = utils.GoGetPackage(projectPath, p.DBDriverMap[p.DBDriver].packageName)
		if err != nil {
			log.Printf("Could not install go dependency for chosen driver %v\n", err)
			return err
		}

		err = p.CreatePath(internalDatabasePath, projectPath)
		if err != nil {
			log.Printf("Error creating path: %s", internalDatabasePath)
			return err
		}

		err = p.CreateFileWithInjection(internalDatabasePath, projectPath, "database.go", "database")
		if err != nil {
			log.Printf("Error injecting database.go file: %v", err)
			return err
		}

		if p.DBDriver != "sqlite" {
			err = p.CreateFileWithInjection(internalDatabasePath, projectPath, "database_test.go", "integration-tests")
			if err != nil {
				log.Printf("Error injecting database_test.go file: %v", err)
				return err
			}
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
				return err
			}
		} else {
			fmt.Println(" SQLite doesn't support docker-compose.yml configuration")
		}
	}

	// Install the godotenv package
	err = utils.GoGetPackage(projectPath, godotenvPackage)

	if err != nil {
		log.Printf("Could not install go dependency %v\n", err)

		return err
	}

	if p.DBDriver == flags.Scylla {
		replace := fmt.Sprintf("%s=%s", gocqlDriver[0], scyllaDriver)
		err = utils.GoModReplace(projectPath, replace)
		if err != nil {
			log.Printf("Could not replace go dependency %v\n", err)
			return err
		}
	}

	err = p.CreatePath(cmdApiPath, projectPath)
	if err != nil {
		log.Printf("Error creating path: %s", projectPath)
		return err
	}

	err = p.CreateFileWithInjection(cmdApiPath, projectPath, "main.go", "main")
	if err != nil {
		return err
	}

	makeFile, err := os.Create(filepath.Join(projectPath, "Makefile"))
	if err != nil {
		return err
	}

	defer makeFile.Close()

	// inject makefile template
	makeFileTemplate := template.Must(template.New("makefile").Parse(string(backend.MakeTemplate())))
	err = makeFileTemplate.Execute(makeFile, p)
	if err != nil {
		return err
	}

	readmeFile, err := os.Create(filepath.Join(projectPath, "README.md"))
	if err != nil {
		return err
	}
	defer readmeFile.Close()

	// inject readme template
	readmeFileTemplate := template.Must(template.New("readme").Parse(string(backend.ReadmeTemplate())))
	err = readmeFileTemplate.Execute(readmeFile, p)
	if err != nil {
		return err
	}

	err = p.CreatePath(internalServerPath, projectPath)
	if err != nil {
		log.Printf("Error creating path: %s", internalServerPath)
		return err
	}

	if p.FrontendFramework == flags.React {
		if err := p.CreateViteReactProject(projectPath); err != nil {
			return fmt.Errorf("failed to set up React project: %w", err)
		}
	}

	if p.FrontendOptions[string(flags.Tailwind)] && p.FrontendFramework == flags.Htmx {
		tailwindConfigFile, err := os.Create(fmt.Sprintf("%s/tailwind.config.js", projectPath))
		if err != nil {
			return err
		}
		defer tailwindConfigFile.Close()

		tailwindConfigTemplate := frontend.TailwindConfigTemplate()
		err = os.WriteFile(fmt.Sprintf("%s/tailwind.config.js", projectPath), tailwindConfigTemplate, 0o644)
		if err != nil {
			return err
		}

		err = os.MkdirAll(fmt.Sprintf("%s/%s/assets/css", projectPath, cmdWebPath), 0o755)
		if err != nil {
			return err
		}

		err = os.MkdirAll(fmt.Sprintf("%s/%s/styles", projectPath, cmdWebPath), 0o755)
		if err != nil {
			return fmt.Errorf("failed to create styles directory: %w", err)
		}

		inputCssFile, err := os.Create(fmt.Sprintf("%s/%s/styles/input.css", projectPath, cmdWebPath))
		if err != nil {
			return err
		}
		defer inputCssFile.Close()

		inputCssTemplate := frontend.InputCssTemplate()
		err = os.WriteFile(fmt.Sprintf("%s/%s/styles/input.css", projectPath, cmdWebPath), inputCssTemplate, 0o644)
		if err != nil {
			return err
		}

		outputCssFile, err := os.Create(fmt.Sprintf("%s/%s/assets/css/output.css", projectPath, cmdWebPath))
		if err != nil {
			return err
		}
		defer outputCssFile.Close()

		outputCssTemplate := frontend.OutputCssTemplate()
		err = os.WriteFile(fmt.Sprintf("%s/%s/assets/css/output.css", projectPath, cmdWebPath), outputCssTemplate, 0o644)
		if err != nil {
			return err
		}
	}

	if p.FrontendFramework == flags.Htmx {
		// create folders and hello world file
		err = p.CreatePath(cmdWebPath, projectPath)
		if err != nil {
			return err
		}
		helloTemplFile, err := os.Create(fmt.Sprintf("%s/%s/hello.templ", projectPath, cmdWebPath))
		if err != nil {
			return err
		}
		defer helloTemplFile.Close()

		// inject hello.templ template
		helloTemplTemplate := template.Must(template.New("hellotempl").Parse((string(frontend.HelloTemplTemplate()))))
		err = helloTemplTemplate.Execute(helloTemplFile, p)
		if err != nil {
			return err
		}

		baseTemplFile, err := os.Create(fmt.Sprintf("%s/%s/base.templ", projectPath, cmdWebPath))
		if err != nil {
			return err
		}
		defer baseTemplFile.Close()

		baseTemplTemplate := template.Must(template.New("basetempl").Parse((string(frontend.BaseTemplTemplate()))))
		err = baseTemplTemplate.Execute(baseTemplFile, p)
		if err != nil {
			return err
		}

		err = os.MkdirAll(fmt.Sprintf("%s/%s/assets/js", projectPath, cmdWebPath), 0o755)
		if err != nil {
			return err
		}

		htmxMinJsFile, err := os.Create(fmt.Sprintf("%s/%s/assets/js/htmx.min.js", projectPath, cmdWebPath))
		if err != nil {
			return err
		}
		defer htmxMinJsFile.Close()

		htmxMinJsTemplate := frontend.HtmxJSTemplate()
		err = os.WriteFile(fmt.Sprintf("%s/%s/assets/js/htmx.min.js", projectPath, cmdWebPath), htmxMinJsTemplate, 0o644)
		if err != nil {
			return err
		}

		efsFile, err := os.Create(fmt.Sprintf("%s/%s/efs.go", projectPath, cmdWebPath))
		if err != nil {
			return err
		}
		defer efsFile.Close()

		efsTemplate := template.Must(template.New("efs").Parse((string(frontend.EfsTemplate()))))
		err = efsTemplate.Execute(efsFile, p)
		if err != nil {
			return err
		}
		err = utils.GoGetPackage(projectPath, templPackage)
		if err != nil {
			log.Printf("Could not install go dependency %v\n", err)
			return err
		}

		helloGoFile, err := os.Create(fmt.Sprintf("%s/%s/hello.go", projectPath, cmdWebPath))
		if err != nil {
			return err
		}
		defer efsFile.Close()

		if p.BackendFramework == "fiber" {
			helloGoTemplate := template.Must(template.New("efs").Parse((string(frontend.HelloFiberGoTemplate()))))
			err = helloGoTemplate.Execute(helloGoFile, p)
			if err != nil {
				return err
			}
			err = utils.GoGetPackage(projectPath, []string{"github.com/gofiber/fiber/v2/middleware/adaptor"})
			if err != nil {
				log.Printf("Could not install go dependency %v\n", err)
				return err
			}
		} else {
			helloGoTemplate := template.Must(template.New("efs").Parse((string(frontend.HelloGoTemplate()))))
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
			return err
		}

		err = p.CreateFileWithInjection(gitHubActionPath, projectPath, "release.yml", "releaser")
		if err != nil {
			log.Printf("Error injecting release.yml file: %v", err)
			return err
		}

		err = p.CreateFileWithInjection(gitHubActionPath, projectPath, "go-test.yml", "go-test")
		if err != nil {
			log.Printf("Error injecting go-test.yml file: %v", err)
			return err
		}

		err = p.CreateFileWithInjection(root, projectPath, ".goreleaser.yml", "releaser-config")
		if err != nil {
			log.Printf("Error injecting .goreleaser.yml file: %v", err)
			return err
		}
	}

	// if the websocket option is checked, a websocket dependency needs to
	// be added to the routes depending on the backend choosen.
	// Only fiber uses a different websocket library, the other frameworks
	// all work with the same one
	if p.AdvancedOptions[string(flags.Websocket)] {
		p.CreateWebsocketImports(projectPath)
	}

	if p.AdvancedOptions[string(flags.Docker)] {
		Dockerfile, err := os.Create(filepath.Join(projectPath, "Dockerfile"))
		if err != nil {
			return err
		}
		defer Dockerfile.Close()

		// inject Docker template
		dockerfileTemplate := template.Must(template.New("Dockerfile").Parse(string(advanced.Dockerfile())))
		err = dockerfileTemplate.Execute(Dockerfile, p)
		if err != nil {
			return err
		}

		if p.DBDriver == "none" || p.DBDriver == "sqlite" {

			Dockercompose, err := os.Create(filepath.Join(projectPath, "docker-compose.yml"))
			if err != nil {
				return err
			}
			defer Dockercompose.Close()

			// inject DockerCompose template
			dockerComposeTemplate := template.Must(template.New("docker-compose.yml").Parse(string(advanced.DockerCompose())))
			err = dockerComposeTemplate.Execute(Dockercompose, p)
			if err != nil {
				return err
			}
		}
	}

	err = p.CreateFileWithInjection(internalServerPath, projectPath, "routes.go", "routes")
	if err != nil {
		log.Printf("Error injecting routes.go file: %v", err)
		return err
	}

	err = p.CreateFileWithInjection(internalServerPath, projectPath, "routes_test.go", "tests")
	if err != nil {
		return err
	}

	err = p.CreateFileWithInjection(internalServerPath, projectPath, "server.go", "server")
	if err != nil {
		log.Printf("Error injecting server.go file: %v", err)
		return err
	}

	err = p.CreateFileWithInjection(root, projectPath, ".env", "env")
	if err != nil {
		log.Printf("Error injecting .env file: %v", err)
		return err
	}

	// Create .air.toml file
	airTomlFile, err := os.Create(filepath.Join(projectPath, ".air.toml"))
	if err != nil {
		return err
	}

	defer airTomlFile.Close()

	// inject air.toml template
	airTomlTemplate := template.Must(template.New("airtoml").Parse(string(backend.AirTomlTemplate())))
	err = airTomlTemplate.Execute(airTomlFile, p)
	if err != nil {
		return err
	}

	err = utils.GoTidy(projectPath)
	if err != nil {
		log.Printf("Could not go tidy in new project %v\n", err)
		return err
	}

	err = utils.GoFmt(projectPath)
	if err != nil {
		log.Printf("Could not gofmt in new project %v\n", err)
		return err
	}

	if p.GitOptions != flags.Skip {
		nameSet, err := utils.CheckGitConfig("user.name")
		if err != nil {
			return err
		}
		emailSet, err := utils.CheckGitConfig("user.email")
		if err != nil {
			return err
		}

		if !nameSet {
			panic("\nGIT CONFIG ISSUE: user.name is not set in git config.\n")
		}

		if !emailSet {
			panic("\nGIT CONFIG ISSUE: user.email is not set in git config.\n")
		}

		// Initialize git repo
		err = utils.ExecuteCmd("git", []string{"init"}, projectPath)
		if err != nil {
			log.Printf("Error initializing git repo: %v", err)
			return err
		}

		// Git add files
		err = utils.ExecuteCmd("git", []string{"add", "."}, projectPath)
		if err != nil {
			log.Printf("Error adding files to git repo: %v", err)
			return err
		}

		// Create gitignore
		gitignoreFile, err := os.Create(filepath.Join(projectPath, ".gitignore"))
		if err != nil {
			return err
		}
		defer gitignoreFile.Close()

		// inject gitignore template
		gitignoreTemplate := template.Must(template.New(".gitignore").Parse(string(backend.GitIgnoreTemplate())))
		err = gitignoreTemplate.Execute(gitignoreFile, p)
		if err != nil {
			return err
		}

		if p.GitOptions == flags.Commit {
			// Git commit files
			err = utils.ExecuteCmd("git", []string{"commit", "-m", "Initial commit"}, projectPath)
			if err != nil {
				log.Printf("Error committing files to git repo: %v", err)
				return err
			}
		}
	}
	return nil
}
