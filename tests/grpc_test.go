package tests

import (
	"log"
	"testing"

	"github.com/tigrannajaryan/exp-otelproto/grpc_unary"

	"github.com/tigrannajaryan/exp-otelproto/grpc_stream"

	"github.com/tigrannajaryan/exp-otelproto/core"
	"github.com/tigrannajaryan/exp-otelproto/traceprotobuf"
)

func benchmarkImpl(
	name string,
	clientFactory func() core.Client,
	serverFactory func() core.Server,
	generatorFactory func() core.Generator,
	b *testing.B,
) {
	const batchCountInner = 60000
	batchCountTotal := 0
	cpuTime := 0.0
	wallTime := 0.0

	b.Run(name,
		func(b *testing.B) {
			// Reset counters at the beginning of test.
			cpuTime = 0
			wallTime = 0
			batchCountTotal = 0

			for i := 0; i < b.N; i++ {
				core.BenchmarkLocalDelivery(
					clientFactory,
					serverFactory,
					generatorFactory,
					b,
					func(cpuSecs float64, wallSecs float64) {
						cpuTime += cpuSecs
						wallTime += wallSecs
					},
					batchCountInner,
				)
				batchCountTotal += batchCountInner
			}
		})

	log.Printf("%10s: CPU time %.3f sec, wall time %.3f sec, total %d span batches, %.0f batches/cpusec\n",
		name, cpuTime, wallTime, batchCountTotal, float64(batchCountTotal)/cpuTime)
}

func BenchmarkGRPC(b *testing.B) {
	for i := 0; i < 5; i++ {

		benchmarkImpl("GRPCUnary",
			func() core.Client { return &grpc_unary.Client{} },
			func() core.Server { return &grpc_unary.Server{} },
			func() core.Generator { return &traceprotobuf.Generator{} },
			b,
		)
		benchmarkImpl("GRPCStream",
			func() core.Client { return &grpc_stream.Client{} },
			func() core.Server { return &grpc_stream.Server{} },
			func() core.Generator { return &traceprotobuf.Generator{} },
			b,
		)
		log.Printf("========")
	}
}
