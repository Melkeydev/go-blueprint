package template

// CobraTemplates contains the methods used for building
// an app that uses [github.com/spf13/cobra] for building a CLI app.
type CobraTemplates struct{}

func (c CobraTemplates) Main() []byte {
	return CobraMain()
}

func (c CobraTemplates) Server() []byte {
	return []byte("")
}

func (c CobraTemplates) Routes() []byte {
	return []byte("")
}

// MakeCobraCMDRoot returns a byte slice that represents
// the cmd/root.go file when using Cobra.
func MakeCobraCMDRoot() []byte {
	return []byte(`package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "example",
	Short: "example is an awesome way to show you how to use cobra",
	Long: "This is an example of how to use cobra to build a CLI app. This will be a long description of the app.",
Run: func(cmd *cobra.Command, args []string) {
	// Do Stuff Here
	fmt.Println("Hello Example!")
},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
`)
}
