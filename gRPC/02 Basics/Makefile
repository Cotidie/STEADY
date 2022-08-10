# Syntax
# - $(), ${}: Use variable and shell commands
# - $@: the name of the target (generate_greet, build...)
# - := simple assignment. only evaluate when first encountered
# .PHONY: Cut association with physical files/directories

MODULE := $(shell head -1 go.mod | awk '{print $$2}')

# Directory Names
BIN_DIR := bin
PROTO_DIR := proto
CLIENT_DIR := client
SERVER_DIR := server

# Protoc Options
GO_OPT = --go_out=. --go_opt=module=$(MODULE)
GO_GRPC_OPT = --go-grpc_out=. --go-grpc_opt=module=${MODULE}

# Shell Commands
RM_F_CMD = rm -f
RM_RF_CMD = ${RM_F_CMD} -r

PROJECTS = greet

.PHONY: greet

all: ${PROJECTS}

${PROJECTS}: 
	protoc -I $@/${PROTO_DIR} \
		${GO_OPT} ${GO_GRPC_OPT} \
		$@/${PROTO_DIR}/*.proto
	go build -o ./$(BIN_DIR)/$@/${SERVER_DIR} ./greet/${SERVER_DIR}
	go build -o ./$(BIN_DIR)/$@/${CLIENT_DIR} ./greet/${CLIENT_DIR}

clean: clean_greet clean_server clean_client
	${RM_RF_CMD} ${BIN_DIR}/

clean_greet:
	${RM_F_CMD} greet/${PROTO_DIR}/*.pb.go

clean_server:
	${RM_F_CMD} server/${PROTO_DIR}/*.pb.go

clean_client:
	${RM_F_CMD} greet/${PROTO_DIR}/*.pb.go