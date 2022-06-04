package controller

import (
	"context"
	"go-ops/internal/model"
	"go-ops/internal/peer"
	"go-ops/internal/service"
	v1 "go-ops/pkg/api/v1"
	"go-ops/pkg/message"
	"sync"
	"time"

	"github.com/gogf/gf/v2/os/glog"
)

var ScritptTask = &scritptTask{}

type scritptTask struct {
}

func (c *scritptTask) CreateAsync(ctx context.Context, req *v1.ScriptTaskReq) (res *v1.ScriptTaskRes, err error) {

	createFunc := func(peerid string, scriptJob *model.ScriptJob) error {
		err = peer.SendMsgAsync(peer.GetOspPeer().PNet, scriptJob, peerid, "")
		if err != nil {
			glog.Errorf(ctx, "create script task send msg async err:%v", err)
		}
		return err
	}
	taskid, err := service.Task().CreateScriptTask(ctx, &req.ScriptTask, createFunc)
	if err != nil {
		glog.Errorf(ctx, "create script task  async err:%v", err)
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

				glog.Errorf(ctx, "create script task sync send msg with timeout err:%v", err)

				resCmd := model.ResCmd{Err: err.Error(), Code: model.CodeFailed}
				resResponse := &model.ResponseResCmd{Jobid: scriptJob.Jobid, PeerId: peerid, ResCmd: resCmd}

				res.List = append(res.List, &v1.ScriptTaskExecItem{ResponseResCmd: resResponse, Status: "failed"})
				err0 := service.Task().UpdateSubScriptTask(ctx, resResponse)
				if err0 != nil {
					glog.Errorf(ctx, "create script task sync update sub script task err:%v", err0)
					return
				}

				return
			}

			v, err := message.JSONCodec.Decode(r)
			if err != nil {
				glog.Errorf(ctx, "create script task sync json decode err:%v", err)
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
		glog.Errorf(ctx, "create script task sync err:%v", err)
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

	err = service.Task().UpdataScriptTaskStatus(ctx, taskid, res.Status)
	if err != nil {
		glog.Errorf(ctx, "update script task status err:%v", err)
		return
	}
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
		glog.Errorf(ctx, "scripttask get task info err:%v", err)
		return
	}
	val, err := message.JSONCodec.Decode(r)
	if err != nil {
		glog.Errorf(ctx, "scripttask get task info  json decode err:%v", err)
		return
	}
	res.TaskInfo = val.(*model.TaskInfo)
	res.PeerId = req.PeerId
	return

}

func (s *scritptTask) GetScriptTaskInfo(ctx context.Context, req *v1.ScriptTaskInfoReq) (res *v1.ScriptTaskExecRes, err error) {
	return service.Task().GetScriptTask(ctx, req.TaskId)
}
