package main

import (
	"log"

	"github.com/tigrannajaryan/exp-otelproto/grpc_protobuf_impl"

	"github.com/tigrannajaryan/exp-otelproto/core"
	"github.com/tigrannajaryan/exp-otelproto/traceprotobuf"
)

func main() {
	log.Println("GRPC/Protobuf Experiment.")

	core.RunTest(&grpc_protobuf_impl.Client{}, &grpc_protobuf_impl.Server{}, &traceprotobuf.Generator{})
}
