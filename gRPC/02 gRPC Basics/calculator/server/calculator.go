package main

import (
	"context"
	"fmt"
	"io"
	"log"

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

func (s *Server) Average(stream pb.CalculatorService_AverageServer) error {
	fmt.Printf("Average function invoked!")

	numbers := []int32{}
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			average := float32(getSum(numbers)) / float32(len(numbers))
			response := pb.AverageResponse{Result: average}
			return stream.SendAndClose(&response)
		}
		if err != nil {
			log.Fatalf("Failed to read stream: %v\n", err)
		}

		numbers = append(numbers, req.Number)
	}
}

func getSum(array []int32) int {
	sum := 0

	for _, num := range array {
		sum += int(num)
	}

	return sum
}
