package steps

import (
	"testing"

	"github.com/melkeydev/go-blueprint/cmd/flags"
)

func TestWorkerInAdvancedSteps(t *testing.T) {
	// Initialize steps with default values
	steps := InitSteps(flags.Framework("gin"), flags.Database("mysql"))

	// Check that advanced steps exist
	advancedStep, exists := steps.Steps["advanced"]
	if !exists {
		t.Fatal("Advanced step should exist in steps")
	}

	// Check that advanced step has the correct properties
	if advancedStep.StepName != "Advanced Features" {
		t.Errorf("Expected step name 'Advanced Features', got '%s'", advancedStep.StepName)
	}

	if advancedStep.Headers != "Which advanced features do you want?" {
		t.Errorf("Expected headers 'Which advanced features do you want?', got '%s'", advancedStep.Headers)
	}

	// Check that worker option exists in advanced features
	found := false
	var workerItem Item
	for _, option := range advancedStep.Options {
		if option.Flag == "Worker" {
			found = true
			workerItem = option
			break
		}
	}

	if !found {
		t.Error("Worker option not found in advanced features")
	}

	// Verify worker item properties
	if workerItem.Title != "Background Worker" {
		t.Errorf("Expected worker title 'Background Worker', got '%s'", workerItem.Title)
	}

	expectedDesc := "Add background worker implementation using Asynq with Redis"
	if workerItem.Desc != expectedDesc {
		t.Errorf("Expected worker description '%s', got '%s'", expectedDesc, workerItem.Desc)
	}
}

func TestAllAdvancedFeaturesInSteps(t *testing.T) {
	// Initialize steps
	steps := InitSteps(flags.Framework("gin"), flags.Database("mysql"))
	advancedStep := steps.Steps["advanced"]

	// Create a map of expected features based on the flags
	expectedFeatures := map[string]bool{
		"React":        true,
		"Htmx":         true,
		"GitHubAction": true,
		"Websocket":    true,
		"Tailwind":     true,
		"Docker":       true,
		"Worker":       true,
	}

	// Check that all expected features are present
	foundFeatures := make(map[string]bool)
	for _, option := range advancedStep.Options {
		foundFeatures[option.Flag] = true
	}

	for feature := range expectedFeatures {
		if !foundFeatures[feature] {
			t.Errorf("Expected feature '%s' not found in advanced steps", feature)
		}
	}

	// Check that worker is specifically included
	if !foundFeatures["Worker"] {
		t.Error("Worker feature should be included in advanced steps")
	}
}

func TestStepsInitialization(t *testing.T) {
	// Test with different framework and database combinations
	testCases := []struct {
		framework flags.Framework
		database  flags.Database
	}{
		{flags.Framework("gin"), flags.Database("mysql")},
		{flags.Framework("chi"), flags.Database("postgres")},
		{flags.Framework("fiber"), flags.Database("none")},
	}

	for _, tc := range testCases {
		steps := InitSteps(tc.framework, tc.database)

		// Verify that steps are properly initialized
		if steps == nil {
			t.Errorf("Steps should not be nil for framework=%s, database=%s", tc.framework, tc.database)
			continue
		}

		// Verify that advanced step exists regardless of framework/database
		advancedStep, exists := steps.Steps["advanced"]
		if !exists {
			t.Errorf("Advanced step should exist for framework=%s, database=%s", tc.framework, tc.database)
			continue
		}

		// Verify that worker option exists in all cases
		found := false
		for _, option := range advancedStep.Options {
			if option.Flag == "Worker" {
				found = true
				break
			}
		}

		if !found {
			t.Errorf("Worker option should exist for framework=%s, database=%s", tc.framework, tc.database)
		}
	}
}

func TestWorkerOptionProperties(t *testing.T) {
	steps := InitSteps(flags.Framework("gin"), flags.Database("mysql"))
	advancedStep := steps.Steps["advanced"]

	// Find the worker option
	var workerOption Item
	found := false
	for _, option := range advancedStep.Options {
		if option.Flag == "Worker" {
			workerOption = option
			found = true
			break
		}
	}

	if !found {
		t.Fatal("Worker option not found")
	}

	// Test all properties of the worker option
	testCases := []struct {
		property string
		expected string
		actual   string
	}{
		{"Flag", "Worker", workerOption.Flag},
		{"Title", "Background Worker", workerOption.Title},
		{"Desc", "Add background worker implementation using Asynq with Redis", workerOption.Desc},
	}

	for _, tc := range testCases {
		if tc.actual != tc.expected {
			t.Errorf("Worker option %s: expected '%s', got '%s'", tc.property, tc.expected, tc.actual)
		}
	}
}

func TestStepsStructure(t *testing.T) {
	steps := InitSteps(flags.Framework("gin"), flags.Database("mysql"))

	// Verify that Steps structure is properly initialized
	if steps.Steps == nil {
		t.Fatal("Steps.Steps map should not be nil")
	}

	// Check that essential steps exist
	essentialSteps := []string{"framework", "driver", "advanced", "git"}
	for _, stepName := range essentialSteps {
		if _, exists := steps.Steps[stepName]; !exists {
			t.Errorf("Essential step '%s' should exist", stepName)
		}
	}

	// Verify advanced step has options
	advancedStep := steps.Steps["advanced"]
	if len(advancedStep.Options) == 0 {
		t.Error("Advanced step should have options")
	}

	// Verify that all options have required properties
	for i, option := range advancedStep.Options {
		if option.Flag == "" {
			t.Errorf("Option %d should have a Flag", i)
		}
		if option.Title == "" {
			t.Errorf("Option %d should have a Title", i)
		}
		if option.Desc == "" {
			t.Errorf("Option %d should have a Description", i)
		}
	}
}
