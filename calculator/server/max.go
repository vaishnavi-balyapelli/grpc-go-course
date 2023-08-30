package main

import (
	"io"
	"log"

	pb "github.com/vaishnavi-balyapelli/grpc-go-course/calculator/proto"
)

func (s *CalculatorServer) Max(stream pb.CalculatorService_MaxServer) error {
	log.Printf("Max Streaming Server invoked")
	var max int64 = 0
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("Error while receiving the request")
		}
		if number := req.InputNumber; number > max {
			max = number
			err := stream.Send(&pb.MaxResponse{
				MaxResp: max,
			})
			if err != nil {
				log.Fatalf("Error while sending response to client %v\n", err)
			}
		}
	}
}
