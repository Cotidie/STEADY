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

## Error
### make: ... is up to date
```bash
.PHONY: greet

# won't work if there's a physical file named with 'greet'
greet: ...
```
- **Cause**: There is a physical file/directory name same with the name of a command
- **Solve**: Use ```.PHONY``` keyword to cut association with physical files