package main

import (
	"context"
	"log"

	pb "github.com/vaishnavi-balyapelli/grpc-go-course/calculator/proto"
)

func (s *CalculatorServer) Sum(ctx context.Context, req *pb.SumRequest) (*pb.SumResponse, error) {
	log.Printf("Sum function was invoked with %v\n", req)
	return &pb.SumResponse{
		SumResult: req.Num1 + req.Num2,
	}, nil
}
