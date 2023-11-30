/*
Go blueprint version
*/
package cmd

import (
	"fmt"
	"runtime/debug"
	"time"

	"github.com/spf13/cobra"
)

// GoBlueprintVersion is the version of the cli to be overwritten by goreleaser in the CI run with the version of the release in github
var GoBlueprintVersion string

// Go Blueprint needs to be built in a specific way to provide useful version information.
// First we try to get the version from ldflags embedded into GoBlueprintVersion.
// Then we try to get the version from from the go.mod build info. 
// If Go Blueprint is installed with a specific version tag or using @latest then that version will be included in bi.Main.Version. 
// This won't give any version info when running 'go install' with the source code locally.
// Finally we try to get the version from other embedded VCS info.
func getGoBlueprintVersion() string {
	noVersionAvailable := "No version info available for this build, run 'go-blueprint help version' for additional info"
	
	if len(GoBlueprintVersion) != 0 {
		return GoBlueprintVersion
	}

	bi, ok := debug.ReadBuildInfo()
	if !ok {
		return noVersionAvailable
	}

	// If no main version is available, Go defaults it to (devel)
	if bi.Main.Version != "(devel)" {
		return bi.Main.Version
	}

	var vcsRevision string
	var vcsTime time.Time
	for _, setting := range bi.Settings {
		switch setting.Key {
		case "vcs.revision":
			vcsRevision = setting.Value
		case "vcs.time":
			vcsTime, _ = time.Parse(time.RFC3339, setting.Value)
		}
	}

	if vcsRevision != "" {
		return fmt.Sprintf("%s, (%s)", vcsRevision, vcsTime)
	}

	return noVersionAvailable
}

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Display application version information.",
	Long: `
The version command provides information about the application's version.

Go Blueprint requires version information to be embedded at compile time.
For detailed version information, Go Blueprint needs to be built as specified in the README installation instructions.
If Go Blueprint is built within a version control repository and other version info isn't available,
the revision hash will be used instead.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		version := getGoBlueprintVersion()
		fmt.Printf("Go Blueprint CLI version: %v\n", version)
	},
}
