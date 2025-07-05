package flags

import (
	"testing"
)

func TestAdvancedFeaturesWorker(t *testing.T) {
	// Test that Worker constant is properly defined
	if Worker != "worker" {
		t.Errorf("Worker constant expected 'worker', got '%s'", Worker)
	}
}

func TestAllowedAdvancedFeaturesContainsWorker(t *testing.T) {
	// Test that Worker is included in AllowedAdvancedFeatures
	found := false
	for _, feature := range AllowedAdvancedFeatures {
		if feature == Worker {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("Worker feature not found in AllowedAdvancedFeatures slice")
	}
}

func TestAdvancedFeaturesSetWorker(t *testing.T) {
	var features AdvancedFeatures

	// Test valid worker feature
	err := features.Set("worker")
	if err != nil {
		t.Errorf("Expected no error when setting 'worker' feature, got: %v", err)
	}

	// Check if worker was added to the slice
	found := false
	for _, feature := range features {
		if feature == "worker" {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("Worker feature was not added to the AdvancedFeatures slice")
	}
}

func TestAdvancedFeaturesSetInvalidFeature(t *testing.T) {
	var features AdvancedFeatures

	// Test invalid feature
	err := features.Set("invalid-feature")
	if err == nil {
		t.Errorf("Expected error when setting invalid feature, got nil")
	}

	// Check that the error message contains the allowed values
	expectedSubstring := "advanced Feature to use. Allowed values:"
	if err != nil && !contains(err.Error(), expectedSubstring) {
		t.Errorf("Error message should contain '%s', got: %s", expectedSubstring, err.Error())
	}
}

func TestAdvancedFeaturesSetMultipleFeatures(t *testing.T) {
	var features AdvancedFeatures

	// Test setting multiple features including worker
	testFeatures := []string{"worker", "docker", "htmx"}

	for _, feature := range testFeatures {
		err := features.Set(feature)
		if err != nil {
			t.Errorf("Expected no error when setting '%s' feature, got: %v", feature, err)
		}
	}

	// Verify all features were added
	if len(features) != len(testFeatures) {
		t.Errorf("Expected %d features, got %d", len(testFeatures), len(features))
	}

	// Verify worker is in the list
	found := false
	for _, feature := range features {
		if feature == "worker" {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("Worker feature not found in the features list after setting multiple features")
	}
}

func TestAdvancedFeaturesString(t *testing.T) {
	var features AdvancedFeatures

	// Test empty features
	result := features.String()
	if result != "" {
		t.Errorf("Expected empty string for empty features, got: '%s'", result)
	}

	// Test with worker feature
	features.Set("worker")
	result = features.String()
	if result != "worker" {
		t.Errorf("Expected 'worker' for single feature, got: '%s'", result)
	}

	// Test with multiple features
	features.Set("docker")
	result = features.String()
	expected := "worker,docker"
	if result != expected {
		t.Errorf("Expected '%s' for multiple features, got: '%s'", expected, result)
	}
}

func TestAdvancedFeaturesType(t *testing.T) {
	var features AdvancedFeatures

	// Test Type method
	result := features.Type()
	expected := "AdvancedFeatures"
	if result != expected {
		t.Errorf("Expected '%s' from Type(), got: '%s'", expected, result)
	}
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
