package system

import (
	"fmt"
	"os/exec"
	"strings"
)

func newExecCmd(name string, args ...string) *exec.Cmd {
	fmt.Println("cmd:", name, strings.Join(args, " "))
	return exec.Command(name, args...)
}

func mergeEnv(sysEnv []string, cmdEnv map[string]string) []string {
	var env []string
	for k, v := range cmdEnv {
		env = append(env, k+"="+v)
	}
	for _, s := range sysEnv {
		if n := strings.IndexByte(s, '='); n != -1 {
			k := s[:n] // key
			if _, found := cmdEnv[k]; !found {
				env = append(env, s)
			}
		}
	}
	return env
}
