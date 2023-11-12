package frameworks

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

type Framework string

const (
	Chi             Framework = "chi"
	Gin             Framework = "gin"
	Fiber           Framework = "fiber"
	GorillaMux      Framework = "gorilla/mux"
	HttpRouter      Framework = "httprouter"
	StandardLibrary Framework = "standard-library"
	Echo            Framework = "echo"
)

var AllowedProjectTypes = []string{string(Chi), string(Gin), string(Fiber), string(GorillaMux), string(HttpRouter), string(StandardLibrary), string(Echo)}

func (f Framework) String() string {
	return string(f)
}

func (f *Framework) Type() string {
	return "Framework"
}

func (f *Framework) Set(value string) error {
	switch value {
	case Chi.String(), Gin.String(), Fiber.String(), GorillaMux.String(), HttpRouter.String(), StandardLibrary.String(), Echo.String():
		*f = Framework(value)
		return nil
	default:
		return fmt.Errorf("Framework to use. Allowed values: %s", strings.Join(AllowedProjectTypes, ", "))
	}
}

func FrameworkCompletion(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	return AllowedProjectTypes, cobra.ShellCompDirectiveDefault
}
