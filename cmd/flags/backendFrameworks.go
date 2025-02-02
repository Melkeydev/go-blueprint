package flags

import (
	"fmt"
	"strings"
)

type BackendFramework string

// These are all the current backends supported. If you want to add one, you
// can simply copy and paste a line here. Do not forget to also add it into the
// AllowedBackendTypes slice too!
const (
	Chi             BackendFramework = "chi"
	Gin             BackendFramework = "gin"
	Fiber           BackendFramework = "fiber"
	GorillaMux      BackendFramework = "gorilla/mux"
	StandardLibrary BackendFramework = "standard-library"
	Echo            BackendFramework = "echo"
)

var AllowedBackendFrameworkTypes = []string{string(Chi), string(Gin), string(Fiber), string(GorillaMux), string(StandardLibrary), string(Echo)}

func (f BackendFramework) String() string {
	return string(f)
}

func (f *BackendFramework) Type() string {
	return "BackendFramework"
}

func (f *BackendFramework) Set(value string) error {
	// Contains isn't available in 1.20 yet
	// if AllowedBackendTypes.Contains(value) {
	for _, project := range AllowedBackendFrameworkTypes {
		if project == value {
			*f = BackendFramework(value)
			return nil
		}
	}

	return fmt.Errorf("BackendFramework to use. Allowed values: %s", strings.Join(AllowedBackendFrameworkTypes, ", "))
}
