// Code generated by protoc-gen-go. DO NOT EDIT.
// source: dto/dto.proto

package dto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type HelloReq struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloReq) Reset()         { *m = HelloReq{} }
func (m *HelloReq) String() string { return proto.CompactTextString(m) }
func (*HelloReq) ProtoMessage()    {}
func (*HelloReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_52258b0a6a265645, []int{0}
}

func (m *HelloReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloReq.Unmarshal(m, b)
}
func (m *HelloReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloReq.Marshal(b, m, deterministic)
}
func (m *HelloReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloReq.Merge(m, src)
}
func (m *HelloReq) XXX_Size() int {
	return xxx_messageInfo_HelloReq.Size(m)
}
func (m *HelloReq) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloReq.DiscardUnknown(m)
}

var xxx_messageInfo_HelloReq proto.InternalMessageInfo

func (m *HelloReq) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type Error struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Domain               string   `protobuf:"bytes,3,opt,name=domain,proto3" json:"domain,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_52258b0a6a265645, []int{1}
}

func (m *Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Error.Unmarshal(m, b)
}
func (m *Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Error.Marshal(b, m, deterministic)
}
func (m *Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error.Merge(m, src)
}
func (m *Error) XXX_Size() int {
	return xxx_messageInfo_Error.Size(m)
}
func (m *Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Error proto.InternalMessageInfo

func (m *Error) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Error) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *Error) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

type HelloResp struct {
	Err                  *Error   `protobuf:"bytes,1,opt,name=err,proto3" json:"err,omitempty"`
	Epoch                string   `protobuf:"bytes,2,opt,name=epoch,proto3" json:"epoch,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloResp) Reset()         { *m = HelloResp{} }
func (m *HelloResp) String() string { return proto.CompactTextString(m) }
func (*HelloResp) ProtoMessage()    {}
func (*HelloResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_52258b0a6a265645, []int{2}
}

func (m *HelloResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloResp.Unmarshal(m, b)
}
func (m *HelloResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloResp.Marshal(b, m, deterministic)
}
func (m *HelloResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloResp.Merge(m, src)
}
func (m *HelloResp) XXX_Size() int {
	return xxx_messageInfo_HelloResp.Size(m)
}
func (m *HelloResp) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloResp.DiscardUnknown(m)
}

var xxx_messageInfo_HelloResp proto.InternalMessageInfo

func (m *HelloResp) GetErr() *Error {
	if m != nil {
		return m.Err
	}
	return nil
}

func (m *HelloResp) GetEpoch() string {
	if m != nil {
		return m.Epoch
	}
	return ""
}

func init() {
	proto.RegisterType((*HelloReq)(nil), "dto.data.HelloReq")
	proto.RegisterType((*Error)(nil), "dto.data.Error")
	proto.RegisterType((*HelloResp)(nil), "dto.data.HelloResp")
}

func init() { proto.RegisterFile("dto/dto.proto", fileDescriptor_52258b0a6a265645) }

var fileDescriptor_52258b0a6a265645 = []byte{
	// 218 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x34, 0x8f, 0x31, 0x4f, 0xc3, 0x30,
	0x10, 0x85, 0x55, 0x42, 0x4a, 0x7b, 0x08, 0x81, 0x4e, 0x08, 0x65, 0x60, 0x80, 0x4c, 0x2c, 0x4d,
	0x4a, 0xf9, 0x07, 0x08, 0x24, 0x16, 0x96, 0x8c, 0x6c, 0xae, 0x7d, 0x72, 0x22, 0xd5, 0x3e, 0x73,
	0x76, 0xfb, 0xfb, 0x91, 0x4d, 0xba, 0xbd, 0xa7, 0x4f, 0xf7, 0xde, 0x3d, 0xb8, 0x31, 0x89, 0x7b,
	0x93, 0xb8, 0x0b, 0xc2, 0x89, 0x71, 0x95, 0xa5, 0x51, 0x49, 0xb5, 0x8f, 0xb0, 0xfa, 0xa2, 0xc3,
	0x81, 0x07, 0xfa, 0xc5, 0x3b, 0xa8, 0x5c, 0xb4, 0xcd, 0xe2, 0x69, 0xf1, 0xb2, 0x1e, 0xb2, 0x6c,
	0xbf, 0xa1, 0xfe, 0x14, 0x61, 0x41, 0x84, 0x4b, 0xcd, 0x86, 0x0a, 0xab, 0x87, 0xa2, 0xb1, 0x81,
	0x2b, 0x47, 0x31, 0x2a, 0x4b, 0xcd, 0x45, 0x39, 0x39, 0x5b, 0x7c, 0x80, 0xa5, 0x61, 0xa7, 0x26,
	0xdf, 0x54, 0x05, 0xcc, 0xae, 0xfd, 0x80, 0xf5, 0x5c, 0x16, 0x03, 0x3e, 0x43, 0x45, 0x22, 0x25,
	0xf1, 0x7a, 0x77, 0xdb, 0x9d, 0x3f, 0xea, 0x4a, 0xe1, 0x90, 0x19, 0xde, 0x43, 0x4d, 0x81, 0xf5,
	0x38, 0xe7, 0xff, 0x9b, 0xf7, 0xdd, 0xcf, 0xd6, 0x4e, 0x69, 0x3c, 0xee, 0x3b, 0xcd, 0xae, 0x3f,
	0x29, 0x3f, 0x4e, 0x74, 0x34, 0x7e, 0xfb, 0xda, 0x47, 0x92, 0xd3, 0xa4, 0x69, 0x93, 0x28, 0xa6,
	0xde, 0x4a, 0xd0, 0x1b, 0x4b, 0x3e, 0xcf, 0xde, 0x2f, 0xcb, 0xee, 0xb7, 0xbf, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x32, 0xc9, 0xb3, 0xf5, 0x08, 0x01, 0x00, 0x00,
}
