package utils

import (
	"bytes"
	"os/exec"
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

func InitGoMod(projectName string, appDir string) error {
	if err := ExecuteCmd("go",
		[]string{"mod", "init", projectName},
		appDir); err != nil {
		return err
	}

	return nil
}

func GoGetPackage(appDir, packageName string) error {
	if err := ExecuteCmd("go",
		[]string{"get", "-u", packageName},
		appDir); err != nil {
		return err
	}

	return nil
}

func GoFmt(appDir string) error {
	if err := ExecuteCmd("gofmt",
		[]string{"-s", "-w", "."},
		appDir); err != nil {
		return err
	}

	return nil
}
