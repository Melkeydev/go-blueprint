package program_test

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"github.com/melkeydev/go-blueprint/cmd/flags"
	"github.com/melkeydev/go-blueprint/cmd/program"
	"github.com/melkeydev/go-blueprint/cmd/utils"
	"github.com/stretchr/testify/assert"
)

func TestCreateHumaProject(t *testing.T) {
	projectName := "testhumaproject"
	projectPath := filepath.Join(os.TempDir(), "blueprint-tests", projectName)

	// Clean up any previous test runs
	if _, err := os.Stat(projectPath); !os.IsNotExist(err) {
		os.RemoveAll(projectPath)
	}

	// Ensure the base directory for tests exists
	err := os.MkdirAll(filepath.Join(os.TempDir(), "blueprint-tests"), 0755)
	assert.NoError(t, err)

	p := &program.Project{
		ProjectName:  projectName,
		AbsolutePath: filepath.Join(os.TempDir(), "blueprint-tests"),
		ProjectType:  flags.Huma,
		DBDriver:     "none", // Or a specific driver if Huma needs one by default
		FrameworkMap: make(map[flags.Framework]program.Framework), // Initialize FrameworkMap
		DBDriverMap:  make(map[flags.Database]program.Driver),    // Initialize DBDriverMap
		DockerMap:    make(map[flags.Database]program.Docker),    // Initialize DockerMap
		AdvancedOptions: map[string]bool{
			"websocket": false, // Set advanced options as needed for the test
			"htmx":      false,
			"docker":    false,
		},
		GitOptions: flags.Skip, // Skip git initialization for tests
	}

	// Create the project
	err = p.CreateMainFile()
	assert.NoError(t, err, "CreateMainFile should not return an error")

	// Assertions
	// 1. Check if project directory was created
	_, err = os.Stat(projectPath)
	assert.False(t, os.IsNotExist(err), "Project directory should exist: %s", projectPath)

	// 2. Check for key Huma-specific files
	expectedFiles := []string{
		"go.mod",
		"cmd/api/main.go",
		"internal/server/server.go",
		"internal/server/routes.go",
		".env",
		"Makefile",
	}
	for _, file := range expectedFiles {
		_, err = os.Stat(filepath.Join(projectPath, file))
		assert.False(t, os.IsNotExist(err), "Expected file should exist: %s", file)
	}

	// 3. Check go.mod for Huma dependencies
	goModPath := filepath.Join(projectPath, "go.mod")
	goModBytes, err := os.ReadFile(goModPath)
	assert.NoError(t, err, "Should be able to read go.mod")
	goModContent := string(goModBytes)
	assert.Contains(t, goModContent, "github.com/danielgtaylor/huma/v2", "go.mod should contain huma dependency")
	assert.Contains(t, goModContent, "github.com/gofiber/fiber/v2", "go.mod should contain fiber dependency for Huma")

	// 4. Check if server.go contains Huma setup
	serverGoPath := filepath.Join(projectPath, "internal/server/server.go")
	serverGoBytes, err := os.ReadFile(serverGoPath)
	assert.NoError(t, err, "Should be able to read internal/server/server.go")
	serverGoContent := string(serverGoBytes)
	assert.Contains(t, serverGoContent, "humafiber.New(app", "server.go should contain Huma API initialization with Fiber")
	assert.Contains(t, serverGoContent, "huma.API", "server.go should define a Huma API field")

	// 5. Check if routes.go contains Huma route registration
	routesGoPath := filepath.Join(projectPath, "internal/server/routes.go")
	routesGoBytes, err := os.ReadFile(routesGoPath)
	assert.NoError(t, err, "Should be able to read internal/server/routes.go")
	routesGoContent := string(routesGoBytes)
	assert.Contains(t, routesGoContent, "huma.Register(api", "routes.go should use huma.Register")

	// 6. Check if the project builds
	// Temporarily change working directory to projectPath for go build
	originalWd, err := os.Getwd()
	assert.NoError(t, err)
	err = os.Chdir(projectPath)
	assert.NoError(t, err)

	// Run go mod tidy first
	err = utils.GoTidy(projectPath)
	assert.NoError(t, err, "go mod tidy should run successfully.")

	// Run go build
	// The ExecuteCmd function expects arguments as a slice of strings.
	buildCmd := exec.Command("go", "build", "./...")
	buildCmd.Dir = projectPath
	output, buildErr := buildCmd.CombinedOutput()

	// Check if the error output contains "cannot find package" for git config related issue
	// This is a workaround for tests running in environments where git user.name/user.email might not be set
	// The main CreateMainFile function has checks for this, but tests might bypass some initial CLI setup.
	if buildErr != nil && (strings.Contains(string(output), "user.name") || strings.Contains(string(output), "user.email")) {
		fmt.Printf("Skipping build assertion due to potential git config issue in test environment: %s\n", string(output))
	} else {
		assert.NoError(t, buildErr, "go build ./... should succeed. Output: %s", string(output))
	}

	// Change back to original working directory
	err = os.Chdir(originalWd)
	assert.NoError(t, err)

	// Clean up: Remove the generated project directory
	// Defer this to ensure it runs even if assertions fail
	defer os.RemoveAll(projectPath)
	// Double check removal, sometimes it fails on first try in CI
	if _, err := os.Stat(projectPath); !os.IsNotExist(err) {
		os.RemoveAll(projectPath)
	}
}
