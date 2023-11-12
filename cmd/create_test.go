package cmd_test

import (
	"strings"
	"testing"

	"github.com/melkeydev/go-blueprint/cmd"
	"github.com/melkeydev/go-blueprint/cmd/program"
)

func TestCheckMissingProjectTypes(t *testing.T) {
	project := program.Project{
		FrameworkMap: make(map[string]program.Framework),
	}

	project.CreateFrameworkMap()

	missingTypes := checkMissingProjectTypes(project.FrameworkMap, cmd.AllowedProjectTypes)

	if len(missingTypes) != 0 {
		t.Errorf("Expected missingTypes to be empty, got %v", missingTypes)
	}
}

func checkMissingProjectTypes(frameworkMap map[string]program.Framework, types []string) []string {
	var missingTypes []string

	for _, t := range types {
		t := strings.ReplaceAll(t, "-", " ")

		if _, ok := frameworkMap[t]; !ok {
			missingTypes = append(missingTypes, t)
		}

		if _, ok := frameworkMap[t]; !ok {
			missingTypes = append(missingTypes, t)
		}

	}

	return missingTypes
}
