package main

import (
	"context"
	"fmt"
	"go-grpc-example2/unary/greeting/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"net"
	"sync"
)

type server struct {
	proto.UnimplementedGreetingServiceServer
	mu sync.Mutex
}

func (s *server) Greet(ctx context.Context, in *proto.GreetRequest) (*proto.GreetingResponse, error) {
	firstName := in.GetGreeting().GetFirstName()

	if firstName == "1" {
		fmt.Println("time out")
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("gak boleh 1"),
		)
	}

	//time.Sleep(5 * time.Second)
	if ctx.Err() == context.DeadlineExceeded {
		fmt.Println("time out")
		return nil, status.Errorf(
			codes.DeadlineExceeded,
			fmt.Sprintf("Cannot devide by"),
		)
	} else {
		fmt.Println("gak time out")

		result := "hello : " + firstName

		resp := proto.GreetingResponse{Result: result}

		return &resp, nil
	}

}

func main() {
	fmt.Println("server running...")

	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()

	proto.RegisterGreetingServiceServer(s, &server{})
	// mengaktifkan reflection
	// agar bisa digunakan untuk pengujian dengan evans
	reflection.Register(s)

	err = s.Serve(listen)
	if err != nil {
		return
	}
}
