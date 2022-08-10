# gRPC
![APIs](https://www.altexsoft.com/media/2020/05/word-image-53.png)  
gRPC is a API architecture developed by Google. It uses **protocol buffers** and **HTTP/2** for communication.
- Faster than REST API (not stateless, binary encoded) 
- One TCP connection for streamings
- Can be bidirectional between Client and Server

## Table of Contents
- [Basics](#basics)
	- [Types](#types)
	- [Implementation](#implementation)
		- [Protocol Buffer](#protocol-buffer)
		- [Server](#server)
		- [Client](#client)
- [Implementation](#implementation-1)
	- [Unary](#unary)
		- [Server](#server-1)
		- [Client](#client-1)
	- [Server Streaming](#server-streaming)
		- [Server](#server-2)
		- [Client](#client-2)
	- [Client Streaming](#client-streaming)
    - [Bi-directional](#bi-directional)
- [Error](#error)
	- [make: ... is up to date](#make--is-up-to-date)

## Basics
### Types
![types](https://miro.medium.com/max/1838/0*HWPuAyVrZndcocTF.png)
- **Unary**: Client and Server exchange one request and response
- **Server streaming**: Client sends one request and server returns one or more responses (HTTP/2)
  - ex) real-time communication
- **Client Streaming**: Client one or more requests 
  - ex) uploading, updating information..
- **Bi-directional Streaming**: Client and Server both can send multiple requests/responses

### Implementation
#### Protocol Buffer
Protocol buffer supports easy implementation over gRPC. ```service``` object compiles into gRPC server interface.
```proto
service GreetService {
    // Unary Type
    rpc Greet (GreetRequest) returns (GreetResponse);
	// Server Streaming Type
	rpc GreetManyTImes (GreetRequest) returns (stream GreetResponse);
}
```

#### Server
APIs provided by service interface should be implemented somewhere
```go
// pb: compiled package of proto file above
type Server struct {
  pb.GreetServiceServer
}

func main() {
 	// Listen establishes Transport and Session layers
	lis, err := net.Listen("tcp", addr)
	// gRPC Server establishes Presentation and Application layers
	s := grpc.NewServer()
	// Register API to gRPC server instance
	pb.RegisterGreetServiceServer(s, &Server{})
	err = s.Serve(lis)
}
```

#### Client
```go
func main() {
  	// Establish connection to Server using HTTP/2
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	defer conn.Close()

	// Creates gRPC client that implements APIs
	client := pb.NewGreetServiceClient(conn)
	doGreet(client)
}
```

## Implementation
Each function signiture can be found on .pb.go file
### Unary
#### Server
```go
// Context has metainfo from the client 
func (s *Server) Greet(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	return &pb.GreetResponse{
		Result: "Hello " + in.FirstName,
	}, nil
}
```

#### Client
```go
func doGreet(client pb.GreetServiceClient) {
  	// No meaningful information is provided yet
	ctx := context.Background()
	request := pb.GreetRequest{FirstName: "Wonseok"}

  	// Call Greet() procedure in back-end side
	res, err := client.Greet(ctx, &request)
	fmt.Printf("Response Received! %v\n", res.Result)
}
```
### Server Streaming
- Client waits if there's a delay in server-side
- Client uses buffer if there's a delay in client-side
#### Server
```go
func (s *Server) GreetManyTImes(in *pb.GreetRequest, stream pb.GreetService_GreetManyTImesServer) error {
	firstName := in.FirstName

	for i := 1; i <= 10; i++ {
		response := pb.GreetResponse{
			Result: fmt.Sprintf("Hello %s, number: %d", firstName, i),
		}
		err := stream.Send(&response)
	}
	return nil
}
```

#### Client
```go
func doGreetManyTimes(client pb.GreetServiceClient) {
	ctx := context.Background()
	request := pb.GreetRequest{FirstName: "Wonseok"}
	stream, err := client.GreetManyTImes(ctx, &request)

	for {
		// Receive() works like a channel
		msg, err := stream.Recv()
		// io.EOF means end of streaming
		if err == io.EOF {
			break
		}
		fmt.Println(msg.Result)
	}
}
```
### Client Streaming

### Bi-directional

## Error
### make: ... is up to date
```bash
.PHONY: greet

# won't work if there's a physical file named with 'greet'
greet: ...
```
- **Cause**: There is a physical file/directory name same with the name of a command
- **Solve**: Use ```.PHONY``` keyword to cut association with physical files