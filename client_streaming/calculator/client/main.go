package main

import (
	"context"
	"fmt"
	"go-grpc-example2/client_streaming/calculator/proto"
	"google.golang.org/grpc"
	"log"
	"time"
)

func main() {
	conn, err := grpc.Dial(":1234", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := proto.NewCalculatorOperationClient(conn)

	rata2(client)
}

func rata2(client proto.CalculatorOperationClient) {
	rataRataClient, err := client.Average(context.Background())

	if err != nil {
		log.Fatalf("errer client.go client stream %v", err)
	}

	//generate req
	numberList := []float64{2, 1, 3, 4, 5}
	req := proto.AverageReq{Number: 0}

	// send req
	for _, number := range numberList {
		fmt.Println("send req: ", number)
		req.Number = number
		rataRataClient.Send(&req)
		time.Sleep(1 * time.Second)
	}

	//receive resp
	resp, err := rataRataClient.CloseAndRecv()
	if err != nil {
		log.Fatalf("errer client stream while receiving response, err: %v", err)
	}

	fmt.Println("Sum : ", resp.GetSum())
	fmt.Println("n : ", resp.GetN())
	fmt.Println("Average: ", resp.GetResult())
}
