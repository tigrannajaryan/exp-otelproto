.PHONY: all gen build

all: gen build

gen:
	protoc -I/usr/local/include -I traceprotobuf/ traceprotobuf/trace.proto --go_out=plugins=grpc:traceprotobuf
	protoc -I/usr/local/include -I traceprotobuf/ traceprotobuf/resource.proto --go_out=plugins=grpc:traceprotobuf
	protoc -I/usr/local/include -I traceflatbuffers/ traceflatbuffers/trace.proto --go_out=plugins=grpc:traceflatbuffers
	flatc --gen-object-api --go traceflatbuffers/trace.fbs

build:
	go build -o bin/grpc_protobuf cmd/grpc_protobuf/main.go
	go build -o bin/grpc_protobuf_agent cmd/grpc_protobuf_agent/main.go
	go build -o bin/benchmark cmd/benchmark/main.go

benchmark:
	./runbenchmarks.sh

run:
	go run cmd/grpc-protobuf.go