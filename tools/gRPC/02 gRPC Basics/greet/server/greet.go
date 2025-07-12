package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	pb "github.com/Cotidie/STEADY/gRPC/Basics/greet/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

func (s *Server) GreetLong(stream pb.GreetService_GreetLongServer) error {
	fmt.Printf("GreetLong function is invoked")
	names := []string{}

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			response := pb.GreetResponse{Result: fmt.Sprintf("%v", names)}
			return stream.SendAndClose(&response)
		}
		if err != nil {
			log.Fatalf("Failed to read client stream %v\n", err)
		}

		names = append(names, msg.FirstName)
	}
}

func (s *Server) GreetEveryone(stream pb.GreetService_GreetEveryoneServer) error {
	fmt.Println("GreetEveryone function invoked")

	names := []string{}
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Failed to receive: %v\n", err)
		}
		names = append(names, req.FirstName)
	}

	response := pb.GreetResponse{}
	for _, name := range names {
		sayHello := "Hello, " + name
		response.Result = sayHello

		err := stream.Send(&response)
		if err != nil {
			log.Fatalf("Failed to send: %v\n", err)
		}
	}

	return nil
}

func (s *Server) GreetWithDeadline(ctx context.Context, req *pb.GreetRequest) (*pb.GreetResponse, error) {
	fmt.Printf("GreetWithDeadline is invoked with: %v\n", req)

	// Error cases are
	for i := 0; i < 3; i++ {
		if ctx.Err() == context.DeadlineExceeded {
			fmt.Println("The deadline has expired!")
			return nil, status.Error(codes.DeadlineExceeded, codes.DeadlineExceeded.String())
		}

		time.Sleep(1 * time.Second)
	}

	return &pb.GreetResponse{
		Result: "Hello " + req.FirstName,
	}, nil
}
