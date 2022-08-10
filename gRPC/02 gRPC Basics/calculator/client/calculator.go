package main

import (
	"context"
	"fmt"
	"io"
	"log"

	pb "github.com/Cotidie/STEADY/gRPC/Basics/calculator/proto"
)

func doSum(client pb.CalculatorServiceClient) {
	ctx := context.Background()
	request := pb.SumRequest{
		Operand1: 3,
		Operand2: 10,
	}

	res, err := client.Sum(ctx, &request)
	if err != nil {
		log.Fatalf("Failed to call Calculate: %v\n", err)
	}

	fmt.Printf("%d + %d = %d\n", 3, 10, res.Result)
}

func doPrimes(client pb.CalculatorServiceClient, number int32) {
	ctx := context.Background()
	request := pb.PrimeRequest{Number: number}

	stream, err := client.Primes(ctx, &request)
	if err != nil {
		log.Fatalf("Failed to call Prime: %v\n", err)
	}

	primes := []int64{}
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Failed to get message: %v\n", err)
		}

		primes = append(primes, msg.Result)
	}

	fmt.Printf("Primes for %d: %v\n", number, primes)
}
