package core

import (
	"log"
	"time"
)

func onBatchReceive(batch SpanBatch) {
	log.Printf("Server received a batch")
}

func Run(clnt Client, srv Server, gen Generator) {

	go srv.Listen("0.0.0.0:3465", onBatchReceive)

	// Hack: wait for serve to start.
	time.Sleep(time.Millisecond * 100)

	clnt.Connect("localhost:3465")

	batch := gen.GenerateBatch()

	clnt.Export(batch)
}
