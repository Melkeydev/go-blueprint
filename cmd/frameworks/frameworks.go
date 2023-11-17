package frameworks

import (
	"fmt"
	"strings"
)

type Framework string

// These are all the current frameworks supported, if you want to add one you
// can simply copy and past a line here. Do not forget to also add it into the
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

// This interface is required on a type to make it useable as a flag
//
//	type Value interface {
//		String() string
//		Set(string) error
//		Type() string
//	}
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
