// Package utils provides extra utility
// for the program
package utils

import (
	"bytes"
	"fmt"
	"os/exec"
	"regexp"
	"strings"

	"github.com/spf13/pflag"
)

const ProgramName = "go-blueprint"

// NonInteractiveCommand creates the command string from a flagSet
// to be used for getting the equivalent non-interactive shell command
func NonInteractiveCommand(use string, flagSet *pflag.FlagSet) string {
	nonInteractiveCommand := fmt.Sprintf("%s %s", ProgramName, use)

	// Helper function to get flag prefix (-n or --name)
	getFlagPrefix := func(flag *pflag.Flag) string {
		if flag.Shorthand != "" {
			return fmt.Sprintf("-%s", flag.Shorthand)
		}
		return fmt.Sprintf("--%s", flag.Name)
	}

	// name flag
	flagSet.VisitAll(func(flag *pflag.Flag) {
		if flag.Name == "name" && flag.Value.String() != "" {
			nonInteractiveCommand = fmt.Sprintf("%s %s %s", nonInteractiveCommand,
				getFlagPrefix(flag), flag.Value.String())
		}
	})

	// main flags (excluding name, git, feature, frontend-advanced, and advanced)
	flagSet.VisitAll(func(flag *pflag.Flag) {
		if flag.Name != "help" && flag.Name != "feature" &&
			flag.Name != "frontend-advanced" && flag.Name != "git" &&
			flag.Name != "name" && flag.Name != "advanced" {
			if flag.Value.Type() == "bool" {
				if flag.Value.String() == "true" {
					nonInteractiveCommand = fmt.Sprintf("%s %s", nonInteractiveCommand,
						getFlagPrefix(flag))
				}
			} else {
				if flag.Value.String() != "" {
					nonInteractiveCommand = fmt.Sprintf("%s %s %s", nonInteractiveCommand,
						getFlagPrefix(flag), flag.Value.String())
				}
			}
		}
	})

	// frontend-advanced flags
	flagSet.VisitAll(func(flag *pflag.Flag) {
		if flag.Name == "frontend-advanced" && flag.Value.String() != "" {
			nonInteractiveCommand = fmt.Sprintf("%s %s %s", nonInteractiveCommand,
				getFlagPrefix(flag), flag.Value.String())
		}
	})

	// advanced flag and features together
	var hasAdvanced bool
	flagSet.VisitAll(func(flag *pflag.Flag) {
		if flag.Name == "advanced" && flag.Value.String() == "true" {
			hasAdvanced = true
			nonInteractiveCommand = fmt.Sprintf("%s %s", nonInteractiveCommand,
				getFlagPrefix(flag))
		}
	})

	if hasAdvanced {
		flagSet.VisitAll(func(flag *pflag.Flag) {
			if flag.Name == "feature" {
				featureFlags := strings.Split(flag.Value.String(), ",")
				for _, k := range featureFlags {
					if k != "" {
						nonInteractiveCommand = fmt.Sprintf("%s --feature %s", nonInteractiveCommand, k)
					}
				}
			}
		})
	}

	// git flag
	flagSet.VisitAll(func(flag *pflag.Flag) {
		if flag.Name == "git" && flag.Value.String() != "" {
			nonInteractiveCommand = fmt.Sprintf("%s %s %s", nonInteractiveCommand,
				getFlagPrefix(flag), flag.Value.String())
		}
	})

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

// GoModReplace runs "go mod edit -replace" in the selected
// replace_payload e.g: github.com/gocql/gocql=github.com/scylladb/gocql@v1.14.4
func GoModReplace(appDir string, replace string) error {
	if err := ExecuteCmd("go",
		[]string{"mod", "edit", "-replace", replace},
		appDir,
	); err != nil {
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

// ValidateModuleName returns true if it's a valid module name.
// It allows any number of / and . characters in between.
func ValidateModuleName(moduleName string) bool {
	matched, _ := regexp.MatchString("^[a-zA-Z0-9_-]+(?:[\\/.][a-zA-Z0-9_-]+)*$", moduleName)
	return matched
}

// GetRootDir returns the project directory name from the module path.
// Returns the last token by splitting the moduleName with /
func GetRootDir(moduleName string) string {
	tokens := strings.Split(moduleName, "/")
	return tokens[len(tokens)-1]
}
