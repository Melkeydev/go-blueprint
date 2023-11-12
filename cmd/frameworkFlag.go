package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

type framework string

const (
	Chi             framework = "chi"
	Gin             framework = "gin"
	fiber           framework = "fiber"
	GorillaMux      framework = "gorilla/mux"
	HttpRouter      framework = "httprouter"
	StandardLibrary framework = "standard-library"
	Echo            framework = "echo"
)

func (f *framework) String() string {
	return string(*f)
}

func (f *framework) Type() string {
	return "framework"
}

func (f *framework) Set(value string) error {
	switch value {
	case "chi", "gin", "fiber", "gorilla/mux", "httprouter", "standard-library", "echo":
		*f = framework(value)
		return nil
	default:
		return fmt.Errorf("Framework to use. Allowed values: %s", strings.Join(allowedProjectTypes, ", "))
	}
}

func frameworkCompletion(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
    return allowedProjectTypes, cobra.ShellCompDirectiveDefault
}
