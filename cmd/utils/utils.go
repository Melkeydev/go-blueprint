// Package utils provides extra utility
// for the program
package utils

import (
	"bytes"
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

const ProgramName = "go-blueprint"

// NonInteractiveCommand creates the command string from a flagSet
// to be used for getting the equivalent non-interactive shell command
func NonInteractiveCommand(cmd *cobra.Command) string {
	nonInteractiveCommand := fmt.Sprintf("%s %s", ProgramName, cmd.Name())

	visitFn := func(flag *pflag.Flag) {
		if flag.Name != "help" && flag.Name != "name" {
			nonInteractiveCommand = fmt.Sprintf("%s --%s %s", nonInteractiveCommand, flag.Name, flag.Value.String())
		}
	}

	flagSet := cmd.Flags()
	flagSet.SortFlags = false

	flagSet.VisitAll(visitFn)

	flag := cmd.Flag("name")
	nonInteractiveCommand = fmt.Sprintf("%s --%s %s", nonInteractiveCommand, flag.Name, flag.Value.String())

	return nonInteractiveCommand
}

func HasChangedFlag(flagSet *pflag.FlagSet) bool {
	hasChangedFlag := false
	flagSet.Visit(func(_ *pflag.Flag) {
		hasChangedFlag = true
	})
	return hasChangedFlag
}

// ExecuteCmd provides a shorthand way to run a shell command
func ExecuteCmd(name string, args []string, dir string) error {
	command := exec.Command(name, args...)
	command.Dir = dir
	var out bytes.Buffer
	command.Stdout = &out
	if err := command.Run(); err != nil {
		return err
	}
	return nil
}

// InitGoMod initializes go.mod with the given project name
// in the selected directory
func InitGoMod(projectName string, appDir string) error {
	if err := ExecuteCmd("go",
		[]string{"mod", "init", projectName},
		appDir); err != nil {
		return err
	}

	return nil
}

// GoGetPackage runs "go get" for a given package in the
// selected directory
func GoGetPackage(appDir string, packages []string) error {
	for _, packageName := range packages {
		if err := ExecuteCmd("go",
			[]string{"get", "-u", packageName},
			appDir); err != nil {
			return err
		}
	}

	return nil
}

// GoFmt runs "gofmt" in a selected directory using the
// simplify and overwrite flags
func GoFmt(appDir string) error {
	if err := ExecuteCmd("gofmt",
		[]string{"-s", "-w", "."},
		appDir); err != nil {
		return err
	}

	return nil
}
