package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/vaishnavi-balyapelli/grpc-go-course/greet/proto"
)

func doGreetEveryOne(c pb.GreetServiceClient) {
	log.Println("doGreetEveryOne is invoked")

	stream, err := c.GreetEveryOne(context.Background())
	if err != nil {
		log.Fatalf("Error while creating the stream ")
	}
	reqs := []*pb.GreetRequest{
		{FirstName: "Vaishnavi"},
		{FirstName: "test"},
		{FirstName: "Nikkie"},
	}

	watc := make(chan struct{})

	go func() {
		for _, req := range reqs {
			log.Printf("Send request %v\n", req)
			stream.Send(req)
			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()

	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Printf("Error while receiving the response %v\n", err)
			}
			log.Printf("Recieved response %v\n", res.Result)
		}
		close(watc)
	}()

	<-watc
}
