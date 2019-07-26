package main

import (
	"flag"
	"fmt"

	"github.com/tigrannajaryan/exp-otelproto/core"
	"github.com/tigrannajaryan/exp-otelproto/grpc_stream"
	"github.com/tigrannajaryan/exp-otelproto/grpc_stream_lb"
	"github.com/tigrannajaryan/exp-otelproto/grpc_stream_lb_async"
	"github.com/tigrannajaryan/exp-otelproto/grpc_unary"
	"github.com/tigrannajaryan/exp-otelproto/traceprotobuf"
)

func main() {
	options := core.Options{}

	var protocol string
	flag.StringVar(&protocol, "protocol", "", "protocol to benchmark (unary,streamsync,streamlbsync,streamlbasync)")

	flag.IntVar(&options.Batches, "batches", 100, "total batches to send")
	flag.IntVar(&options.SpansPerBatch, "spansperbatch", 100, "spans per batch")
	flag.IntVar(&options.AttrPerSpan, "attrperspan", 2, "attributes per span")

	flag.Parse()

	switch protocol {
	case "unary":
		benchmarkGRPCUnary(options)
	case "streamsync":
		benchmarkGRPCStreamNoLB(options)
	case "streamlbtimedsync":
		benchmarkGRPCStreamLBTimedSync(options)
	case "streamlbalwayssync":
		benchmarkGRPCStreamLBAlwaysSync(options)
	case "streamlbasync":
		benchmarkGRPCStreamLBAsync(options)
	default:
		flag.Usage()
	}
}

func benchmarkGRPCUnary(options core.Options) {
	benchmarkImpl(
		"GRPC/Unary",
		options,
		func() core.Client { return &grpc_unary.Client{} },
		func() core.Server { return &grpc_unary.Server{} },
		func() core.Generator { return &traceprotobuf.Generator{} },
	)
}

func benchmarkGRPCStreamLBTimedSync(options core.Options) {
	benchmarkImpl(
		"GRPC/Stream/LBTimed/Sync",
		options,
		func() core.Client { return &grpc_stream_lb.Client{} },
		func() core.Server { return &grpc_stream_lb.Server{} },
		func() core.Generator { return &traceprotobuf.Generator{} },
	)
}

func benchmarkGRPCStreamLBAlwaysSync(options core.Options) {
	benchmarkImpl(
		"GRPC/Stream/LBAlways/Sync",
		options,
		func() core.Client { return &grpc_stream_lb.Client{ReopenAfterEveryRequest: true} },
		func() core.Server { return &grpc_stream_lb.Server{} },
		func() core.Generator { return &traceprotobuf.Generator{} },
	)
}

func benchmarkGRPCStreamLBAsync(options core.Options) {
	benchmarkImpl(
		"GRPC/Stream/LBTimed/Async",
		options,
		func() core.Client { return &grpc_stream_lb_async.Client{} },
		func() core.Server { return &grpc_stream_lb_async.Server{} },
		func() core.Generator { return &traceprotobuf.Generator{} },
	)
}

func benchmarkGRPCStreamNoLB(options core.Options) {
	benchmarkImpl(
		"GRPC/Stream/NoLB",
		options,
		func() core.Client { return &grpc_stream.Client{} },
		func() core.Server { return &grpc_stream.Server{} },
		func() core.Generator { return &traceprotobuf.Generator{} },
	)
}

func benchmarkImpl(
	name string,
	options core.Options,
	clientFactory func() core.Client,
	serverFactory func() core.Server,
	generatorFactory func() core.Generator,
) {
	cpuSecs, wallSecs := core.BenchmarkLocalDelivery(
		clientFactory,
		serverFactory,
		generatorFactory,
		options,
	)

	fmt.Printf("%-25s %5d batches, %4d spans/batch, %6d spans, CPU time %4.1f sec, wall time %4.1f sec, %4.0f batches/cpusec, %4.0f batches/wallsec\n",
		name,
		options.Batches,
		options.SpansPerBatch,
		options.Batches*options.SpansPerBatch,
		cpuSecs,
		wallSecs,
		float64(options.Batches)/cpuSecs,
		float64(options.Batches)/wallSecs)
}
