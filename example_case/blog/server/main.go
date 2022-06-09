package main

import (
	"context"
	"fmt"
	"go-grpc-example2/example_case/blog/helper"
	"go-grpc-example2/example_case/blog/proto"
	"go.mongodb.org/mongo-driver/bson"
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
	fmt.Println("read from mongo")

	idFromHex, err := primitive.ObjectIDFromHex(id.Id)
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"Cannot parse ID",
		)
	}

	tmp := &BlogItem{}
	filter := bson.M{"_id": idFromHex}

	result := s.Collection.FindOne(ctx, filter)
	if err2 := result.Decode(tmp); err2 != nil {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("id %v not found, error : %v", idFromHex, err2))
	}

	return documentToBlog(tmp), nil

}

func (s *server) UpdateBlog(ctx context.Context, blog *proto.Blog) (*emptypb.Empty, error) {
	fmt.Println("update data mongo")

	//getId
	idFromHex, err := primitive.ObjectIDFromHex(blog.GetId())
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("cant parse id, error : %v", err),
		)
	}

	//update data in mongodb
	updateResult, err := s.Collection.UpdateOne(
		ctx,
		bson.M{"_id": idFromHex},
		bson.M{"$set": &BlogItem{
			AuthorID: blog.GetAuthorId(),
			Title:    blog.GetTitle(),
			Content:  blog.GetContent(),
		}},
	)

	//logic if you got error when update data in mongoDB
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("failed update, error : %v", err),
		)
	}

	//logic if id not found
	if updateResult.MatchedCount == 0 {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("failed to found id %v, error : %v", idFromHex, err),
		)
	}

	return &emptypb.Empty{}, nil
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
	mongoDB, err := mongo.NewClient(options.Client().ApplyURI("mongodb://root:root@localhost:27017/"))
	helper.PrintError(err)

	err = mongoDB.Connect(context.Background())
	helper.PrintError(err)

	collection := mongoDB.Database("blogBD").Collection("blog")

	listen, err := net.Listen("tcp", ":50051")
	helper.PrintError(err)

	s := grpc.NewServer()

	proto.RegisterBlogServiceServer(s, &server{Collection: collection})

	err = s.Serve(listen)
	if err != nil {
		return
	}
}
