package main

import (
	"context"
	"log"
	"net"

	pb "github.com/vaishnavi-balyapelli/grpc-go-course/blog/proto"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

var collection *mongo.Collection

var addr string = "0.0.0.0:50051"

type BlogServer struct {
	pb.BlogServiceServer
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

	collection = client.Database("blogdb").Collection("blog")

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to Listen %v\n", err)
	}
	log.Printf("Listening on port %s\n", addr)

	bs := grpc.NewServer()
	pb.RegisterBlogServiceServer(bs, &BlogServer{})

	if err := bs.Serve(lis); err != nil {
		log.Fatalf("Failed to serve %v\n", err)
	}

}
