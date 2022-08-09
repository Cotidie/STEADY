# Protocol Buffer
| https://developers.google.com/protocol-buffers

## Table of Contents
- [Overview](#overview)
  - [Setting](#setting)
- [Basics](#basics)
  - [Type](#type)
  - [Tag (Field Number)](#tag-field-number)
  - [Import](#import)
  - [Compiler](#compiler)
- [Golang](#golang)
  - [Packages](#packages)
  - [Compile](#compile)
- [Errors](#errors)
  - [go: could not create module cache](#go-could-not-create-module-cache)
  - [Cannot use <type> as type <type2> in struct literal](#cannot-use-type-as-type-type2-in-struct-literal)

## Overview
![Protocol Buffer](https://i.imgur.com/bBXCEpw.png)

Protocol Buffer is a format of data structure for serializing and deserializing suggested by Google. It has several advantages over JSON, XML etc. Protocol Buffer is commonly used in **gRPC**.
- Data is **strongly typed**
- Can have **comments** over schema
- Stored into **binary**, not string

### Setting
* ```vscode-proto3```: vscode extension for syntax highlighting .proto files
* ```Protoc```: .proto file compiler

### Data Evolution
Protocol Buffer provides backwards/forwards compatibility with tags, reserved keywords and defaults.
- Renaming a field will work as long as **the tag is intact**
- Removing a field will work if **the tag is reserved**
```proto
// before
message Profile {
  int32 id = 1;
  string name = 2;
}

// after
message Profile {
  reserved 2;
  reserved "name";  // optional, won't use the field 'name'
  
  int32 id = 1;
}
```

## Basics
```proto3
// Decide which protocol buffer syntax to use
syntax = "proto3";

// parent of /google directory should specified with -I option to the compiler
import "google/protobuf.timestamp.proto";

// what package name to use. variable name is reserved
option go_package = "proto";

// Enum
enum EyeColor {
    EYE_COLOR_UNSPECIFIED = 0;
    EYE_COLOR_GREEN = 1;
    EYE_COLOR_BROWN = 2;
}

// Define a message schema(ref. Class)
message User {
    // Nested Message
    message Status {}

    // <type> <field> = <field number>
    uint32 id = 1;
    EyeColor eye_color = 2;
    repeated string phones = 3;
    Status status = 4;
}
```
- it won't compile if there is a unpopulated field
  
### Type
Default value of Number types is 0. Put ```repeated``` keyword in the front to indicate an array type.
- **(s)int32/64**
- **uint32/64**
- **(s)fixed32/64**: fixed amount of bytes
- **float, double**
- **bool**: true/false. default to 0
- **string**: UTF-8/ASCII encoded string. default to ""
- **bytes**: raw data bytes. default to NULL
- **enum**: user defined type. default to its first field, with field number 0

### Tag (Field Number)
Field numbers are unique numbers that each field has as a binary identifier.
- range 1 to 15 takes one byte to encode
    => Use for frequently occuring data
- range 16 to 2047 takes two bytes.
- range 19000 to 19999 is reserved by compiler

### OneOf
OneOf type allows optional member which can be either of different types.
```proto
// proto file
message Result {
  oneof result {
    string message = 1;
    uint32 id = 2;
  }
}
```
```go
// go file
message := proto.Result_Message{Message: "msg"}
result := proto.Result{Result: &message}
```

### Import
Import path is relative to the project path or the workspace path and it must specify the full path of the filename including its extension(.proto). **Package** names can be defined as namespaces. 
- -I(--proto_path) can be specified multiple times
- proto files in the same directory should have the same package name
```proto3
// Project Structure
project
  |-- basics
    - account.proto
  |-- custom
    - import.proto

// account.proto
package basic;

message Account {}

// parent of /basics directory should specified with -I option to the compiler
import "basics/account.proto";

message Import {
  // If not defined package in account.proto, it should be just Account.
  basic.Account account = 1;
}
```

### Compiler
```bash
# 1. Find the latest release and download
# -O: Keeps file name as remote file name
# -L: Allows any redirections
curl -OL https://github.com/google/protobuf/releases/download/v3.5.1/protoc-3.5.1-linux-x86_64.zip

# 2. Unzip the file to ./protoc3 (-d option)
unzip protoc-3.5.1-linux-x86_64.zip -d protoc3

# 3. Add the path ~/protoc3 to $PATH
export PATH=$PATH:~/protoc3
```

## Golang
### Packages
```go
go get google.golang.org/protobuf/cmd/protoc-gen-go@latest
```

### Compile
```bash
# -I spectifies import path
# --go_out specifies where to reside output files
protoc -I proto --go_out proto address/*.proto
```

### Usage
```go
// import the directory (package name: proto)
import (
  pb "github.com/Cotidie/STEADY/gRPC/protocol_buffer/golang/proto"
)

// created protocol buffers are always pointer types to each other
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
    EyeColor:        pb.EyeColor_EYE_COLOR_GREEN,
	}
}
```

### I/O
- Disk
  - Write: Serialize(proto.Marshal) -> Save(ioutil.WriteFile)
  - Read: Read(ioutil.ReadFile) -> Deserialize(proto.Unmarshal)
- Json
  - Write: Serialize(protojson.Marshal)
  - Read: Deserialize(protojson.Unmarshal)


## Errors
### go: could not create module cache
```bash
protoc -I proto --go_opt=module=github.com/Cotidie/STEADY/gRPC/protocol_buffer/golang --go_out=. proto/*.proto
go build .
go: could not create module cache: stat /usr/bin/go/bin/pkg/mod: not a directory
go: could not create module cache: stat /usr/bin/go/bin/pkg/mod: not a directory
make: *** [Makefile:6: build] Error 1
```
- **Context**: When compiling golang file
- **Cause**: Go compiler automatically finds modules in ```$GOPATH/pkg/mod```, but this directory doesn't exist
- **Solve**: Check $GOPATH env variable and set it to the right directory

### Cannot use <type> as type <type2> in struct literal
```bash
protoc -I proto --go_opt=module=github.com/Cotidie/STEADY/gRPC/protocol_buffer/golang --go_out=. proto/*.proto
go build .
# github.com/Cotidie/STEADY/gRPC/protocol_buffer/golang
./main.go:29:20: cannot use dummy_one (variable of type "github.com/Cotidie/STEADY/gRPC/protocol_buffer/golang/proto".Dummy) as type *"github.com/Cotidie/STEADY/gRPC/protocol_buffer/golang/proto".Dummy in struct literal
./main.go:30:20: cannot use []pb.Dummy{…} (value of type []"github.com/Cotidie/STEADY/gRPC/protocol_buffer/golang/proto".Dummy) as type []*"github.com/Cotidie/STEADY/gRPC/protocol_buffer/golang/proto".Dummy in struct literal
make: *** [Makefile:6: build] Error 2
```
- **Context**: When creating a complex protocol buffer object
  - **struct literal** means newly allocated struct value itself
- **Cause**: the types for the struct fields don't match. Protoc always returns pointer types when importing another message type.
- **Solve**: Provide pointers to protocol buffer objects