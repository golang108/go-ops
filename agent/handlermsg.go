package agent

import (
	"encoding/json"
	"osp/internal/model"

	"github.com/luxingwen/pnet"
	"github.com/luxingwen/pnet/log"
)

func (ospAgent *OspAgent) HandlerFunc(msg interface{}, msgID []byte, srcID, rpath string, pn *pnet.PNet) {
	switch v := msg.(type) {
	case *model.ScriptJob:
		ospAgent.CreateScriptTask(v, srcID, msgID, rpath, pn)
	default:
		b, _ := json.Marshal(v)
		log.Error("msg handler not found,msg: %s", string(b))
	}
}
