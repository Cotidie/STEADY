package main

import (
	"context"
	"fmt"

	pb "github.com/Cotidie/STEADY/gRPC/Basics/calculator/proto"
)

func (s *Server) Sum(ctx context.Context, request *pb.SumRequest) (*pb.CalculateResposne, error) {
	operand1 := request.Operand1
	operand2 := request.Operand2
	calculated := operand1 + operand2

	return &pb.CalculateResposne{
		Result: int64(calculated),
	}, nil
}

func (s *Server) Primes(request *pb.PrimeRequest, stream pb.CalculatorService_PrimesServer) error {
	fmt.Printf("Primes function invoked!")
	response := pb.CalculateResposne{}
	number := request.Number

	var divide int32 = 2
	for divide <= number {
		if number%divide == 0 {
			response.Result = int64(divide)
			err := stream.Send(&response)
			if err != nil {
				fmt.Println("Failed to send", divide)
			}
			number = number / divide
		} else {
			divide = divide + 1
		}
		fmt.Printf("number: %d, divide: %d\n", number, divide)
	}

	fmt.Println("All Primes were sent!")
	return nil
}
