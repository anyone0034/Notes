// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ProdService.proto

package prod

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

type ProdRequest struct {
	ProdId               int32    `protobuf:"varint,1,opt,name=prod_id,json=prodId,proto3" json:"prod_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProdRequest) Reset()         { *m = ProdRequest{} }
func (m *ProdRequest) String() string { return proto.CompactTextString(m) }
func (*ProdRequest) ProtoMessage()    {}
func (*ProdRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_50db98fd6a3e2ab5, []int{0}
}

func (m *ProdRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProdRequest.Unmarshal(m, b)
}
func (m *ProdRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProdRequest.Marshal(b, m, deterministic)
}
func (m *ProdRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProdRequest.Merge(m, src)
}
func (m *ProdRequest) XXX_Size() int {
	return xxx_messageInfo_ProdRequest.Size(m)
}
func (m *ProdRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ProdRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ProdRequest proto.InternalMessageInfo

func (m *ProdRequest) GetProdId() int32 {
	if m != nil {
		return m.ProdId
	}
	return 0
}

type ProdResponse struct {
	ProdStock            int32    `protobuf:"varint,1,opt,name=prod_stock,json=prodStock,proto3" json:"prod_stock,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProdResponse) Reset()         { *m = ProdResponse{} }
func (m *ProdResponse) String() string { return proto.CompactTextString(m) }
func (*ProdResponse) ProtoMessage()    {}
func (*ProdResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_50db98fd6a3e2ab5, []int{1}
}

func (m *ProdResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProdResponse.Unmarshal(m, b)
}
func (m *ProdResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProdResponse.Marshal(b, m, deterministic)
}
func (m *ProdResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProdResponse.Merge(m, src)
}
func (m *ProdResponse) XXX_Size() int {
	return xxx_messageInfo_ProdResponse.Size(m)
}
func (m *ProdResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ProdResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ProdResponse proto.InternalMessageInfo

func (m *ProdResponse) GetProdStock() int32 {
	if m != nil {
		return m.ProdStock
	}
	return 0
}

func init() {
	proto.RegisterType((*ProdRequest)(nil), "prod.ProdRequest")
	proto.RegisterType((*ProdResponse)(nil), "prod.ProdResponse")
}

func init() { proto.RegisterFile("ProdService.proto", fileDescriptor_50db98fd6a3e2ab5) }

var fileDescriptor_50db98fd6a3e2ab5 = []byte{
	// 118 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x0c, 0x28, 0xca, 0x4f,
	0x09, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x29,
	0x28, 0xca, 0x4f, 0x51, 0x52, 0xe3, 0xe2, 0x06, 0x49, 0x05, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97,
	0x08, 0x89, 0x73, 0xb1, 0x83, 0x84, 0xe3, 0x33, 0x53, 0x24, 0x18, 0x15, 0x18, 0x35, 0x58, 0x83,
	0xd8, 0x40, 0x5c, 0xcf, 0x14, 0x25, 0x5d, 0x2e, 0x1e, 0x88, 0xba, 0xe2, 0x82, 0xfc, 0xbc, 0xe2,
	0x54, 0x21, 0x59, 0x2e, 0x2e, 0xb0, 0xc2, 0xe2, 0x92, 0xfc, 0xe4, 0x6c, 0xa8, 0x5a, 0x4e, 0x90,
	0x48, 0x30, 0x48, 0x20, 0x89, 0x0d, 0x6c, 0x87, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x20, 0x2d,
	0xce, 0xa7, 0x78, 0x00, 0x00, 0x00,
}
