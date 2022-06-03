package main

import (
	"go-grpc-example2/client_streaming/calculator/proto"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	"sync"
)

type server struct {
	proto.UnimplementedCalculatorOperationServer
	mutex sync.Mutex
}

func (s *server) Average(server proto.CalculatorOperation_AverageServer) error {
	var sum float64 = 0
	var n uint32 = 0

	//loop untuk menerima semua data stream
	for {
		number, err := server.Recv()

		//masuk kondisi ini jika udah selesai membaca semua data stream kemudian mengembalikan nilai(resp)
		if err == io.EOF {
			f := sum / float64(n)

			//generate resp
			resp := proto.AverageResp{
				Result: f,
				Sum:    sum,
				N:      n,
			}

			return server.SendAndClose(&resp)
		}

		//if any error is found
		if err != nil {
			panic(err)
		}

		//logic
		sum += number.GetNumber()
		n++
	}

}

func main() {
	con, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	proto.RegisterCalculatorOperationServer(s, &server{})
	log.Printf("server client streaming listening at %v", con.Addr())
	if err := s.Serve(con); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
