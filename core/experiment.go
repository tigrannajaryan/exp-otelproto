package core

import (
	"log"
)

func onBatchReceive(batch SpanBatch) {
	log.Printf("Server received a batch")
}

func RunTest(clnt Client, srv Server, gen Generator) {

	// Connect to Agent
	clnt.Connect("localhost:3465")

	// Listen locally for Agent's forwarded data
	go srv.Listen("0.0.0.0:4848", onBatchReceive)

	// Generate and send a batch
	batch := gen.GenerateBatch()
	clnt.Export(batch)
}

func RunAgent(clnt Client, srv Server, listenAddress, destination string) {

	log.Printf("Agent: listening on %s", listenAddress)
	log.Printf("Agent: forwarding to %s", destination)

	err := clnt.Connect(destination)
	if err != nil {
		log.Fatalf("Cannot connection to %v: %v", destination, err)
	}

	srv.Listen(listenAddress, func(batch SpanBatch) {
		log.Printf("Agent: forwarding span batch")
		clnt.Export(batch)
	})
}
