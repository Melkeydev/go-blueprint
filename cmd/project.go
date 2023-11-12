package cmd

import (
	"fmt"
	"strings"

	"github.com/melkeydev/go-blueprint/cmd/program"
	"github.com/melkeydev/go-blueprint/cmd/steps"
	"github.com/melkeydev/go-blueprint/cmd/ui/textinput"
	"github.com/melkeydev/go-blueprint/cmd/utils"
	"github.com/spf13/cobra"
)

func NewProject(cmd *cobra.Command) (*program.Project, steps.Options, bool) {
	options := steps.Options{
		ProjectName: &textinput.Output{},
	}

	flagName := utils.FlagValue(cmd, "name")
	flagFramework := utils.FlagValue(cmd, "framework")

	if flagFramework != "" {
		isValid := utils.IsValidProjectType(flagFramework, AllowedProjectTypes)
		if !isValid {
			cobra.CheckErr(fmt.Errorf("Project type '%s' is not valid. Valid types are: %s", flagFramework, strings.Join(AllowedProjectTypes, ", ")))
		}
	}

	isInteractive := !utils.HasChangedFlag(cmd.Flags())

	project := &program.Project{
		FrameworkMap: make(map[string]program.Framework),
		ProjectName:  flagName,
		ProjectType:  strings.ReplaceAll(flagFramework, "-", " "),
	}

	return project, options, isInteractive
}
