// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package torrent

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// RPCClient is the client API for RPC service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RPCClient interface {
	MessageStream(ctx context.Context, opts ...grpc.CallOption) (RPC_MessageStreamClient, error)
}

type rPCClient struct {
	cc grpc.ClientConnInterface
}

func NewRPCClient(cc grpc.ClientConnInterface) RPCClient {
	return &rPCClient{cc}
}

func (c *rPCClient) MessageStream(ctx context.Context, opts ...grpc.CallOption) (RPC_MessageStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_RPC_serviceDesc.Streams[0], "/torrent.RPC/MessageStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &rPCMessageStreamClient{stream}
	return x, nil
}

type RPC_MessageStreamClient interface {
	Send(*Message) error
	Recv() (*Message, error)
	grpc.ClientStream
}

type rPCMessageStreamClient struct {
	grpc.ClientStream
}

func (x *rPCMessageStreamClient) Send(m *Message) error {
	return x.ClientStream.SendMsg(m)
}

func (x *rPCMessageStreamClient) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// RPCServer is the server API for RPC service.
// All implementations must embed UnimplementedRPCServer
// for forward compatibility
type RPCServer interface {
	MessageStream(RPC_MessageStreamServer) error
	mustEmbedUnimplementedRPCServer()
}

// UnimplementedRPCServer must be embedded to have forward compatible implementations.
type UnimplementedRPCServer struct {
}

func (UnimplementedRPCServer) MessageStream(RPC_MessageStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method MessageStream not implemented")
}
func (UnimplementedRPCServer) mustEmbedUnimplementedRPCServer() {}

// UnsafeRPCServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RPCServer will
// result in compilation errors.
type UnsafeRPCServer interface {
	mustEmbedUnimplementedRPCServer()
}

func RegisterRPCServer(s grpc.ServiceRegistrar, srv RPCServer) {
	s.RegisterService(&_RPC_serviceDesc, srv)
}

func _RPC_MessageStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RPCServer).MessageStream(&rPCMessageStreamServer{stream})
}

type RPC_MessageStreamServer interface {
	Send(*Message) error
	Recv() (*Message, error)
	grpc.ServerStream
}

type rPCMessageStreamServer struct {
	grpc.ServerStream
}

func (x *rPCMessageStreamServer) Send(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

func (x *rPCMessageStreamServer) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _RPC_serviceDesc = grpc.ServiceDesc{
	ServiceName: "torrent.RPC",
	HandlerType: (*RPCServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "MessageStream",
			Handler:       _RPC_MessageStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "torrent.proto",
}
