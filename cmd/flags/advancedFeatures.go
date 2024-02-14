package flags

import (
	"fmt"
	"strings"
)

type AdvancedFeature string

type AdvancedFeatures []string

const (
	Htmx              AdvancedFeature = "htmx"
	GoProjectWorkflow AdvancedFeature = "githubaction"
)

var AllowedAdvancedFeatures = []string{string(Htmx), string(GoProjectWorkflow)}

func (f AdvancedFeature) String() string {
	return string(f)
}

func (f AdvancedFeatures) String() string {
	return strings.Join(f, ",")
}

func (f *AdvancedFeatures) Type() string {
	return "AdvancedFeature"
}

func (f *AdvancedFeatures) Set(value string) error {
	// Contains isn't available in 1.20 yet
	// if AllowedProjectTypes.Contains(value) {
	for _, advancedFeature := range AllowedAdvancedFeatures {
		if advancedFeature == value {
			*f = append(*f, advancedFeature)
			return nil
		}
	}

	return fmt.Errorf("Advanced Feature to use. Allowed values: %s", strings.Join(AllowedAdvancedFeatures, ", "))
}
