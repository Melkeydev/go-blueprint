// Package utils provides extra utility
// for the program
package utils

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"

	"github.com/spf13/pflag"
)

const ProgramName = "go-blueprint"

// NonInteractiveCommand creates the command string from a flagSet
// to be used for getting the equivalent non-interactive shell command
func NonInteractiveCommand(use string, flagSet *pflag.FlagSet) string {
	nonInteractiveCommand := fmt.Sprintf("%s %s", ProgramName, use)

	visitFn := func(flag *pflag.Flag) {
		if flag.Name != "help" {
			if flag.Name == "feature" {
				featureFlagsString := ""
				// Creates string representation for the feature flags to be
				// concatenated with the nonInteractiveCommand
				for _, k := range strings.Split(flag.Value.String(), ",") {
					if k != "" {
						featureFlagsString += fmt.Sprintf(" --feature %s", k)
					}
				}
				nonInteractiveCommand += featureFlagsString
			} else {
				nonInteractiveCommand = fmt.Sprintf("%s --%s %s", nonInteractiveCommand, flag.Name, flag.Value.String())
			}
		}
	}

	flagSet.SortFlags = false
	flagSet.VisitAll(visitFn)

	return nonInteractiveCommand
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

func GoTidy(appDir string) error {
	err := ExecuteCmd("go", []string{"mod", "tidy"}, appDir)
	if err != nil {
		return err
	}
	return nil
}

func PnpmPackageInstall(appDir string) error {
	err := ExecuteCmd("pnpm", []string{"install"}, appDir)
	if err != nil {
		return err
	}
	return nil
}

func CheckGitConfig(key string) (bool, error) {
	cmd := exec.Command("git", "config", "--get", key)
	if err := cmd.Run(); err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			// The command failed to run.
			if exitError.ExitCode() == 1 {
				// The 'git config --get' command returns 1 if the key was not found.
				return false, nil
			}
		}
		// Some other error occurred.
		return false, err
	}
	// The command ran successfully, so the key is set.
	return true, nil
}
