package main

import (
	"context"
	"fmt"
	"log"

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
