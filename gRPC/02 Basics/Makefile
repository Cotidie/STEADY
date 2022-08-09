MODULE = $(shell head -1 go.mod | awk '{print $$2}')
PROTO_DIR = proto
GO_OUT = --go_out=. --go_opt=module=${MODULE}
GO_GRPC_OUT = --go-grpc_out=. --go-grpc_opt=module=${MODULE}

RM_F_CMD = rm -f
RM_RF_CMD = ${RM_F_CMD} -r

generate_greet: 
	protoc -I greet/proto \
		${GO_OUT} ${GO_GRPC_OUT} \
		greet/proto/*.proto

build:
	generate_greet
	go build .

clean: clean_greet

clean_greet:
	${RM_F_CMD} greet/${PROTO_DIR}/*.pb.go