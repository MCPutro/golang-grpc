package main

import (
	"go-grpc-example2/server_streaming/greeting/proto"
	"google.golang.org/grpc"
	"log"
	"net"
	"strconv"
	"sync"
	"time"
)

type server struct {
	proto.UnimplementedGreetingServer
	sync.Mutex
}

func (s *server) GreetingManyTime(req *proto.GreetingReq, timeServer proto.Greeting_GreetingManyTimeServer) error {
	name := req.GetName()
	n := req.GetN()

	for i := 0; i < int(n); i++ {
		msg := "hello " + name + "-" + strconv.Itoa(i)

		resp := &proto.GreetingResp{Message: msg}

		err := timeServer.Send(resp)
		if err != nil {
			panic(err)
		}
		time.Sleep(500 * time.Millisecond)
	}

	return nil
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
