.PHONY: all gen build

all: gen build

gen:
	protoc -I/usr/local/include -I traceprotobuf/ traceprotobuf/trace.proto --go_out=plugins=grpc:traceprotobuf
	protoc -I/usr/local/include -I traceprotobuf/ traceprotobuf/resource.proto --go_out=plugins=grpc:traceprotobuf

build:
	go build -o bin/grpc-protobuf cmd/grpc-protobuf.go
	go build -o bin/grpc-protobuf-agent cmd/grpc-protobuf-agent.go

run:
	go run cmd/grpc-protobuf.go