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

var AllowedProjectTypes = []string{"chi", "gin", "fiber", "gorilla/mux", "httprouter", "standard-library", "echo"}

var (
	logoStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("#01FAC6")).Bold(true)
	tipMsgStyle    = lipgloss.NewStyle().PaddingLeft(1).Foreground(lipgloss.Color("190")).Italic(true)
	separatorStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	endingMsgStyle = lipgloss.NewStyle().PaddingLeft(1).Foreground(lipgloss.Color("170")).Bold(true)
)

func init() {
	rootCmd.AddCommand(createCmd)

	createCmd.Flags().StringP("name", "n", "", "Name of project to create")
	createCmd.Flags().StringP("framework", "f", "", fmt.Sprintf("Framework to use. Allowed values: %s", strings.Join(AllowedProjectTypes, ", ")))
}

// createCmd defines the "create" command for the CLI
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a Go project and don't worry about the structure",
	Long:  "Go Blueprint is a CLI tool that allows you to focus on the actual Go code, and not the project structure. Perfect for someone new to the Go language",

	Run: func(cmd *cobra.Command, args []string) {
		project, options, isInteractive := NewProject(cmd)

		if project.ProjectName == "" {
			createNamePrompt(cmd, project, &options)
		}

		if project.ProjectType == "" {
			createFrameworkPrompt(cmd, project, &options)
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

		fmt.Println(separatorStyle.Render(strings.Repeat("-", 80)))
		fmt.Println(endingMsgStyle.Render("\nNext step:  navigate into the newly created project with:"))
		fmt.Println(endingMsgStyle.Render(fmt.Sprintf("• cd %s\n", project.ProjectName)))

		if isInteractive {
			nonInteractiveCommand := utils.NonInteractiveCommand(cmd.Flags())
			fmt.Println(tipMsgStyle.Render("Tip: go-blueprint supports non-interactive commands:"))
			fmt.Println(tipMsgStyle.Italic(false).Render(fmt.Sprintf("• %s\n", nonInteractiveCommand)))
		}
	},
}

func createNamePrompt(cmd *cobra.Command, project *program.Project, options *steps.Options) {
	tprogram := tea.NewProgram(textinput.InitialTextInputModel(options.ProjectName, "What is the name of your project?", project))
	if _, err := tprogram.Run(); err != nil {
		log.Printf("Name of project contains an error: %v", err)
		cobra.CheckErr(err)
	}
	project.ExitCLI(tprogram)

	project.ProjectName = options.ProjectName.Output
	cmd.Flag("name").Value.Set(project.ProjectName)
}

func createFrameworkPrompt(cmd *cobra.Command, project *program.Project, options *steps.Options) {
	steps := steps.InitSteps(options)
	fmt.Printf("%s\n", logoStyle.Render(logo))

	for _, step := range steps.Steps {
		s := &multiInput.Selection{}
		tprogram := tea.NewProgram(multiInput.InitialModelMulti(step.Options, s, step.Headers, project))
		if _, err := tprogram.Run(); err != nil {
			cobra.CheckErr(err)
		}
		project.ExitCLI(tprogram)

		*step.Field = s.Choice
	}

	project.ProjectType = strings.ToLower(options.ProjectType)
	cmd.Flag("framework").Value.Set(project.ProjectType)
}
