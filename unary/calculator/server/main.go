package main

import (
	"context"
	"fmt"
	"go-grpc-example2/unary/calculator/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"
	"log"
	"net"
	"sync"
)

type Server struct {
	proto.UnimplementedCalculatorOperationServer
	mutex sync.Mutex
}

func (s *Server) Add(ctx context.Context, request *proto.AddRequest) (*proto.AddResponse, error) {
	varA := request.GetVarA()
	varB := request.GetVarB()

	resp := proto.AddResponse{Result: varA + varB}

	return &resp, nil
}

func (s *Server) Subtract(ctx context.Context, request *proto.SubtractRequest) (*proto.SubtractResponse, error) {
	varA := request.GetVarA()
	varB := request.GetVarB()

	resp := proto.SubtractResponse{Result: varA - varB}

	return &resp, nil
}

func (s *Server) Multiply(ctx context.Context, request *proto.MultiplyRequest) (*proto.MultiplyResponse, error) {
	varA := request.GetVarA()
	varB := request.GetVarB()

	resp := proto.MultiplyResponse{Result: varA * varB}

	return &resp, nil
}

func (s *Server) Divide(ctx context.Context, request *proto.DivideRequest) (*proto.DivideResponse, error) {
	varA := request.GetVarA()
	varB := request.GetVarB()

	//error handler
	if varB == 0 {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Cannot devide by %v", varB),
		)
	}

	resp := proto.DivideResponse{Result: varA / varB}

	return &resp, nil
}

func main() {
	fmt.Println("server running...")

	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	var opt []grpc.ServerOption
	tls := true

	if tls {
		certFile := "ssl/server.crt"
		keyFile := "ssl/server.pem"
		creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)

		if err != nil {
			log.Fatalf("failed load cert %v", err)
			return
		}

		opt = append(opt, grpc.Creds(creds))
	}

	s := grpc.NewServer(opt...)

	proto.RegisterCalculatorOperationServer(s, &Server{})

	err = s.Serve(listen)
	if err != nil {
		return
	}
}
