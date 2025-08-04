package main

import (
	"log"
	"net"

	pb "github.com/Cotidie/STEADY/gRPC/Basics/greet/proto"
	"google.golang.org/grpc"
)

var addr string = "0.0.0.0:50051"

type Server struct {
	pb.GreetServiceServer
}

type ServerBoth struct {
	pb.BothServiceServer
}

func main() {
	// Listen establishes Transport and Session layers
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen on: %v\n", err)
	}

	log.Printf("Listening on %s\n", addr)

	// gRPC Server establishes Presentation and Application layers
	s := grpc.NewServer()
	pb.RegisterGreetServiceServer(s, &Server{})
	pb.RegisterBothServiceServer(s, &ServerBoth{})
	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}
