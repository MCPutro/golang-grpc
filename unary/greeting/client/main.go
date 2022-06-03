package main

import (
	"context"
	"fmt"
	"go-grpc-example2/unary/greeting/proto"
	"google.golang.org/grpc"
	"time"
)

func unary(client proto.GreetingServiceClient) {
	fmt.Println("unary example")
	request := &proto.GreetRequest{Greeting: &proto.Greeting{
		FirstName: "11",
		LastName:  "b",
	}}

	//call function rpc Greet
	ctx, cancelFunc := context.WithDeadline(context.Background(), time.Now())
	defer cancelFunc()

	greet, err := client.Greet(ctx, request)
	if err != nil {
		// panic(err)
		fmt.Println(">>", err.Error())

		return
	}

	fmt.Println(greet)
}

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer func(conn *grpc.ClientConn) {
		err2 := conn.Close()
		if err2 != nil {
			panic(err2)
		}
	}(conn)

	client := proto.NewGreetingServiceClient(conn)

	unary(client)
}
