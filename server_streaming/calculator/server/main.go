package main

import (
	"fmt"
	"go-grpc-example2/server_streaming/calculator/proto"
	"google.golang.org/grpc"
	"log"
	"net"
	"sync"
)

type server struct {
	proto.UnimplementedCalculatorServer
	sync.Mutex
}

func (s *server) Prime(req *proto.PrimeReq, stream proto.Calculator_PrimeServer) error {
	N := req.GetNumber()
	k := int64(2)

	for N > 1 {
		if N%k == 0 {
			err := stream.Send(&proto.PrimeResp{Number: k})
			if err != nil {
				panic(err)
			}
			N /= k
		} else {
			k++
		}
	}
	return nil
}

func main() {

	con, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	proto.RegisterCalculatorServer(s, &server{})
	log.Printf("server client streaming listening at %v", con.Addr())
	if err := s.Serve(con); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}

func main2() {

	N := int64(21)
	k := int64(2)

	for N > 1 {
		if N%k == 0 {
			fmt.Println(k)
			//stream.Send(&pb.PrimeResponse{
			//	Result: divisor,
			//})

			N /= k
		} else {
			k++
		}
	}
}
