package cmdrunner

import (
	"fmt"
	"os"
	"path"

	"io/ioutil"

	"osp/pkg/errors"
	ospsys "osp/pkg/system"
)

const (
	fileOpenFlag int         = os.O_RDWR | os.O_CREATE | os.O_TRUNC
	fileOpenPerm os.FileMode = os.FileMode(0640)
)

type ScriptCmdRunner struct {
	cmdRunner ospsys.CmdRunner
	baseDir   string
}

type ScriptExecErr struct {
	result *CmdResult
	err    error
}

func (f ScriptExecErr) Error() string {
	stdoutTitle := "Stdout"

	stderrTitle := "Stderr"

	return fmt.Sprintf("Command exited with %d; %s: %s, %s: %s,Err:%v",
		f.result.ExitStatus,
		stdoutTitle,
		string(f.result.Stdout),
		stderrTitle,
		string(f.result.Stderr),
		f.err,
	)
}

func NewScriptCmdRunner(
	cmdRunner ospsys.CmdRunner,
	baseDir string,
) CmdRunner {
	return ScriptCmdRunner{
		cmdRunner: cmdRunner,
		baseDir:   baseDir,
	}
}

func (f ScriptCmdRunner) RunCommand(jobId string, cmd ospsys.Command) (*CmdResult, error) {
	logsDir := path.Join(f.baseDir, jobId)

	err := os.MkdirAll(logsDir, os.FileMode(0750))
	if err != nil {
		return nil, errors.WrapErrorf(err, "Creating log dir for job %s", jobId)
	}

	stdoutPath := path.Join(logsDir, fmt.Sprintf("%s.stdout.log", jobId))
	stderrPath := path.Join(logsDir, fmt.Sprintf("%s.stderr.log", jobId))

	stdoutFile, err := os.OpenFile(stdoutPath, fileOpenFlag, fileOpenPerm)
	if err != nil {
		return nil, errors.WrapErrorf(err, "Opening stdout for task %s", jobId)
	}
	defer func() {
		_ = stdoutFile.Close()
	}()

	cmd.Stdout = stdoutFile

	stderrFile, err := os.OpenFile(stderrPath, fileOpenFlag, fileOpenPerm)
	if err != nil {
		return nil, errors.WrapErrorf(err, "Opening stderr for task %s", jobId)
	}
	defer func() {
		_ = stderrFile.Close()
	}()

	cmd.Stderr = stderrFile

	// Stdout/stderr are redirected to the files
	_, _, exitStatus, runErr := f.cmdRunner.RunComplexCommand(cmd)

	stdout, err := ioutil.ReadFile(stdoutPath)
	if err != nil {
		return nil, errors.WrapErrorf(err, "ReadFile stdout for task %s", jobId)
	}

	stderr, err := ioutil.ReadFile(stderrPath)
	if err != nil {
		return nil, errors.WrapErrorf(err, "ReadFile stderr for task %s", jobId)
	}

	result := &CmdResult{
		Stdout:     stdout,
		Stderr:     stderr,
		ExitStatus: exitStatus,
	}

	if runErr != nil {
		return result, ScriptExecErr{result: result, err: runErr}
	}

	return result, nil
}
