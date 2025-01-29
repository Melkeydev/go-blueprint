package program

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/melkeydev/go-blueprint/cmd/flags"
	"github.com/melkeydev/go-blueprint/cmd/template/frontend"
)

func (p *Project) CreateViteReactProject(projectPath string) error {
	if err := checkNpmInstalled(); err != nil {
		return err
	}

	originalDir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get current directory: %w", err)
	}
	defer func() {
		if err := os.Chdir(originalDir); err != nil {
			fmt.Fprintf(os.Stderr, "failed to change back to original directory: %v\n", err)
		}
	}()

	// change into the project directory to run vite command
	err = os.Chdir(projectPath)
	if err != nil {
		fmt.Println("failed to change into project directory: %w", err)
	}

	// the interactive vite command will not work as we can't interact with it
	fmt.Println("Installing create-vite (using cache if available)...")
	cmd := exec.Command("npm", "create", "vite@latest", "frontend", "--",
		"--template", "react-ts",
		"--prefer-offline",
		"--no-fund")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to use create-vite: %w", err)
	}

	frontendPath := filepath.Join(projectPath, "frontend")
	if err := os.MkdirAll(frontendPath, 0755); err != nil {
		return fmt.Errorf("failed to create frontend directory: %w", err)
	}

	if err := os.Chdir(frontendPath); err != nil {
		return fmt.Errorf("failed to change to frontend directory: %w", err)
	}

	srcDir := filepath.Join(frontendPath, "src")
	if err := os.MkdirAll(srcDir, 0755); err != nil {
		return fmt.Errorf("failed to create src directory: %w", err)
	}

	if err := os.WriteFile(filepath.Join(srcDir, "App.tsx"), frontend.ReactAppfile(), 0644); err != nil {
		return fmt.Errorf("failed to write App.tsx template: %w", err)
	}

	// Create the global `.env` file from the template
	err = p.CreateFileWithInjection("", projectPath, ".env", "env")
	if err != nil {
		return fmt.Errorf("failed to create global .env file: %w", err)
	}

	// Read from the global `.env` file and create the frontend-specific `.env`
	globalEnvPath := filepath.Join(projectPath, ".env")
	vitePort := "8080" // Default fallback

	// Read the global .env file
	if data, err := os.ReadFile(globalEnvPath); err == nil {
		lines := strings.Split(string(data), "\n")
		for _, line := range lines {
			if strings.HasPrefix(line, "PORT=") {
				vitePort = strings.SplitN(line, "=", 2)[1] // Get the backend port value
				break
			}
		}
	}

	// Use a template to generate the frontend .env file
	frontendEnvContent := fmt.Sprintf("VITE_PORT=%s\n", vitePort)
	if err := os.WriteFile(filepath.Join(frontendPath, ".env"), []byte(frontendEnvContent), 0644); err != nil {
		return fmt.Errorf("failed to create frontend .env file: %w", err)
	}

	// Handle Tailwind configuration if selected
	if p.FrontendOptions[string(flags.Tailwind)] && p.FrontendFramework == flags.React {
		fmt.Println("Installing Tailwind dependencies (using cache if available)...")
		cmd := exec.Command("npm", "install",
			"--prefer-offline",
			"--no-fund",
			"tailwindcss", "postcss", "autoprefixer")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			return fmt.Errorf("failed to install Tailwind: %w", err)
		}

		fmt.Println("Initializing Tailwind...")
		cmd = exec.Command("npx", "--prefer-offline", "tailwindcss", "init", "-p")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			return fmt.Errorf("failed to initialize Tailwind: %w", err)
		}

		// use the tailwind config file
		err = os.WriteFile("tailwind.config.js", frontend.ReactTailwindConfigTemplate(), 0644)
		if err != nil {
			return fmt.Errorf("failed to write tailwind config: %w", err)
		}

		srcDir := filepath.Join(frontendPath, "src")
		if err := os.MkdirAll(srcDir, 0755); err != nil {
			return fmt.Errorf("failed to create src directory: %w", err)
		}

		err = os.WriteFile(filepath.Join(srcDir, "index.css"), frontend.InputCssTemplateReact(), 0644)
		if err != nil {
			return fmt.Errorf("failed to update index.css: %w", err)
		}

		if err := os.WriteFile(filepath.Join(srcDir, "App.tsx"), frontend.ReactTailwindAppfile(), 0644); err != nil {
			return fmt.Errorf("failed to write App.tsx template: %w", err)
		}

		if err := os.Remove(filepath.Join(srcDir, "App.css")); err != nil {
			// Don't return error if file doesn't exist
			if !os.IsNotExist(err) {
				return fmt.Errorf("failed to remove App.css: %w", err)
			}
		}
	}

	return nil
}

func (p *Project) CreateHtmxTemplates() {
	routesPlaceHolder := ""
	importsPlaceHolder := ""
	if p.FrontendFramework == flags.Htmx {
		routesPlaceHolder += string(p.BackendMap[p.ProjectType].templater.HtmxTemplRoutes())
		importsPlaceHolder += string(p.BackendMap[p.ProjectType].templater.HtmxTemplImports())
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
	p.FrontendTemplates.TemplateRoutes = routeBuffer.String()
	p.FrontendTemplates.TemplateImports = importBuffer.String()
}
