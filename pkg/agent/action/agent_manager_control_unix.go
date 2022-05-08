package action

import (
	"os/exec"
	"path/filepath"
)

func getControlName() string {
	return "control"
}

func control(workdir, args string) error {
	controlScript := getControlName()

	cmd := exec.Command("."+string(filepath.Separator)+controlScript, args)

	cmd.Dir = workdir

	return cmd.Start()
}
