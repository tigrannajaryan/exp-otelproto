package main

import (
	"log"

	"github.com/tigrannajaryan/exp-otelproto/core"
	"github.com/tigrannajaryan/exp-otelproto/grpcimpl"
	"github.com/tigrannajaryan/exp-otelproto/traceprotobuf"
)

func main() {
	log.Println("GRPC/Protobuf Experiment.")

	core.RunTest(&grpcimpl.Client{}, &grpcimpl.Server{}, &traceprotobuf.Generator{})
}
