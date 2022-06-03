package main

import (
	"fmt"
	"go-grpc-example2/bi_direction/greet/proto"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	"sync"
)

type server struct {
	proto.UnimplementedGreetBiDirectionServer
	sync.Mutex
}

func (s *server) GreetEveryone(everyoneServer proto.GreetBiDirection_GreetEveryoneServer) error {
	for {
		rec, err := everyoneServer.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("error while reading Bi-directional stream :%v", err)
		}

		msg := "hello " + rec.GetGreeting().GetFirstName() + rec.GetGreeting().GetLastName() + "!"

		err = everyoneServer.Send(&proto.GreetResponse{Msg: msg})

		if err != nil {
			log.Fatalf("error while send data to client :%v", err)
		}
	}

}

//func (s *server) mustEmbedUnimplementedGreetingBiDirectionServer() {
//	//TODO implement me
//	panic("implement me")
//}

func main() {
	fmt.Println("server running...")

	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()

	proto.RegisterGreetBiDirectionServer(s, &server{})

	err = s.Serve(listen)
	if err != nil {
		return
	}
}
