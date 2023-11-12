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

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "example",
	Short: "example is an awesome way to show you how to use cobra",
	Long:  "This is an example of how to use cobra to build a CLI app. This will be a long description of the root command / main application.",
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		fmt.Println("Hello Example, (you invoked the root command)!")

		// this is an example of how to get a flag value
		flag, err := cmd.Flags().GetString("flag")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println("The value of the flag is:", flag)

	},
}

// pingCmd represents the ping command, which is a subcommand of the root command, usually this would be put in its own file called cmd/ping.go
var pingCmd = &cobra.Command{
	Use:   "ping",
	Short: "ping is an example of a subcommand, it pings!",
	Long:  "The ping command is a subcommand of the root command. It will print 'pong' to the terminal. It must be invoked explicitly because it is a subcommand.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("pong (you invoked the ping command)!")
	},
}

// init is called before the rootCmd is executed, and is used to add subcommands (like ping) and flags
func init() {
	rootCmd.AddCommand(pingCmd)

	// this is an example of a persistent flag, it will be passed to all subcommands
	rootCmd.PersistentFlags().StringP("flag", "f", "default", "This is an example of a flag. It has a default value of 'default'.")

	// this is an example of a local flag, it will only be passed to the root command
	rootCmd.Flags().StringP("local-flag", "l", "default", "This is an example of a local flag. It has a default value of 'default'.")

}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
`)
}
