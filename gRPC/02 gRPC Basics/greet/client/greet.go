package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	pb "github.com/Cotidie/STEADY/gRPC/Basics/greet/proto"
)

// Call Greet() procedure in back-end side
func doGreet(client pb.GreetServiceClient) {
	ctx := context.Background()
	request := pb.GreetRequest{FirstName: "Wonseok"}

	res, err := client.Greet(ctx, &request)
	if err != nil {
		log.Fatalf("Couldn't Call Greet Procedure: %v\n", err)
	}

	fmt.Printf("Response Received! %v\n", res.Result)
}

func doGreetManyTimes(client pb.GreetServiceClient) {
	ctx := context.Background()
	request := pb.GreetRequest{FirstName: "Wonseok"}

	stream, err := client.GreetManyTImes(ctx, &request)
	if err != nil {
		log.Fatalln("Failed to send request")
	}

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error while reading streaming")
			continue
		}

		fmt.Println(msg.Result)
		time.Sleep(1 * time.Second)
	}
}

func doGreetLong(client pb.GreetServiceClient) {
	log.Println("doGreetLong invoked")

	reqs := []*pb.GreetRequest{
		{FirstName: "Wonseok"},
		{FirstName: "Finn"},
		{FirstName: "Kyle"},
	}
	ctx := context.Background()
	stream, err := client.GreetLong(ctx)
	if err != nil {
		log.Fatalf("Failed to connect to server: %v\n", err)
	}

	for _, req := range reqs {
		log.Printf("Sending req: %v\n", req)
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Failed to close stream: %v\n", err)
	}

	fmt.Printf("GreetLong finished with: %v\n", res.Result)
}

func doGreetEveryone(client pb.GreetServiceClient) {
	fmt.Printf("doGreetEveryone function invoked\n")
	reqs := []*pb.GreetRequest{
		{FirstName: "Wonseok"},
		{FirstName: "Finn"},
		{FirstName: "Kyle"},
	}

	// Open stream
	ctx := context.Background()
	stream, err := client.GreetEveryone(ctx)
	if err != nil {
		log.Fatalf("Failed to open stream: %v\n", err)
	}

	// Channel to wait for goroutine to be finished
	waitc := make(chan struct{})

	// Stream to server
	go func() {
		for _, req := range reqs {
			stream.Send(req)
			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
	}()

	// Receive from server
	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Failed to receive data from server: %v\n", err)
			}

			fmt.Println(res.Result)
		}

		close(waitc)
	}()

	<-waitc
}
