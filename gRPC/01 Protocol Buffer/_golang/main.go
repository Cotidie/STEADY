package main

import (
	"fmt"
	"io/ioutil"
	"log"

	pb "github.com/Cotidie/STEADY/gRPC/protocol_buffer/golang/proto"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func doSimple() *pb.Simple {
	return &pb.Simple{
		Id:         42,
		IsSimple:   true,
		Name:       "A name",
		SimpleList: []int32{1, 2, 3, 4, 5, 6},
	}
}

func doComplex() *pb.Complex {
	dummy_one := &pb.Dummy{
		Id:   1,
		Name: "one",
	}
	dummy_two := &pb.Dummy{
		Id:   2,
		Name: "two",
	}
	oneOfId := &pb.Result_Id{Id: 32}
	maps := map[string]int32{
		"a": 1,
		"b": 2,
	}

	return &pb.Complex{
		OneDummy:        dummy_one,
		MultipleDummies: []*pb.Dummy{dummy_one, dummy_two},
		EyeColor:        pb.EyeColor_EYE_COLOR_GREEN,
		Result:          &pb.Result{Result: oneOfId},
		Maps:            maps,
	}
}

// Serialize: proto.Marshal
// Save: ioutil.WriteFile
func writeToFile(fname string, pbf proto.Message) {
	out, err := proto.Marshal(pbf)
	if err != nil {
		log.Fatalln("Cannot serialize the pb object")
		return
	}

	err = ioutil.WriteFile(fname, out, 0644)
	if err != nil {
		log.Fatalln("Error in writing to file")
		return
	}

	fmt.Println("Data has been written!")
}

func ReadFromFile(fname string, pbf proto.Message) {
	in, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatalln("Couldn't read file")
		return
	}

	err = proto.Unmarshal(in, pbf)
	if err != nil {
		log.Fatalln("Couldn't unmarshal protocol buffer bytes")
		return
	}

	fmt.Println("Data has been retrieved!")
}

func writeToJSON(fname string, pbf proto.Message) {
	out, err := protojson.Marshal(pbf)
	if err != nil {
		log.Fatalln("Couldn't marshal protobuf")
		return
	}

	fname = fmt.Sprintf("%s.json", fname)
	err = ioutil.WriteFile(fname, out, 0644)
	if err != nil {
		log.Fatalln("Error in writing to file")
		return
	}

	fmt.Println("Data has been saved!")
}

func readFromJson(fname string, pbf proto.Message) {
	fname = fmt.Sprintf("%s.json", fname)
	in, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatalln("Couldn't read file")
		return
	}

	err = protojson.Unmarshal(in, pbf)
	if err != nil {
		log.Fatalln("Couldn't unmarshal protocol buffer bytes")
		return
	}

	fmt.Println("Data bas been retrieved!")
}

func doFile(fname string) {
	objComplex := doComplex()
	writeToFile(fname, objComplex)

	readComplex := &pb.Complex{}
	ReadFromFile(fname, readComplex)

	fmt.Printf("Read Message %v", readComplex)
}

func doJson(fname string) {
	objComplex := doComplex()
	writeToJSON(fname, objComplex)

	readComplex := &pb.Complex{}
	readFromJson(fname, readComplex)

	fmt.Println(readComplex)
}

func main() {
	doJson("ComplexProto")
}
