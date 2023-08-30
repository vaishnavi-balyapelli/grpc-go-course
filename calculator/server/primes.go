package main

import (
	"log"

	pb "github.com/vaishnavi-balyapelli/grpc-go-course/calculator/proto"
)

func (s *CalculatorServer) Primes(in *pb.PrimesRequest, stream pb.CalculatorService_PrimesServer) error {
	log.Printf("Primes function was invoked %v\n", in)
	number := in.Number
	divisor := int64(2)

	for number > 1 {
		if number%uint64(divisor) == 0 {
			stream.Send(&pb.PrimesResponse{
				Result: uint64(divisor),
			})
			number /= uint64(divisor)
		} else {
			divisor++
		}
	}
	return nil

}
