package main

import (
	"io"
	"log"

	pb "github.com/vaishnavi-balyapelli/grpc-go-course/calculator/proto"
)

func (s *CalculatorServer) Average(stream pb.CalculatorService_AverageServer) error {
	log.Println("Average function is invoked")
	var sum float64
	count := 0
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.AverageResponse{
				Average: float32(sum) / float32(count),
			})
		}
		if err != nil {
			log.Fatalf("Error while reading the request %v\n", err)
		}
		sum += float64(req.InputNumber)
		count++
	}
}
