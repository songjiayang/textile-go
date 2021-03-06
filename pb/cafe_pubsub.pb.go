// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cafe_pubsub.proto

package pb

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

type CafePubSubContactQuery struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	FindId               string   `protobuf:"bytes,2,opt,name=findId,proto3" json:"findId,omitempty"`
	FindAddress          string   `protobuf:"bytes,3,opt,name=findAddress,proto3" json:"findAddress,omitempty"`
	FindUsername         string   `protobuf:"bytes,4,opt,name=findUsername,proto3" json:"findUsername,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CafePubSubContactQuery) Reset()         { *m = CafePubSubContactQuery{} }
func (m *CafePubSubContactQuery) String() string { return proto.CompactTextString(m) }
func (*CafePubSubContactQuery) ProtoMessage()    {}
func (*CafePubSubContactQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_cafe_pubsub_05f629cd2b86bfba, []int{0}
}
func (m *CafePubSubContactQuery) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CafePubSubContactQuery.Unmarshal(m, b)
}
func (m *CafePubSubContactQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CafePubSubContactQuery.Marshal(b, m, deterministic)
}
func (dst *CafePubSubContactQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CafePubSubContactQuery.Merge(dst, src)
}
func (m *CafePubSubContactQuery) XXX_Size() int {
	return xxx_messageInfo_CafePubSubContactQuery.Size(m)
}
func (m *CafePubSubContactQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_CafePubSubContactQuery.DiscardUnknown(m)
}

var xxx_messageInfo_CafePubSubContactQuery proto.InternalMessageInfo

func (m *CafePubSubContactQuery) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *CafePubSubContactQuery) GetFindId() string {
	if m != nil {
		return m.FindId
	}
	return ""
}

func (m *CafePubSubContactQuery) GetFindAddress() string {
	if m != nil {
		return m.FindAddress
	}
	return ""
}

func (m *CafePubSubContactQuery) GetFindUsername() string {
	if m != nil {
		return m.FindUsername
	}
	return ""
}

type CafePubSubContactQueryResult struct {
	Id                   string     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Contacts             []*Contact `protobuf:"bytes,2,rep,name=contacts,proto3" json:"contacts,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *CafePubSubContactQueryResult) Reset()         { *m = CafePubSubContactQueryResult{} }
func (m *CafePubSubContactQueryResult) String() string { return proto.CompactTextString(m) }
func (*CafePubSubContactQueryResult) ProtoMessage()    {}
func (*CafePubSubContactQueryResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_cafe_pubsub_05f629cd2b86bfba, []int{1}
}
func (m *CafePubSubContactQueryResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CafePubSubContactQueryResult.Unmarshal(m, b)
}
func (m *CafePubSubContactQueryResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CafePubSubContactQueryResult.Marshal(b, m, deterministic)
}
func (dst *CafePubSubContactQueryResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CafePubSubContactQueryResult.Merge(dst, src)
}
func (m *CafePubSubContactQueryResult) XXX_Size() int {
	return xxx_messageInfo_CafePubSubContactQueryResult.Size(m)
}
func (m *CafePubSubContactQueryResult) XXX_DiscardUnknown() {
	xxx_messageInfo_CafePubSubContactQueryResult.DiscardUnknown(m)
}

var xxx_messageInfo_CafePubSubContactQueryResult proto.InternalMessageInfo

func (m *CafePubSubContactQueryResult) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *CafePubSubContactQueryResult) GetContacts() []*Contact {
	if m != nil {
		return m.Contacts
	}
	return nil
}

func init() {
	proto.RegisterType((*CafePubSubContactQuery)(nil), "CafePubSubContactQuery")
	proto.RegisterType((*CafePubSubContactQueryResult)(nil), "CafePubSubContactQueryResult")
}

func init() { proto.RegisterFile("cafe_pubsub.proto", fileDescriptor_cafe_pubsub_05f629cd2b86bfba) }

var fileDescriptor_cafe_pubsub_05f629cd2b86bfba = []byte{
	// 196 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4c, 0x4e, 0x4c, 0x4b,
	0x8d, 0x2f, 0x28, 0x4d, 0x2a, 0x2e, 0x4d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x97, 0xe2, 0xce,
	0xcd, 0x4f, 0x49, 0xcd, 0x81, 0x70, 0x94, 0xda, 0x18, 0xb9, 0xc4, 0x9c, 0x13, 0xd3, 0x52, 0x03,
	0x4a, 0x93, 0x82, 0x4b, 0x93, 0x9c, 0xf3, 0xf3, 0x4a, 0x12, 0x93, 0x4b, 0x02, 0x4b, 0x53, 0x8b,
	0x2a, 0x85, 0xf8, 0xb8, 0x98, 0x32, 0x53, 0x24, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0x98, 0x32,
	0x53, 0x84, 0xc4, 0xb8, 0xd8, 0xd2, 0x32, 0xf3, 0x52, 0x3c, 0x53, 0x24, 0x98, 0xc0, 0x62, 0x50,
	0x9e, 0x90, 0x02, 0x17, 0x37, 0x88, 0xe5, 0x98, 0x92, 0x52, 0x94, 0x5a, 0x5c, 0x2c, 0xc1, 0x0c,
	0x96, 0x44, 0x16, 0x12, 0x52, 0xe2, 0xe2, 0x01, 0x71, 0x43, 0x8b, 0x53, 0x8b, 0xf2, 0x12, 0x73,
	0x53, 0x25, 0x58, 0xc0, 0x4a, 0x50, 0xc4, 0x94, 0x42, 0xb8, 0x64, 0xb0, 0xbb, 0x23, 0x28, 0xb5,
	0xb8, 0x34, 0xa7, 0x04, 0xc3, 0x35, 0x2a, 0x5c, 0x1c, 0xc9, 0x10, 0x55, 0xc5, 0x12, 0x4c, 0x0a,
	0xcc, 0x1a, 0xdc, 0x46, 0x1c, 0x7a, 0x50, 0x6d, 0x41, 0x70, 0x19, 0x27, 0x96, 0x28, 0xa6, 0x82,
	0xa4, 0x24, 0x36, 0xb0, 0x5f, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xda, 0x52, 0xfe, 0x8e,
	0x0d, 0x01, 0x00, 0x00,
}
