package controller

import (
	"context"
	"go-ops/internal/model"
	"go-ops/internal/peer"
	"go-ops/internal/service"
	v1 "go-ops/pkg/api/v1"
)

var DownloadFileTask = &downloadFileTask{}

type downloadFileTask struct {
}

func (self *downloadFileTask) Create(ctx context.Context, req *v1.DownloadFileReq) (res *v1.DownloadfileRes, err error) {

	createFunc := func(peerid string, job *model.DownloadFileJob) error {
		err = peer.SendMsgAsync(peer.GetOspPeer().PNet, job, peerid, "")
		return err
	}
	taskid, err := service.Task().CreateFileDownload(ctx, req, createFunc)
	if err != nil {
		return
	}
	res = new(v1.DownloadfileRes)
	res.Taskid = taskid
	return
}

func (self *downloadFileTask) Details(ctx context.Context, req *v1.DownloadFileDetailsReq) (res *v1.DownloadfileRes, err error) {
	return service.Task().GetFileDownloadTaskInfo(ctx, req.Taskid)
}
