package flags

import (
	"fmt"
	"strings"
)

type Backend string

// These are all the current backends supported. If you want to add one, you
// can simply copy and paste a line here. Do not forget to also add it into the
// AllowedBackedTypes slice too!
const (
	Chi             Backend = "chi"
	Gin             Backend = "gin"
	Fiber           Backend = "fiber"
	GorillaMux      Backend = "gorilla/mux"
	StandardLibrary Backend = "standard-library"
	Echo            Backend = "echo"
)

var AllowedBackedTypes = []string{string(Chi), string(Gin), string(Fiber), string(GorillaMux), string(StandardLibrary), string(Echo)}

func (f Backend) String() string {
	return string(f)
}

func (f *Backend) Type() string {
	return "Backend"
}

func (f *Backend) Set(value string) error {
	// Contains isn't available in 1.20 yet
	// if AllowedBackedTypes.Contains(value) {
	for _, project := range AllowedBackedTypes {
		if project == value {
			*f = Backend(value)
			return nil
		}
	}

	return fmt.Errorf("Backend to use. Allowed values: %s", strings.Join(AllowedBackedTypes, ", "))
}
