package main

import (
	"log"

	pb "github.com/vaishnavi-balyapelli/grpc-go-course/blog/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr string = "localhost:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Error while connecting to server %v\n", err)
	}
	defer conn.Close()
	bc := pb.NewBlogServiceClient(conn)
	id := createBlog(bc)
	readBlog(bc, id)
	//readBlog(bc, "NonExistant ID")
	updateBlog(bc, id)
	listBlog(bc)
	deleteBlog(bc, id)
}
