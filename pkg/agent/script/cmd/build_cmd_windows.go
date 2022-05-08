package cmd

import (
	ospenv "go-ops/pkg/agent/script/pathenv"
	ospsys "go-ops/pkg/system"
)

func BuildCommand(path string) ospsys.Command {
	return ospsys.Command{
		Name: "powershell",
		Args: []string{path},
		Env: map[string]string{
			"PATH": ospenv.Path(),
		},
	}
}
