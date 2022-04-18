package cmd

import (
	ospenv "go-ops/agent/script/pathenv"
	ospsys "go-ops/pkg/system"
)

func BuildCommand(path string) ospsys.Command {
	return ospsys.Command{
		Name: path,
		Env: map[string]string{
			"PATH": ospenv.Path(),
		},
	}
}
