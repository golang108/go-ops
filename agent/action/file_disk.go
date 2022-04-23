package action

import (
	"context"
	"go-ops/internal/model"
	"io/ioutil"
	"os"
	"path/filepath"
)

type fileDisk struct {
}

func FileDisk() *fileDisk {
	return &fileDisk{}
}

func (self *fileDisk) GetDirInfo(ctx context.Context, spath string) (r []*model.FileInfo, err error) {
	infos, err := ioutil.ReadDir(spath)
	if err != nil {
		return
	}
	for _, item := range infos {
		fi := &model.FileInfo{
			Name:    item.Name(),
			ModTime: item.ModTime(),
			Size:    item.Size(),
			Path:    spath,
		}
		if item.IsDir() {
			fi.Type = "dir"
		} else {
			fi.Path = filepath.Join(spath, item.Name())
		}
		r = append(r, fi)
	}
	return
}

func (self *fileDisk) Remove(ctx context.Context, filename string) (err error) {
	err = os.Remove(filename)
	return
}

func (self *fileDisk) Move(ctx context.Context, src, dest string) (err error) {
	err = os.Rename(src, dest)
	return
}

func (self *fileDisk) CreateDir(ctx context.Context, spath string) (r []*model.FileInfo, err error) {

	if err = os.MkdirAll(spath, os.ModePerm); err != nil {
		return
	}

	return self.GetDirInfo(ctx, spath)
}
