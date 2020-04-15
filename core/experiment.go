package core

import (
	"fmt"
	"log"
	"net"
	"os"
	"sync"
	"time"

	"github.com/shirou/gopsutil/process"
)

func onBatchReceive(batch ExportRequest, spanCount int) {
	log.Printf("Server received a batch of %v spans", spanCount)
}

func RunTest(clnt Client, srv Server, gen SpanGenerator) {

	// Listen locally for Agent's forwarded data
	go srv.Listen("0.0.0.0:3465", onBatchReceive)

	// Connect to Agent
	clnt.Connect("localhost:3465")

	// Generate and send a batch
	for i := 0; i < 2; i++ {
		batch := gen.GenerateSpanBatch(100, 2, 0)
		clnt.Export(batch)
	}
}

type Options struct {
	Batches       int
	SpansPerBatch int
	AttrPerSpan   int
}

func BenchmarkLocalDelivery(
	clientFactory func() Client,
	serverFactory func() Server,
	generatorFactory func() SpanGenerator,
	options Options,
) (cpuSecs float64, wallSecs float64) {
	// Create client, server and generator from factories
	clnt := clientFactory()
	srv := serverFactory()
	gen := generatorFactory()

	// Find a local address for delivery.
	endpoint := GetAvailableLocalAddress()

	// Create a WaitGroup to count sent/received Batches.
	wg := sync.WaitGroup{}

	// Server listen locally.
	go srv.Listen(endpoint, func(batch ExportRequest, spanCount int) {
		// Count delivered batch.
		wg.Done()
	})

	// Client connect to the server.
	if err := clnt.Connect(endpoint); err != nil {
		log.Fatalf("cannot connect: %v", err)
	}

	// Begin measuring CPU time.
	proc, err := process.NewProcess(int32(os.Getpid()))
	if err != nil {
		log.Fatal(err)
	}

	batchCount := options.Batches
	const maxBatchesToRotate = 3000
	if batchCount > maxBatchesToRotate {
		// We will use at most 100 batches and will rotate through them when sending.
		// This is a hacky solution to limit the amount of memory we use for long running
		// tests that send many thousands of batches.
		// Note that protocols may modify the batch that is being sent (primarily the
		// request ID field) so we cannot use the same batch for sending while it is
		// still being sent by previously issued "Export" call.
		// Limiting the number of batches in memory only works because concurrency of
		// protocols is limited. This means that a batch will be reused after its
		// sending was completed.
		// Note that this is not a guarantee of correctness and may still fail
		// (and protocols may detect and fail when this happens). This approach should
		// never be used in a production code, but is acceptable for a benchmark.
		batchCount = maxBatchesToRotate
	}

	var batches []ExportRequest
	for i := 0; i < batchCount; i++ {
		batches = append(batches, gen.GenerateSpanBatch(options.SpansPerBatch, options.AttrPerSpan, 0))
	}

	startCPUTimes, err := proc.Times()
	if err != nil {
		log.Fatal(err)
	}

	// Begin measuring wall time.
	startWallTime := time.Now()

	// Generate and send Batches.
	for i := 0; i < options.Batches; i++ {
		// Count sent batch.
		wg.Add(1)

		batch := batches[i%batchCount]
		//batch := gen.GenerateSpanBatch(options.SpansPerBatch, options.AttrPerSpan, 0)

		clnt.Export(batch)
	}

	// Wait until all Batches are delivered.
	wg.Wait()

	clnt.Shutdown()

	// Measure used CPU time.
	endCPUTimes, err := proc.Times()
	if err != nil {
		log.Fatal(err)
	}
	cpuSecs = endCPUTimes.Total() - startCPUTimes.Total()

	// Measure used wall time.
	endWallTime := time.Now()
	wallSecs = endWallTime.Sub(startWallTime).Seconds()

	// Stop the server.
	srv.Stop()

	return
}

func LoadGenerator(
	clientFactory func() Client,
	generatorFactory func() SpanGenerator,
	serverEndpoint string,
	spansPerSecond int,
) {
	// Create client, server and generator from factories
	clnt := clientFactory()
	gen := generatorFactory()

	// Client connect to the server.
	clnt.Connect(serverEndpoint)

	// Generate and send Batches.
	totalSpans := 0
	for {
		startTime := time.Now()
		ch := time.After(1 * time.Second)
		batch := gen.GenerateSpanBatch(spansPerSecond, 10, 0)
		clnt.Export(batch)
		<-ch
		wallSecs := time.Now().Sub(startTime).Seconds()
		totalSpans += spansPerSecond
		actualSpansPerSecond := float64(spansPerSecond) / wallSecs
		fmt.Printf("Total spans sent %v, current rate %.1f spans/sec\n", totalSpans, actualSpansPerSecond)
	}
}

func RunServer(srv Server, listenAddress string, onReceive func(spanCount int)) {

	log.Printf("Server: listening on %s", listenAddress)

	totalSpans := 0
	prevTime := time.Now()

	srv.Listen(listenAddress, func(batch ExportRequest, spanCount int) {
		t := time.Now()
		d := t.Sub(prevTime)
		prevTime = t

		rate := float64(spanCount) / d.Seconds()

		totalSpans += spanCount
		log.Printf("Server: total spans received %v, current rate %.1f", totalSpans, rate)

		onReceive(spanCount)
	})
}

func RunAgent(clnt Client, srv Server, listenAddress, destination string) {

	log.Printf("Agent: listening on %s", listenAddress)
	log.Printf("Agent: forwarding to %s", destination)

	err := clnt.Connect(destination)
	if err != nil {
		log.Fatalf("Cannot connection to %v: %v", destination, err)
	}

	srv.Listen(listenAddress, func(batch ExportRequest, spanCount int) {
		log.Printf("Agent: forwarding %d span batch", spanCount)
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
