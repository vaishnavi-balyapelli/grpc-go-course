package main

import (
	"context"
	"log"

	pb "github.com/vaishnavi-balyapelli/grpc-go-course/blog/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (bs *BlogServer) UpdateBlog(ctx context.Context, in *pb.Blog) (*emptypb.Empty, error) {
	log.Printf("UpdateBlog Invoked with %v\n", in)

	oid, err := primitive.ObjectIDFromHex(in.Id)
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"Cannot parse ID, Invalid objectID",
		)
	}
	data := &BlogItem{
		AuthorId: in.AuthorId,
		Title:    in.Title,
		Content:  in.Content,
	}
	res, err := collection.UpdateOne(
		ctx,
		bson.M{"_id": oid},
		bson.M{"$set": data},
	)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			"Could not Update",
		)
	}
	if res.MatchedCount == 0 {
		return nil, status.Errorf(
			codes.NotFound,
			"Cannot find a blog with the given ID",
		)
	}
	return &emptypb.Empty{}, nil
}
