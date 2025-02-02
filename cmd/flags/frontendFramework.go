package flags

import (
	"fmt"
	"strings"
)

type FrontendFramework string

// These are all the current frameworks supported. If you want to add one, you
// can simply copy and paste a line here. Do not forget to also add it into the
// AllowedFrontedTypes slice too!
const (
	Htmx  FrontendFramework = "htmx"
	React FrontendFramework = "react"
)

var AllowedFrontendTypes = []string{string(Htmx), string(React)}

func (f FrontendFramework) String() string {
	return string(f)
}

func (f *FrontendFramework) Type() string {
	return "Frontendframework"
}

func (f *FrontendFramework) Set(value string) error {
	// Contains isn't available in 1.20 yet
	// if AllowedFrontedTypes.Contains(value) {
	for _, frontendFrameworks := range AllowedFrontendTypes {
		if frontendFrameworks == value {
			*f = FrontendFramework(value)
			return nil
		}
	}

	return fmt.Errorf("Frontend framework to use. Allowed values: %s", strings.Join(AllowedFrontendTypes, ", "))
}
