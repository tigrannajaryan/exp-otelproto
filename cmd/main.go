package main

import (
	"fmt"
	"log"

	"github.com/tigrannajaryan/exp-otelproto/core"
	"github.com/tigrannajaryan/exp-otelproto/grpcimpl"
	"github.com/tigrannajaryan/exp-otelproto/tracerprotobuf"
)

func onBatchReceive(batch core.SpanBatch) {
	log.Printf("Server received a batch")
}

func main() {
	fmt.Println("OpenTelemetry Protocol Experiment")

	srv := &grpcimpl.Server{}
	go srv.Listen("0.0.0.0:3465", onBatchReceive)

	clnt := grpcimpl.Client{}
	clnt.Connect("localhost:3465")

	gen := &tracerprotobuf.Generator{}
	batch := gen.GenerateBatch()

	clnt.SendBatch(batch)
}
