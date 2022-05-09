package agent

import (
	"context"
	"fmt"
	"go-ops/internal/model"
	"go-ops/internal/peer"
	"go-ops/pkg/agent/action"
	"path/filepath"

	"github.com/luxingwen/pnet"
)

func (self *OspAgent) ListFileInfo(req *model.PeerListFileInfo, srcId string, msgID []byte, rpath string, pn *pnet.PNet) {
	ctx := context.Background()
	fileInfos, err := action.FileDisk().GetDirInfo(ctx, req.Path)
	if err != nil {
		err = peer.SendMsgReplay(pn, &model.PeerListFileInfoRes{List: fileInfos, Path: req.Path, Error: err.Error()}, msgID, srcId, rpath)
		return
	}

	err = peer.SendMsgReplay(pn, &model.PeerListFileInfoRes{List: fileInfos, Path: req.Path}, msgID, srcId, rpath)
	if err != nil {
		fmt.Println("send msg replay err:", err)
	}
}

func (self *OspAgent) MoveFile(req *model.PeerMoveFile, srcId string, msgID []byte, rpath string, pn *pnet.PNet) {
	ctx := context.Background()
	err := action.FileDisk().Move(ctx, req.Src, req.Dst)
	if err != nil {
		err = peer.SendMsgReplay(pn, &model.PeerListFileInfoRes{Error: err.Error()}, msgID, srcId, rpath)

		return
	}

	bpath := filepath.Clean(req.Dst)

	fileInfos, err := action.FileDisk().GetDirInfo(ctx, bpath)
	if err != nil {
		err = peer.SendMsgReplay(pn, &model.PeerListFileInfoRes{Error: err.Error()}, msgID, srcId, rpath)

		return
	}

	err = peer.SendMsgReplay(pn, &model.PeerListFileInfoRes{List: fileInfos, Path: bpath}, msgID, srcId, rpath)
	if err != nil {
		fmt.Println("send msg replay err:", err)
	}
}

func (self *OspAgent) CreateDir(req *model.PeerNewDir, srcId string, msgID []byte, rpath string, pn *pnet.PNet) {
	ctx := context.Background()
	fileInfos, err := action.FileDisk().CreateDir(ctx, req.Path)
	if err != nil {
		err = peer.SendMsgReplay(pn, &model.PeerListFileInfoRes{Error: err.Error()}, msgID, srcId, rpath)
		return
	}

	err = peer.SendMsgReplay(pn, &model.PeerListFileInfoRes{List: fileInfos, Path: req.Path}, msgID, srcId, rpath)
	if err != nil {
		fmt.Println("send msg replay err:", err)
	}
}

func (self *OspAgent) RemoveFile(req *model.PeerDeleteFile, srcId string, msgID []byte, rpath string, pn *pnet.PNet) {
	ctx := context.Background()
	err := action.FileDisk().Remove(ctx, req.Path)
	if err != nil {
		err = peer.SendMsgReplay(pn, &model.PeerListFileInfoRes{Error: err.Error()}, msgID, srcId, rpath)
		fmt.Println("remove file err:", err)
		return
	}

	bpath := filepath.Dir(req.Path)

	fileInfos, err := action.FileDisk().GetDirInfo(ctx, bpath)
	if err != nil {
		fmt.Println("获取文件夹失败:", err)
		err = peer.SendMsgReplay(pn, &model.PeerListFileInfoRes{Error: err.Error()}, msgID, srcId, rpath)
		return
	}

	err = peer.SendMsgReplay(pn, &model.PeerListFileInfoRes{List: fileInfos, Path: bpath}, msgID, srcId, rpath)
	if err != nil {
		fmt.Println("send msg replay err:", err)
	}
}
