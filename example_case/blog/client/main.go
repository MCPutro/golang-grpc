package main

import (
	"context"
	"fmt"
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

	fmt.Println(createBlog(client))

}

func createBlog(c proto.BlogServiceClient) string {
	log.Println("---createBlog was invoked---")

	blog := &proto.Blog{
		AuthorId: "Clement",
		Title:    "My First Blog",
		Content:  "Content of the first blog",
	}

	res, err := c.CreateBlog(context.Background(), blog)

	if err != nil {
		log.Fatalf("Unexpected error: %v\n", err)
	}

	log.Printf("Blog has been created: %v\n", res)
	return res.Id
}
