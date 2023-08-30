package main

import (
	"log"
	"net"

	pb "github.com/vaishnavi-balyapelli/grpc-go-course/calculator/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var addr string = "0.0.0.0:50051"

type CalculatorServer struct {
	pb.CalculatorServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to Listen %v\n", err)
	}
	log.Printf("Listening on port %s\n", addr)

	cs := grpc.NewServer()
	pb.RegisterCalculatorServiceServer(cs, &CalculatorServer{})

	if err := cs.Serve(lis); err != nil {
		log.Fatalf("Failed to serve %v\n", err)
	}
	reflection.Register(cs)

}
