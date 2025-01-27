package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"
	"sync"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/melkeydev/go-blueprint/cmd/flags"
	"github.com/melkeydev/go-blueprint/cmd/program"
	"github.com/melkeydev/go-blueprint/cmd/steps"
	"github.com/melkeydev/go-blueprint/cmd/ui/multiInput"
	"github.com/melkeydev/go-blueprint/cmd/ui/multiSelect"
	"github.com/melkeydev/go-blueprint/cmd/ui/spinner"
	"github.com/melkeydev/go-blueprint/cmd/ui/textinput"
	"github.com/melkeydev/go-blueprint/cmd/utils"
	"github.com/spf13/cobra"
)

const logo = `

 ____  _                       _       _
|  _ \| |                     (_)     | |
| |_) | |_   _  ___ _ __  _ __ _ _ __ | |_
|  _ <| | | | |/ _ \ '_ \| '__| | '_ \| __|
| |_) | | |_| |  __/ |_) | |  | | | | | |
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
	var flagFramework flags.Framework
	var flagDBDriver flags.Database
	var frontendFrameworks flags.FrontendFramework
	var frontendAdvanced flags.FrontendAdvanced
	var advancedFeatures flags.AdvancedFeatures
	var flagGit flags.Git
	rootCmd.AddCommand(createCmd)

	createCmd.Flags().StringP("name", "n", "", "Name of project to create")
	createCmd.Flags().VarP(&flagFramework, "framework", "b", fmt.Sprintf("Framework to use. Allowed values: %s", strings.Join(flags.AllowedProjectTypes, ", ")))
	createCmd.Flags().VarP(&flagDBDriver, "driver", "d", fmt.Sprintf("Database drivers to use. Allowed values: %s", strings.Join(flags.AllowedDBDrivers, ", ")))
	createCmd.Flags().BoolP("frontend", "f", false, "Get prompts for frontend frameworks")
	createCmd.Flags().Var(&frontendFrameworks, "frontend-framework", fmt.Sprintf("Frontend framework to use. Allowed values: %s", strings.Join(flags.AllowedFrontendTypes, ", ")))
	createCmd.Flags().Var(&frontendAdvanced, "frontend-advanced", fmt.Sprintf("Frontend framework advanced features to use. Allowed values: %s", strings.Join(flags.AllowedFrontendAdvanced, ", ")))
	createCmd.Flags().BoolP("advanced", "a", false, "Get prompts for advanced features")
	createCmd.Flags().Var(&advancedFeatures, "feature", fmt.Sprintf("Advanced feature to use. Allowed values: %s", strings.Join(flags.AllowedAdvancedFeatures, ", ")))
	createCmd.Flags().VarP(&flagGit, "git", "g", fmt.Sprintf("Git to use. Allowed values: %s", strings.Join(flags.AllowedGitsOptions, ", ")))
}

type Options struct {
	ProjectName       *textinput.Output
	ProjectType       *multiInput.Selection
	DBDriver          *multiInput.Selection
	FrontendFramework *multiInput.Selection
	FrontendAdvanced  *multiSelect.Selection
	Advanced          *multiSelect.Selection
	Git               *multiInput.Selection
}

// createCmd defines the "create" command for the CLI
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a Go project and don't worry about the structure",
	Long:  "Go Blueprint is a CLI tool that allows you to focus on the actual Go code, and not the project structure. Perfect for someone new to the Go language",

	Run: func(cmd *cobra.Command, args []string) {
		var tprogram *tea.Program
		var err error

		isInteractive := false
		flagName := cmd.Flag("name").Value.String()

		if flagName != "" && !utils.ValidateModuleName(flagName) {
			err = fmt.Errorf("'%s' is not a valid module name. Please choose a different name", flagName)
			cobra.CheckErr(textinput.CreateErrorInputModel(err).Err())
		}

		rootDirName := utils.GetRootDir(flagName)
		if rootDirName != "" && doesDirectoryExistAndIsNotEmpty(rootDirName) {
			err = fmt.Errorf("directory '%s' already exists and is not empty. Please choose a different name", rootDirName)
			cobra.CheckErr(textinput.CreateErrorInputModel(err).Err())
		}

		// VarP already validates the contents of the framework flag.
		// If this flag is filled, it is always valid
		flagFramework := flags.Framework(cmd.Flag("framework").Value.String())
		flagDBDriver := flags.Database(cmd.Flag("driver").Value.String())
		flagFrontendFremwork := flags.FrontendFramework(cmd.Flag("frontend-framework").Value.String())
		flagGit := flags.Git(cmd.Flag("git").Value.String())

		options := Options{
			ProjectName:       &textinput.Output{},
			ProjectType:       &multiInput.Selection{},
			DBDriver:          &multiInput.Selection{},
			FrontendFramework: &multiInput.Selection{},
			FrontendAdvanced: &multiSelect.Selection{
				Choices: make(map[string]bool),
			},
			Advanced: &multiSelect.Selection{
				Choices: make(map[string]bool),
			},
			Git: &multiInput.Selection{},
		}

		project := &program.Project{
			ProjectName:       flagName,
			ProjectType:       flagFramework,
			DBDriver:          flagDBDriver,
			FrameworkMap:      make(map[flags.Framework]program.Framework),
			DBDriverMap:       make(map[flags.Database]program.Driver),
			FrontendFramework: flagFrontendFremwork,
			FrontendOptions:   make(map[string]bool),
			AdvancedOptions:   make(map[string]bool),
			GitOptions:        flagGit,
		}

		steps := steps.InitSteps(flagFramework, flagDBDriver, flagFrontendFremwork, flagGit)
		fmt.Printf("%s\n", logoStyle.Render(logo))

		// Advanced option steps:
		flagAdvanced, err := cmd.Flags().GetBool("advanced")
		if err != nil {
			log.Fatal("failed to retrieve advanced flag")
		}

		if flagAdvanced {
			fmt.Println(tipMsgStyle.Render("*** You are in advanced mode ***\n\n"))
		}

		flagFrontend, err := cmd.Flags().GetBool("frontend")
		if err != nil {
			log.Fatal("failed to retrieve frontend flag")
		}

		if flagFrontend {
			fmt.Println(tipMsgStyle.Render("*** You are in Frontend mode ***\n\n"))
		}

		if project.ProjectName == "" {
			isInteractive = true
			tprogram := tea.NewProgram(textinput.InitialTextInputModel(options.ProjectName, "What is the name of your project?", project))
			if _, err := tprogram.Run(); err != nil {
				log.Printf("Name of project contains an error: %v", err)
				cobra.CheckErr(textinput.CreateErrorInputModel(err).Err())
			}

			if options.ProjectName.Output != "" && !utils.ValidateModuleName(options.ProjectName.Output) {
				err = fmt.Errorf("'%s' is not a valid module name. Please choose a different name", options.ProjectName.Output)
				cobra.CheckErr(textinput.CreateErrorInputModel(err).Err())
			}

			rootDirName = utils.GetRootDir(options.ProjectName.Output)
			if doesDirectoryExistAndIsNotEmpty(rootDirName) {
				err = fmt.Errorf("directory '%s' already exists and is not empty. Please choose a different name", rootDirName)
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
			isInteractive = true
			step := steps.Steps["framework"]
			tprogram = tea.NewProgram(multiInput.InitialModelMulti(step.Options, options.ProjectType, step.Headers, project))
			if _, err := tprogram.Run(); err != nil {
				cobra.CheckErr(textinput.CreateErrorInputModel(err).Err())
			}
			project.ExitCLI(tprogram)

			step.Field = options.ProjectType.Choice

			// this type casting is always safe since the user interface can
			// only pass strings that can be cast to a flags.Framework instance
			project.ProjectType = flags.Framework(strings.ToLower(options.ProjectType.Choice))
			err := cmd.Flag("framework").Value.Set(project.ProjectType.String())
			if err != nil {
				log.Fatal("failed to set the framework flag value", err)
			}
		}

		if project.DBDriver == "" {
			isInteractive = true
			step := steps.Steps["driver"]
			tprogram = tea.NewProgram(multiInput.InitialModelMulti(step.Options, options.DBDriver, step.Headers, project))
			if _, err := tprogram.Run(); err != nil {
				cobra.CheckErr(textinput.CreateErrorInputModel(err).Err())
			}
			project.ExitCLI(tprogram)

			// this type casting is always safe since the user interface can
			// only pass strings that can be cast to a flags.Database instance
			project.DBDriver = flags.Database(strings.ToLower(options.DBDriver.Choice))
			err := cmd.Flag("driver").Value.Set(project.DBDriver.String())
			if err != nil {
				log.Fatal("failed to set the driver flag value", err)
			}
		}

		if flagAdvanced {

			featureFlags := cmd.Flag("feature").Value.String()

			if featureFlags != "" {
				featuresFlagValues := strings.Split(featureFlags, ",")
				for _, key := range featuresFlagValues {
					project.AdvancedOptions[key] = true
				}
			} else {
				isInteractive = true
				step := steps.Steps["advanced"]
				tprogram = tea.NewProgram((multiSelect.InitialModelMultiSelect(step.Options, options.Advanced, step.Headers, project)))
				if _, err := tprogram.Run(); err != nil {
					cobra.CheckErr(textinput.CreateErrorInputModel(err).Err())
				}
				project.ExitCLI(tprogram)
				for key, opt := range options.Advanced.Choices {
					project.AdvancedOptions[strings.ToLower(key)] = opt
					err := cmd.Flag("feature").Value.Set(strings.ToLower(key))
					if err != nil {
						log.Fatal("failed to set the feature flag value", err)
					}
				}
				if err != nil {
					log.Fatal(err)
				}
			}

		}

		if flagFrontend {

			if project.FrontendFramework == "" {
				isInteractive = true
				step := steps.Steps["frontend-framework"]
				tprogram = tea.NewProgram(multiInput.InitialModelMulti(step.Options, options.FrontendFramework, step.Headers, project))
				if _, err := tprogram.Run(); err != nil {
					cobra.CheckErr(textinput.CreateErrorInputModel(err).Err())
				}
				project.ExitCLI(tprogram)

				project.FrontendFramework = flags.FrontendFramework(strings.ToLower(options.FrontendFramework.Choice))
				err := cmd.Flag("frontend-framework").Value.Set(project.FrontendFramework.String())
				if err != nil {
					log.Fatal("failed to set the frontend flag value", err)
				}
			}

			featureFrontend := cmd.Flag("frontend-advanced").Value.String()

			if featureFrontend != "" {
				featuresFrontendFlagValues := strings.Split(featureFrontend, ",")
				for _, key := range featuresFrontendFlagValues {
					project.AdvancedOptions[key] = true
				}
			} else {
				isInteractive = true
				step := steps.Steps["frontend-advanced"]
				tprogram = tea.NewProgram((multiSelect.InitialModelMultiSelect(step.Options, options.FrontendAdvanced, step.Headers, project)))
				if _, err := tprogram.Run(); err != nil {
					cobra.CheckErr(textinput.CreateErrorInputModel(err).Err())
				}
				project.ExitCLI(tprogram)
				for key, opt := range options.FrontendAdvanced.Choices {
					project.FrontendOptions[strings.ToLower(key)] = opt
					err := cmd.Flag("frontend-advanced").Value.Set(strings.ToLower(key))
					if err != nil {
						log.Fatal("failed to set the advanced fronted features flag value", err)
					}
				}
				if err != nil {
					log.Fatal(err)
				}
			}

		}

		if project.GitOptions == "" {
			isInteractive = true
			step := steps.Steps["git"]
			tprogram = tea.NewProgram(multiInput.InitialModelMulti(step.Options, options.Git, step.Headers, project))
			if _, err := tprogram.Run(); err != nil {
				cobra.CheckErr(textinput.CreateErrorInputModel(err).Err())
			}
			project.ExitCLI(tprogram)

			project.GitOptions = flags.Git(strings.ToLower(options.Git.Choice))
			err := cmd.Flag("git").Value.Set(project.GitOptions.String())
			if err != nil {
				log.Fatal("failed to set the git flag value", err)
			}
		}

		currentWorkingDir, err := os.Getwd()
		if err != nil {
			log.Printf("could not get current working directory: %v", err)
			cobra.CheckErr(textinput.CreateErrorInputModel(err).Err())
		}
		project.AbsolutePath = currentWorkingDir

		spinner := tea.NewProgram(spinner.InitialModelNew())

		// add synchronization to wait for spinner to finish
		wg := sync.WaitGroup{}
		wg.Add(1)
		go func() {
			defer wg.Done()
			if _, err := spinner.Run(); err != nil {
				cobra.CheckErr(err)
			}
		}()

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("The program encountered an unexpected issue and had to exit. The error was:", r)
				fmt.Println("If you continue to experience this issue, please post a message on our GitHub page or join our Discord server for support.")
				if releaseErr := spinner.ReleaseTerminal(); releaseErr != nil {
					log.Printf("Problem releasing terminal: %v", releaseErr)
				}
			}
		}()

		// This calls the templates
		err = project.CreateMainFile()
		if err != nil {
			if releaseErr := spinner.ReleaseTerminal(); releaseErr != nil {
				log.Printf("Problem releasing terminal: %v", releaseErr)
			}
			log.Printf("Problem creating files for project. %v", err)
			cobra.CheckErr(textinput.CreateErrorInputModel(err).Err())
		}

		fmt.Println(endingMsgStyle.Render("\nNext steps:"))
		fmt.Println(endingMsgStyle.Render(fmt.Sprintf("• cd into the newly created project with: `cd %s`\n", utils.GetRootDir(project.ProjectName))))

		if options.FrontendFramework.Choice == "React" {
			fmt.Println(endingMsgStyle.Render("• cd into frontend\n"))
			fmt.Println(endingMsgStyle.Render("• npm install\n"))
			fmt.Println(endingMsgStyle.Render("• npm run dev\n"))
		}

		if options.FrontendAdvanced.Choices["Tailwind"] {
			fmt.Println(endingMsgStyle.Render("• Install the tailwind standalone cli if you haven't already, grab the executable for your platform from the latest release on GitHub\n"))
			fmt.Println(endingMsgStyle.Render("• More info about the Tailwind CLI: https://tailwindcss.com/blog/standalone-cli\n"))
		}

		if options.FrontendFramework.Choice == "Htmx" {
			fmt.Println(endingMsgStyle.Render("• Install the templ cli if you haven't already by running `go install github.com/a-h/templ/cmd/templ@latest`\n"))
			fmt.Println(endingMsgStyle.Render("• Generate templ function files by running `templ generate`\n"))
		}

		if isInteractive {
			nonInteractiveCommand := utils.NonInteractiveCommand(cmd.Use, cmd.Flags())
			fmt.Println(tipMsgStyle.Render("Tip: Repeat the equivalent Blueprint with the following non-interactive command:"))
			fmt.Println(tipMsgStyle.Italic(false).Render(fmt.Sprintf("• %s\n", nonInteractiveCommand)))
		}
		err = spinner.ReleaseTerminal()
		if err != nil {
			log.Printf("Could not release terminal: %v", err)
			cobra.CheckErr(err)
		}
	},
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
