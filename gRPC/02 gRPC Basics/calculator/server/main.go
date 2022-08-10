package main

import (
	"fmt"
	"log"
	"net"

	pb "github.com/Cotidie/STEADY/gRPC/Basics/calculator/proto"
	"google.golang.org/grpc"
)

var addr string = "0.0.0.0:50052"

type Server struct {
	pb.CalculatorServiceServer
}

func main() {
	conn, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen on: %v\n", addr)
	}
	fmt.Printf("Listening on %v...\n", addr)

	server := grpc.NewServer()
	pb.RegisterCalculatorServiceServer(server, &Server{})
	err = server.Serve(conn)
	if err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}
