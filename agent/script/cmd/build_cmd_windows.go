package cmd

import (
	ospenv "osp/agent/script/pathenv"
	ospsys "osp/pkg/system"
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
