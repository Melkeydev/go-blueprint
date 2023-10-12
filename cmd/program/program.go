package program

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

type Project struct {
	ProjectName  string
	Exit         bool
	AbsolutePath string
}

func (p *Project) ExitCLI(tprogram *tea.Program) {
	if p.Exit {
		// logo render here
		tprogram.ReleaseTerminal()
		os.Exit(1)
	}
}

func (p *Project) CreateAPIProject() {
	appDir := fmt.Sprintf("%s/%s", p.AbsolutePath, p.ProjectName)
	if _, err := os.Stat(p.AbsolutePath); err == nil {
		if err := os.Mkdir(appDir, 0755); err != nil {
			cobra.CheckErr(err)
		}
	}

	// Get the current working directory
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error getting current working directory: %v\n", err)
		return
	}

	scriptRelPath := filepath.Join(currentDir, "cmd", "scripts", "create_structure.sh")

	// Check if the script file exists
	if _, err := os.Stat(scriptRelPath); os.IsNotExist(err) {
		fmt.Printf("Script file '%s' does not exist\n", scriptRelPath)
		return
	}

	cmd := exec.Command("bash", scriptRelPath)
	cmd.Dir = appDir

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Printf("Error executing Bash script: %v\n", err)
		return
	}

	fmt.Println("Project structure created successfully in", appDir)
}
