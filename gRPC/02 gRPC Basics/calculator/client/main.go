package main

import (
	"log"

	pb "github.com/Cotidie/STEADY/gRPC/Basics/calculator/proto"
	"google.golang.org/grpc"
)

var addr string = "0.0.0.0:50052"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to dial: %v\n", err)
	}
	defer conn.Close()

	client := pb.NewCalculatorServiceClient(conn)
	doCalculate(client)
}
