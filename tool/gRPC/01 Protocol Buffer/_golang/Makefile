# $() means using variable, $(shell <command>) means executing a shell command
PACKAGE = $(shell head -1 go.mod | awk '{print $$2}')

# make build -> uses generate command
build: generate 
	go build .

# go_opt=module will strip out the module name from the package name
generate:
	protoc -I proto --go_opt=module=${PACKAGE} --go_out=. proto/*.proto

clean:
	rm proto/*.pb.go
	rm golang