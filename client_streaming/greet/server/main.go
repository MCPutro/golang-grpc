package main

import (
	"go-grpc-example2/client_streaming/greet/proto"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	"sync"
)

type server struct {
	proto.UnimplementedGreetingServer
	sync.Mutex
}

func (s *server) GreetingAll(server proto.Greeting_GreetingAllServer) error {
	nameList := "Hello, "

	for {
		nameRecv, err := server.Recv()

		//has received all nameList then return the message
		if err == io.EOF {
			greetingResp := &proto.GreetingResp{Msg: nameList}
			return server.SendAndClose(greetingResp)
		}

		//if any error when recv data
		if err != nil {
			panic(err.Error())
		}

		//logic
		nameList += nameRecv.GetFirstName() + " " + nameRecv.GetLastName() + "!, "
	}
}

func main() {
	con, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	proto.RegisterGreetingServer(s, &server{})
	log.Printf("server client streaming listening at %v", con.Addr())
	if err := s.Serve(con); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
