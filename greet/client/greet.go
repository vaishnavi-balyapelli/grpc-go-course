package main

import (
	"context"
	"log"

	pb "github.com/vaishnavi-balyapelli/grpc-go-course/greet/proto"
)

func doGreet(c pb.GreetServiceClient) {
	log.Println("doGreet was invoked")
	res, err := c.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "Vaishnavi",
	})

	if err != nil {
		log.Fatalf("Error while Greeting %v\n", err)
	}

	log.Printf("Greeting: %s\n", res.Result)
}
