package program

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
	"os/exec"

	tea "github.com/charmbracelet/bubbletea"
	tpl "github.com/melkeydev/go-blueprint/cmd/template"
	"github.com/spf13/cobra"
)

type Project struct {
	ProjectName  string
	Exit         bool
	AbsolutePath string
}

const (
	chiPackage     = "github.com/go-chi/chi/v5"
	gorillaPackage = "github.com/gorilla/mux"
	routerPackage  = "github.com/julienschmidt/httprouter"
	ginPackage     = "github.com/gin-gonic/gin"
	fiberPacker    = "github.com/gofiber/fiber/v2"
)

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

func goGetPackage(appDir, packageName string) {
	if err := executeCmd("go",
		[]string{"get", "-u", packageName},
		appDir); err != nil {
		cobra.CheckErr(err)
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
		fmt.Printf("this is the err %v\n", err)
		return err
	}

	makeFile, err := os.Create(fmt.Sprintf("%s/Makefile", projectPath))
	if err != nil {
		return err
	}

	defer makeFile.Close()

	// inject makefile template
	makeFileTemplate := template.Must(template.New("makefile").Parse(string(tpl.MakeTemplate())))
	err = makeFileTemplate.Execute(makeFile, p)
	if err != nil {
		return err
	}

	// create /internal/server
	if _, err := os.Stat(fmt.Sprintf("%s/internal/server", projectPath)); os.IsNotExist(err) {
		err := os.MkdirAll(fmt.Sprintf("%s/internal/server", projectPath), 0751)
		if err != nil {
			fmt.Printf("Error creating directory %v\n", err)
		}
	}

	serverFile, err := os.Create(fmt.Sprintf("%s/internal/server/server.go", projectPath))
	if err != nil {
		return err
	}

	serverFileTemplate := template.Must(template.New("server").Parse(string(tpl.MakeHTTPServer())))
	err = serverFileTemplate.Execute(serverFile, p)
	if err != nil {
		return err
	}

	defer serverFile.Close()

	routesFile, err := os.Create(fmt.Sprintf("%s/internal/server/routes.go", projectPath))
	if err != nil {
		return err
	}

	routesFileTemplate := template.Must(template.New("routes").Parse(string(tpl.MakeHTTPRoutes())))
	err = routesFileTemplate.Execute(routesFile, p)
	if err != nil {
		return err
	}

	defer routesFile.Close()

	return nil
}
