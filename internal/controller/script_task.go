package controller

import (
	"context"
	v1 "osp/api/v1"
	"osp/internal/model"
	"osp/peer"
	"osp/pkg/message"
	"osp/service"
	"sync"
	"time"
)

var ScritptTask = &scritptTask{}

type scritptTask struct {
}

func (c *scritptTask) CreateAsync(ctx context.Context, req *v1.ScriptTaskReq) (res *v1.ScriptTaskRes, err error) {

	createFunc := func(peerid string, scriptJob *model.ScriptJob) error {
		err = peer.SendMsgAsync(peer.GetOspPeer().PNet, scriptJob, peerid, "")
		return err
	}
	taskid, err := service.Task().CreateScriptTask(ctx, &req.ScriptTask, createFunc)
	if err != nil {
		return
	}
	res = new(v1.ScriptTaskRes)
	res.TaskId = taskid
	return
}

func (c *scritptTask) CreateSync(ctx context.Context, req *v1.ScriptTaskSyncReq) (res *v1.ScriptTaskExecRes, err error) {

	res = &v1.ScriptTaskExecRes{}
	wg := sync.WaitGroup{}

	createFunc := func(peerid string, scriptJob *model.ScriptJob) error {
		wg.Add(1)
		go func() {
			defer func() {
				wg.Done()
			}()

			scriptJob.RunMode = "sync"

			r, err := peer.SendMsgSyncWithTimeout(peer.GetOspPeer().PNet, scriptJob, peerid, "", time.Duration(req.Content.Timeout*int(time.Second))+time.Second*20)
			if err != nil {
				resCmd := model.ResCmd{Err: err.Error(), Code: model.CodeFailed}
				resResponse := &model.ResponseResCmd{Jobid: scriptJob.Jobid, PeerId: peerid, ResCmd: resCmd}

				res.List = append(res.List, &v1.ScriptTaskExecItem{ResponseResCmd: resResponse, Status: "failed"})
				service.Task().UpdateSubScriptTask(ctx, resResponse)

				return
			}

			v, err := message.JSONCodec.Decode(r)
			if err != nil {
				resCmd := model.ResCmd{Err: err.Error()}
				resResponse := &model.ResponseResCmd{Jobid: scriptJob.Jobid, PeerId: peerid, ResCmd: resCmd}
				res.List = append(res.List, &v1.ScriptTaskExecItem{ResponseResCmd: resResponse, Status: "failed"})
				service.Task().UpdateSubScriptTask(ctx, resResponse)

				return
			}

			val := v.(*model.ResponseResCmd)

			service.Task().UpdateSubScriptTask(ctx, val)
			status := "done"

			if val.ResCmd.Code != model.CodeSuccess {
				status = "failed"
			}
			res.List = append(res.List, &v1.ScriptTaskExecItem{ResponseResCmd: val, Status: status})
		}()
		return nil
	}

	taskid, err := service.Task().CreateScriptTask(ctx, &req.ScriptTask, createFunc)
	if err != nil {
		return
	}

	wg.Wait()

	failedcnt := 0
	for _, item := range res.List {
		if item.Status == "failed" {
			failedcnt++
		}
	}
	res.Status = "failed"
	if failedcnt == 0 {
		res.Status = "done"
	}

	service.Task().UpdataScriptTaskStatus(ctx, taskid, res.Status)
	res.TaskId = taskid

	return
}

func (c *scritptTask) CancelTask(ctx context.Context, req *v1.ScriptTaskCancelReq) (res *v1.ScriptTaskCancelRes, err error) {

	res = new(v1.ScriptTaskCancelRes)
	for _, item := range req.Tasks {
		job := &model.ScriptJobCancel{
			Jobid: item.Jobid,
		}
		r, err := peer.SendMsgSync(peer.GetOspPeer().PNet, job, item.PeerId, "")
		resItem := &v1.ScriptTaskCancel{PeerId: item.PeerId, Jobid: item.Jobid}
		if err != nil {
			resItem.Msg = err.Error()
			res.List = append(res.List, resItem)
			continue
		}
		resItem.Msg = string(r)

		res.List = append(res.List, resItem)
	}

	return
}

func (s *scritptTask) GetTaskInfo(ctx context.Context, req *v1.PeerScriptTaskInfoReq) (res *v1.ScriptTaskInfoRes, err error) {
	res = new(v1.ScriptTaskInfoRes)
	r, err := peer.SendMsgSync(peer.GetOspPeer().PNet, &model.GetTaskInfo{TaskId: req.TaskId}, req.PeerId, "")
	if err != nil {
		return
	}

	val, err := message.JSONCodec.Decode(r)
	if err != nil {
		return
	}

	res.TaskInfo = val.(*model.TaskInfo)
	res.PeerId = req.PeerId
	return

}

func (s *scritptTask) GetScriptTaskInfo(ctx context.Context, req *v1.ScriptTaskInfoReq) (res *v1.ScriptTaskExecRes, err error) {
	return service.Task().GetScriptTask(ctx, req.TaskId)
}
