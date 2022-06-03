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

// GreetingBiDirectionClient is the client API for GreetingBiDirection service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GreetingBiDirectionClient interface {
	GreetEveryone(ctx context.Context, opts ...grpc.CallOption) (GreetingBiDirection_GreetEveryoneClient, error)
}

type greetingBiDirectionClient struct {
	cc grpc.ClientConnInterface
}

func NewGreetingBiDirectionClient(cc grpc.ClientConnInterface) GreetingBiDirectionClient {
	return &greetingBiDirectionClient{cc}
}

func (c *greetingBiDirectionClient) GreetEveryone(ctx context.Context, opts ...grpc.CallOption) (GreetingBiDirection_GreetEveryoneClient, error) {
	stream, err := c.cc.NewStream(ctx, &GreetingBiDirection_ServiceDesc.Streams[0], "/proto.GreetingBiDirection/GreetEveryone", opts...)
	if err != nil {
		return nil, err
	}
	x := &greetingBiDirectionGreetEveryoneClient{stream}
	return x, nil
}

type GreetingBiDirection_GreetEveryoneClient interface {
	Send(*GreetRequest) error
	Recv() (*GreetingResponse, error)
	grpc.ClientStream
}

type greetingBiDirectionGreetEveryoneClient struct {
	grpc.ClientStream
}

func (x *greetingBiDirectionGreetEveryoneClient) Send(m *GreetRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *greetingBiDirectionGreetEveryoneClient) Recv() (*GreetingResponse, error) {
	m := new(GreetingResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GreetingBiDirectionServer is the server API for GreetingBiDirection service.
// All implementations must embed UnimplementedGreetingBiDirectionServer
// for forward compatibility
type GreetingBiDirectionServer interface {
	GreetEveryone(GreetingBiDirection_GreetEveryoneServer) error
	mustEmbedUnimplementedGreetingBiDirectionServer()
}

// UnimplementedGreetingBiDirectionServer must be embedded to have forward compatible implementations.
type UnimplementedGreetingBiDirectionServer struct {
}

func (UnimplementedGreetingBiDirectionServer) GreetEveryone(GreetingBiDirection_GreetEveryoneServer) error {
	return status.Errorf(codes.Unimplemented, "method GreetEveryone not implemented")
}
func (UnimplementedGreetingBiDirectionServer) mustEmbedUnimplementedGreetingBiDirectionServer() {}

// UnsafeGreetingBiDirectionServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GreetingBiDirectionServer will
// result in compilation errors.
type UnsafeGreetingBiDirectionServer interface {
	mustEmbedUnimplementedGreetingBiDirectionServer()
}

func RegisterGreetingBiDirectionServer(s grpc.ServiceRegistrar, srv GreetingBiDirectionServer) {
	s.RegisterService(&GreetingBiDirection_ServiceDesc, srv)
}

func _GreetingBiDirection_GreetEveryone_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GreetingBiDirectionServer).GreetEveryone(&greetingBiDirectionGreetEveryoneServer{stream})
}

type GreetingBiDirection_GreetEveryoneServer interface {
	Send(*GreetingResponse) error
	Recv() (*GreetRequest, error)
	grpc.ServerStream
}

type greetingBiDirectionGreetEveryoneServer struct {
	grpc.ServerStream
}

func (x *greetingBiDirectionGreetEveryoneServer) Send(m *GreetingResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *greetingBiDirectionGreetEveryoneServer) Recv() (*GreetRequest, error) {
	m := new(GreetRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GreetingBiDirection_ServiceDesc is the grpc.ServiceDesc for GreetingBiDirection service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GreetingBiDirection_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.GreetingBiDirection",
	HandlerType: (*GreetingBiDirectionServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GreetEveryone",
			Handler:       _GreetingBiDirection_GreetEveryone_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "greeting.proto",
}
