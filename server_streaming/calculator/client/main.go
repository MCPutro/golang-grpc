package main

import (
	"context"
	"fmt"
	"go-grpc-example2/server_streaming/calculator/proto"
	"google.golang.org/grpc"
	"io"
)

func main() {
	conn, err := grpc.Dial("localhost:1234", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := proto.NewCalculatorClient(conn)

	req := proto.PrimeReq{Number: 1200}

	prime, err := client.Prime(context.Background(), &req)

	for {
		primeResp, err := prime.Recv()

		if err == io.EOF {
			fmt.Println("break")
			break
		}

		if err != nil {
			panic(err)
		}

		fmt.Println(primeResp.Number)
	}
}
