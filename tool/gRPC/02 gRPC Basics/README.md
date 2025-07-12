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
    - [Server Streaming](#server-streaming)
    - [Client Streaming](#client-streaming)
    - [Bi-directional](#bi-directional)
  - [Features](#features)
    - [Error Handling](#error-handling)
    - [Deadline](#deadline)
    - [SSL Security](#ssl-security)
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

### Initialize
#### Protocol Buffer
Protocol buffer supports easy implementation over gRPC. ```service``` object compiles into gRPC server interface.
```proto
service GreetService {
    // Unary Type
    rpc Greet (GreetRequest) returns (GreetResponse);
	// Server Streaming Type
	rpc GreetManyTImes (GreetRequest) returns (stream GreetResponse);
	// Client Streaming Type
	rpc GreetLong (stream GreetRequest) returns (GreetResponse);
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
- Server ends streaming with return
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
	// Close the channel with sending io.EOF
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
#### Server
```go
func (s *Server) GreetLong(stream pb.GreetService_GreetLongServer) error {
	for {
		msg, err := stream.Recv()
		// Client closed streaming
		if err == io.EOF {
			response := pb.GreetResponse{Result: fmt.Sprintf("%v", names)}
			return stream.SendAndClose(&response)
		}
		// msg.Message...
	}
}
```
#### Client
```go
func doGreetLong(client pb.GreetServiceClient) {
	reqs := []*pb.GreetRequest{...	}
	ctx := context.Background()
	stream, err := client.GreetLong(ctx)

	for _, req := range reqs {
		stream.Send(req)
	}
	res, err := stream.CloseAndRecv()
}
```

### Bi-directional
- Use goroutine to send/receive simultaneously
- Use channel to wait for the execution

#### Server
```go
func (s *Server) GreetEveryone(stream pb.GreetService_GreetEveryoneServer) error {
	names := []string{}
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		}
		names = append(names, req.FirstName)
	}

	response := pb.GreetResponse{}
	for _, name := range names {
		response.Result = "Hello, " + name
		err := stream.Send(&response)
	}

	return nil
}
```

#### Client
```go
func doGreetEveryone(client pb.GreetServiceClient) {
	// Open stream
	ctx := context.Background()
	stream, err := client.GreetEveryone(ctx)
	// Channel to wait for goroutine to be finished
	waitc := make(chan struct{})

	// Stream to server
	go func() {
		for _, req := range reqs {
			stream.Send(req)
		}
		stream.CloseSend()
	}()

	// Receive from server
	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			// res.Result...
		}
		// wake up main goroutine
		close(waitc)
	}()
	// Wait for all the goroutines to be executed
	<-waitc
}
```

## Features
### Error Handling
 Clients can receive errors on server-side produced with gRPC package. ```codes``` package is for gRPC error codes, ```status``` is for gRPC error instances. ```error``` types can be parsed into gRPC errors with ```status.FromError()``` method.
 - Complete list of error codes can be found [here](https://grpc.github.io/grpc/core/md_doc_statuscodes.html)
 - Client can set deadline with context
```go
// Server Side: returns an error if input is negative
func (s *Server) Sqrt(ctx context.Context, req *pb.SqrtRequest) (*pb.SqrtResponse, error) {
	number := req.Number
	if number < 0 {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Received a negative number: %d", number),
		)
	}

	// Check deadline
	go func() {
		for {
			if ctx.Err() == context.DeadlineExceeded {...}
		}	
	}()

	// ....
}

// Client Side
func doSqrt(client pb.CalculatorServiceClient) {
	ctx := context.Background()
	req := pb.SqrtRequest{Number: -5}
	res, err := client.Sqrt(ctx, &req)

	if err != nil {
		// ok bool represents whether the error is produced from grpc package
		// e.Message(), e.Code()
		e, ok := status.FromError(err)
		if !ok {
			log.Fatalf("A non gRPC error: %v\n", err)
		}
		return
	}
	// ...
}
```

## Error
### make: ... is up to date
```bash
.PHONY: greet

# won't work if there's a physical file named with 'greet'
greet: ...
```
- **Cause**: There is a physical file/directory name same with the name of a command
- **Solve**: Use ```.PHONY``` keyword to cut association with physical files