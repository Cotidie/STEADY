package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/Cotidie/STEADY/gRPC/Basics/greet/proto"
)

func (s *Server) Greet(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("Greet function was invoked with %v\n", in)
	return &pb.GreetResponse{
		Result: "Hello " + in.FirstName,
	}, nil
}

func (s *Server) GreetManyTImes(in *pb.GreetRequest, stream pb.GreetService_GreetManyTImesServer) error {
	log.Printf("GreetManyTimes function was invoked with %v\n", in)
	firstName := in.FirstName

	for i := 1; i <= 10; i++ {
		response := pb.GreetResponse{
			Result: fmt.Sprintf("Hello %s, number: %d", firstName, i),
		}
		err := stream.Send(&response)
		if err != nil {
			fmt.Printf("Failed to send: %dth response", i)
		}
	}

	fmt.Println("GreetManyTimes Completed!")
	return nil
}
