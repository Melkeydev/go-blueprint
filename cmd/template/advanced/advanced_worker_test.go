package advanced

import (
	"strings"
	"testing"
)

func TestWorkerTemplatesExist(t *testing.T) {
	// Test that worker templates are properly embedded and accessible

	// Test WorkerMainTemplate
	mainTemplate := WorkerMainTemplate()
	if len(mainTemplate) == 0 {
		t.Error("WorkerMainTemplate should not be empty")
	}

	// Test WorkerHelloWorldTaskTemplate
	taskTemplate := WorkerHelloWorldTaskTemplate()
	if len(taskTemplate) == 0 {
		t.Error("WorkerHelloWorldTaskTemplate should not be empty")
	}

	// Test WorkerEnvTemplate
	envTemplate := WorkerEnvTemplate()
	if len(envTemplate) == 0 {
		t.Error("WorkerEnvTemplate should not be empty")
	}
}

func TestWorkerMainTemplateContent(t *testing.T) {
	template := WorkerMainTemplate()
	templateStr := string(template)

	// Check for essential components in the main worker template
	expectedContent := []string{
		"package main",
		"github.com/hibiken/asynq",
		"github.com/joho/godotenv",
		"{{.ProjectName}}/cmd/worker/tasks",
		"asynq.NewServeMux()",
		"tasks.HandleHelloWorldTask",
		"server := asynq.NewServer(",
		"mux.HandleFunc",
		"getEnv(",
		"getEnvInt(",
	}

	for _, expected := range expectedContent {
		if !strings.Contains(templateStr, expected) {
			t.Errorf("WorkerMainTemplate missing expected content: %s", expected)
		}
	}
}

func TestWorkerTaskTemplateContent(t *testing.T) {
	template := WorkerHelloWorldTaskTemplate()
	templateStr := string(template)

	// Check for essential components in the task template
	expectedContent := []string{
		"package tasks",
		"TypeHelloWorld",
		"HelloWorldPayload",
		"NewHelloWorldTask",
		"HandleHelloWorldTask",
		"github.com/hibiken/asynq",
		"context.Context",
		"json.Marshal",
		"json.Unmarshal",
	}

	for _, expected := range expectedContent {
		if !strings.Contains(templateStr, expected) {
			t.Errorf("WorkerHelloWorldTaskTemplate missing expected content: %s", expected)
		}
	}
}

func TestWorkerEnvTemplateContent(t *testing.T) {
	template := WorkerEnvTemplate()
	templateStr := string(template)

	// Check for essential environment variables
	expectedContent := []string{
		"Worker Configuration",
		"REDIS_ADDR",
		"REDIS_PASSWORD",
		"REDIS_DB",
		"localhost:6379",
	}

	for _, expected := range expectedContent {
		if !strings.Contains(templateStr, expected) {
			t.Errorf("WorkerEnvTemplate missing expected content: %s", expected)
		}
	}
}

func TestWorkerTemplateConsistency(t *testing.T) {
	// Test that templates reference each other correctly

	mainTemplate := string(WorkerMainTemplate())
	taskTemplate := string(WorkerHelloWorldTaskTemplate())

	// Check that main template references the task correctly
	if !strings.Contains(mainTemplate, "tasks.TypeHelloWorld") {
		t.Error("Main template should reference tasks.TypeHelloWorld")
	}

	if !strings.Contains(mainTemplate, "tasks.HandleHelloWorldTask") {
		t.Error("Main template should reference tasks.HandleHelloWorldTask")
	}

	// Check that task template exports the expected symbols
	if !strings.Contains(taskTemplate, "TypeHelloWorld") {
		t.Error("Task template should define TypeHelloWorld constant")
	}

	if !strings.Contains(taskTemplate, "HandleHelloWorldTask") {
		t.Error("Task template should define HandleHelloWorldTask function")
	}

	// Check that the task type constant matches between templates
	// The main template should use the same task type as defined in the task template
	if !strings.Contains(taskTemplate, `TypeHelloWorld = "hello_world"`) {
		t.Error("Task template should define TypeHelloWorld as 'hello_world'")
	}
}

func TestWorkerTemplateProjectNamePlaceholder(t *testing.T) {
	// Test that templates contain the correct project name placeholder

	mainTemplate := string(WorkerMainTemplate())

	// Check for the Go template placeholder
	expectedPlaceholder := "{{.ProjectName}}/cmd/worker/tasks"
	if !strings.Contains(mainTemplate, expectedPlaceholder) {
		t.Errorf("Main template should contain project name placeholder: %s", expectedPlaceholder)
	}

	// Ensure no hardcoded project names exist
	forbiddenContent := []string{
		"test-project",
		"example-project",
		"my-project",
	}

	for _, forbidden := range forbiddenContent {
		if strings.Contains(mainTemplate, forbidden) {
			t.Errorf("Main template should not contain hardcoded project name: %s", forbidden)
		}
	}
}

func TestWorkerTemplateValidGoSyntax(t *testing.T) {
	// Test that templates contain valid Go syntax structure

	templates := map[string][]byte{
		"main": WorkerMainTemplate(),
		"task": WorkerHelloWorldTaskTemplate(),
	}

	for name, template := range templates {
		templateStr := string(template)

		// Check for basic Go file structure
		if !strings.Contains(templateStr, "package ") {
			t.Errorf("%s template should contain package declaration", name)
		}

		if !strings.Contains(templateStr, "import (") {
			t.Errorf("%s template should contain import statement", name)
		}

		// Count braces to ensure they're balanced (simple check)
		openBraces := strings.Count(templateStr, "{") - strings.Count(templateStr, "{{")
		closeBraces := strings.Count(templateStr, "}") - strings.Count(templateStr, "}}")

		if openBraces != closeBraces {
			t.Errorf("%s template has unbalanced braces: %d open, %d close", name, openBraces, closeBraces)
		}
	}
}

func TestWorkerEnvTemplateValidFormat(t *testing.T) {
	template := WorkerEnvTemplate()
	templateStr := string(template)

	// Split into lines and check format
	lines := strings.Split(templateStr, "\n")

	for i, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue // Skip empty lines
		}

		// Check for comment lines
		if strings.HasPrefix(line, "#") {
			continue // Comments are valid
		}

		// Check for key=value format
		if !strings.Contains(line, "=") {
			t.Errorf("Line %d should be in KEY=VALUE format or be a comment: %s", i+1, line)
		}

		// Check that keys are uppercase
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			key := parts[0]
			if key != strings.ToUpper(key) {
				t.Errorf("Environment variable key should be uppercase: %s", key)
			}
		}
	}
}
