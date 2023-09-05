package main

import (
	"context"
	"log"

	pb "github.com/vaishnavi-balyapelli/grpc-go-course/blog/proto"
)

func createBlog(c pb.BlogServiceClient) string {
	log.Printf("----createBlog Invoked----")
	blog := &pb.Blog{
		AuthorId: "Vaishnavi",
		Title:    "golang-grpc-blog",
		Content:  "golang-grpc-blog-content",
	}

	res, err := c.CreateBlog(context.Background(), blog)

	if err != nil {
		log.Fatalf("Unexpected Error %v\n", err)
	}

	log.Printf("Blog created %s\n", res.Id)
	return res.Id
}
