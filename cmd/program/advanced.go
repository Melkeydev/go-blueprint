package program

import (
	"bytes"
	"log"
	"strings"
	"text/template"

	"github.com/melkeydev/go-blueprint/cmd/flags"
	"github.com/melkeydev/go-blueprint/cmd/utils"
)

func (p *Project) CreateWebsocketImports(appDir string) {
	websocketDependency := []string{"github.com/coder/websocket"}
	if p.ProjectType == flags.Fiber {
		websocketDependency = []string{"github.com/gofiber/contrib/websocket"}
	}

	// Websockets require a different package depending on what backend is
	// choosen. The application calls go mod tidy at the end so we don't
	// have to here
	err := utils.GoGetPackage(appDir, websocketDependency)
	if err != nil {
		log.Fatal(err)
	}

	importsPlaceHolder := string(p.BackendMap[p.ProjectType].templater.WebsocketImports())

	importTmpl, err := template.New("imports").Parse(importsPlaceHolder)
	if err != nil {
		log.Fatalf("CreateWebsocketImports failed to create template: %v", err)
	}
	var importBuffer bytes.Buffer
	err = importTmpl.Execute(&importBuffer, p)
	if err != nil {
		log.Fatalf("CreateWebsocketImports failed write template: %v", err)
	}
	newImports := strings.Join([]string{string(p.AdvancedTemplates.TemplateImports), importBuffer.String()}, "\n")
	p.AdvancedTemplates.TemplateImports = newImports
}
