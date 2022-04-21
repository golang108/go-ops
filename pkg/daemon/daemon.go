package daemon

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
)

const ENV_NAME = "LXW_DAEMON_IDX"

//运行时调用background的次数
var runIdx int = 0

// auto Start 是否自动重启
func Background(autoStart bool, f func()) (cmd *exec.Cmd, err error) {

	var runIdx int = 0

	envIdx, err := strconv.Atoi(os.Getenv(ENV_NAME))
	if err != nil {
		envIdx = 0
	}
	if runIdx <= envIdx { //子进程, 退出
		if autoStart {
			AutoRestart(autoStart, f)
			return
		}
		return  nil
	}

	//设置子进程环境变量
	env := os.Environ()
	env = append(env, fmt.Sprintf("%s=%d", ENV_NAME, runIdx))

	//启动子进程
	cmd,err = startProc(os.Args, env)
	if err != nil {
		log.Println(os.Getpid(), "启动子进程失败:", err)
		return nil, err
	} 

	return cmd, nil
}


// 用守护进程的方式 运行
func runDaemon(autoStart bool, f func()) (err error) {

	cmd,err:=Background(false, f)
	if err!=nil{

	}
	return
}

func startProc(args, env []string) (*exec.Cmd, error) {
	cmd := &exec.Cmd{
		Path:        args[0],
		Args:        args,
		Env:         env,
		SysProcAttr: NewSysProcAttr(),
	}

	err := cmd.Start()
	if err != nil {
		return nil, err
	}

	return cmd, nil
}
