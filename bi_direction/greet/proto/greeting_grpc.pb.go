// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: greeting.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// GreetBiDirectionClient is the client API for GreetBiDirection service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GreetBiDirectionClient interface {
	GreetEveryone(ctx context.Context, opts ...grpc.CallOption) (GreetBiDirection_GreetEveryoneClient, error)
}

type greetBiDirectionClient struct {
	cc grpc.ClientConnInterface
}

func NewGreetBiDirectionClient(cc grpc.ClientConnInterface) GreetBiDirectionClient {
	return &greetBiDirectionClient{cc}
}

func (c *greetBiDirectionClient) GreetEveryone(ctx context.Context, opts ...grpc.CallOption) (GreetBiDirection_GreetEveryoneClient, error) {
	stream, err := c.cc.NewStream(ctx, &GreetBiDirection_ServiceDesc.Streams[0], "/proto.GreetBiDirection/GreetEveryone", opts...)
	if err != nil {
		return nil, err
	}
	x := &greetBiDirectionGreetEveryoneClient{stream}
	return x, nil
}

type GreetBiDirection_GreetEveryoneClient interface {
	Send(*GreetRequest) error
	Recv() (*GreetResponse, error)
	grpc.ClientStream
}

type greetBiDirectionGreetEveryoneClient struct {
	grpc.ClientStream
}

func (x *greetBiDirectionGreetEveryoneClient) Send(m *GreetRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *greetBiDirectionGreetEveryoneClient) Recv() (*GreetResponse, error) {
	m := new(GreetResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GreetBiDirectionServer is the server API for GreetBiDirection service.
// All implementations must embed UnimplementedGreetBiDirectionServer
// for forward compatibility
type GreetBiDirectionServer interface {
	GreetEveryone(GreetBiDirection_GreetEveryoneServer) error
	mustEmbedUnimplementedGreetBiDirectionServer()
}

// UnimplementedGreetBiDirectionServer must be embedded to have forward compatible implementations.
type UnimplementedGreetBiDirectionServer struct {
}

func (UnimplementedGreetBiDirectionServer) GreetEveryone(GreetBiDirection_GreetEveryoneServer) error {
	return status.Errorf(codes.Unimplemented, "method GreetEveryone not implemented")
}
func (UnimplementedGreetBiDirectionServer) mustEmbedUnimplementedGreetBiDirectionServer() {}

// UnsafeGreetBiDirectionServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GreetBiDirectionServer will
// result in compilation errors.
type UnsafeGreetBiDirectionServer interface {
	mustEmbedUnimplementedGreetBiDirectionServer()
}

func RegisterGreetBiDirectionServer(s grpc.ServiceRegistrar, srv GreetBiDirectionServer) {
	s.RegisterService(&GreetBiDirection_ServiceDesc, srv)
}

func _GreetBiDirection_GreetEveryone_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GreetBiDirectionServer).GreetEveryone(&greetBiDirectionGreetEveryoneServer{stream})
}

type GreetBiDirection_GreetEveryoneServer interface {
	Send(*GreetResponse) error
	Recv() (*GreetRequest, error)
	grpc.ServerStream
}

type greetBiDirectionGreetEveryoneServer struct {
	grpc.ServerStream
}

func (x *greetBiDirectionGreetEveryoneServer) Send(m *GreetResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *greetBiDirectionGreetEveryoneServer) Recv() (*GreetRequest, error) {
	m := new(GreetRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GreetBiDirection_ServiceDesc is the grpc.ServiceDesc for GreetBiDirection service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GreetBiDirection_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.GreetBiDirection",
	HandlerType: (*GreetBiDirectionServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GreetEveryone",
			Handler:       _GreetBiDirection_GreetEveryone_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "greeting.proto",
}
