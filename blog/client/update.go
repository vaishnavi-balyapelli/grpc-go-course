package main

import (
	"context"
	"log"

	pb "github.com/vaishnavi-balyapelli/grpc-go-course/blog/proto"
)

func updateBlog(bs pb.BlogServiceClient, id string) {
	log.Println("--- UpdateBlog Invoked---")
	updatedData := &pb.Blog{
		Id:       id,
		AuthorId: "Vaishnavi Balyapelli",
		Title:    "Updated Title",
		Content:  "Updated Content",
	}
	_, err := bs.UpdateBlog(context.Background(), updatedData)

	if err != nil {
		log.Fatalf("Unable to update Blog %v\n", err)
	}
	log.Printf("Blog updated  ---> %v\n ", updatedData)
}
