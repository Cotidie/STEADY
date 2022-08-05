package main

import (
	"fmt"

	pb "github.com/Cotidie/STEADY/gRPC/protocol_buffer/golang/proto"
)

func doSimple() *pb.Simple {
	return &pb.Simple{
		Id:         42,
		IsSimple:   true,
		Name:       "A name",
		SimpleList: []int32{1, 2, 3, 4, 5, 6},
	}
}

func doComplex() pb.Complex {
	dummy_one := &pb.Dummy{
		Id:   1,
		Name: "one",
	}
	dummy_two := &pb.Dummy{
		Id:   2,
		Name: "two",
	}

	return pb.Complex{
		OneDummy:        dummy_one,
		MultipleDummies: []*pb.Dummy{dummy_one, dummy_two},
	}
}

func main() {
	fmt.Println(doComplex())
}
