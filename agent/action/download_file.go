package action

import (
	"context"
	"io"
	"net/http"
	"os"
	"osp/internal/model"
	"path/filepath"
)

func DownloadFile(ctx context.Context, req *model.DownloadFileInfo) (res *model.DownloadFileRes) {

	savePath := filepath.Dir(req.Path)

	res = new(model.DownloadFileRes)
	res.Filename = req.Filename
	if _, err := os.Stat(savePath); err != nil {
		if !req.AutoCreatePath {
			res.Msg = "文件夹不存在"
			res.Code = model.CodeFailed
			return
		}
		if err = os.MkdirAll(savePath, os.ModePerm); err != nil {
			res.Msg = "Create dir err:" + err.Error()
			res.Code = model.CodeFailed
			return
		}
	}

	if _, err := os.Stat(req.Path); err == nil && !req.Replace {

		res.Msg = "文件已经存在且不需要覆盖"
		res.Code = model.CodeNotRun
		return
	}

	hreq, err := http.NewRequestWithContext(ctx, "GET", req.Address, nil)
	if err != nil {
		res.Code = model.CodeDownloadFailed
		res.Msg = err.Error()
		return
	}

	client := &http.Client{}
	resp, err := client.Do(hreq)
	if err != nil {
		res.Code = model.CodeDownloadFailed
		res.Msg = err.Error()
		return
	}

	defer resp.Body.Close()

	f, err := os.Create(req.Path)
	if err != nil {
		res.Msg = "Create file err:" + err.Error()
		res.Code = model.CodeFailed
		return
	}

	defer f.Close()

	io.Copy(f, resp.Body)

	res.Code = model.CodeSuccess
	res.Msg = "sucess"
	return
}
