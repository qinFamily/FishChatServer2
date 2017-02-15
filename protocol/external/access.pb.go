// Code generated by protoc-gen-go.
// source: access.proto
// DO NOT EDIT!

/*
Package external is a generated protocol buffer package.

It is generated from these files:
	access.proto
	base.proto
	error.proto
	gateway.proto

It has these top-level messages:
	ReqLogin
	ResLogin
	ReqLogout
	ResLogout
	ReqPing
	ReqSendP2PMsg
	ResSendP2PMsg
	ResNotify
	ReqSyncMsg
	OffsetMsg
	ResSyncMsg
	ReqAcceptP2PMsgAck
	ResAcceptP2PMsgAck
	ReqSendGroupMsg
	ResSendGroupMsg
	Base
	Error
	ReqAccessServer
	ResSelectAccessServerForClient
*/
package external

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// login
type ReqLogin struct {
	Cmd    uint32 `protobuf:"varint,1,opt,name=cmd" json:"cmd,omitempty"`
	CmdStr string `protobuf:"bytes,2,opt,name=cmdStr" json:"cmdStr,omitempty"`
	UID    int64  `protobuf:"varint,3,opt,name=UID" json:"UID,omitempty"`
	Token  string `protobuf:"bytes,4,opt,name=token" json:"token,omitempty"`
}

func (m *ReqLogin) Reset()                    { *m = ReqLogin{} }
func (m *ReqLogin) String() string            { return proto.CompactTextString(m) }
func (*ReqLogin) ProtoMessage()               {}
func (*ReqLogin) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ReqLogin) GetCmd() uint32 {
	if m != nil {
		return m.Cmd
	}
	return 0
}

func (m *ReqLogin) GetCmdStr() string {
	if m != nil {
		return m.CmdStr
	}
	return ""
}

func (m *ReqLogin) GetUID() int64 {
	if m != nil {
		return m.UID
	}
	return 0
}

func (m *ReqLogin) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type ResLogin struct {
	Cmd     uint32 `protobuf:"varint,1,opt,name=cmd" json:"cmd,omitempty"`
	CmdStr  string `protobuf:"bytes,2,opt,name=cmdStr" json:"cmdStr,omitempty"`
	ErrCode uint32 `protobuf:"varint,3,opt,name=errCode" json:"errCode,omitempty"`
	ErrStr  string `protobuf:"bytes,4,opt,name=errStr" json:"errStr,omitempty"`
}

func (m *ResLogin) Reset()                    { *m = ResLogin{} }
func (m *ResLogin) String() string            { return proto.CompactTextString(m) }
func (*ResLogin) ProtoMessage()               {}
func (*ResLogin) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ResLogin) GetCmd() uint32 {
	if m != nil {
		return m.Cmd
	}
	return 0
}

func (m *ResLogin) GetCmdStr() string {
	if m != nil {
		return m.CmdStr
	}
	return ""
}

func (m *ResLogin) GetErrCode() uint32 {
	if m != nil {
		return m.ErrCode
	}
	return 0
}

func (m *ResLogin) GetErrStr() string {
	if m != nil {
		return m.ErrStr
	}
	return ""
}

// logout
type ReqLogout struct {
	Cmd    uint32 `protobuf:"varint,1,opt,name=cmd" json:"cmd,omitempty"`
	CmdStr string `protobuf:"bytes,2,opt,name=cmdStr" json:"cmdStr,omitempty"`
	UID    int64  `protobuf:"varint,3,opt,name=UID" json:"UID,omitempty"`
}

func (m *ReqLogout) Reset()                    { *m = ReqLogout{} }
func (m *ReqLogout) String() string            { return proto.CompactTextString(m) }
func (*ReqLogout) ProtoMessage()               {}
func (*ReqLogout) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ReqLogout) GetCmd() uint32 {
	if m != nil {
		return m.Cmd
	}
	return 0
}

func (m *ReqLogout) GetCmdStr() string {
	if m != nil {
		return m.CmdStr
	}
	return ""
}

func (m *ReqLogout) GetUID() int64 {
	if m != nil {
		return m.UID
	}
	return 0
}

type ResLogout struct {
	Cmd     uint32 `protobuf:"varint,1,opt,name=cmd" json:"cmd,omitempty"`
	CmdStr  string `protobuf:"bytes,2,opt,name=cmdStr" json:"cmdStr,omitempty"`
	ErrCode uint32 `protobuf:"varint,3,opt,name=errCode" json:"errCode,omitempty"`
	ErrStr  string `protobuf:"bytes,4,opt,name=errStr" json:"errStr,omitempty"`
}

func (m *ResLogout) Reset()                    { *m = ResLogout{} }
func (m *ResLogout) String() string            { return proto.CompactTextString(m) }
func (*ResLogout) ProtoMessage()               {}
func (*ResLogout) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ResLogout) GetCmd() uint32 {
	if m != nil {
		return m.Cmd
	}
	return 0
}

func (m *ResLogout) GetCmdStr() string {
	if m != nil {
		return m.CmdStr
	}
	return ""
}

func (m *ResLogout) GetErrCode() uint32 {
	if m != nil {
		return m.ErrCode
	}
	return 0
}

func (m *ResLogout) GetErrStr() string {
	if m != nil {
		return m.ErrStr
	}
	return ""
}

// ping
type ReqPing struct {
	Cmd    uint32 `protobuf:"varint,1,opt,name=cmd" json:"cmd,omitempty"`
	CmdStr string `protobuf:"bytes,2,opt,name=cmdStr" json:"cmdStr,omitempty"`
	UID    int64  `protobuf:"varint,3,opt,name=UID" json:"UID,omitempty"`
}

func (m *ReqPing) Reset()                    { *m = ReqPing{} }
func (m *ReqPing) String() string            { return proto.CompactTextString(m) }
func (*ReqPing) ProtoMessage()               {}
func (*ReqPing) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *ReqPing) GetCmd() uint32 {
	if m != nil {
		return m.Cmd
	}
	return 0
}

func (m *ReqPing) GetCmdStr() string {
	if m != nil {
		return m.CmdStr
	}
	return ""
}

func (m *ReqPing) GetUID() int64 {
	if m != nil {
		return m.UID
	}
	return 0
}

// p2p msg
type ReqSendP2PMsg struct {
	Cmd       uint32 `protobuf:"varint,1,opt,name=cmd" json:"cmd,omitempty"`
	CmdStr    string `protobuf:"bytes,2,opt,name=cmdStr" json:"cmdStr,omitempty"`
	SourceUID int64  `protobuf:"varint,3,opt,name=sourceUID" json:"sourceUID,omitempty"`
	TargetUID int64  `protobuf:"varint,4,opt,name=targetUID" json:"targetUID,omitempty"`
	MsgID     string `protobuf:"bytes,5,opt,name=msgID" json:"msgID,omitempty"`
	Msg       string `protobuf:"bytes,6,opt,name=msg" json:"msg,omitempty"`
}

func (m *ReqSendP2PMsg) Reset()                    { *m = ReqSendP2PMsg{} }
func (m *ReqSendP2PMsg) String() string            { return proto.CompactTextString(m) }
func (*ReqSendP2PMsg) ProtoMessage()               {}
func (*ReqSendP2PMsg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *ReqSendP2PMsg) GetCmd() uint32 {
	if m != nil {
		return m.Cmd
	}
	return 0
}

func (m *ReqSendP2PMsg) GetCmdStr() string {
	if m != nil {
		return m.CmdStr
	}
	return ""
}

func (m *ReqSendP2PMsg) GetSourceUID() int64 {
	if m != nil {
		return m.SourceUID
	}
	return 0
}

func (m *ReqSendP2PMsg) GetTargetUID() int64 {
	if m != nil {
		return m.TargetUID
	}
	return 0
}

func (m *ReqSendP2PMsg) GetMsgID() string {
	if m != nil {
		return m.MsgID
	}
	return ""
}

func (m *ReqSendP2PMsg) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type ResSendP2PMsg struct {
	Cmd       uint32 `protobuf:"varint,1,opt,name=cmd" json:"cmd,omitempty"`
	CmdStr    string `protobuf:"bytes,2,opt,name=cmdStr" json:"cmdStr,omitempty"`
	ErrCode   uint32 `protobuf:"varint,3,opt,name=errCode" json:"errCode,omitempty"`
	ErrStr    string `protobuf:"bytes,4,opt,name=errStr" json:"errStr,omitempty"`
	SourceUID int64  `protobuf:"varint,5,opt,name=sourceUID" json:"sourceUID,omitempty"`
	TargetUID int64  `protobuf:"varint,6,opt,name=targetUID" json:"targetUID,omitempty"`
	MsgID     string `protobuf:"bytes,7,opt,name=msgID" json:"msgID,omitempty"`
	Msg       string `protobuf:"bytes,8,opt,name=msg" json:"msg,omitempty"`
}

func (m *ResSendP2PMsg) Reset()                    { *m = ResSendP2PMsg{} }
func (m *ResSendP2PMsg) String() string            { return proto.CompactTextString(m) }
func (*ResSendP2PMsg) ProtoMessage()               {}
func (*ResSendP2PMsg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *ResSendP2PMsg) GetCmd() uint32 {
	if m != nil {
		return m.Cmd
	}
	return 0
}

func (m *ResSendP2PMsg) GetCmdStr() string {
	if m != nil {
		return m.CmdStr
	}
	return ""
}

func (m *ResSendP2PMsg) GetErrCode() uint32 {
	if m != nil {
		return m.ErrCode
	}
	return 0
}

func (m *ResSendP2PMsg) GetErrStr() string {
	if m != nil {
		return m.ErrStr
	}
	return ""
}

func (m *ResSendP2PMsg) GetSourceUID() int64 {
	if m != nil {
		return m.SourceUID
	}
	return 0
}

func (m *ResSendP2PMsg) GetTargetUID() int64 {
	if m != nil {
		return m.TargetUID
	}
	return 0
}

func (m *ResSendP2PMsg) GetMsgID() string {
	if m != nil {
		return m.MsgID
	}
	return ""
}

func (m *ResSendP2PMsg) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

// notify
type ResNotify struct {
	Cmd       uint32 `protobuf:"varint,1,opt,name=cmd" json:"cmd,omitempty"`
	CmdStr    string `protobuf:"bytes,2,opt,name=cmdStr" json:"cmdStr,omitempty"`
	ErrCode   uint32 `protobuf:"varint,3,opt,name=errCode" json:"errCode,omitempty"`
	ErrStr    string `protobuf:"bytes,4,opt,name=errStr" json:"errStr,omitempty"`
	CurrentID int64  `protobuf:"varint,5,opt,name=currentID" json:"currentID,omitempty"`
}

func (m *ResNotify) Reset()                    { *m = ResNotify{} }
func (m *ResNotify) String() string            { return proto.CompactTextString(m) }
func (*ResNotify) ProtoMessage()               {}
func (*ResNotify) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *ResNotify) GetCmd() uint32 {
	if m != nil {
		return m.Cmd
	}
	return 0
}

func (m *ResNotify) GetCmdStr() string {
	if m != nil {
		return m.CmdStr
	}
	return ""
}

func (m *ResNotify) GetErrCode() uint32 {
	if m != nil {
		return m.ErrCode
	}
	return 0
}

func (m *ResNotify) GetErrStr() string {
	if m != nil {
		return m.ErrStr
	}
	return ""
}

func (m *ResNotify) GetCurrentID() int64 {
	if m != nil {
		return m.CurrentID
	}
	return 0
}

// sync msg
type ReqSyncMsg struct {
	Cmd       uint32 `protobuf:"varint,1,opt,name=cmd" json:"cmd,omitempty"`
	CmdStr    string `protobuf:"bytes,2,opt,name=cmdStr" json:"cmdStr,omitempty"`
	UID       int64  `protobuf:"varint,3,opt,name=UID" json:"UID,omitempty"`
	CurrentID int64  `protobuf:"varint,4,opt,name=currentID" json:"currentID,omitempty"`
}

func (m *ReqSyncMsg) Reset()                    { *m = ReqSyncMsg{} }
func (m *ReqSyncMsg) String() string            { return proto.CompactTextString(m) }
func (*ReqSyncMsg) ProtoMessage()               {}
func (*ReqSyncMsg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *ReqSyncMsg) GetCmd() uint32 {
	if m != nil {
		return m.Cmd
	}
	return 0
}

func (m *ReqSyncMsg) GetCmdStr() string {
	if m != nil {
		return m.CmdStr
	}
	return ""
}

func (m *ReqSyncMsg) GetUID() int64 {
	if m != nil {
		return m.UID
	}
	return 0
}

func (m *ReqSyncMsg) GetCurrentID() int64 {
	if m != nil {
		return m.CurrentID
	}
	return 0
}

type OffsetMsg struct {
	SourceUID int64  `protobuf:"varint,1,opt,name=sourceUID" json:"sourceUID,omitempty"`
	TargetUID int64  `protobuf:"varint,2,opt,name=targetUID" json:"targetUID,omitempty"`
	GroupID   int64  `protobuf:"varint,3,opt,name=groupID" json:"groupID,omitempty"`
	MsgType   string `protobuf:"bytes,4,opt,name=msgType" json:"msgType,omitempty"`
	MsgID     string `protobuf:"bytes,5,opt,name=msgID" json:"msgID,omitempty"`
	Msg       string `protobuf:"bytes,6,opt,name=msg" json:"msg,omitempty"`
}

func (m *OffsetMsg) Reset()                    { *m = OffsetMsg{} }
func (m *OffsetMsg) String() string            { return proto.CompactTextString(m) }
func (*OffsetMsg) ProtoMessage()               {}
func (*OffsetMsg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *OffsetMsg) GetSourceUID() int64 {
	if m != nil {
		return m.SourceUID
	}
	return 0
}

func (m *OffsetMsg) GetTargetUID() int64 {
	if m != nil {
		return m.TargetUID
	}
	return 0
}

func (m *OffsetMsg) GetGroupID() int64 {
	if m != nil {
		return m.GroupID
	}
	return 0
}

func (m *OffsetMsg) GetMsgType() string {
	if m != nil {
		return m.MsgType
	}
	return ""
}

func (m *OffsetMsg) GetMsgID() string {
	if m != nil {
		return m.MsgID
	}
	return ""
}

func (m *OffsetMsg) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type ResSyncMsg struct {
	Cmd     uint32       `protobuf:"varint,1,opt,name=cmd" json:"cmd,omitempty"`
	CmdStr  string       `protobuf:"bytes,2,opt,name=cmdStr" json:"cmdStr,omitempty"`
	ErrCode uint32       `protobuf:"varint,3,opt,name=errCode" json:"errCode,omitempty"`
	ErrStr  string       `protobuf:"bytes,4,opt,name=errStr" json:"errStr,omitempty"`
	Msgs    []*OffsetMsg `protobuf:"bytes,5,rep,name=Msgs" json:"Msgs,omitempty"`
}

func (m *ResSyncMsg) Reset()                    { *m = ResSyncMsg{} }
func (m *ResSyncMsg) String() string            { return proto.CompactTextString(m) }
func (*ResSyncMsg) ProtoMessage()               {}
func (*ResSyncMsg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *ResSyncMsg) GetCmd() uint32 {
	if m != nil {
		return m.Cmd
	}
	return 0
}

func (m *ResSyncMsg) GetCmdStr() string {
	if m != nil {
		return m.CmdStr
	}
	return ""
}

func (m *ResSyncMsg) GetErrCode() uint32 {
	if m != nil {
		return m.ErrCode
	}
	return 0
}

func (m *ResSyncMsg) GetErrStr() string {
	if m != nil {
		return m.ErrStr
	}
	return ""
}

func (m *ResSyncMsg) GetMsgs() []*OffsetMsg {
	if m != nil {
		return m.Msgs
	}
	return nil
}

// p2p msg accept ack
type ReqAcceptP2PMsgAck struct {
	Cmd       uint32 `protobuf:"varint,1,opt,name=cmd" json:"cmd,omitempty"`
	CmdStr    string `protobuf:"bytes,2,opt,name=cmdStr" json:"cmdStr,omitempty"`
	SourceUID int64  `protobuf:"varint,3,opt,name=sourceUID" json:"sourceUID,omitempty"`
	TargetUID int64  `protobuf:"varint,4,opt,name=targetUID" json:"targetUID,omitempty"`
	MsgID     string `protobuf:"bytes,5,opt,name=msgID" json:"msgID,omitempty"`
}

func (m *ReqAcceptP2PMsgAck) Reset()                    { *m = ReqAcceptP2PMsgAck{} }
func (m *ReqAcceptP2PMsgAck) String() string            { return proto.CompactTextString(m) }
func (*ReqAcceptP2PMsgAck) ProtoMessage()               {}
func (*ReqAcceptP2PMsgAck) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *ReqAcceptP2PMsgAck) GetCmd() uint32 {
	if m != nil {
		return m.Cmd
	}
	return 0
}

func (m *ReqAcceptP2PMsgAck) GetCmdStr() string {
	if m != nil {
		return m.CmdStr
	}
	return ""
}

func (m *ReqAcceptP2PMsgAck) GetSourceUID() int64 {
	if m != nil {
		return m.SourceUID
	}
	return 0
}

func (m *ReqAcceptP2PMsgAck) GetTargetUID() int64 {
	if m != nil {
		return m.TargetUID
	}
	return 0
}

func (m *ReqAcceptP2PMsgAck) GetMsgID() string {
	if m != nil {
		return m.MsgID
	}
	return ""
}

type ResAcceptP2PMsgAck struct {
	Cmd     uint32 `protobuf:"varint,1,opt,name=cmd" json:"cmd,omitempty"`
	CmdStr  string `protobuf:"bytes,2,opt,name=cmdStr" json:"cmdStr,omitempty"`
	ErrCode uint32 `protobuf:"varint,3,opt,name=errCode" json:"errCode,omitempty"`
	ErrStr  string `protobuf:"bytes,4,opt,name=errStr" json:"errStr,omitempty"`
}

func (m *ResAcceptP2PMsgAck) Reset()                    { *m = ResAcceptP2PMsgAck{} }
func (m *ResAcceptP2PMsgAck) String() string            { return proto.CompactTextString(m) }
func (*ResAcceptP2PMsgAck) ProtoMessage()               {}
func (*ResAcceptP2PMsgAck) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *ResAcceptP2PMsgAck) GetCmd() uint32 {
	if m != nil {
		return m.Cmd
	}
	return 0
}

func (m *ResAcceptP2PMsgAck) GetCmdStr() string {
	if m != nil {
		return m.CmdStr
	}
	return ""
}

func (m *ResAcceptP2PMsgAck) GetErrCode() uint32 {
	if m != nil {
		return m.ErrCode
	}
	return 0
}

func (m *ResAcceptP2PMsgAck) GetErrStr() string {
	if m != nil {
		return m.ErrStr
	}
	return ""
}

// group
type ReqSendGroupMsg struct {
	Cmd       uint32 `protobuf:"varint,1,opt,name=cmd" json:"cmd,omitempty"`
	CmdStr    string `protobuf:"bytes,2,opt,name=cmdStr" json:"cmdStr,omitempty"`
	SourceUID int64  `protobuf:"varint,3,opt,name=sourceUID" json:"sourceUID,omitempty"`
	GroupID   int64  `protobuf:"varint,4,opt,name=groupID" json:"groupID,omitempty"`
	MsgID     string `protobuf:"bytes,5,opt,name=msgID" json:"msgID,omitempty"`
	Msg       string `protobuf:"bytes,6,opt,name=msg" json:"msg,omitempty"`
}

func (m *ReqSendGroupMsg) Reset()                    { *m = ReqSendGroupMsg{} }
func (m *ReqSendGroupMsg) String() string            { return proto.CompactTextString(m) }
func (*ReqSendGroupMsg) ProtoMessage()               {}
func (*ReqSendGroupMsg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *ReqSendGroupMsg) GetCmd() uint32 {
	if m != nil {
		return m.Cmd
	}
	return 0
}

func (m *ReqSendGroupMsg) GetCmdStr() string {
	if m != nil {
		return m.CmdStr
	}
	return ""
}

func (m *ReqSendGroupMsg) GetSourceUID() int64 {
	if m != nil {
		return m.SourceUID
	}
	return 0
}

func (m *ReqSendGroupMsg) GetGroupID() int64 {
	if m != nil {
		return m.GroupID
	}
	return 0
}

func (m *ReqSendGroupMsg) GetMsgID() string {
	if m != nil {
		return m.MsgID
	}
	return ""
}

func (m *ReqSendGroupMsg) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type ResSendGroupMsg struct {
	Cmd     uint32 `protobuf:"varint,1,opt,name=cmd" json:"cmd,omitempty"`
	CmdStr  string `protobuf:"bytes,2,opt,name=cmdStr" json:"cmdStr,omitempty"`
	ErrCode uint32 `protobuf:"varint,3,opt,name=errCode" json:"errCode,omitempty"`
	ErrStr  string `protobuf:"bytes,4,opt,name=errStr" json:"errStr,omitempty"`
}

func (m *ResSendGroupMsg) Reset()                    { *m = ResSendGroupMsg{} }
func (m *ResSendGroupMsg) String() string            { return proto.CompactTextString(m) }
func (*ResSendGroupMsg) ProtoMessage()               {}
func (*ResSendGroupMsg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

func (m *ResSendGroupMsg) GetCmd() uint32 {
	if m != nil {
		return m.Cmd
	}
	return 0
}

func (m *ResSendGroupMsg) GetCmdStr() string {
	if m != nil {
		return m.CmdStr
	}
	return ""
}

func (m *ResSendGroupMsg) GetErrCode() uint32 {
	if m != nil {
		return m.ErrCode
	}
	return 0
}

func (m *ResSendGroupMsg) GetErrStr() string {
	if m != nil {
		return m.ErrStr
	}
	return ""
}

func init() {
	proto.RegisterType((*ReqLogin)(nil), "external.ReqLogin")
	proto.RegisterType((*ResLogin)(nil), "external.ResLogin")
	proto.RegisterType((*ReqLogout)(nil), "external.ReqLogout")
	proto.RegisterType((*ResLogout)(nil), "external.ResLogout")
	proto.RegisterType((*ReqPing)(nil), "external.ReqPing")
	proto.RegisterType((*ReqSendP2PMsg)(nil), "external.ReqSendP2PMsg")
	proto.RegisterType((*ResSendP2PMsg)(nil), "external.ResSendP2PMsg")
	proto.RegisterType((*ResNotify)(nil), "external.ResNotify")
	proto.RegisterType((*ReqSyncMsg)(nil), "external.ReqSyncMsg")
	proto.RegisterType((*OffsetMsg)(nil), "external.offsetMsg")
	proto.RegisterType((*ResSyncMsg)(nil), "external.ResSyncMsg")
	proto.RegisterType((*ReqAcceptP2PMsgAck)(nil), "external.ReqAcceptP2PMsgAck")
	proto.RegisterType((*ResAcceptP2PMsgAck)(nil), "external.ResAcceptP2PMsgAck")
	proto.RegisterType((*ReqSendGroupMsg)(nil), "external.ReqSendGroupMsg")
	proto.RegisterType((*ResSendGroupMsg)(nil), "external.ResSendGroupMsg")
}

func init() { proto.RegisterFile("access.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 468 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xc4, 0x55, 0xcd, 0x6e, 0x13, 0x31,
	0x10, 0x96, 0x9b, 0xcd, 0xdf, 0x40, 0x04, 0x5a, 0x10, 0xf2, 0x81, 0x43, 0xb4, 0x17, 0x72, 0xca,
	0xa1, 0x3c, 0x41, 0x45, 0x51, 0x55, 0x89, 0xa2, 0xc8, 0x85, 0x1b, 0x97, 0xe0, 0xcc, 0x5a, 0x51,
	0xd9, 0xf5, 0xc6, 0xe3, 0x48, 0xe4, 0x01, 0xb8, 0xc3, 0x19, 0x71, 0xe2, 0x81, 0x78, 0x25, 0x34,
	0xde, 0xfc, 0xb0, 0x05, 0xaa, 0x3a, 0x01, 0xf5, 0xe6, 0x99, 0xf1, 0xce, 0x7c, 0xdf, 0x37, 0xe3,
	0x59, 0xb8, 0x3f, 0xd5, 0x1a, 0x89, 0xc6, 0x95, 0xb3, 0xde, 0xa6, 0x3d, 0xfc, 0xe8, 0xd1, 0x95,
	0xd3, 0x0f, 0xd9, 0x3b, 0xe8, 0x29, 0x5c, 0xbc, 0xb2, 0x66, 0x5e, 0xa6, 0x0f, 0xa1, 0xa5, 0x8b,
	0x99, 0x14, 0x43, 0x31, 0x1a, 0x28, 0x3e, 0xa6, 0x4f, 0xa0, 0xa3, 0x8b, 0xd9, 0xa5, 0x77, 0xf2,
	0x68, 0x28, 0x46, 0x7d, 0xb5, 0xb6, 0xf8, 0xe6, 0xdb, 0xf3, 0x53, 0xd9, 0x1a, 0x8a, 0x51, 0x4b,
	0xf1, 0x31, 0x7d, 0x0c, 0x6d, 0x6f, 0xaf, 0xb0, 0x94, 0x49, 0xb8, 0x58, 0x1b, 0x59, 0xce, 0xd9,
	0x29, 0x36, 0xbb, 0x84, 0x2e, 0x3a, 0xf7, 0xc2, 0xce, 0x30, 0x54, 0x18, 0xa8, 0x8d, 0xc9, 0x5f,
	0xa0, 0x73, 0xfc, 0x45, 0x5d, 0x66, 0x6d, 0x65, 0x67, 0xd0, 0xaf, 0x59, 0xd8, 0xa5, 0x3f, 0x84,
	0x46, 0x66, 0x38, 0x11, 0x45, 0x27, 0x8a, 0x47, 0xfc, 0x12, 0xba, 0x0a, 0x17, 0x93, 0x79, 0x69,
	0x0e, 0xc2, 0xfb, 0x4d, 0xc0, 0x40, 0xe1, 0xe2, 0x12, 0xcb, 0xd9, 0xe4, 0x78, 0x72, 0x41, 0x31,
	0xd9, 0x9e, 0x42, 0x9f, 0xec, 0xd2, 0x69, 0xdc, 0xe5, 0xdc, 0x39, 0x38, 0xea, 0xa7, 0xce, 0xa0,
	0xe7, 0x68, 0x52, 0x47, 0xb7, 0x0e, 0x6e, 0x77, 0x41, 0xe6, 0xfc, 0x54, 0xb6, 0xeb, 0x76, 0x07,
	0x83, 0x6b, 0x17, 0x64, 0x64, 0x27, 0xf8, 0xf8, 0x98, 0xfd, 0x08, 0xf8, 0x68, 0x2f, 0x7c, 0xd1,
	0xa2, 0x36, 0x19, 0xb5, 0x6f, 0x64, 0xd4, 0xf9, 0x2b, 0xa3, 0xee, 0x1f, 0x18, 0xf5, 0x76, 0x8c,
	0x3e, 0x89, 0x30, 0x22, 0xaf, 0xad, 0x9f, 0xe7, 0xab, 0xff, 0xcd, 0x46, 0x2f, 0x9d, 0xc3, 0xd2,
	0xef, 0xd8, 0x6c, 0x1d, 0x59, 0x0e, 0xc0, 0x8d, 0x5f, 0x95, 0x3a, 0x4e, 0xd5, 0xdf, 0x9f, 0x6e,
	0xa3, 0x4e, 0x72, 0xbd, 0xce, 0x77, 0x01, 0x7d, 0x9b, 0xe7, 0x84, 0x9e, 0xeb, 0x34, 0x14, 0x16,
	0x37, 0x2a, 0x7c, 0x74, 0x5d, 0x61, 0x09, 0x5d, 0xe3, 0xec, 0xb2, 0xda, 0x56, 0xdf, 0x98, 0x1c,
	0x29, 0xc8, 0xbc, 0x59, 0x55, 0xb8, 0x96, 0x60, 0x63, 0xde, 0x7a, 0xce, 0xbe, 0x08, 0x96, 0x83,
	0xe2, 0xe5, 0x88, 0x6f, 0xcb, 0x33, 0x48, 0x2e, 0xc8, 0x90, 0x6c, 0x0f, 0x5b, 0xa3, 0x7b, 0xc7,
	0x8f, 0xc6, 0x9b, 0x55, 0x3a, 0xde, 0xaa, 0xa4, 0xc2, 0x85, 0xec, 0xb3, 0x80, 0x54, 0xe1, 0xe2,
	0x44, 0x6b, 0xac, 0x7c, 0x3d, 0xfd, 0x27, 0xfa, 0xea, 0x2e, 0x1f, 0x68, 0x56, 0x31, 0x22, 0xda,
	0x1f, 0x51, 0xfc, 0x9e, 0xfb, 0x2a, 0xe0, 0xc1, 0x7a, 0x41, 0x9d, 0x71, 0xb7, 0xff, 0xe5, 0x8a,
	0xfa, 0x65, 0xa0, 0x92, 0xe6, 0x40, 0xdd, 0x76, 0x6c, 0x0a, 0x06, 0x47, 0x7b, 0x82, 0x8b, 0x16,
	0xe3, 0x7d, 0x27, 0xfc, 0x7d, 0x9f, 0xff, 0x0c, 0x00, 0x00, 0xff, 0xff, 0xd2, 0xad, 0x13, 0x4c,
	0x8d, 0x07, 0x00, 0x00,
}
