package flags

import (
	"fmt"
	"strings"
)

type Framework string

// These are all the current frameworks supported. If you want to add one, you
// can simply copy and paste a line here. Do not forget to also add it into the
// AllowedProjectTypes slice too!
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
	// Contains isn't available in 1.20 yet
	// if AllowedProjectTypes.Contains(value) {
	for _, project := range AllowedProjectTypes {
		if project == value {
			*f = Framework(value)
			return nil
		}
	}

	return fmt.Errorf("Framework to use. Allowed values: %s", strings.Join(AllowedProjectTypes, ", "))
}
