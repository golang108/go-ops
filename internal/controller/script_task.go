package controller

import (
	"context"
	"fmt"
	v1 "osp/api/v1"
	"osp/internal/model"
	"osp/peer"
	"osp/pkg/message"

	"github.com/google/uuid"
)

var ScritptTask = &scritptTask{}

type scritptTask struct {
}

func (c *scritptTask) CreateAsync(ctx context.Context, req *v1.ScriptTaskReq) (res *v1.ScriptTaskRes, err error) {

	taskId := uuid.New().String()
	res = &v1.ScriptTaskRes{
		TaskId: taskId,
	}

	for _, item := range req.Hostinfos {
		scriptJob := model.ScriptJob{
			Jobid:  uuid.New().String(),
			Script: req.Content,
		}
		err = peer.SendMsgAsync(peer.GetOspPeer().PNet, scriptJob, item, "")
		if err != nil {
			return
		}
	}
	return
}

func (c *scritptTask) CreateSync(ctx context.Context, req *v1.ScriptTaskSyncReq) (res *v1.ScriptTaskRes, err error) {

	taskId := uuid.New().String()
	res = &v1.ScriptTaskRes{
		TaskId: taskId,
	}

	for _, item := range req.Hostinfos {
		scriptJob := model.ScriptJob{
			Jobid:   uuid.New().String(),
			Script:  req.Content,
			RunMode: "sync",
		}
		r, err := peer.SendMsgSync(peer.GetOspPeer().PNet, scriptJob, item, "")
		if err != nil {
			return nil, err
		}

		v, err := message.JSONCodec.Decode(r)
		if err != nil {
			return nil, err
		}

		fmt.Println("v:", v)

	}

	return
}
