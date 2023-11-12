package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/melkeydev/go-blueprint/cmd/frameworks"
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
	logoStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("#01FAC6")).Bold(true)
	tipMsgStyle    = lipgloss.NewStyle().PaddingLeft(1).Foreground(lipgloss.Color("190")).Italic(true)
	endingMsgStyle = lipgloss.NewStyle().PaddingLeft(1).Foreground(lipgloss.Color("170")).Bold(true)
)

func init() {
	var frameworkFlag frameworks.Framework

	rootCmd.AddCommand(createCmd)

	createCmd.Flags().StringP("name", "n", "", "Name of project to create")
	createCmd.Flags().VarP(&frameworkFlag, "framework", "f", fmt.Sprintf("Framework to use. Allowed values: %s", strings.Join(frameworks.AllowedProjectTypes, ", ")))

	if err := createCmd.RegisterFlagCompletionFunc("framework", frameworks.FrameworkCompletion); err != nil {
		log.Fatal(err)
	}
}

// createCmd defines the "create" command for the CLI
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a Go project and don't worry about the structure",
	Long:  "Go Blueprint is a CLI tool that allows you to focus on the actual Go code, and not the project structure. Perfect for someone new to the Go language",

	Run: func(cmd *cobra.Command, args []string) {
		var tprogram *tea.Program

		options := steps.Options{
			ProjectName: &textinput.Output{},
		}

		isInteractive := !utils.HasChangedFlag(cmd.Flags())

		flagName := cmd.Flag("name").Value.String()

		_flagFramework := cmd.Flag("framework").Value.String()

		if _flagFramework != "" {
			isValid := isValidProjectType(_flagFramework, frameworks.AllowedProjectTypes)
			if !isValid {
				cobra.CheckErr(fmt.Errorf("Project type '%s' is not valid. Valid types are: %s", _flagFramework, strings.Join(frameworks.AllowedProjectTypes, ", ")))
			}
		}

		flagFramework := frameworks.Framework(_flagFramework)

		project := &program.Project{
			FrameworkMap: make(map[string]program.Framework),
			ProjectName:  flagName,
			ProjectType:  flagFramework,
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
			_ = cmd.Flag("name").Value.Set(project.ProjectName)
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

			project.ProjectType = options.ProjectType
			_ = cmd.Flag("framework").Value.Set(project.ProjectType.String())
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
