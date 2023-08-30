package main

import (
	"context"
	"log"
	"time"

	pb "github.com/vaishnavi-balyapelli/grpc-go-course/greet/proto"
)

func doLongGreet(c pb.GreetServiceClient) {
	log.Println("Long Greet was invoked")

	reqs := []*pb.GreetRequest{
		{FirstName: "Vaishnavi"},
		{FirstName: "Mario"},
		{FirstName: "Luigi"},
	}
	stream, err := c.LongGreet(context.Background())
	if err != nil {
		log.Printf("Error while calling Long Greet %v\n", err)
	}
	for _, req := range reqs {
		log.Printf("Sending Request %s\n", req)
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receiving response from LongGreet %v\n", err)
	}
	log.Printf("LongGreet %s\n", res.Result)
}
