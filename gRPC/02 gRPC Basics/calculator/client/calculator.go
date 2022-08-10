package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/Cotidie/STEADY/gRPC/Basics/calculator/proto"
)

func doCalculate(client pb.CalculatorServiceClient) {
	ctx := context.Background()
	request := pb.CalculatorRequest{
		Operand1: 3,
		Operand2: 10,
	}

	res, err := client.Calculate(ctx, &request)
	if err != nil {
		log.Fatalf("Failed to call Calculate: %v\n", err)
	}

	fmt.Printf("%d + %d = %d", 3, 10, res.Result)
}
