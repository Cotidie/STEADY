package main

import (
	"context"
	"fmt"

	pb "github.com/Cotidie/STEADY/gRPC/Basics/greet/proto"
)

func (s *ServerBoth) TestBoth(ctx context.Context, req *pb.BothRequest) (*pb.BothResponse, error) {
	fmt.Printf("TestBoth invoked! \n")

	for idx, both := range req.Data {
		fmt.Printf("idx: %v, content: %v\n", idx, both)
	}

	return &pb.BothResponse{Message: "Yas!"}, nil
}
