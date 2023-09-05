package main

import (
	"context"
	"io"
	"log"

	pb "github.com/vaishnavi-balyapelli/grpc-go-course/blog/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

func listBlog(bs pb.BlogServiceClient) {
	log.Println("----ListBlog client Invoked----")

	stream, err := bs.ListBlogs(context.Background(), &emptypb.Empty{})

	if err != nil {
		log.Fatalf("Error occured while Listing the docs %v\n", err)
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("An unknown error occurred %v\n", err)
		}
		log.Printf("ListBlogs --- %v\n", res)
	}

}
