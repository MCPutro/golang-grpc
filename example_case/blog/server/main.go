package main

import (
	"context"
	"fmt"
	"go-grpc-example2/example_case/blog/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
	"net"
)

type server struct {
	//proto.BlogServiceServer
	proto.UnimplementedBlogServiceServer
	Collection *mongo.Collection
}

func (s *server) CreateBlog(ctx context.Context, in *proto.Blog) (*proto.BlogId, error) {
	log.Printf("CreateBlog was invoked with %v\n", in)

	data := BlogItem{
		AuthorID: in.AuthorId,
		Title:    in.Title,
		Content:  in.Content,
	}

	res, err := s.Collection.InsertOne(ctx, data)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error: %v", err),
		)
	}

	oid, ok := res.InsertedID.(primitive.ObjectID)

	if !ok {
		return nil, status.Errorf(
			codes.Internal,
			"Cannot convert to OID",
		)
	}

	return &proto.BlogId{
		Id: oid.Hex(),
	}, nil
}

func (s *server) ReadBlog(ctx context.Context, id *proto.BlogId) (*proto.Blog, error) {
	//TODO implement me
	panic("implement me")
}

func (s *server) UpdateBlog(ctx context.Context, blog *proto.Blog) (*emptypb.Empty, error) {
	//TODO implement me
	panic("implement me")
}

func (s *server) DeleteBlog(ctx context.Context, id *proto.BlogId) (*emptypb.Empty, error) {
	//TODO implement me
	panic("implement me")
}

func (s *server) ListBlogs(empty *emptypb.Empty, blogsServer proto.BlogService_ListBlogsServer) error {
	//TODO implement me
	panic("implement me")
}

func main() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://root:root@localhost:27017/"))
	if err != nil {
		log.Fatal(err)
	}

	err = client.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	collection := client.Database("blogBD").Collection("blog")

	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()

	proto.RegisterBlogServiceServer(s, &server{Collection: collection})

	err = s.Serve(listen)
	if err != nil {
		return
	}
}
