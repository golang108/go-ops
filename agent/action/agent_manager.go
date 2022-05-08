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
	"os"
	"path/filepath"
	"runtime"

	"github.com/toolkits/file"
)

const (
	AGENT_RUNNING = "running"
	AGENT_STOPPED = "stopped"
	AGENT_DELETED = "deleted"
)

type agentManager struct {
	agentDir string
}

var sAgentManager *agentManager

func AgentManager() *agentManager {
	return sAgentManager
}

func (self *agentManager) getAgentTar(a *model.AgentInfo) string {
	return fmt.Sprintf("%s-%s-%s-%s.tar.gz", a.Name, runtime.GOOS, runtime.GOARCH, a.Version)
}

func (self *agentManager) getAgentTarMd5(a *model.AgentInfo) string {
	return fmt.Sprintf("%s-%s-%s-%s.tar.gz.md5", a.Name, runtime.GOOS, runtime.GOARCH, a.Version)
}

func (self *agentManager) getAgentBackupName(version string, a *model.AgentInfo) string {
	return fmt.Sprintf("%s-%s", a.Name, version)
}

func (self *agentManager) Download(ctx context.Context, a *model.AgentInfo) (err error) {
	tarName := self.getAgentTar(a)
	md5Name := self.getAgentTarMd5(a)
	downloadUrl := a.UrlAddress + "/" + tarName
	tmpDir := self.agentDir + string(filepath.Separator) + "tmp"
	tmpFile := tmpDir + string(filepath.Separator) + tarName
	md5fileName := tmpDir + string(filepath.Separator) + md5Name
	md5fileDownloadUrl := a.UrlAddress + "/" + md5Name

	file.EnsureDir(tmpDir)

	err = Download(ctx, tmpFile, downloadUrl)
	if err != nil {
		return
	}

	err = Download(ctx, md5fileName, md5fileDownloadUrl)
	if err != nil {
		return
	}

	md5Content, err := ioutil.ReadFile(md5fileName)

	if err != nil {
		return
	}

	filecontent, err := ioutil.ReadFile(tmpFile)
	if err != nil {
		return
	}

	fileMd5 := fmt.Sprintf("%x", md5.Sum(filecontent))

	if fileMd5 != string(md5Content) {
		err = errors.New("check md5 failed")
		return
	}
	return
}

func (self *agentManager) Untar(a *model.AgentInfo) (err error) {
	tarName := self.getAgentTar(a)
	dstDir := self.agentDir + string(filepath.Separator) + a.Name
	tmpDir := self.agentDir + string(filepath.Separator) + "tmp"
	tmpFile := tmpDir + string(filepath.Separator) + tarName

	file.EnsureDir(dstDir)

	content, err := ioutil.ReadFile(tmpFile)
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

	err = os.Rename(tmpDir+string(filepath.Separator)+a.Name, dstDir)
	if err != nil {
		return
	}

	os.RemoveAll(tmpDir)

	return

}

func (self *agentManager) Backup(a *model.AgentInfo, version string) (err error) {
	agentDir := self.agentDir + string(filepath.Separator) + a.Name

	if !file.IsExist(agentDir) {
		return nil
	}

	status, err := self.Status(a)
	if err != nil {
		return
	}

	if status == AGENT_RUNNING {
		err = self.Stop(a)
		if err != nil {
			return err
		}
	}

	if version == "" {

		version, err = self.Version(a)
		if err != nil {

			os.RemoveAll(agentDir)

			return nil
		}
	}

	backupName := self.getAgentBackupName(version, a)

	backupDir := self.agentDir + string(filepath.Separator) + "backup" + string(filepath.Separator) + backupName

	file.EnsureDir(self.agentDir + string(filepath.Separator) + "backup")

	if file.IsExist(backupDir) {
		os.RemoveAll(backupDir)
	}

	err = os.Rename(agentDir, backupDir)

	if err != nil {
		return
	}

	return

}

func (self *agentManager) GetAgentStatus(a *model.AgentInfo) (agentStatus *model.AgentInfo) {
	controlExist := self.ControlScriptCheck(a)
	if !controlExist {
		a.Status = "deleted"
		return a
	}

	status, err := self.Status(a)
	if err != nil {
		a.Status = "unknow"
	}

	a.Status = status
	version, err := self.Version(a)
	if err != nil {
		a.Version = "unknow"
	}
	a.Version = version
	return a
}

func (self *agentManager) Install(ctx context.Context, a *model.AgentInfo) (err error) {
	err = self.Download(ctx, a)
	if err != nil {
		return
	}

	err = self.Backup(a, "")
	if err != nil {
		return
	}

	err = self.Untar(a)
	if err != nil {
		return
	}
	err = self.Start(a)
	return

}

func (self *agentManager) CheckAgentStatus(a *model.AgentInfo) (r *model.AgentInfo) {
	switch a.Status {
	case AGENT_RUNNING:

	case AGENT_DELETED:
	case AGENT_STOPPED:
		self.Stop(a)

	}
	return
}
