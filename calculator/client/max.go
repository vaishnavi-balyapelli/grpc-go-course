package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/vaishnavi-balyapelli/grpc-go-course/calculator/proto"
)

func doMax(c pb.CalculatorServiceClient) {
	log.Println("doMax function invoked")

	stream, err := c.Max(context.Background())
	if err != nil {
		log.Fatalf("Error while Opening stream  %v\n", err)
	}
	watc := make(chan struct{})

	go func() {
		reqs := []int64{1, 9, 8, 90, 7, 6, 88}
		for _, req := range reqs {
			log.Printf("Sending Number %d\n", req)
			stream.Send(&pb.MaxRequest{
				InputNumber: req,
			})
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
				log.Fatalf("Error while receiving the stream %v\n", err)
				break
			}

			log.Printf("The maximum is %d\n", res.MaxResp)
		}
		close(watc)

	}()
	<-watc
}
