package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/vaishnavi-balyapelli/grpc-go-course/blog/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (bs *BlogServer) DeleteBlog(ctx context.Context, in *pb.BlogId) (*emptypb.Empty, error) {
	log.Printf("Delete Blog was invoked with ID %v\n", in)

	oid, err := primitive.ObjectIDFromHex(in.Id)

	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			"Could not Parse ID",
		)
	}

	res, err := collection.DeleteOne(context.Background(), bson.M{"_id": oid})

	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Could not delete %v\n", err),
		)
	}

	if res.DeletedCount == 0 {
		return nil, status.Errorf(
			codes.NotFound,
			"Blog was NotFound",
		)

	}

	log.Printf("Deleted the document with ID %v\n", res)
	return &emptypb.Empty{}, nil

}
