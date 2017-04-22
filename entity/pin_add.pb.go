// Code generated by protoc-gen-go.
// source: pin_add.proto
// DO NOT EDIT!

/*
Package entity is a generated protocol buffer package.

It is generated from these files:
	pin_add.proto
	pin_delete.proto
	pin_update.proto
	response_pin.proto
	response_user.proto
	user_get.proto

It has these top-level messages:
	AddRequest
	DeleteRequest
	PinDeleteResponse
	UpdateRequest
	PinResponse
	UserResponse
	UserIdRequest
	UserTokenRequest
*/
package entity

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

type AddRequest struct {
	Title       string   `protobuf:"bytes,1,opt,name=title" json:"title,omitempty"`
	UserId      int32    `protobuf:"varint,2,opt,name=userId" json:"userId,omitempty"`
	Url         string   `protobuf:"bytes,3,opt,name=url" json:"url,omitempty"`
	Timestamp   int64    `protobuf:"varint,4,opt,name=timestamp" json:"timestamp,omitempty"`
	Description string   `protobuf:"bytes,5,opt,name=description" json:"description,omitempty"`
	Tags        []string `protobuf:"bytes,6,rep,name=tags" json:"tags,omitempty"`
}

func (m *AddRequest) Reset()                    { *m = AddRequest{} }
func (m *AddRequest) String() string            { return proto.CompactTextString(m) }
func (*AddRequest) ProtoMessage()               {}
func (*AddRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *AddRequest) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *AddRequest) GetUserId() int32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *AddRequest) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *AddRequest) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *AddRequest) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *AddRequest) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func init() {
	proto.RegisterType((*AddRequest)(nil), "entity.AddRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for AddPin service

type AddPinClient interface {
	Execute(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*PinResponse, error)
}

type addPinClient struct {
	cc *grpc.ClientConn
}

func NewAddPinClient(cc *grpc.ClientConn) AddPinClient {
	return &addPinClient{cc}
}

func (c *addPinClient) Execute(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*PinResponse, error) {
	out := new(PinResponse)
	err := grpc.Invoke(ctx, "/entity.AddPin/Execute", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for AddPin service

type AddPinServer interface {
	Execute(context.Context, *AddRequest) (*PinResponse, error)
}

func RegisterAddPinServer(s *grpc.Server, srv AddPinServer) {
	s.RegisterService(&_AddPin_serviceDesc, srv)
}

func _AddPin_Execute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddPinServer).Execute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/entity.AddPin/Execute",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddPinServer).Execute(ctx, req.(*AddRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AddPin_serviceDesc = grpc.ServiceDesc{
	ServiceName: "entity.AddPin",
	HandlerType: (*AddPinServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Execute",
			Handler:    _AddPin_Execute_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pin_add.proto",
}

func init() { proto.RegisterFile("pin_add.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 223 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x4c, 0x8f, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0x5d, 0xd3, 0xac, 0x64, 0x44, 0x90, 0x51, 0x64, 0x29, 0x1e, 0x96, 0x9e, 0x72, 0xca,
	0x41, 0x3d, 0x0b, 0x3d, 0x78, 0xf0, 0x56, 0xf6, 0x05, 0x4a, 0xec, 0x0c, 0x32, 0xd0, 0x6e, 0xd6,
	0xdd, 0x09, 0xe8, 0xdb, 0xf8, 0xa8, 0x62, 0x52, 0xa9, 0xb7, 0xff, 0xff, 0x66, 0x06, 0xe6, 0x83,
	0xab, 0x24, 0x71, 0xdb, 0x13, 0x75, 0x29, 0x0f, 0x3a, 0xa0, 0xe5, 0xa8, 0xa2, 0x5f, 0x4b, 0xcc,
	0x5c, 0xd2, 0x10, 0x0b, 0x6f, 0x93, 0xc4, 0x79, 0xb6, 0xfa, 0x36, 0x00, 0x6b, 0xa2, 0xc0, 0x1f,
	0x23, 0x17, 0xc5, 0x5b, 0xa8, 0x55, 0x74, 0xcf, 0xce, 0x78, 0xd3, 0x36, 0x61, 0x2e, 0x78, 0x07,
	0x76, 0x2c, 0x9c, 0x5f, 0xc9, 0x9d, 0x7b, 0xd3, 0xd6, 0xe1, 0xd8, 0xf0, 0x1a, 0xaa, 0x31, 0xef,
	0x5d, 0x35, 0xed, 0xfe, 0x46, 0xbc, 0x87, 0x46, 0xe5, 0xc0, 0x45, 0xfb, 0x43, 0x72, 0x0b, 0x6f,
	0xda, 0x2a, 0x9c, 0x00, 0x7a, 0xb8, 0x24, 0x2e, 0xbb, 0x2c, 0x49, 0x65, 0x88, 0xae, 0x9e, 0xee,
	0xfe, 0x23, 0x44, 0x58, 0x68, 0xff, 0x5e, 0x9c, 0xf5, 0x55, 0xdb, 0x84, 0x29, 0x3f, 0x3c, 0x83,
	0x5d, 0x13, 0x6d, 0x24, 0xe2, 0x13, 0x5c, 0xbc, 0x7c, 0xf2, 0x6e, 0x54, 0x46, 0xec, 0x66, 0xa9,
	0xee, 0xf4, 0xfc, 0xf2, 0xe6, 0x8f, 0x6d, 0x24, 0x86, 0xa3, 0xea, 0xea, 0xec, 0xcd, 0x4e, 0xa6,
	0x8f, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xa1, 0xc9, 0x33, 0x00, 0x16, 0x01, 0x00, 0x00,
}
