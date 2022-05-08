package action

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"context"
	"crypto/md5"
	"errors"
	"fmt"
	"go-ops/internal/model"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	"github.com/toolkits/file"
)

func Download(ctx context.Context, toFile, urlAddress string) (err error) {
	hreq, err := http.NewRequestWithContext(ctx, "GET", urlAddress, nil)
	if err != nil {
		return
	}

	client := &http.Client{}
	resp, err := client.Do(hreq)
	if err != nil {

		return
	}

	defer resp.Body.Close()

	f, err := os.Create(toFile)
	if err != nil {
		return
	}

	defer f.Close()

	io.Copy(f, resp.Body)

	return
}

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

// 解压
func Untar(filename string) (err error) {

	dir := filepath.Dir(filename)

	tmpDir := dir + string(filepath.Separator) + "tmp"

	file.EnsureDir(tmpDir)

	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}
	bufReader := bytes.NewReader(content)

	unGzipReader, err := gzip.NewReader(bufReader)
	if err != nil {
		return
	}

	tarReader := tar.NewReader(unGzipReader)

	for {
		fileHeader, err := tarReader.Next()
		if err == io.EOF {
			break
		}

		if fileHeader == nil {
			err = errors.New("tar.gz is bad")
			return err
		}

		fileInfo := fileHeader.FileInfo()

		path := tmpDir + string(filepath.Separator) + fileHeader.Name
		if fileInfo.IsDir() {
			err = os.MkdirAll(path, os.ModeDir|0666)
			if err != nil {
				return err
			}

			err = os.Chmod(path, fileInfo.Mode())
			if err != nil {
				return err
			}
			continue
		}

		f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
		if err != nil {
			return err
		}

		_, err = io.Copy(f, tarReader)
		if err != nil {
			f.Close()
			return err
		}
		f.Close()

		err = os.Chmod(path, fileInfo.Mode())
		if err != nil {
			return err
		}
	}

	err = os.Rename(tmpDir, dir)
	if err != nil {
		return
	}

	os.RemoveAll(tmpDir)

	return

}

func CheckFileMd5(filename, md5str string) (err error) {
	filecontent, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}

	fileMd5 := fmt.Sprintln("%x", md5.Sum(filecontent))

	if fileMd5 != md5str {
		err = errors.New("check md5 failed")
		return
	}
	return
}
