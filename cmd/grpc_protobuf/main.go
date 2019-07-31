package main

import (
	"log"

	"github.com/tigrannajaryan/exp-otelproto/core"
	"github.com/tigrannajaryan/exp-otelproto/encodings/traceprotobuf"
	"github.com/tigrannajaryan/exp-otelproto/protoimpls/grpc_unary"
)

func main() {
	log.Println("GRPC/Protobuf Experiment.")

	core.RunTest(&grpc_unary.Client{}, &grpc_unary.Server{}, &traceprotobuf.Generator{})
}
