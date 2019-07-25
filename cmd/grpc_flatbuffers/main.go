package main

import (
	"log"

	"github.com/tigrannajaryan/exp-otelproto/core"
	"github.com/tigrannajaryan/exp-otelproto/grpc_flatbuffers_impl"
	"github.com/tigrannajaryan/exp-otelproto/traceflatbuffers"
)

func main() {
	log.Println("GRPC/Flatbuffers Experiment.")

	core.RunTest(&grpc_flatbuffers_impl.Client{}, &grpc_flatbuffers_impl.Server{}, &traceflatbuffers.Generator{})
}
