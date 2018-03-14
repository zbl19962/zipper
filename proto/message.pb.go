// Copyright (C) 2017, Zipper Team.  All rights reserved.
//
// This file is part of zipper
//
// The zipper is free software: you can use, copy, modify,
// and distribute this software for any purpose with or
// without fee is hereby granted, provided that the above
// copyright notice and this permission notice appear in all copies.
//
// The zipper is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// ISC License for more details.
//
// You should have received a copy of the ISC License
// along with this program.  If not, see <https://opensource.org/licenses/isc>.


// Code generated by protoc-gen-go. DO NOT EDIT.
// source: message.proto

package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type MsgType int32

const (
	MsgType_BC_GetBlocksMsg     MsgType = 0
	MsgType_BC_GetInvMsg        MsgType = 1
	MsgType_BC_GetDataMsg       MsgType = 2
	MsgType_BC_OnBlockMsg       MsgType = 3
	MsgType_BC_OnTransactionMsg MsgType = 4
	MsgType_BC_OnStatusMSg      MsgType = 5
)

var MsgType_name = map[int32]string{
	0: "BC_GetBlocksMsg",
	1: "BC_GetInvMsg",
	2: "BC_GetDataMsg",
	3: "BC_OnBlockMsg",
	4: "BC_OnTransactionMsg",
	5: "BC_OnStatusMSg",
}
var MsgType_value = map[string]int32{
	"BC_GetBlocksMsg":     0,
	"BC_GetInvMsg":        1,
	"BC_GetDataMsg":       2,
	"BC_OnBlockMsg":       3,
	"BC_OnTransactionMsg": 4,
	"BC_OnStatusMSg":      5,
}

func (x MsgType) String() string {
	return proto1.EnumName(MsgType_name, int32(x))
}
func (MsgType) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

type InvType int32

const (
	InvType_block       InvType = 0
	InvType_transaction InvType = 1
)

var InvType_name = map[int32]string{
	0: "block",
	1: "transaction",
}
var InvType_value = map[string]int32{
	"block":       0,
	"transaction": 1,
}

func (x InvType) String() string {
	return proto1.EnumName(InvType_name, int32(x))
}
func (InvType) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

type GetBlocksMsg struct {
	Version       uint32   `protobuf:"varint,1,opt,name=version" json:"version,omitempty"`
	LocatorHashes []string `protobuf:"bytes,2,rep,name=locatorHashes" json:"locatorHashes,omitempty"`
	HashStop      string   `protobuf:"bytes,3,opt,name=hashStop" json:"hashStop,omitempty"`
}

func (m *GetBlocksMsg) Reset()                    { *m = GetBlocksMsg{} }
func (m *GetBlocksMsg) String() string            { return proto1.CompactTextString(m) }
func (*GetBlocksMsg) ProtoMessage()               {}
func (*GetBlocksMsg) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *GetBlocksMsg) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *GetBlocksMsg) GetLocatorHashes() []string {
	if m != nil {
		return m.LocatorHashes
	}
	return nil
}

func (m *GetBlocksMsg) GetHashStop() string {
	if m != nil {
		return m.HashStop
	}
	return ""
}

type StatusMsg struct {
	Version     uint32 `protobuf:"varint,1,opt,name=Version" json:"Version,omitempty"`
	StartHeight uint32 `protobuf:"varint,2,opt,name=StartHeight" json:"StartHeight,omitempty"`
}

func (m *StatusMsg) Reset()                    { *m = StatusMsg{} }
func (m *StatusMsg) String() string            { return proto1.CompactTextString(m) }
func (*StatusMsg) ProtoMessage()               {}
func (*StatusMsg) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *StatusMsg) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *StatusMsg) GetStartHeight() uint32 {
	if m != nil {
		return m.StartHeight
	}
	return 0
}

type GetInvMsg struct {
	Type  InvType  `protobuf:"varint,1,opt,name=type,enum=proto.InvType" json:"type,omitempty"`
	Hashs []string `protobuf:"bytes,2,rep,name=hashs" json:"hashs,omitempty"`
}

func (m *GetInvMsg) Reset()                    { *m = GetInvMsg{} }
func (m *GetInvMsg) String() string            { return proto1.CompactTextString(m) }
func (*GetInvMsg) ProtoMessage()               {}
func (*GetInvMsg) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *GetInvMsg) GetType() InvType {
	if m != nil {
		return m.Type
	}
	return InvType_block
}

func (m *GetInvMsg) GetHashs() []string {
	if m != nil {
		return m.Hashs
	}
	return nil
}

type OnBlockMsg struct {
	Block *Block `protobuf:"bytes,1,opt,name=block" json:"block,omitempty"`
}

func (m *OnBlockMsg) Reset()                    { *m = OnBlockMsg{} }
func (m *OnBlockMsg) String() string            { return proto1.CompactTextString(m) }
func (*OnBlockMsg) ProtoMessage()               {}
func (*OnBlockMsg) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func (m *OnBlockMsg) GetBlock() *Block {
	if m != nil {
		return m.Block
	}
	return nil
}

type OnTransactionMsg struct {
	Transaction *Transaction `protobuf:"bytes,1,opt,name=transaction" json:"transaction,omitempty"`
}

func (m *OnTransactionMsg) Reset()                    { *m = OnTransactionMsg{} }
func (m *OnTransactionMsg) String() string            { return proto1.CompactTextString(m) }
func (*OnTransactionMsg) ProtoMessage()               {}
func (*OnTransactionMsg) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

func (m *OnTransactionMsg) GetTransaction() *Transaction {
	if m != nil {
		return m.Transaction
	}
	return nil
}

type GetDataMsg struct {
	InvList []*GetInvMsg `protobuf:"bytes,1,rep,name=InvList" json:"InvList,omitempty"`
}

func (m *GetDataMsg) Reset()                    { *m = GetDataMsg{} }
func (m *GetDataMsg) String() string            { return proto1.CompactTextString(m) }
func (*GetDataMsg) ProtoMessage()               {}
func (*GetDataMsg) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{5} }

func (m *GetDataMsg) GetInvList() []*GetInvMsg {
	if m != nil {
		return m.InvList
	}
	return nil
}

func init() {
	proto1.RegisterType((*GetBlocksMsg)(nil), "proto.GetBlocksMsg")
	proto1.RegisterType((*StatusMsg)(nil), "proto.StatusMsg")
	proto1.RegisterType((*GetInvMsg)(nil), "proto.GetInvMsg")
	proto1.RegisterType((*OnBlockMsg)(nil), "proto.OnBlockMsg")
	proto1.RegisterType((*OnTransactionMsg)(nil), "proto.OnTransactionMsg")
	proto1.RegisterType((*GetDataMsg)(nil), "proto.GetDataMsg")
	proto1.RegisterEnum("proto.MsgType", MsgType_name, MsgType_value)
	proto1.RegisterEnum("proto.InvType", InvType_name, InvType_value)
}

func init() { proto1.RegisterFile("message.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 384 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x51, 0xc1, 0x8a, 0xdb, 0x30,
	0x10, 0x5d, 0xc7, 0xeb, 0xa6, 0x1e, 0xc7, 0x59, 0xad, 0xb6, 0x50, 0x93, 0x93, 0x31, 0x2d, 0x98,
	0x1c, 0x96, 0x92, 0xf6, 0xd0, 0xf3, 0xb6, 0xc5, 0x09, 0x74, 0x59, 0xb0, 0x43, 0xaf, 0x45, 0x09,
	0xc2, 0x76, 0x9b, 0x4a, 0xc6, 0x9a, 0x1a, 0x72, 0xee, 0x8f, 0x17, 0x4b, 0x72, 0x62, 0x9f, 0xc4,
	0xbc, 0x99, 0xf7, 0x66, 0xde, 0x13, 0x84, 0x7f, 0xb8, 0x52, 0xac, 0xe4, 0x8f, 0x4d, 0x2b, 0x51,
	0x52, 0x4f, 0x3f, 0xab, 0xe0, 0x70, 0x92, 0xc7, 0xdf, 0x06, 0x5b, 0xdd, 0x63, 0xcb, 0x84, 0x62,
	0x47, 0xac, 0xa5, 0x30, 0x50, 0xf2, 0x0b, 0x16, 0x19, 0xc7, 0xa7, 0x7e, 0x48, 0x3d, 0xab, 0x92,
	0x46, 0x30, 0xef, 0x78, 0xab, 0x6a, 0x29, 0x22, 0x27, 0x76, 0xd2, 0x30, 0x1f, 0x4a, 0xfa, 0x0e,
	0xc2, 0x93, 0x3c, 0x32, 0x94, 0xed, 0x96, 0xa9, 0x8a, 0xab, 0x68, 0x16, 0xbb, 0xa9, 0x9f, 0x4f,
	0x41, 0xba, 0x82, 0xd7, 0x15, 0x53, 0x55, 0x81, 0xb2, 0x89, 0xdc, 0xd8, 0x49, 0xfd, 0xfc, 0x52,
	0x27, 0x19, 0xf8, 0x05, 0x32, 0xfc, 0x3b, 0x2c, 0xfa, 0x31, 0x5d, 0x64, 0x4b, 0x1a, 0x43, 0x50,
	0x20, 0x6b, 0x71, 0xcb, 0xeb, 0xb2, 0xc2, 0x68, 0xa6, 0xbb, 0x63, 0x28, 0xf9, 0x06, 0x7e, 0xc6,
	0x71, 0x27, 0xba, 0x5e, 0x28, 0x81, 0x5b, 0x3c, 0x37, 0x5c, 0xab, 0x2c, 0x37, 0x4b, 0xe3, 0xeb,
	0x71, 0x27, 0xba, 0xfd, 0xb9, 0xe1, 0xb9, 0xee, 0xd1, 0x37, 0xe0, 0xf5, 0x57, 0x0c, 0x37, 0x9b,
	0x22, 0xf9, 0x00, 0xf0, 0x22, 0xb4, 0x75, 0xa3, 0xe3, 0xe9, 0xac, 0xb4, 0x50, 0xb0, 0x59, 0x58,
	0x21, 0xdd, 0xcf, 0x4d, 0x2b, 0xd9, 0x02, 0x79, 0x11, 0xfb, 0x6b, 0x88, 0x3d, 0xef, 0x13, 0x04,
	0xa3, 0x58, 0x2d, 0x9b, 0x5a, 0xf6, 0x68, 0x36, 0x1f, 0x8f, 0x25, 0x9f, 0x01, 0x32, 0x8e, 0x5f,
	0x19, 0xb2, 0x5e, 0x63, 0x0d, 0xf3, 0x9d, 0xe8, 0xbe, 0xd7, 0x0a, 0x23, 0x27, 0x76, 0xd3, 0x60,
	0x43, 0x2c, 0xff, 0x62, 0x33, 0x1f, 0x06, 0xd6, 0xff, 0x1c, 0x98, 0x3f, 0xab, 0xb2, 0x77, 0x47,
	0x1f, 0xe0, 0xee, 0xe9, 0xcb, 0xcf, 0xf1, 0x07, 0x92, 0x1b, 0x4a, 0x60, 0x61, 0x40, 0xc3, 0x24,
	0x0e, 0xbd, 0x87, 0xd0, 0x20, 0x76, 0x1f, 0x99, 0x59, 0xe8, 0x6a, 0x9f, 0xb8, 0xf4, 0x2d, 0x3c,
	0x68, 0x68, 0xea, 0x8f, 0xdc, 0x52, 0x0a, 0x4b, 0xdd, 0xb0, 0x9f, 0x57, 0x94, 0xc4, 0x5b, 0xbf,
	0xd7, 0x17, 0xeb, 0x23, 0x7c, 0x1b, 0x1c, 0xb9, 0xa1, 0x77, 0x93, 0x2c, 0x88, 0x73, 0x78, 0xa5,
	0x6d, 0x7c, 0xfc, 0x1f, 0x00, 0x00, 0xff, 0xff, 0x4a, 0x03, 0xe8, 0xe6, 0x9d, 0x02, 0x00, 0x00,
}
