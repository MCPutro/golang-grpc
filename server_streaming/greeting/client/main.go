package main

import (
	"context"
	"fmt"
	"go-grpc-example2/server_streaming/greeting/proto"
	"google.golang.org/grpc"
	"io"
	"time"
)

func main() {
	conn, err := grpc.Dial("localhost:1234", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := proto.NewGreetingClient(conn)

	serverStream(client)
}

func serverStream(client proto.GreetingClient) {
	fmt.Println("server stream example from client")

	req := proto.GreetingReq{
		Name: "aku",
		N:    10,
	}

	greetManyTime, err := client.GreetingManyTime(context.Background(), &req)
	if err != nil {
		panic(err)
	}
	for {
		msg, err := greetManyTime.Recv()
		if err == io.EOF {
			fmt.Println("break")
			break
		}

		if err != nil {
			fmt.Println("panik")
			panic(err)
		}

		fmt.Println("data :", msg.GetMessage())
		for i := 0; i < 20; i++ {
			fmt.Print("-")
			time.Sleep(100 * time.Millisecond)
		}
		fmt.Println(" Saved")
	}

	fmt.Println("DONE!")
}
