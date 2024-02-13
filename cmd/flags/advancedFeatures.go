package flags

import (
	"fmt"
	"strings"
)

type AdvancedFeature string

const (
	Htmx              AdvancedFeature = "htmx"
	GoProjectWorkflow AdvancedFeature = "go-project-workflow"
)

var AllowedAdvancedFeatures = []string{string(Htmx), string(GoProjectWorkflow)}

func (f AdvancedFeature) String() string {
	return string(f)
}

func (f *AdvancedFeature) Type() string {
	return "AdvancedFeature"
}

func (f *AdvancedFeature) Set(value string) error {
	// Contains isn't available in 1.20 yet
	// if AllowedProjectTypes.Contains(value) {
	for _, advancedFeature := range AllowedAdvancedFeatures {
		if advancedFeature == value {
			*f = AdvancedFeature(value)
			return nil
		}
	}

	return fmt.Errorf("Advanced Feature to use. Allowed values: %s", strings.Join(AllowedAdvancedFeatures, ", "))
}
