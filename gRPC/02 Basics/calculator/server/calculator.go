package main

import (
	"context"

	pb "github.com/Cotidie/STEADY/gRPC/Basics/calculator/proto"
)

func (s *Server) Calculate(ctx context.Context, request *pb.CalculatorRequest) (*pb.CalculatorResposne, error) {
	operand1 := request.Operand1
	operand2 := request.Operand2
	calculated := operand1 + operand2

	return &pb.CalculatorResposne{
		Result: int64(calculated),
	}, nil
}
