package main

import (
	"context"
	"fmt"
	"go-grpc-example2/bi_direction/greet/proto"
	"google.golang.org/grpc"
	"io"
	"log"
	"time"
)

func GreetingEveryBody(client proto.GreetBiDirectionClient) {
	fmt.Println("client bi-directions")

	greetEveryone, err := client.GreetEveryone(context.Background())
	if err != nil {
		log.Fatalf("error create Bi-directional stream :%v", err)
	}

	requests := []*proto.GreetRequest{
		{
			Greeting: &proto.Greet{
				FirstName: "a",
				LastName:  "1",
			},
		},
		{
			Greeting: &proto.Greet{
				FirstName: "b",
				LastName:  "2",
			},
		},
		{
			Greeting: &proto.Greet{
				FirstName: "c",
				LastName:  "3",
			},
		},
		{
			Greeting: &proto.Greet{
				FirstName: "d",
				LastName:  "4",
			},
		},
	}

	channel := make(chan struct{})

	//send req
	go func() {
		for _, greetRequest := range requests {
			log.Printf("send req : %v\n", greetRequest)
			greetEveryone.Send(greetRequest)
			time.Sleep(1 * time.Second)
		}
		greetEveryone.CloseSend()
	}()

	//receive response
	go func() {
		for {
			response, err2 := greetEveryone.Recv()

			if err2 == io.EOF {
				break
			}

			if err2 != nil {
				log.Printf("error while receiving : %v\n", response)
				break
			}

			//print response
			log.Printf("receive resp : %v\n", response)
		}
		close(channel)
	}()

	<-channel
}

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := proto.NewGreetBiDirectionClient(conn)

	GreetingEveryBody(client)
}
