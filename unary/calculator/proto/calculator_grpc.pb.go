// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: calculator.proto

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

// CalculatorOperationClient is the client API for CalculatorOperation service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CalculatorOperationClient interface {
	Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddResponse, error)
	Subtract(ctx context.Context, in *SubtractRequest, opts ...grpc.CallOption) (*SubtractResponse, error)
	Multiply(ctx context.Context, in *MultiplyRequest, opts ...grpc.CallOption) (*MultiplyResponse, error)
	Divide(ctx context.Context, in *DivideRequest, opts ...grpc.CallOption) (*DivideResponse, error)
}

type calculatorOperationClient struct {
	cc grpc.ClientConnInterface
}

func NewCalculatorOperationClient(cc grpc.ClientConnInterface) CalculatorOperationClient {
	return &calculatorOperationClient{cc}
}

func (c *calculatorOperationClient) Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddResponse, error) {
	out := new(AddResponse)
	err := c.cc.Invoke(ctx, "/proto.calculatorOperation/add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calculatorOperationClient) Subtract(ctx context.Context, in *SubtractRequest, opts ...grpc.CallOption) (*SubtractResponse, error) {
	out := new(SubtractResponse)
	err := c.cc.Invoke(ctx, "/proto.calculatorOperation/subtract", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calculatorOperationClient) Multiply(ctx context.Context, in *MultiplyRequest, opts ...grpc.CallOption) (*MultiplyResponse, error) {
	out := new(MultiplyResponse)
	err := c.cc.Invoke(ctx, "/proto.calculatorOperation/multiply", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calculatorOperationClient) Divide(ctx context.Context, in *DivideRequest, opts ...grpc.CallOption) (*DivideResponse, error) {
	out := new(DivideResponse)
	err := c.cc.Invoke(ctx, "/proto.calculatorOperation/divide", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CalculatorOperationServer is the server API for CalculatorOperation service.
// All implementations must embed UnimplementedCalculatorOperationServer
// for forward compatibility
type CalculatorOperationServer interface {
	Add(context.Context, *AddRequest) (*AddResponse, error)
	Subtract(context.Context, *SubtractRequest) (*SubtractResponse, error)
	Multiply(context.Context, *MultiplyRequest) (*MultiplyResponse, error)
	Divide(context.Context, *DivideRequest) (*DivideResponse, error)
	mustEmbedUnimplementedCalculatorOperationServer()
}

// UnimplementedCalculatorOperationServer must be embedded to have forward compatible implementations.
type UnimplementedCalculatorOperationServer struct {
}

func (UnimplementedCalculatorOperationServer) Add(context.Context, *AddRequest) (*AddResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (UnimplementedCalculatorOperationServer) Subtract(context.Context, *SubtractRequest) (*SubtractResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Subtract not implemented")
}
func (UnimplementedCalculatorOperationServer) Multiply(context.Context, *MultiplyRequest) (*MultiplyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Multiply not implemented")
}
func (UnimplementedCalculatorOperationServer) Divide(context.Context, *DivideRequest) (*DivideResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Divide not implemented")
}
func (UnimplementedCalculatorOperationServer) mustEmbedUnimplementedCalculatorOperationServer() {}

// UnsafeCalculatorOperationServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CalculatorOperationServer will
// result in compilation errors.
type UnsafeCalculatorOperationServer interface {
	mustEmbedUnimplementedCalculatorOperationServer()
}

func RegisterCalculatorOperationServer(s grpc.ServiceRegistrar, srv CalculatorOperationServer) {
	s.RegisterService(&CalculatorOperation_ServiceDesc, srv)
}

func _CalculatorOperation_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorOperationServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.calculatorOperation/add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorOperationServer).Add(ctx, req.(*AddRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalculatorOperation_Subtract_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubtractRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorOperationServer).Subtract(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.calculatorOperation/subtract",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorOperationServer).Subtract(ctx, req.(*SubtractRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalculatorOperation_Multiply_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MultiplyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorOperationServer).Multiply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.calculatorOperation/multiply",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorOperationServer).Multiply(ctx, req.(*MultiplyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalculatorOperation_Divide_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DivideRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorOperationServer).Divide(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.calculatorOperation/divide",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorOperationServer).Divide(ctx, req.(*DivideRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CalculatorOperation_ServiceDesc is the grpc.ServiceDesc for CalculatorOperation service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CalculatorOperation_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.calculatorOperation",
	HandlerType: (*CalculatorOperationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "add",
			Handler:    _CalculatorOperation_Add_Handler,
		},
		{
			MethodName: "subtract",
			Handler:    _CalculatorOperation_Subtract_Handler,
		},
		{
			MethodName: "multiply",
			Handler:    _CalculatorOperation_Multiply_Handler,
		},
		{
			MethodName: "divide",
			Handler:    _CalculatorOperation_Divide_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "calculator.proto",
}
