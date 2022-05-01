package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/gogf/gf/v2/frame/g"

	"github.com/gogf/gf/v2/database/gdb"

	v1 "go-ops/api/v1"
	"go-ops/internal/model"
	"go-ops/internal/model/entity"
	"go-ops/internal/service/internal/dao"
	"go-ops/internal/service/internal/do"

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
	_, err = dao.Task.Ctx(ctx).Data(g.Map{"status": status}).Where("task_id = ?", taskid).Update()
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
		return
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

	mstatus := mtask.Status

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

func (s *sTask) CreateFileDownload(ctx context.Context, req *v1.DownloadFileReq, createTask func(string, *model.DownloadFileJob) error) (taskid string, err error) {
	taskid = uuid.New().String()
	b, err := json.Marshal(req.DownloadFileInfo)
	if err != nil {
		return
	}
	t := entity.Task{
		TaskId:  taskid,
		Content: string(b),
		Type:    "file_download",
		Name:    req.Name,
		Status:  "doing",
		Creater: req.Creater,
	}

	_, err = dao.Task.Ctx(ctx).Data(t).Insert()

	if err != nil {
		return
	}

	for _, item := range req.Files {

		for _, peerid := range req.Peers {

			itaskId := uuid.New().String()

			taskRes := &model.DownloadFileJobRes{
				Jobid:  itaskId,
				PeerId: peerid,
				DownloadFileRes: &model.DownloadFileRes{
					Filename: item.Filename,
				},
			}

			content, _ := json.Marshal(taskRes)

			titem := entity.Task{
				TaskId:   itaskId,
				Content:  string(content),
				Type:     "file_download",
				Name:     peerid,
				Status:   "doing",
				Creater:  req.Creater,
				ParentId: taskid,
			}

			_, err = dao.Task.Ctx(ctx).Data(titem).Insert()
			if err != nil {
				return
			}

			err = createTask(peerid, &model.DownloadFileJob{
				Jobid:            itaskId,
				DownloadFileInfo: item,
			})

			if err != nil {
				taskRes.Code = model.CodeFailed
				taskRes.Msg = err.Error()
				content, _ = json.Marshal(taskRes)
				titem.Content = string(content)
				titem.Status = "failed"
				_, err = dao.Task.Ctx(ctx).Data(item).Where("task_id = ?", titem.TaskId).Update()
				if err != nil {
					return
				}
			}

		}
	}
	return
}

func (self *sTask) UpdateDownloadFileTask(ctx context.Context, req *model.DownloadFileJobRes) (err error) {

	b, err := json.Marshal(req)
	if err != nil {
		return
	}
	status := "done"
	if req.Code != model.CodeSuccess {
		status = "failed"
	}

	item := do.Task{
		TaskId:  req.Jobid,
		Content: string(b),
		Status:  status,
	}

	_, err = dao.Task.Ctx(ctx).Data(item).Where("task_id = ?", req.Jobid).Update()

	return err
}

func (self *sTask) GetFileDownloadTaskInfo(ctx context.Context, taskid string) (res *v1.DownloadfileRes, err error) {
	var mtask *entity.Task

	err = dao.Task.Ctx(ctx).Where("task_id = ?", taskid).Scan(&mtask)
	if err != nil {
		return
	}

	if mtask == nil {
		fmt.Println("mtask nil")
		return
	}

	var subList []*entity.Task

	err = dao.Task.Ctx(ctx).Where("parent_id = ?", taskid).Scan(&subList)

	if err != nil {
		fmt.Println("err2->", err)
		return
	}

	resFileTaskList := make([]*v1.DownloadfileItem, 0)

	doingcnt := 0
	failedcnt := 0
	for _, item := range subList {

		if item.Status == "doing" {
			doingcnt++
		}

		if item.Status == "failed" {
			failedcnt++
		}

		resTaskFileItem := new(v1.DownloadfileItem)

		json.Unmarshal([]byte(item.Content), resTaskFileItem)
		resTaskFileItem.Status = item.Status
		resFileTaskList = append(resFileTaskList, resTaskFileItem)

	}

	mstatus := mtask.Status

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

			_, err = dao.Task.Ctx(ctx).Data(g.Map{"status": mstatus}).Where("task_id = ?", taskid).Update()

			if err != nil {
				fmt.Println("update err:", err)
			}
		}
	}

	res = &v1.DownloadfileRes{
		Taskid: taskid,
		Status: mstatus,
		List:   resFileTaskList,
	}

	return
}

func (self *sTask) QueryTask(ctx context.Context, req *v1.TaskQueryReq) (res *v1.TaskInfoRes, err error) {

	m := g.Map{"parent_id": ""}

	if req.Name != "" {
		m["name"] = req.Name
	}

	if req.Creater != "" {
		m["creater"] = req.Creater
	}

	if req.TaskID !="" {
		m["task_id"] = req.TaskID
	} 

	tasks := make([]*entity.Task, 0)

	err = dao.Task.Ctx(ctx).Where(m).Page(req.PageNum, req.PageSize).Scan(&tasks)

	if err != nil {
		return
	}

	count, err := dao.Task.Ctx(ctx).Where(m).Count()
	if err != nil {
		return
	}

	res = &v1.TaskInfoRes{
		List: tasks,
	}

	res.Total = count
	res.PageNum = req.PageNum
	res.PageSize = req.PageSize

	if res.Total%res.PageSize > 0 {
		res.PageTotal = 1
	}
	res.PageTotal += res.Total / res.PageSize

	return
}

func (self *sTask) GetTaskInfo(ctx context.Context, req *v1.TaskInfoReq) (taskinfo *v1.TaskInfo, err error) {
	return getTaskInfo(ctx, req.TaskID)
}

func getTaskInfo(ctx context.Context, taskid string) (taskInfo *v1.TaskInfo, err error) {

	var task *entity.Task

	err = dao.Task.Ctx(ctx).Where("task_id = ?", taskid).Scan(&task)
	if err != nil {
		return
	}

	if task == nil {
		return
	}

	subTasks := make([]*entity.Task, 0)

	err = dao.Task.Ctx(ctx).Where("parent_id = ?", taskid).Scan(&subTasks)
	if err != nil {
		return
	}

	taskInfo = &v1.TaskInfo{
		Task:    task,
		Sublist: make([]*v1.TaskInfo, 0),
	}

	for _, item := range subTasks {
		subTask, err := getTaskInfo(ctx, item.TaskId)
		if err != nil {
			continue
		}
		taskInfo.Sublist = append(taskInfo.Sublist, subTask)
	}
	return
}
