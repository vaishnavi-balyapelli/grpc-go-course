package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/vaishnavi-balyapelli/grpc-go-course/blog/proto"
)

func deleteBlog(bs pb.BlogServiceClient, id string) {
	log.Printf("---DeleteBlog was invoked with ID %v---", id)
	_, err := bs.DeleteBlog(context.Background(), &pb.BlogId{Id: id})

	if err != nil {
		log.Fatalf("Error while Deleting the document %v\n", err)
	}

	fmt.Printf("Blog deleted %v\n", &pb.Blog{Id: id})
}
