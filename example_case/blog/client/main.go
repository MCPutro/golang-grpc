package main

import (
	"context"
	"fmt"
	"go-grpc-example2/example_case/blog/helper"
	"go-grpc-example2/example_case/blog/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func updateBlog(c proto.BlogServiceClient, id string) string {
	log.Println("---updateBlog was invoked---")

	blog := &proto.Blog{
		Id:       id,
		AuthorId: "ok ok (update)",
		Title:    "test update",
		Content:  "ha ha ha  (update)",
	}

	res, err := c.UpdateBlog(context.Background(), blog)

	if err != nil {
		return fmt.Sprintf("Unexpected error: %v\n", err)
	}

	return fmt.Sprintf("Blog has been update: %v\n", res)
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

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials())) // grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer func(conn *grpc.ClientConn) {
		err = conn.Close()
		helper.PrintError(err)
	}(conn)

	client := proto.NewBlogServiceClient(conn)

	blogId := createBlog(client)
	findById(client, blogId)
	updateBlog(client, blogId)
}
