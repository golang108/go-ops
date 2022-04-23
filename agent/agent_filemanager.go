package agent

import (
	"context"
	"fmt"
	"go-ops/agent/action"
	"go-ops/internal/model"
	"go-ops/peer"
	"path/filepath"

	"github.com/luxingwen/pnet"
)

func (self *OspAgent) ListFileInfo(req *model.PeerListFileInfo, srcId string, msgID []byte, rpath string, pn *pnet.PNet) {
	ctx := context.Background()
	fileInfos, err := action.FileDisk().GetDirInfo(ctx, req.Path)
	if err != nil {
		return
	}

	err = peer.SendMsgAsync(pn, fileInfos, srcId, rpath)
	if err != nil {
		fmt.Println("send msg replay err:", err)
	}
}

func (self *OspAgent) MoveFile(req *model.PeerMoveFile, srcId string, msgID []byte, rpath string, pn *pnet.PNet) {
	ctx := context.Background()
	err := action.FileDisk().Move(ctx, req.Src, req.Dst)
	if err != nil {
		return
	}

	bpath := filepath.Base(req.Dst)

	fileInfos, err := action.FileDisk().GetDirInfo(ctx, bpath)
	if err != nil {
		return
	}

	err = peer.SendMsgAsync(pn, fileInfos, srcId, rpath)
	if err != nil {
		fmt.Println("send msg replay err:", err)
	}
}

func (self *OspAgent) CreateDir(req *model.PeerNewDir, srcId string, msgID []byte, rpath string, pn *pnet.PNet) {
	ctx := context.Background()
	fileInfos, err := action.FileDisk().CreateDir(ctx, req.Path)
	if err != nil {
		return
	}

	err = peer.SendMsgAsync(pn, fileInfos, srcId, rpath)
	if err != nil {
		fmt.Println("send msg replay err:", err)
	}
}

func (self *OspAgent) RemoveFile(req *model.PeerDeleteFile, srcId string, msgID []byte, rpath string, pn *pnet.PNet) {
	ctx := context.Background()
	err := action.FileDisk().Remove(ctx, req.Path)
	if err != nil {
		return
	}

	bpath := filepath.Base(req.Path)

	fileInfos, err := action.FileDisk().GetDirInfo(ctx, bpath)
	if err != nil {
		return
	}

	err = peer.SendMsgAsync(pn, fileInfos, srcId, rpath)
	if err != nil {
		fmt.Println("send msg replay err:", err)
	}
}
