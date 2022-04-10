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

	"github.com/google/uuid"
)

var ScritptTask = &scritptTask{}

type scritptTask struct {
}

func (c *scritptTask) CreateAsync(ctx context.Context, req *v1.ScriptTaskReq) (res *v1.ScriptTaskRes, err error) {

	createFunc := func(peerid string, scriptJob *model.ScriptJob) error {
		err = peer.SendMsgAsync(peer.GetOspPeer().PNet, scriptJob, peerid, "")
		return err
	}
	taskid, err := service.Task().CreateScriptTask(ctx, req, createFunc)
	if err != nil {
		return
	}
	res = new(v1.ScriptTaskRes)
	res.TaskId = taskid
	return
}

func (c *scritptTask) CreateSync(ctx context.Context, req *v1.ScriptTaskSyncReq) (res *v1.ScriptRes, err error) {

	res = &v1.ScriptRes{}
	wg := sync.WaitGroup{}

	for _, item := range req.Peers {
		wg.Add(1)
		go func(peerid string) {
			defer func() {
				wg.Done()
			}()
			jobid := uuid.New().String()
			scriptJob := model.ScriptJob{
				Jobid:   jobid,
				Script:  req.Content,
				RunMode: "sync",
			}

			r, err := peer.SendMsgSyncWithTimeout(peer.GetOspPeer().PNet, scriptJob, peerid, "", time.Duration(req.Content.Timeout*int(time.Second))+time.Second*20)
			if err != nil {
				resCmd := model.ResCmd{Err: err.Error()}
				resResponse := &model.ResponseResCmd{Jobid: jobid, PeerId: peerid, ResCmd: resCmd}
				res.List = append(res.List, resResponse)
				return
			}

			v, err := message.JSONCodec.Decode(r)
			if err != nil {
				resCmd := model.ResCmd{Err: err.Error()}
				resResponse := &model.ResponseResCmd{Jobid: jobid, PeerId: peerid, ResCmd: resCmd}
				res.List = append(res.List, resResponse)
				return
			}

			val := v.(*model.ResponseResCmd)
			res.List = append(res.List, val)
		}(item)

	}

	wg.Wait()

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

func (s *scritptTask) GetTaskInfo(ctx context.Context, req *v1.ScriptTaskInfoReq) (res *v1.ScriptTaskInfoRes, err error) {
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
