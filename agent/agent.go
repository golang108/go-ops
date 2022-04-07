package agent

import (
	"fmt"
	"osp/agent/model"
	"osp/agent/script"
	"osp/agent/task"
	"osp/peer"

	"github.com/luxingwen/pnet"
)

type OspAgent struct {
	task.Manager
	task.Service
}

func NewOspAgent(workdir string) *OspAgent {

	ospAgent := &OspAgent{
		Manager: task.NewManagerProvider().NewManager(workdir),
		Service: task.NewAsyncTaskService(),
	}
	return ospAgent
}

func (self *OspAgent) CreateScriptTask(s *model.ScriptJob, srcId string, msgID []byte, rpath string, pn *pnet.PNet) {
	startFunc := func() (r interface{}, err error) {
		res := script.NewJobScriptProvider(*s).Run()
		r = &model.ResponseResCmd{
			Jobid:  s.Jobid,
			ResCmd: res,
		}
		return
	}

	endFunc := func(t task.Task) {

		if s.RunMode == "sync" {
			err := peer.SendMsgReplay(pn, t.Value, msgID, srcId, rpath)
			if err != nil {
				fmt.Println("send msg replay err:", err)
			}
			return
		}
		err := peer.SendMsgAsync(pn, t.Value, srcId, rpath)
		if err != nil {
			fmt.Println("send msg replay err:", err)
		}
	}

	t := self.CreateTask(s.Jobid, startFunc, nil, endFunc)
	self.StartTask(t)

}
