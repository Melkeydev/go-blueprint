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
			nonInteractiveCommand = fmt.Sprintf("%s --%s %s", nonInteractiveCommand, flag.Name, flag.Value.String())
		}
	}

	flagSet.SortFlags = false
	flagSet.VisitAll(visitFn)

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

func GoTidy(appDir string) error {
	err := ExecuteCmd("go", []string{"mod", "tidy"}, appDir)
	if err != nil {
		return err
	}
	return nil
}

type LineNumbers struct {
	importLine int
	route      int
	fsRoute    int
}

func AddHTMXImports(framework string, fileBytes []byte, projectName string) []byte {
	lineNumbers := make(map[string]LineNumbers)

	lineNumbers["httprouter"] = LineNumbers{
		importLine: 8,
		route:      14,
		fsRoute:    13,
	}

	// Open the file for reading.
	file := string(fileBytes)
	lines := strings.Split(file, "\n")

	// Insert the new line at the desired index.
	routeIndex := lineNumbers[framework].route - 1 // Change this to the line number where you want to insert - 1
	newLine := `  r.Handler(http.MethodGet, "/web", templ.Handler(web.HelloForm()))
	r.HandlerFunc(http.MethodPost, "/hello", web.HelloWebHandler)`
	lines = append(lines[:routeIndex], append([]string{newLine}, lines[routeIndex:]...)...)

	fsRouteIndex := lineNumbers[framework].fsRoute - 1
	newLine = `fileServer := http.FileServer(http.FS(web.Files))
	r.Handler(http.MethodGet, "/js/*filepath", fileServer)`
	lines = append(lines[:fsRouteIndex], append([]string{newLine}, lines[routeIndex:]...)...)

	importIndex := lineNumbers[framework].importLine - 1 // Change this to the line number where you want to insert - 1
	newLine = fmt.Sprintf("  \"%s/cmd/web\"", projectName)
	newLine += "\n  \"github.com/a-h/templ\""
	lines = append(lines[:importIndex], append([]string{newLine}, lines[importIndex:]...)...)

	// Join the lines back together.
	output := strings.Join(lines, "\n")

	return []byte(output)
}
