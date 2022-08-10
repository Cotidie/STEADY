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
