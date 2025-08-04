package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/Cotidie/STEADY/gRPC/Basics/greet/proto"
)

func makeBoth() *pb.BothRequest {
	list := []*pb.KtBoth{
		{One: &pb.KtOne{Id: 1}},
		{Two: &pb.KtTwo{Name: "Won"}},
		{One: &pb.KtOne{Id: 3}, Two: &pb.KtTwo{Name: "Seok"}},
	}

	return &pb.BothRequest{
		Data: list,
	}

}

func doBoth(client pb.BothServiceClient) {
	req := makeBoth()

	res, err := client.TestBoth(context.Background(), req)
	if err != nil {
		log.Fatalf("Error occured: %v\n", err)
	}

	fmt.Printf("Result: %v\n", res.Message)
}
