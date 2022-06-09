package main

import (
	"context"
	"fmt"
	"go-grpc-example2/example_case/blog/helper"
	"go-grpc-example2/example_case/blog/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
	"io"
	"log"
	"time"
)

func deleteBlog(c proto.BlogServiceClient, id string) {
	fmt.Println("delete data with id ", id)

	blogIdReq := &proto.BlogId{Id: id}

	_, err := c.DeleteBlog(context.Background(), blogIdReq)

	if err != nil {
		log.Fatalf("error delete id %v, error : %v", id, err)
	}

	log.Print("complete delete id ", id)
}

func getList(c proto.BlogServiceClient) {
	log.Println("---listBlog was invoked---")
	stream, err := c.ListBlogs(context.Background(), &emptypb.Empty{})

	if err != nil {
		log.Fatalf("Error while calling ListBlogs: %v\n", err)
	}

	for {
		res, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Something happened: %v\n", err)
		}

		log.Println(res)
		for i := 1; i <= 100; i++ {
			time.Sleep(10 * time.Millisecond)
			fmt.Print("-")
		}
		fmt.Println(" OK")
	}
}

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
	getList(client)
	deleteBlog(client, blogId)
}
