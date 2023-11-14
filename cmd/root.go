/*
Copyright © 2023 Melkey melkeydev@gmail.com
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "go-blueprint",
	Short: "A program to spin up a quick Go project using a popular framework",
	Long: `Go Blueprint is a CLI tool that allows users to spin up a Go project with the corresponding structure seamlessly. 
It also gives the option to integrate with one of the more popular Go frameworks!`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.go-blueprint.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.AddCommand(versionCmd)
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
