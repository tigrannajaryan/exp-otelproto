package core

import (
	"log"
	"net"
	"os"
	"sync"
	"testing"
	"time"

	"github.com/shirou/gopsutil/process"
)

func onBatchReceive(batch SpanBatch) {
	log.Printf("Server received a batch")
}

func RunTest(clnt Client, srv Server, gen Generator) {

	// Listen locally for Agent's forwarded data
	go srv.Listen("0.0.0.0:3465", onBatchReceive)

	// Connect to Agent
	clnt.Connect("localhost:3465")

	// Generate and send a batch
	for i := 0; i < 2; i++ {
		batch := gen.GenerateBatch()
		clnt.Export(batch)
	}
}

func BenchmarkLocalDelivery(
	clientFactory func() Client,
	serverFactory func() Server,
	generatorFactory func() Generator,
	b *testing.B,
	addTime func(cpuSecs float64, wallSecs float64),
	batchCount int,
) {
	// Stop benchmark timer while setting up the test.
	b.StopTimer()

	// Create client, server and generator from factories
	clnt := clientFactory()
	srv := serverFactory()
	gen := generatorFactory()

	// Find a local address for delivery.
	endpoint := GetAvailableLocalAddress()

	// Create a WaitGroup to count sent/received batches.
	wg := sync.WaitGroup{}

	// Server listen locally.
	go srv.Listen(endpoint, func(batch SpanBatch) {
		// Count delivered batch.
		wg.Done()
	})

	// Clietn connect to the server.
	clnt.Connect(endpoint)

	// Begin measuring CPU time.
	proc, err := process.NewProcess(int32(os.Getpid()))
	if err != nil {
		log.Fatal(err)
	}

	startCPUTimes, err := proc.Times()
	if err != nil {
		log.Fatal(err)
	}

	// Begin measuring wall time.
	startWallTime := time.Now()

	// Restart benchmark timer.
	b.StartTimer()

	// Generate and send batches.
	for i := 0; i < batchCount; i++ {
		// Count sent batch.
		wg.Add(1)
		batch := gen.GenerateBatch()
		clnt.Export(batch)
	}

	// Wait until all batches are delivered.
	wg.Wait()

	// Stop benachmark timer.
	b.StopTimer()

	// Measure used CPU time.
	endCPUTimes, err := proc.Times()
	if err != nil {
		log.Fatal(err)
	}
	deltaCPUTime := endCPUTimes.Total() - startCPUTimes.Total()

	// Measure used wall time.
	endWallTime := time.Now()
	deltaWallTime := endWallTime.Sub(startWallTime)

	// Report used times.
	addTime(deltaCPUTime, deltaWallTime.Seconds())

	// Stop the server.
	srv.Stop()
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

// GetAvailableLocalAddress finds an available local port and returns an endpoint
// describing it. The port is available for opening when this function returns
// provided that there is no race by some other code to grab the same port
// immediately.
func GetAvailableLocalAddress() string {
	ln, err := net.Listen("tcp", ":0")
	if err != nil {
		log.Fatalf("failed to get a free local port: %v", err)
	}
	// There is a possible race if something else takes this same port before
	// the test uses it, however, that is unlikely in practice.
	defer ln.Close()
	return ln.Addr().String()
}
