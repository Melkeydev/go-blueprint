package cmd

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/melkeydev/go-blueprint/cmd/program"
	"github.com/melkeydev/go-blueprint/cmd/steps"
	"github.com/melkeydev/go-blueprint/cmd/ui/multiInput"
	"github.com/melkeydev/go-blueprint/cmd/ui/textinput"
	"github.com/spf13/cobra"
	"log"
	"os"
	"strings"
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
	allowedCICD   		= []string{"jenkins", "github-action", "none"}
)

func init() {
	rootCmd.AddCommand(createCmd)

	createCmd.Flags().StringP("name", "n", "", "Name of project to create")
	createCmd.Flags().StringP("framework", "f", "", fmt.Sprintf("Framework to use. Allowed values: %s", strings.Join(allowedProjectTypes, ", ")))
	createCmd.Flags().StringP("cicd", "c", "", fmt.Sprintf("cicd to use. Allowed values: %s", strings.Join(allowedCICD, ", ")))
}

type Options struct {
	ProjectName *textinput.Output
	ProjectType *multiInput.Selection
	CICD 		*multiInput.Selection
}

// createCmd defines the "create" command for the CLI
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a Go project and don't worry about the structure",
	Long:  "Go Blueprint is a CLI tool that allows you to focus on the actual Go code, and not the project structure. Perfect for someone new to the Go language",

	Run: func(cmd *cobra.Command, args []string) {
		var tprogram *tea.Program

		options := Options{
			ProjectName: &textinput.Output{},
			ProjectType: &multiInput.Selection{},
			CICD:		 &multiInput.Selection{},	
		}

		flagName := cmd.Flag("name").Value.String()
		flagFramework := cmd.Flag("framework").Value.String()
		flagCICD := cmd.Flag("cicd").Value.String()

		if flagFramework != "" {
			isValid := isValidProjectType(flagFramework, allowedProjectTypes)
			if !isValid {
				cobra.CheckErr(fmt.Errorf("Project type '%s' is not valid. Valid types are: %s", flagFramework, strings.Join(allowedProjectTypes, ", ")))
			}
		}


		if flagCICD != "" {
			isValid := isValidCICD(flagCICD, allowedCICD)
			if !isValid {
				cobra.CheckErr(fmt.Errorf("CICD piplne '%s' is not valid. Valid types are: %s", flagCICD, strings.Join(allowedCICD, ", ")))
			}
		}

		project := &program.Project{
			ProjectName:  flagName,
			ProjectType:  strings.ReplaceAll(flagFramework, "-", " "),
			CICD:		  flagCICD,
			FrameworkMap: make(map[string]program.Framework),
			CICDMap: 	  make(map[string]program.CICD),
		}

		steps := steps.InitSteps()
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
			step := steps.Steps["framework"]
			tprogram = tea.NewProgram(multiInput.InitialModelMulti(step.Options, options.ProjectType, step.Headers, project))
			if _, err := tprogram.Run(); err != nil {
				cobra.CheckErr(err)
				project.ExitCLI(tprogram)
			}
			project.ProjectType = strings.ToLower(options.ProjectType.Choice)
		}

		if project.CICD == "" {
			step := steps.Steps["cicd"]
			tprogram = tea.NewProgram(multiInput.InitialModelMulti(step.Options, options.CICD, step.Headers, project))
			if _, err := tprogram.Run(); err != nil {
				cobra.CheckErr(err)
				project.ExitCLI(tprogram)
			}
			project.CICD = strings.ToLower(options.CICD.Choice)
		}	

		currentWorkingDir, err := os.Getwd()
		project.AbsolutePath = currentWorkingDir

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

func isValidCICD(input string, allowedCICD []string) bool {
	for _, c := range allowedCICD {
		if input == c {
			return true
		}
	}
	return false
}