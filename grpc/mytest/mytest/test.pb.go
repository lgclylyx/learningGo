// Code generated by protoc-gen-go. DO NOT EDIT.
// source: test.proto

package mytest

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Seed struct {
	Seed                 int32    `protobuf:"varint,1,opt,name=seed,proto3" json:"seed,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Seed) Reset()         { *m = Seed{} }
func (m *Seed) String() string { return proto.CompactTextString(m) }
func (*Seed) ProtoMessage()    {}
func (*Seed) Descriptor() ([]byte, []int) {
	return fileDescriptor_test_02d1db2cc1bde6a0, []int{0}
}
func (m *Seed) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Seed.Unmarshal(m, b)
}
func (m *Seed) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Seed.Marshal(b, m, deterministic)
}
func (dst *Seed) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Seed.Merge(dst, src)
}
func (m *Seed) XXX_Size() int {
	return xxx_messageInfo_Seed.Size(m)
}
func (m *Seed) XXX_DiscardUnknown() {
	xxx_messageInfo_Seed.DiscardUnknown(m)
}

var xxx_messageInfo_Seed proto.InternalMessageInfo

func (m *Seed) GetSeed() int32 {
	if m != nil {
		return m.Seed
	}
	return 0
}

type Reply struct {
	Reply                string   `protobuf:"bytes,1,opt,name=reply,proto3" json:"reply,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Reply) Reset()         { *m = Reply{} }
func (m *Reply) String() string { return proto.CompactTextString(m) }
func (*Reply) ProtoMessage()    {}
func (*Reply) Descriptor() ([]byte, []int) {
	return fileDescriptor_test_02d1db2cc1bde6a0, []int{1}
}
func (m *Reply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Reply.Unmarshal(m, b)
}
func (m *Reply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Reply.Marshal(b, m, deterministic)
}
func (dst *Reply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Reply.Merge(dst, src)
}
func (m *Reply) XXX_Size() int {
	return xxx_messageInfo_Reply.Size(m)
}
func (m *Reply) XXX_DiscardUnknown() {
	xxx_messageInfo_Reply.DiscardUnknown(m)
}

var xxx_messageInfo_Reply proto.InternalMessageInfo

func (m *Reply) GetReply() string {
	if m != nil {
		return m.Reply
	}
	return ""
}

func init() {
	proto.RegisterType((*Seed)(nil), "mytest.Seed")
	proto.RegisterType((*Reply)(nil), "mytest.Reply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// FirstClient is the client API for First service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FirstClient interface {
	GetNum(ctx context.Context, in *Seed, opts ...grpc.CallOption) (*Reply, error)
}

type firstClient struct {
	cc *grpc.ClientConn
}

func NewFirstClient(cc *grpc.ClientConn) FirstClient {
	return &firstClient{cc}
}

func (c *firstClient) GetNum(ctx context.Context, in *Seed, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := c.cc.Invoke(ctx, "/mytest.First/GetNum", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FirstServer is the server API for First service.
type FirstServer interface {
	GetNum(context.Context, *Seed) (*Reply, error)
}

func RegisterFirstServer(s *grpc.Server, srv FirstServer) {
	s.RegisterService(&_First_serviceDesc, srv)
}

func _First_GetNum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Seed)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FirstServer).GetNum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mytest.First/GetNum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FirstServer).GetNum(ctx, req.(*Seed))
	}
	return interceptor(ctx, in, info, handler)
}

var _First_serviceDesc = grpc.ServiceDesc{
	ServiceName: "mytest.First",
	HandlerType: (*FirstServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetNum",
			Handler:    _First_GetNum_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "test.proto",
}

func init() { proto.RegisterFile("test.proto", fileDescriptor_test_02d1db2cc1bde6a0) }

var fileDescriptor_test_02d1db2cc1bde6a0 = []byte{
	// 128 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0x49, 0x2d, 0x2e,
	0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xcb, 0xad, 0x04, 0xf1, 0x94, 0xa4, 0xb8, 0x58,
	0x82, 0x53, 0x53, 0x53, 0x84, 0x84, 0xb8, 0x58, 0x8a, 0x53, 0x53, 0x53, 0x24, 0x18, 0x15, 0x18,
	0x35, 0x58, 0x83, 0xc0, 0x6c, 0x25, 0x59, 0x2e, 0xd6, 0xa0, 0xd4, 0x82, 0x9c, 0x4a, 0x21, 0x11,
	0x2e, 0xd6, 0x22, 0x10, 0x03, 0x2c, 0xcb, 0x19, 0x04, 0xe1, 0x18, 0x19, 0x70, 0xb1, 0xba, 0x65,
	0x16, 0x15, 0x97, 0x08, 0xa9, 0x73, 0xb1, 0xb9, 0xa7, 0x96, 0xf8, 0x95, 0xe6, 0x0a, 0xf1, 0xe8,
	0x41, 0x8c, 0xd5, 0x03, 0x99, 0x29, 0xc5, 0x0b, 0xe3, 0x81, 0x4d, 0x51, 0x62, 0x48, 0x62, 0x03,
	0xdb, 0x6d, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x1a, 0x26, 0x19, 0x8e, 0x89, 0x00, 0x00, 0x00,
}
