package main

import (
	"context"
	"go-grpc-example2/example_case/blog/proto"
	"google.golang.org/grpc"
	"log"
)

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

	client := proto.NewBlogServiceClient(conn)

	blogId := createBlog(client)
	findById(client, blogId)
}

func findById(c proto.BlogServiceClient, id string) *proto.Blog {
	log.Println("---readBlog was invoked---")

	req := &proto.BlogId{Id: id}
	res, err := c.ReadBlog(context.Background(), req)

	if err != nil {
		log.Fatalf("Error happened while reading: %v\n", err)
	}

	log.Printf("Blog was read: %v\n", res)
	return res
}

func createBlog(c proto.BlogServiceClient) string {
	log.Println("---createBlog was invoked---")

	blog := &proto.Blog{
		AuthorId: "ok ok",
		Title:    "test insert",
		Content:  "ha ha ha",
	}

	res, err := c.CreateBlog(context.Background(), blog)

	if err != nil {
		log.Fatalf("Unexpected error: %v\n", err)
	}

	log.Printf("Blog has been created: %v\n", res)
	return res.Id
}
