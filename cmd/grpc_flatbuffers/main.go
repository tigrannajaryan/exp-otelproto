package main

import (
	"log"

	"github.com/tigrannajaryan/exp-otelproto/core"
	"github.com/tigrannajaryan/exp-otelproto/encodings/traceflatbuffers"
	"github.com/tigrannajaryan/exp-otelproto/protoimpls/grpc_flatbuffers_impl"
)

func main() {
	log.Println("GRPC/Flatbuffers Experiment.")

	core.RunTest(&grpc_flatbuffers_impl.Client{}, &grpc_flatbuffers_impl.Server{}, &traceflatbuffers.Generator{})
}
