.PHONY: all gen build

all: gen build

gen:
	protoc -I/usr/local/include -I encodings/traceprotobuf/ encodings/traceprotobuf/trace.proto --go_out=plugins=grpc:encodings/traceprotobuf
	protoc -I/usr/local/include -I encodings/traceprotobuf/ encodings/traceprotobuf/resource.proto --go_out=plugins=grpc:encodings/traceprotobuf

	protoc -I/usr/local/include -I encodings/octraceprotobuf/ encodings/octraceprotobuf/octrace.proto --go_out=plugins=grpc:encodings/octraceprotobuf
	protoc -I/usr/local/include -I encodings/octraceprotobuf/ encodings/octraceprotobuf/resource.proto --go_out=plugins=grpc:encodings/octraceprotobuf

	# FlatBuffers experiment is disabled since it does not provide functionality that we need.
	# protoc -I/usr/local/include -I encodings/traceflatbuffers/ encodings/traceflatbuffers/trace.proto --go_out=plugins=grpc:encodings/traceflatbuffers
	# flatc --gen-object-api --go encodings/traceflatbuffers/trace.fbs

build:
	go build -o bin/grpc_protobuf cmd/grpc_protobuf/main.go
	go build -o bin/grpc_protobuf_agent cmd/grpc_protobuf_agent/main.go
	go build -o bin/benchmark cmd/benchmark/main.go

benchmark:
	./runbenchmarks.sh

run:
	go run cmd/grpc-protobuf.go