package utils

import (
	"bytes"
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

func ExecuteCmd(name string, args []string, dir string) error {
	command := exec.Command(name, args...)
	command.Dir = dir
	var out bytes.Buffer
	command.Stdout = &out
	if err := command.Run(); err != nil {
		return err
	}
	return nil
}

func InitGoMod(projectName string, appDir string) {
	if err := ExecuteCmd("go",
		[]string{"mod", "init", projectName},
		appDir); err != nil {
		cobra.CheckErr(err)
	}
}

func GoGetPackage(appDir, packageName string) {
	fmt.Println("this is the packageName", packageName)
	if err := ExecuteCmd("go",
		[]string{"get", "-u", packageName},
		appDir); err != nil {
		cobra.CheckErr(err)
	}
}
