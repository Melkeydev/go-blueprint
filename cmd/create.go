package cmd

import (
        "fmt"
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
                var errMsg string

                options := steps.Options{
                        ProjectName: &textinput.Output{},
                }

                flagName := cmd.Flag("name").Value.String()
                if doesDirectoryExist(flagName) {
                        errMsg = fmt.Sprintf("Directory '%s' already exists", flagName)
                        cobra.CheckErr(textinput.CreateErrorValue(errMsg).Error())
                }

                flagFramework := cmd.Flag("framework").Value.String()
                if flagFramework != "" {
                        isValid := isValidProjectType(flagFramework, allowedProjectTypes)
                        if !isValid {
                                errMsg = fmt.Sprintf("Project type '%s' is not valid. Valid types are: %s", flagFramework, strings.Join(allowedProjectTypes, ", "))
                                cobra.CheckErr(textinput.CreateErrorValue(errMsg))
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
                                errMsg = fmt.Sprintf("Name of project contains an error: %v", err)
                                cobra.CheckErr(textinput.CreateErrorValue(errMsg).Error())
                        }

                        if doesDirectoryExist(options.ProjectName.Output) {
                                errMsg = fmt.Sprintf("Directory '%s' already exists", options.ProjectName.Output)
                                cobra.CheckErr(textinput.CreateErrorValue(errMsg).Error())
                        }

                        project.ExitCLI(tprogram)

                        project.ProjectName = options.ProjectName.Output
                }

                if project.ProjectType == "" {
                        for _, step := range steps.Steps {
                                s := &multiInput.Selection{}
                                tprogram = tea.NewProgram(multiInput.InitialModelMulti(step.Options, s, step.Headers, project))
                                if _, err := tprogram.Run(); err != nil {
                                        errMsg = err.Error()
                                        cobra.CheckErr(textinput.CreateErrorValue(errMsg).Error())
                                }
                                project.ExitCLI(tprogram)

                                *step.Field = s.Choice
                        }

                        project.ProjectType = strings.ToLower(options.ProjectType)
                }

                currentWorkingDir, err := os.Getwd()
                if err != nil {
                        errMsg = fmt.Sprintf("could not get current working directory: %v", err)
                        cobra.CheckErr(textinput.CreateErrorValue(errMsg).Error())
                }

                project.AbsolutePath = currentWorkingDir

                // This calls the templates
                err = project.CreateMainFile()
                if err != nil {
                        errMsg = fmt.Sprintf("Problem creating files for project. %v", err)
                        cobra.CheckErr(textinput.CreateErrorValue(errMsg).Error())
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

// doesDirectoryExist checks if the directory exists
func doesDirectoryExist(dir string) bool {
        if _, err := os.Stat(dir); !os.IsNotExist(err) {
                return true
        }
        return false
}
