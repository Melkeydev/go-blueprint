package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/melkeydev/go-blueprint/cmd/program"
	"github.com/melkeydev/go-blueprint/cmd/steps"
	"github.com/melkeydev/go-blueprint/cmd/ui/multiInput"
	"github.com/melkeydev/go-blueprint/cmd/ui/textinput"
	"github.com/melkeydev/go-blueprint/cmd/utils"
	"github.com/spf13/cobra"
)

const logo = `

 ____  _                       _       _   
|  _ \| |                     (_)     | |  
| |_) | |_   _  ___ _ __  _ __ _ _ __ | |_ 
|  _ <| | | | |/ _ \ '_ \| '__| | '_ \| __|
| |_) | | |_| |  __/ |_) | |  | | | | | |_ 
|____/|_|\__,_|\___| .__/|_|  |_|_| |_|\__|
				   | |                     
				   |_|                     

`

var (
	logoStyle           = lipgloss.NewStyle().Foreground(lipgloss.Color("#01FAC6")).Bold(true)
	tipMsgStyle         = lipgloss.NewStyle().PaddingLeft(1).Foreground(lipgloss.Color("190")).Italic(true)
	endingMsgStyle      = lipgloss.NewStyle().PaddingLeft(1).Foreground(lipgloss.Color("170")).Bold(true)
	allowedProjectTypes = []string{"chi", "gin", "fiber", "gorilla/mux", "httprouter", "standard-library", "echo"}
)

func init() {
	rootCmd.AddCommand(createCmd)

	createCmd.Flags().StringP("name", "n", "", "Name of project to create")
	createCmd.Flags().StringP("framework", "f", "", fmt.Sprintf("Framework to use. Allowed values: %s", strings.Join(allowedProjectTypes, ", ")))
}

// createCmd defines the "create" command for the CLI
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a Go project and don't worry about the structure",
	Long:  "Go Blueprint is a CLI tool that allows you to focus on the actual Go code, and not the project structure. Perfect for someone new to the Go language",

	Run: func(cmd *cobra.Command, args []string) {
		var tprogram *tea.Program
		var err error

		options := steps.Options{
			ProjectName: &textinput.Output{},
		}

		isInteractive := !utils.HasChangedFlag(cmd.Flags())

		flagName := cmd.Flag("name").Value.String()
		if flagName != "" && doesDirectoryExistAndIsNotEmpty(flagName) {
			err = fmt.Errorf("Directory '%s' already exists and is not empty. Please choose a different name", flagName)
			cobra.CheckErr(textinput.CreateErrorInputModel(err).Err())
		}
		flagFramework := cmd.Flag("framework").Value.String()

		if flagFramework != "" {
			isValid := isValidProjectType(flagFramework, allowedProjectTypes)
			if !isValid {
				err = fmt.Errorf("Project type '%s' is not valid. Valid types are: %s", flagFramework, strings.Join(allowedProjectTypes, ", "))
				cobra.CheckErr(textinput.CreateErrorInputModel(err).Err())
			}
		}

		project := &program.Project{
			FrameworkMap: make(map[string]program.Framework),
			ProjectName:  flagName,
			ProjectType:  strings.ReplaceAll(flagFramework, "-", " "),
		}

		steps := steps.InitSteps(&options)
		fmt.Printf("%s\n", logoStyle.Render(logo))

		if project.ProjectName == "" {
			tprogram := tea.NewProgram(textinput.InitialTextInputModel(options.ProjectName, "What is the name of your project?", project))
			if _, err := tprogram.Run(); err != nil {
				log.Printf("Name of project contains an error: %v", err)
				cobra.CheckErr(textinput.CreateErrorInputModel(err).Err())
			}
			if doesDirectoryExistAndIsNotEmpty(options.ProjectName.Output) {
				err = fmt.Errorf("Directory '%s' already exists and is not empty. Please choose a different name", options.ProjectName.Output)
				cobra.CheckErr(textinput.CreateErrorInputModel(err).Err())
			}
			project.ExitCLI(tprogram)

			project.ProjectName = options.ProjectName.Output
			err := cmd.Flag("name").Value.Set(project.ProjectName)
			if err != nil {
				log.Fatal("failed to set the name flag value", err)
			}
		}

		if project.ProjectType == "" {
			for _, step := range steps.Steps {
				s := &multiInput.Selection{}
				tprogram = tea.NewProgram(multiInput.InitialModelMulti(step.Options, s, step.Headers, project))
				if _, err := tprogram.Run(); err != nil {
					cobra.CheckErr(textinput.CreateErrorInputModel(err).Err())
				}
				project.ExitCLI(tprogram)

				*step.Field = s.Choice
			}

			project.ProjectType = strings.ToLower(options.ProjectType)
			err := cmd.Flag("framework").Value.Set(project.ProjectType)
			if err != nil {
				log.Fatal("failed to set the framework flag value", err)
			}
		}

		currentWorkingDir, err := os.Getwd()
		if err != nil {
			log.Printf("could not get current working directory: %v", err)
			cobra.CheckErr(textinput.CreateErrorInputModel(err).Err())
		}

		project.AbsolutePath = currentWorkingDir

		// This calls the templates
		err = project.CreateMainFile()
		if err != nil {
			log.Printf("Problem creating files for project. %v", err)
			cobra.CheckErr(textinput.CreateErrorInputModel(err).Err())
		}

		fmt.Println(endingMsgStyle.Render("\nNext steps cd into the newly created project with:"))
		fmt.Println(endingMsgStyle.Render(fmt.Sprintf("• cd %s\n", project.ProjectName)))

		if isInteractive {
			nonInteractiveCommand := utils.NonInteractiveCommand(cmd.Flags())
			fmt.Println(tipMsgStyle.Render("Tip: Repeat the equivalent Blueprint with the following non-interactive command:"))
			fmt.Println(tipMsgStyle.Italic(false).Render(fmt.Sprintf("• %s\n", nonInteractiveCommand)))
		}
	},
}

// isValidProjectType checks if the inputted project type matches
// the currently supported list of project types
func isValidProjectType(input string, allowedTypes []string) bool {
	for _, t := range allowedTypes {
		if input == t {
			return true
		}
	}
	return false
}

// doesDirectoryExistAndIsNotEmpty checks if the directory exists and is not empty
func doesDirectoryExistAndIsNotEmpty(name string) bool {
	if _, err := os.Stat(name); err == nil {
		dirEntries, err := os.ReadDir(name)	
		if err != nil {
			log.Printf("could not read directory: %v", err)
			cobra.CheckErr(textinput.CreateErrorInputModel(err))
		}
		if len(dirEntries) > 0 {
			return true
		}
	}
	return false
}
