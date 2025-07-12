package program

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/melkeydev/go-blueprint/cmd/flags"
	"github.com/melkeydev/go-blueprint/cmd/template/advanced"
)

func TestCreateWorkerFiles(t *testing.T) {
	// Create a temporary directory for testing
	tempDir, err := os.MkdirTemp("", "worker-test")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Create a basic go.mod file to prevent go get errors
	goModPath := filepath.Join(tempDir, "go.mod")
	goModContent := "module test-worker-project\n\ngo 1.21\n"
	err = os.WriteFile(goModPath, []byte(goModContent), 0644)
	if err != nil {
		t.Fatalf("Failed to create go.mod: %v", err)
	}

	// Create a project with worker feature enabled
	project := &Project{
		ProjectName:     "test-worker-project",
		AdvancedOptions: make(map[string]bool),
	}
	project.AdvancedOptions[string(flags.Worker)] = true

	// Test CreateWorkerFiles method
	err = project.CreateWorkerFiles(tempDir)
	if err != nil {
		t.Fatalf("CreateWorkerFiles failed: %v", err)
	}

	// Verify that the expected directories were created
	expectedDirs := []string{
		filepath.Join(tempDir, "cmd", "worker"),
		filepath.Join(tempDir, "cmd", "worker", "tasks"),
	}

	for _, dir := range expectedDirs {
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			t.Errorf("Expected directory does not exist: %s", dir)
		}
	}

	// Verify that the expected files were created
	expectedFiles := []string{
		filepath.Join(tempDir, "cmd", "worker", "main.go"),
		filepath.Join(tempDir, "cmd", "worker", "tasks", "hello_world_task.go"),
	}

	for _, file := range expectedFiles {
		if _, err := os.Stat(file); os.IsNotExist(err) {
			t.Errorf("Expected file does not exist: %s", file)
		}
	}
}

func TestCreateWorkerFilesContent(t *testing.T) {
	// Create a temporary directory for testing
	tempDir, err := os.MkdirTemp("", "worker-content-test")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Create a basic go.mod file to prevent go get errors
	goModPath := filepath.Join(tempDir, "go.mod")
	goModContent := "module test-content-project\n\ngo 1.21\n"
	err = os.WriteFile(goModPath, []byte(goModContent), 0644)
	if err != nil {
		t.Fatalf("Failed to create go.mod: %v", err)
	}

	// Create a project with worker feature enabled
	project := &Project{
		ProjectName:     "test-content-project",
		AdvancedOptions: make(map[string]bool),
	}
	project.AdvancedOptions[string(flags.Worker)] = true

	// Test CreateWorkerFiles method
	err = project.CreateWorkerFiles(tempDir)
	if err != nil {
		t.Fatalf("CreateWorkerFiles failed: %v", err)
	}

	// Test main.go content
	mainGoPath := filepath.Join(tempDir, "cmd", "worker", "main.go")
	mainGoContent, err := os.ReadFile(mainGoPath)
	if err != nil {
		t.Fatalf("Failed to read main.go: %v", err)
	}

	// Check for essential imports and functions in main.go
	mainGoStr := string(mainGoContent)
	expectedMainGoContent := []string{
		"package main",
		"github.com/hibiken/asynq",
		"asynq.NewServeMux()",
		"test-content-project/cmd/worker/tasks",
		"tasks.HandleHelloWorldTask",
	}

	for _, expected := range expectedMainGoContent {
		if !contains(mainGoStr, expected) {
			t.Errorf("main.go missing expected content: %s", expected)
		}
	}

	// Test hello_world_task.go content
	taskPath := filepath.Join(tempDir, "cmd", "worker", "tasks", "hello_world_task.go")
	taskContent, err := os.ReadFile(taskPath)
	if err != nil {
		t.Fatalf("Failed to read hello_world_task.go: %v", err)
	}

	taskStr := string(taskContent)
	expectedTaskContent := []string{
		"package tasks",
		"TypeHelloWorld",
		"HandleHelloWorldTask",
		"HelloWorldPayload",
		"github.com/hibiken/asynq",
	}

	for _, expected := range expectedTaskContent {
		if !contains(taskStr, expected) {
			t.Errorf("hello_world_task.go missing expected content: %s", expected)
		}
	}
}

func TestCreateWorkerFilesWithoutWorkerFeature(t *testing.T) {
	// Create a temporary directory for testing
	tempDir, err := os.MkdirTemp("", "no-worker-test")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Create a basic go.mod file to prevent go get errors
	goModPath := filepath.Join(tempDir, "go.mod")
	goModContent := "module test-no-worker-project\n\ngo 1.21\n"
	err = os.WriteFile(goModPath, []byte(goModContent), 0644)
	if err != nil {
		t.Fatalf("Failed to create go.mod: %v", err)
	}

	// Create a project without worker feature enabled
	project := &Project{
		ProjectName:     "test-no-worker-project",
		AdvancedOptions: make(map[string]bool),
	}
	// Explicitly set worker to false
	project.AdvancedOptions[string(flags.Worker)] = false

	// Test CreateWorkerFiles method (should still work if called directly)
	err = project.CreateWorkerFiles(tempDir)
	if err != nil {
		t.Fatalf("CreateWorkerFiles failed even when called directly: %v", err)
	}

	// Verify that files were still created (since the method was called directly)
	workerMainPath := filepath.Join(tempDir, "cmd", "worker", "main.go")
	if _, err := os.Stat(workerMainPath); os.IsNotExist(err) {
		t.Errorf("Worker files should be created when CreateWorkerFiles is called directly")
	}
}

func TestWorkerEnvironmentVariables(t *testing.T) {
	// Test that worker-specific environment variables are properly defined
	// This test verifies the env template exists and contains expected values

	// Since we can't easily test the template execution without a full setup,
	// we'll test that the template function exists and returns content
	envTemplate := getWorkerEnvTemplate()
	if len(envTemplate) == 0 {
		t.Error("Worker environment template should not be empty")
	}

	envStr := string(envTemplate)
	expectedEnvVars := []string{
		"REDIS_ADDR",
		"REDIS_PASSWORD",
		"REDIS_DB",
		"Worker Configuration",
	}

	for _, expected := range expectedEnvVars {
		if !contains(envStr, expected) {
			t.Errorf("Worker env template missing expected content: %s", expected)
		}
	}
}

// Helper function to get worker env template for testing
func getWorkerEnvTemplate() []byte {
	// Import the actual template from the advanced package
	return advanced.WorkerEnvTemplate()
}

// Helper function to check if a string contains a substring
func contains(s, substr string) bool {
	return len(s) >= len(substr) && (len(substr) == 0 || findSubstring(s, substr))
}

func findSubstring(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
