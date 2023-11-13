/*
Go blueprint version
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var GoBlueprintVersion = ""

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Display application version information.",
	Long: `The version command provides information about the application's version.
Use this command to check the current version of the application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Go Blueprint CLI version %v\n", GoBlueprintVersion)
	},
}
