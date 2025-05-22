
package flags

import (
	"fmt"
	"strings"
)

type Builder string

const (
	Make Builder = "make"
	Just Builder = "just"
)

var AllowedBuilders = []string{string(Make), string(Just)}

func (f Builder) String() string {
	return string(f)
}

func (f *Builder) Type() string {
	return "Database"
}

func (f *Builder) Set(value string) error {
	// Contains isn't available in 1.20 yet
	// if AllowedDBDrivers.Contains(value) {
	for _, builder := range AllowedBuilders {
		if builder == value {
			*f = Builder(value)
			return nil
		}
	}

	return fmt.Errorf("Builder to use. Allowed values: %s", strings.Join(AllowedDBDrivers, ", "))
}
