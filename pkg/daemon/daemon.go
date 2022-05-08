package daemon

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"time"
)

const ENV_NAME = "LXW_DAEMON_IDX"

//运行时调用background的次数
var runIdx int = 0

type Daemon struct {
	pid  int
	cpid int
}

func NewDaemon() *Daemon {
	return &Daemon{}
}

func (self *Daemon) Start() {
	Background(true)

	for {
		cmd, err := Background(false)
		if err != nil {
			log.Println("start proc err:", err)
			time.Sleep(time.Second)
			continue

		}

		//	self.cpid = cmd.Process.Pid
		err = cmd.Wait()
		if err != nil {
			log.Println("sub procss exit err:", err)
			time.Sleep(time.Second)
		}
	}
}

// auto Start 是否自动重启
func Background(isExit bool) (cmd *exec.Cmd, err error) {

	var runIdx int = 0

	envIdx, err := strconv.Atoi(os.Getenv(ENV_NAME))
	if err != nil {
		envIdx = 0
	}
	if runIdx > envIdx { //子进程, 退出
		return nil, err
	}

	//设置子进程环境变量
	env := os.Environ()
	env = append(env, fmt.Sprintf("%s=%d", ENV_NAME, runIdx))

	//启动子进程
	cmd, err = startProc(os.Args, env)
	if err != nil {
		log.Println(os.Getpid(), "启动子进程失败:", err)
		return nil, err
	}

	if isExit {
		os.Exit(0)
	}

	return cmd, nil
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
