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
	endingMsgStyle      = lipgloss.NewStyle().PaddingLeft(1).Foreground(lipgloss.Color("170")).Bold(true)
	allowedProjectTypes = []string{"chi", "gin", "fiber", "gorilla/mux", "httprouter", "standard-library", "echo"}
)

func init() {
	rootCmd.AddCommand(createCmd)

	createCmd.Flags().StringP("project-name", "n", "", "Name of project to create")
	createCmd.Flags().StringP("project-type", "t", "", fmt.Sprintf("Type of project to create. Allowed values: %s", strings.Join(allowedProjectTypes, ", ")))
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a Go project and don't worry about the structure",
	Long:  "Go Blueprint is a CLI tool that allows you to focus on the actual Go code, and not the project structure. Perfect for someone new to the Go language",

	Run: func(cmd *cobra.Command, args []string) {
		var tprogram *tea.Program

		options := steps.Options{
			ProjectName: &textinput.Output{},
		}

		flagName := cmd.Flag("project-name").Value.String()
		flagType := cmd.Flag("project-type").Value.String()

		isValid := isValidProjectType(flagType, allowedProjectTypes)
		if !isValid {
			cobra.CheckErr(fmt.Errorf("Project type '%s' is not valid. Valid types are: %s", flagType, strings.Join(allowedProjectTypes, ", ")))
		}

		project := &program.Project{
			FrameworkMap: make(map[string]program.Framework),
			ProjectName:  flagName,
			ProjectType:  strings.ReplaceAll(flagType, "-", " "),
		}

		steps := steps.InitSteps(&options)
		fmt.Printf("%s\n", logoStyle.Render(logo))

		if project.ProjectName == "" {
			tprogram := tea.NewProgram(textinput.InitialTextInputModel(options.ProjectName, "What is the name of your project?", project))

			if _, err := tprogram.Run(); err != nil {
				log.Printf("Name of project contains an error: %v", err)
				cobra.CheckErr(err)
			}
			project.ExitCLI(tprogram)

			project.ProjectName = options.ProjectName.Output
		}

		if project.ProjectType == "" {
			for _, step := range steps.Steps {
				s := &multiInput.Selection{}
				tprogram = tea.NewProgram(multiInput.InitialModelMulti(step.Options, s, step.Headers, project))
				if _, err := tprogram.Run(); err != nil {
					cobra.CheckErr(err)
				}
				project.ExitCLI(tprogram)

				*step.Field = s.Choice
			}

			project.ProjectType = strings.ToLower(options.ProjectType)
		}

		currentWorkingDir, err := os.Getwd()
		if err != nil {
			log.Printf("could not get current working directory: %v", err)
			cobra.CheckErr(err)
		}

		project.AbsolutePath = currentWorkingDir

		// This calls the templates
		err = project.CreateMainFile()
		if err != nil {
			log.Printf("Problem creating files for project. %v", err)
			cobra.CheckErr(err)
		}

		fmt.Println(endingMsgStyle.Render("\nNext steps cd into the newly created project with:"))
		fmt.Println(endingMsgStyle.Render(fmt.Sprintf("• cd %s\n", project.ProjectName)))
	},
}

func isValidProjectType(input string, allowedTypes []string) bool {
	for _, t := range allowedTypes {
		if input == t {
			return true
		}
	}
	return false
}
