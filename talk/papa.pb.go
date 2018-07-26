// Code generated by protoc-gen-go.
// source: papa.proto
// DO NOT EDIT!

/*
Package talk is a generated protocol buffer package.

It is generated from these files:
	papa.proto

It has these top-level messages:
	PapaTask
*/
package talk

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/golang/protobuf/ptypes/empty"

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

type PapaTask struct {
	Id   uint64 `protobuf:"fixed64,1,opt,name=id" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
}

func (m *PapaTask) Reset()                    { *m = PapaTask{} }
func (m *PapaTask) String() string            { return proto.CompactTextString(m) }
func (*PapaTask) ProtoMessage()               {}
func (*PapaTask) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *PapaTask) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *PapaTask) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*PapaTask)(nil), "talk.PapaTask")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Papa service

type PapaClient interface {
	Pull(ctx context.Context, opts ...grpc.CallOption) (Papa_PullClient, error)
}

type papaClient struct {
	cc *grpc.ClientConn
}

func NewPapaClient(cc *grpc.ClientConn) PapaClient {
	return &papaClient{cc}
}

func (c *papaClient) Pull(ctx context.Context, opts ...grpc.CallOption) (Papa_PullClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Papa_serviceDesc.Streams[0], c.cc, "/talk.Papa/Pull", opts...)
	if err != nil {
		return nil, err
	}
	x := &papaPullClient{stream}
	return x, nil
}

type Papa_PullClient interface {
	Send(*PapaTask) error
	Recv() (*PapaTask, error)
	grpc.ClientStream
}

type papaPullClient struct {
	grpc.ClientStream
}

func (x *papaPullClient) Send(m *PapaTask) error {
	return x.ClientStream.SendMsg(m)
}

func (x *papaPullClient) Recv() (*PapaTask, error) {
	m := new(PapaTask)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Papa service

type PapaServer interface {
	Pull(Papa_PullServer) error
}

func RegisterPapaServer(s *grpc.Server, srv PapaServer) {
	s.RegisterService(&_Papa_serviceDesc, srv)
}

func _Papa_Pull_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PapaServer).Pull(&papaPullServer{stream})
}

type Papa_PullServer interface {
	Send(*PapaTask) error
	Recv() (*PapaTask, error)
	grpc.ServerStream
}

type papaPullServer struct {
	grpc.ServerStream
}

func (x *papaPullServer) Send(m *PapaTask) error {
	return x.ServerStream.SendMsg(m)
}

func (x *papaPullServer) Recv() (*PapaTask, error) {
	m := new(PapaTask)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Papa_serviceDesc = grpc.ServiceDesc{
	ServiceName: "talk.Papa",
	HandlerType: (*PapaServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Pull",
			Handler:       _Papa_Pull_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "papa.proto",
}

func init() { proto.RegisterFile("papa.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 143 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0x48, 0x2c, 0x48,
	0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x29, 0x49, 0xcc, 0xc9, 0x96, 0x92, 0x4e, 0xcf,
	0xcf, 0x4f, 0xcf, 0x49, 0xd5, 0x07, 0x8b, 0x25, 0x95, 0xa6, 0xe9, 0xa7, 0xe6, 0x16, 0x94, 0x54,
	0x42, 0x94, 0x28, 0xe9, 0x71, 0x71, 0x04, 0x24, 0x16, 0x24, 0x86, 0x24, 0x16, 0x67, 0x0b, 0xf1,
	0x71, 0x31, 0x65, 0xa6, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0xb0, 0x05, 0x31, 0x65, 0xa6, 0x08, 0x09,
	0x71, 0xb1, 0xe4, 0x25, 0xe6, 0xa6, 0x4a, 0x30, 0x29, 0x30, 0x6a, 0x70, 0x06, 0x81, 0xd9, 0x46,
	0x46, 0x5c, 0x2c, 0x20, 0xf5, 0x42, 0x5a, 0x5c, 0x2c, 0x01, 0xa5, 0x39, 0x39, 0x42, 0x7c, 0x7a,
	0x20, 0x3b, 0xf4, 0x60, 0x66, 0x48, 0xa1, 0xf1, 0x35, 0x18, 0x0d, 0x18, 0x93, 0xd8, 0xc0, 0x56,
	0x19, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xd5, 0x1b, 0xd6, 0x56, 0x9b, 0x00, 0x00, 0x00,
}