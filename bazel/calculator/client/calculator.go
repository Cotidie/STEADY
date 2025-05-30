package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	pb "github.com/Cotidie/STEADY/gRPC/Basics/calculator/proto"
	"google.golang.org/grpc/status"
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

func doAverage(client pb.CalculatorServiceClient) {
	fmt.Println("doAverage function invoked!")

	reqs := []*pb.SingleRequest{
		{Number: 2},
		{Number: 5},
		{Number: 7},
		{Number: 100},
	}
	ctx := context.Background()
	stream, err := client.Average(ctx)
	if err != nil {
		log.Fatalf("Failed to open stream: %v\n", err)
	}

	for _, req := range reqs {
		fmt.Printf("Sending %d...\n", req.Number)
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Printf("Failed to close: %v\n", err)
	}

	fmt.Printf("Average: %v\n", res.Result)
}

func doMax(client pb.CalculatorServiceClient) {
	fmt.Println("doMax function invoked")

	stream, err := client.Max(context.Background())
	if err != nil {
		log.Fatalf("Failed to create stream: %v\n", err)
	}

	numbers := []int32{1, 5, 3, 6, 2, 20}
	reqs := []*pb.SingleRequest{}
	for _, num := range numbers {
		req := &pb.SingleRequest{Number: num}
		reqs = append(reqs, req)
	}

	waitc := make(chan struct{})

	go func() {
		for _, req := range reqs {
			fmt.Printf("## Sending %d...\n", req.Number)
			err = stream.Send(req)
			if err != nil {
				fmt.Printf("Failed to send: %v\n", err)
			}
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
				log.Fatalf("Failed to read from server: %v\n", err)
			}
			fmt.Println("Current max value:", res.Result)
		}
		close(waitc)
	}()

	fmt.Println("doMax function finished...")
	<-waitc
}

func doSqrt(client pb.CalculatorServiceClient) {
	fmt.Println("doSqrt is invoked!")

	ctx := context.Background()
	req := pb.SqrtRequest{Number: -5}
	res, err := client.Sqrt(ctx, &req)

	if err != nil {
		// ok bool represents whether the error is produced from grpc package
		e, ok := status.FromError(err)
		if !ok {
			log.Fatalf("A non gRPC error: %v\n", err)
		}

		fmt.Printf("Error message from server: %v\n", e.Message())
		fmt.Printf("Error code from server: %v\n", e.Code())
		return
	}

	fmt.Printf("doSqrt function finished with: %v\n", res.Result)
}
