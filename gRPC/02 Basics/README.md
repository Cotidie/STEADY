# gRPC
![APIs](https://www.altexsoft.com/media/2020/05/word-image-53.png)
gRPC is a API architecture developed by Google. It uses **protocol buffers** and **HTTP/2** for communication.
- Faster than REST API (not stateless, binary encoded) 
- One TCP connection for streamings
- Can be bidirectional between Client and Server

## Basics
### Types
![types](https://miro.medium.com/max/1838/0*HWPuAyVrZndcocTF.png)
- **Unary**: Client and Server exchange one request and response
- **Server streaming**: Client sends one request and server returns one or more responses (HTTP/2)
  - ex) real-time communication
- **Client Streaming**: Client one or more requests 
  - ex) uploading, updating information..
- **Bi-directional Streaming**: Client and Server both can send multiple requests/responses

### Connection
#### Protocol Buffer
Protocol buffer supports easy implementation over gRPC. ```service``` object compiles into gRPC server interface.
```proto
service GreetService {
    // Unary Type
    rpc Greet (GreetRequest) returns (GreetResponse);
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
func doGreet(client pb.GreetServiceClient) {
  // No meaningful information is provided yet
	ctx := context.Background()
	request := pb.GreetRequest{FirstName: "Wonseok"}

  // Call Greet() procedure in back-end side
	res, err := client.Greet(ctx, &request)
	fmt.Printf("Response Received! %v\n", res.Result)
}

func main() {
  // Establish connection to Server using HTTP/2
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	defer conn.Close()

	// Creates gRPC client that implements APIs
	client := pb.NewGreetServiceClient(conn)
	doGreet(client)
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