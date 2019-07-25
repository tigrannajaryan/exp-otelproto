package main

import (
	"log"

	"github.com/tigrannajaryan/exp-otelproto/core"
	"github.com/tigrannajaryan/exp-otelproto/grpc_unary"
	"github.com/tigrannajaryan/exp-otelproto/traceprotobuf"
)

func main() {
	log.Println("GRPC/Protobuf Experiment.")

	core.RunTest(&grpc_unary.Client{}, &grpc_unary.Server{}, &traceprotobuf.Generator{})
}
