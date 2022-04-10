package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/gogf/gf/v2/frame/g"

	"github.com/gogf/gf/v2/database/gdb"

	v1 "osp/api/v1"
	"osp/internal/model"
	"osp/model/entity"
	"osp/service/internal/dao"
	"osp/service/internal/do"

	"github.com/google/uuid"
)

type (
	sTask struct{}
)

var (
	insTask = sTask{}
)

func Task() *sTask {
	return &insTask
}

func (self *sTask) CreateScriptTask(ctx context.Context, req *v1.ScriptTask, createTask func(peerid string, scriptJob *model.ScriptJob) error) (taskId string, err error) {

	taskId = uuid.New().String()

	b, err := json.Marshal(req.Content)
	if err != nil {
		return
	}

	if req.Creater == "" {
		err = errors.New("creater is empty")
		return
	}

	t := do.Task{
		TaskId:  taskId,
		Content: string(b),
		Type:    "script",
		Name:    req.Name,
		Status:  "doing",
		Creater: req.Creater,
	}

	err = dao.Task.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		_, err = dao.Task.Ctx(ctx).Data(t).Insert()
		return err
	})

	if err != nil {
		return
	}

	for _, peerid := range req.Peers {
		jobid := uuid.New().String()

		item := do.Task{
			TaskId:   jobid,
			Type:     "script",
			Name:     peerid,
			Status:   "doing",
			Creater:  req.Creater,
			ParentId: taskId,
		}

		err = dao.Task.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
			_, err = dao.Task.Ctx(ctx).Data(item).Insert()
			return err
		})

		if err != nil {
			return taskId, err
		}

		scriptJob := model.ScriptJob{
			Jobid:  jobid,
			Script: req.Content,
		}

		err = createTask(peerid, &scriptJob)

		if err != nil {
			resCmd := model.ResCmd{
				Err:  err.Error(),
				Code: model.CodeNotRun,
			}
			self.UpdateSubScriptTask(ctx, &model.ResponseResCmd{
				Jobid:  jobid,
				PeerId: peerid,
				ResCmd: resCmd,
			})

		}
	}
	return
}

func (self *sTask) UpdateSubScriptTask(ctx context.Context, req *model.ResponseResCmd) (err error) {

	b, err := json.Marshal(req)
	if err != nil {
		return
	}
	status := "done"
	if req.ResCmd.Code != model.CodeSuccess {
		status = "failed"
	}

	item := do.Task{
		TaskId:  req.Jobid,
		Content: string(b),
		Status:  status,
	}

	err = dao.Task.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		_, err = dao.Task.Ctx(ctx).Data(item).Where("task_id = ?", req.Jobid).Update()
		return err
	})
	return err
}

func (self *sTask) UpdataScriptTaskStatus(ctx context.Context, taskid string, status string) (err error) {
	_, err = dao.Task.Ctx(ctx).Data("status = ", status).Where("task_id = ?", taskid).Update()
	return
}

func (self *sTask) GetScriptTask(ctx context.Context, taskid string) (r *v1.ScriptTaskExecRes, err error) {
	var mtask *entity.Task

	err = dao.Task.Ctx(ctx).Where("task_id = ?", taskid).Scan(&mtask)
	if err != nil {
		return
	}

	if mtask == nil {
		fmt.Println("mtask nil")
	}

	var subList []*entity.Task

	err = dao.Task.Ctx(ctx).Where("parent_id = ?", taskid).Scan(&subList)

	if err != nil {
		fmt.Println("err2->", err)
		return
	}

	resTaskExecList := make([]*v1.ScriptTaskExecItem, 0)

	doingcnt := 0
	failedcnt := 0
	for _, item := range subList {

		if item.Status == "doing" {
			doingcnt++
		}

		if item.Status == "failed" {
			failedcnt++
		}

		resTaskExecItem := new(v1.ScriptTaskExecItem)

		json.Unmarshal([]byte(item.Content), resTaskExecItem)
		resTaskExecItem.Status = item.Status
		resTaskExecList = append(resTaskExecList, resTaskExecItem)

	}

	mstatus := "doing"

	if mtask == nil {
		fmt.Println("mtask is nil")
	}
	if mtask.Status == "doing" {
		if doingcnt == 0 {
			if failedcnt == 0 {
				mstatus = "done"
			} else {
				mstatus = "failed"
			}

			err = dao.Task.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
				_, err = dao.Task.Ctx(ctx).Data(g.Map{"status": mstatus}).Where("task_id = ?", taskid).Update()
				return err
			})
			if err != nil {
				fmt.Println("update err:", err)
			}
		}
	}

	r = &v1.ScriptTaskExecRes{TaskId: taskid, Status: mstatus, List: resTaskExecList}
	return

}
