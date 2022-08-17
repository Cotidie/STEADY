module github.com/Cotidie/STEADY/gRPC/Basics

go 1.18

require (
	google.golang.org/grpc v1.48.0
	google.golang.org/protobuf v1.28.1
)

require (
	github.com/golang/protobuf v1.5.2 // indirect
	golang.org/x/net v0.0.0-20220812174116-3211cb980234 // indirect
	golang.org/x/sys v0.0.0-20220811171246-fbc7d0a398ab // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20220815135757-37a418bb8959 // indirect
)

replace github.com/Cotidie/STEADY/gRPC/Basics/calculator/client => ./calculator/client
