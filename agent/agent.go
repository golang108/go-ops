package agent

import (
	"context"
	"fmt"

	"osp/agent/action"
	"osp/agent/script"
	"osp/agent/task"
	"osp/internal/model"
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

	ctx, cancel := context.WithCancel(context.Background())

	startFunc := func() (r interface{}, err error) {

		res := script.NewJobScriptProvider(ctx, *s).Run()
		r = &model.ResponseResCmd{
			Jobid:  s.Jobid,
			ResCmd: res,
			PeerId: pn.GetLocalNode().Id,
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
		fmt.Println("我来到了这里")
		err := peer.SendMsgAsync(pn, t.Value, srcId, rpath)
		if err != nil {
			fmt.Println("send msg replay err:", err)
		}
	}

	c := func(t task.Task) error {
		cancel()
		return nil
	}

	t := self.CreateTask(s.Jobid, s, startFunc, c, endFunc)
	self.StartTask(t)
}

func (self *OspAgent) CancelcriptTask(taskid string, srcId string, msgID []byte, rpath string, pn *pnet.PNet) {
	task, ok := self.FindTaskWithID(taskid)
	if !ok {
		pn.SendBytesRelayReply(msgID, []byte("没有找到task："+taskid), srcId)
		return
	}

	task.Cancel()
	pn.SendBytesRelayReply(msgID, []byte("任务已经取消："+taskid), srcId)
	return
}

func (self *OspAgent) GetTaskInfo(taskid string, srcId string, msgID []byte, rpath string, pn *pnet.PNet) {
	task, ok := self.FindTaskWithID(taskid)
	res := model.TaskInfo{}
	if !ok {
		res.Err = "找不到任务:" + taskid
		err := peer.SendMsgReplay(pn, res, msgID, srcId, rpath)
		if err != nil {
			fmt.Println("send msg replay err:", err)
		}
		return
	}

	res.Req = task.Req
	res.Value = task.Value
	res.Status = string(task.State)

	peer.SendMsgReplay(pn, res, msgID, srcId, rpath)

}

func (self *OspAgent) DownloadFile(req *model.DownloadFileJob, srcId string, msgID []byte, rpath string, pn *pnet.PNet) {
	ctx, cancel := context.WithCancel(context.Background())

	startFunc := func() (r interface{}, err error) {

		res := action.DownloadFile(ctx, req.DownloadFileInfo)
		resJob := &model.DownloadFileJobRes{
			Jobid:           req.Jobid,
			PeerId:          pn.GetLocalNode().GetId(),
			DownloadFileRes: res,
		}
		r = resJob
		return
	}

	endFunc := func(t task.Task) {
		err := peer.SendMsgAsync(pn, t.Value, srcId, rpath)
		if err != nil {
			fmt.Println("send msg replay err:", err)
		}
	}

	c := func(t task.Task) error {
		cancel()
		return nil
	}

	t := self.CreateTask(req.Jobid, req, startFunc, c, endFunc)
	self.StartTask(t)
}
