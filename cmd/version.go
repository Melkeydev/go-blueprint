/*
Go blueprint version
*/
package cmd

import (
	"fmt"
	"runtime/debug"

	"github.com/spf13/cobra"
)

// GoBlueprintVersion is the version of the cli to be overwritten by goreleaser in the CI run with the version of the release in github
var GoBlueprintVersion string

func getGoBlueprintVersion() string {
	if len(GoBlueprintVersion) != 0 {
		return GoBlueprintVersion
	}
	if info, ok := debug.ReadBuildInfo(); ok {
		return info.Main.Version
	}
	return "unknown"
}

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Display application version information.",
	Long: `The version command provides information about the application's version.
Use this command to check the current version of the application.`,
	Run: func(cmd *cobra.Command, args []string) {
		version := getGoBlueprintVersion()
		fmt.Printf("Go Blueprint CLI version %v\n", version)
	},
}
