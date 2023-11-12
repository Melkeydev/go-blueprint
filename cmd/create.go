package cmd

import (
	"fmt"
	"log"
	"os"
	"path"
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
    createCmd.Flags().StringP("path", "p", "", "Where the project should be created")
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
            ProjectPath: &textinput.Output{},
		}

		isInteractive := !utils.HasChangedFlag(cmd.Flags())

		flagName := cmd.Flag("name").Value.String()
		flagFramework := cmd.Flag("framework").Value.String()
        flagPath := cmd.Flag("path").Value.String()

		if flagFramework != "" {
			isValid := isValidProjectType(flagFramework, allowedProjectTypes)
			if !isValid {
				cobra.CheckErr(fmt.Errorf("Project type '%s' is not valid. Valid types are: %s", flagFramework, strings.Join(allowedProjectTypes, ", ")))
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
				cobra.CheckErr(err)
			}
			project.ExitCLI(tprogram)

			project.ProjectName = options.ProjectName.Output
			cmd.Flag("name").Value.Set(project.ProjectName)
		}

        if flagPath == "" {
            tprogram = tea.NewProgram(textinput.InitialTextInputModel(options.ProjectPath, "Where should the project be created? (use '.' for current directory)", project))
            if _, err := tprogram.Run(); err != nil {
                cobra.CheckErr(err)
            }
            cmd.Flag("path").Value.Set(options.ProjectPath.Output)
            flagPath = options.ProjectPath.Output
            project.ExitCLI(tprogram)
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
			cmd.Flag("framework").Value.Set(project.ProjectType)
		}

		currentWorkingDir, err := os.Getwd()
		if err != nil {
			log.Printf("could not get current working directory: %v", err)
			cobra.CheckErr(err)
		}

        var next_steps_cd_path string
        
        if flagPath == "." {
            project.AbsolutePath = currentWorkingDir
        } else if strings.HasPrefix(flagPath, ".") {
            project.AbsolutePath = fmt.Sprintf("%s%s", currentWorkingDir, strings.TrimPrefix(flagPath, ".")) 
        } else if strings.HasPrefix(flagPath, "~") {
            project.AbsolutePath = strings.Replace(flagPath, "~", os.Getenv("HOME"), -1)
        } else if strings.HasPrefix(flagPath, "/") || strings.HasPrefix(flagPath, "\\") {
            project.AbsolutePath = flagPath 
        } else {
            project.AbsolutePath = path.Join(currentWorkingDir, flagPath)
        }

        if strings.Replace(project.AbsolutePath, os.Getenv("HOME"), "", -1) == "" {
            next_steps_cd_path = project.ProjectName
        } else {
            next_steps_cd_path = path.Join(strings.Replace(flagPath, currentWorkingDir, ".", -1), project.ProjectName)
        }

		// This calls the templates
		err = project.CreateMainFile()
		if err != nil {
			log.Printf("Problem creating files for project. %v", err)
			cobra.CheckErr(err)
		}

		fmt.Println(endingMsgStyle.Render("\nNext steps cd into the newly created project with:"))
		fmt.Println(endingMsgStyle.Render(fmt.Sprintf("• cd %s\n", next_steps_cd_path)))

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
