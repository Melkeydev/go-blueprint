package flags

import (
	"fmt"
	"strings"
)

type FrontendAdvanced []string

const (
	Tailwind string = "tailwind"
)

var AllowedFrontendAdvanced = []string{string(Tailwind)}

func (f FrontendAdvanced) String() string {
	return strings.Join(f, ",")
}

func (f *FrontendAdvanced) Type() string {
	return "FrontendAdvanced"
}

func (f *FrontendAdvanced) Set(value string) error {
	// Contains isn't available in 1.20 yet
	// if FrontendAdvanced.Contains(value) {
	for _, frontendAdvancedFeature := range AllowedFrontendAdvanced {
		if frontendAdvancedFeature == value {
			*f = append(*f, frontendAdvancedFeature)
			return nil
		}
	}

	return fmt.Errorf("advanced Feature to use. Allowed values: %s", strings.Join(AllowedFrontendAdvanced, ", "))
}
