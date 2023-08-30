package main

import (
	"log"

	pb "github.com/vaishnavi-balyapelli/grpc-go-course/calculator/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr string = "localhost:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Error while connecting to server %v\n", err)
	}
	defer conn.Close()
	c := pb.NewCalculatorServiceClient(conn)
	//doSum(c)
	//doPrimes(c)
	//doAverage(c)
	//doMax(c)
	doSqrt(c, 9)
}
