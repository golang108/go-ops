package message

import (
	"fmt"
	"path"
	"reflect"
	"strings"
)

// 消息元信息
type MessageMeta struct {
	//Codec Codec        // 消息用到的编码
	Type reflect.Type // 消息类型, 注册时使用指针类型

	ID uint32 // 消息ID (二进制协议中使用)

}

func (self *MessageMeta) TypeName() string {

	if self == nil {
		return ""
	}

	return self.Type.Name()
}

func (self *MessageMeta) FullName() string {

	if self == nil {
		return ""
	}

	var sb strings.Builder
	sb.WriteString(path.Base(self.Type.PkgPath()))
	sb.WriteString(".")
	sb.WriteString(self.Type.Name())

	return sb.String()
}

// 创建meta类型的实例
func (self *MessageMeta) NewType() interface{} {
	if self.Type == nil {
		return nil
	}

	return reflect.New(self.Type).Interface()
}

var (
	// 消息元信息与消息名称，消息ID和消息类型的关联关系
	metaByFullName = map[string]*MessageMeta{}
	metaByID       = map[uint32]*MessageMeta{}
	metaByType     = map[reflect.Type]*MessageMeta{}
)

// 注册消息元信息
func RegisterMessage(meta *MessageMeta) *MessageMeta {

	// 注册时, 统一为非指针类型
	if meta.Type.Kind() == reflect.Ptr {
		meta.Type = meta.Type.Elem()
	}

	if _, ok := metaByType[meta.Type]; ok {
		panic(fmt.Sprintf("Duplicate message meta register by type: %d name: %s", meta.ID, meta.Type.Name()))
	} else {
		metaByType[meta.Type] = meta
	}

	if _, ok := metaByFullName[meta.FullName()]; ok {
		panic(fmt.Sprintf("Duplicate message meta register by fullname: %s", meta.FullName()))
	} else {
		metaByFullName[meta.FullName()] = meta
	}

	if meta.ID == 0 {
		panic("message meta require 'ID' field: " + meta.TypeName())
	}

	if prev, ok := metaByID[meta.ID]; ok {
		panic(fmt.Sprintf("Duplicate message meta register by id: %d type: %s, pre type: %s", meta.ID, meta.TypeName(), prev.TypeName()))
	} else {
		metaByID[meta.ID] = meta
	}
	return meta
}

// 根据类型查找消息元信息
func MessageMetaByType(t reflect.Type) *MessageMeta {

	if t == nil {
		return nil
	}

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	if v, ok := metaByType[t]; ok {
		return v
	}

	return nil
}

// 根据消息对象获得消息元信息
func MessageMetaByMsg(msg interface{}) *MessageMeta {

	if msg == nil {
		return nil
	}

	return MessageMetaByType(reflect.TypeOf(msg))
}

// 根据id查找消息元信息
func MessageMetaByID(id uint32) *MessageMeta {
	if v, ok := metaByID[id]; ok {
		return v
	}

	return nil
}
