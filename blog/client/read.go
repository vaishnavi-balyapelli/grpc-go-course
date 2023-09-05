package main

import (
	"context"
	"log"

	pb "github.com/vaishnavi-balyapelli/grpc-go-course/blog/proto"
)

func readBlog(bs pb.BlogServiceClient, id string) *pb.Blog {
	log.Println("-----ReadBlog Client Invoked-----")
	req := &pb.BlogId{Id: id}
	res, err := bs.ReadBlog(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while reading the blog %v\n", err)
	}
	log.Printf("Blog was read %v\n", res)
	return res
}
