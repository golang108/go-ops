package action

import (
	"os/exec"
	"path/filepath"
)

func getControlName() string {
	return "control.ps1"
}

func control(workdir, args string) (err error) {
	controlScript := getControlName()

	cmd := exec.Command("powershell", "-File", "."+string(filepath.Separator)+controlScript, args)
	cmd.Dir = workdir
	return cmd.Start()
}
