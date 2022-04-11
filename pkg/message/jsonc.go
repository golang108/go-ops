package message

import (
	"encoding/json"
	"osp/pkg/schema"

	"github.com/luxingwen/pnet/log"

	"github.com/gogo/protobuf/proto"
)

type jsonCodec struct {
}

// 编码器的名称
func (self *jsonCodec) Name() string {
	return "json"
}

// 将结构体编码为JSON的字节数组
func (self *jsonCodec) Encode(msgObj interface{}) (data []byte, err error) {

	v := MessageMetaByMsg(msgObj)
	if v == nil {
		log.Error("msg is nil meta:", msgObj)
		return
	}

	b, err := json.Marshal(msgObj)
	if err != nil {
		log.Error("json marshal err:", err)
		return
	}
	smsg := &schema.Msg{Id: v.ID, Data: b}

	data, err = proto.Marshal(smsg)

	if err != nil {
		log.Error("proto marshal err:", err)
		return
	}
	return

}

// 将JSON的字节数组解码为结构体
func (self *jsonCodec) Decode(data []byte) (r interface{}, err error) {

	rmsg := &schema.Msg{}
	err = proto.Unmarshal(data, rmsg)
	if err != nil {
		log.Error("proto unmarshal err:", err)
		return nil, err
	}

	r = MessageMetaByID(rmsg.Id).NewType()

	err = json.Unmarshal(rmsg.Data, r)
	return
}

var JSONCodec = new(jsonCodec)
