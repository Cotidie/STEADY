package main

import (
	"log"

	pb "github.com/Cotidie/STEADY/gRPC/Basics/greet/proto"
	"google.golang.org/grpc"
)

var addr string = "localhost:50051"

func main() {
	// Establish connection to Server using HTTP/2
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}
	defer conn.Close()

	// Creates gRPC client that implements APIs
	// greetClient := pb.NewGreetServiceClient(conn)
	bothClient := pb.NewBothServiceClient(conn)
	//doGreetWithDeadline(greetClient, 5*time.Second)
	doBoth(bothClient)
}
