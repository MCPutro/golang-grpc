package main

import (
	"context"
	"fmt"
	"go-grpc-example2/client_streaming/greet/proto"
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

	client := proto.NewGreetingClient(conn)

	greetingAll, err := client.GreetingAll(context.Background())

	fmt.Println(greetingAll)

	//generate req
	reqs := []*proto.GreetingReq{
		{
			FirstName: "a",
			LastName:  "A",
		},
		{
			FirstName: "b",
			LastName:  "B",
		},
		{
			FirstName: "c",
			LastName:  "C",
		},
	}

	//send req
	for _, req := range reqs {
		greetingAll.Send(req)
		time.Sleep(500 * time.Millisecond)
	}

	//recv resp
	resp, err := greetingAll.CloseAndRecv()
	if err != nil {
		log.Fatalf("errer client stream while receiving response, err: %v", err)
	}

	fmt.Println(resp)
}
