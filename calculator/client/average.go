package main

import (
	"context"
	"log"
	"time"

	pb "github.com/vaishnavi-balyapelli/grpc-go-course/calculator/proto"
)

func doAverage(c pb.CalculatorServiceClient) {
	log.Printf("doAverage was invoked")

	reqs := []*pb.AverageRequest{
		{InputNumber: 5},
		{InputNumber: 9},
		{InputNumber: 7},
	}
	stream, err := c.Average(context.Background())
	if err != nil {
		log.Fatalf("Error while calling the Average %v\n", err)
	}
	for _, req := range reqs {
		log.Printf("Sending the request %v\n", req)
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while recieving the response %v\n", err)
	}
	log.Printf("The Average is %v\n", res.Average)
}
