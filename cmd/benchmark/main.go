package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
	"runtime/pprof"

	"github.com/tigrannajaryan/exp-otelproto/core"
	"github.com/tigrannajaryan/exp-otelproto/encodings/octraceprotobuf"
	"github.com/tigrannajaryan/exp-otelproto/encodings/traceprotobuf"
	"github.com/tigrannajaryan/exp-otelproto/protoimpls/grpc_oc"
	"github.com/tigrannajaryan/exp-otelproto/protoimpls/grpc_stream"
	"github.com/tigrannajaryan/exp-otelproto/protoimpls/grpc_stream_lb"
	"github.com/tigrannajaryan/exp-otelproto/protoimpls/grpc_stream_lb_async"
	"github.com/tigrannajaryan/exp-otelproto/protoimpls/grpc_unary"
	"github.com/tigrannajaryan/exp-otelproto/protoimpls/ws_stream_async"
	"github.com/tigrannajaryan/exp-otelproto/protoimpls/ws_stream_sync"
)

func main() {
	options := core.Options{}

	ballastSizeBytes := uint64(500) * 1024 * 1024
	ballast := make([]byte, ballastSizeBytes)

	var protocol string
	flag.StringVar(&protocol, "protocol", "",
		"protocol to benchmark (opencensus,ocack,unary,streamsync,streamlbtimedsync,streamlbalwayssync,streamlbasync,wsstreamsync,wsstreamasync,wsstreamasynczlib)")

	flag.IntVar(&options.Batches, "batches", 100, "total batches to send")
	flag.IntVar(&options.SpansPerBatch, "spansperbatch", 100, "spans per batch")
	flag.IntVar(&options.AttrPerSpan, "attrperspan", 2, "attributes per span")
	var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

	flag.Parse()

	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	switch protocol {
	case "opencensus":
		benchmarkGRPCOpenCensus(options)
	case "ocack":
		benchmarkGRPCOpenCensusWithAck(options)
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
	case "wsstreamsync":
		benchmarkWSStreamSync(options)
	case "wsstreamasync":
		benchmarkWSStreamAsync(options, traceprotobuf.CompressionMethod_NONE)
	case "wsstreamasynczlib":
		benchmarkWSStreamAsync(options, traceprotobuf.CompressionMethod_ZLIB)
	default:
		flag.Usage()
	}

	runtime.KeepAlive(ballast)
}

func benchmarkGRPCOpenCensus(options core.Options) {
	benchmarkImpl(
		"GRPC/OpenCensus",
		options,
		func() core.Client { return &grpc_oc.Client{} },
		func() core.Server { return &grpc_oc.Server{} },
		func() core.Generator { return octraceprotobuf.NewGenerator() },
	)
}

func benchmarkGRPCOpenCensusWithAck(options core.Options) {
	benchmarkImpl(
		"GRPC/OpenCensusWithAck",
		options,
		func() core.Client { return &grpc_oc.Client{WaitForAck: true} },
		func() core.Server { return &grpc_oc.Server{SendAck: true} },
		func() core.Generator { return octraceprotobuf.NewGenerator() },
	)
}

func benchmarkGRPCUnary(options core.Options) {
	benchmarkImpl(
		"GRPC/Unary",
		options,
		func() core.Client { return &grpc_unary.Client{} },
		func() core.Server { return &grpc_unary.Server{} },
		func() core.Generator { return traceprotobuf.NewGenerator() },
	)
}

func benchmarkGRPCStreamLBTimedSync(options core.Options) {
	benchmarkImpl(
		"GRPC/Stream/LBTimed/Sync",
		options,
		func() core.Client { return &grpc_stream_lb.Client{} },
		func() core.Server { return &grpc_stream_lb.Server{} },
		func() core.Generator { return traceprotobuf.NewGenerator() },
	)
}

func benchmarkGRPCStreamLBAlwaysSync(options core.Options) {
	benchmarkImpl(
		"GRPC/Stream/LBAlways/Sync",
		options,
		func() core.Client { return &grpc_stream_lb.Client{ReopenAfterEveryRequest: true} },
		func() core.Server { return &grpc_stream_lb.Server{} },
		func() core.Generator { return traceprotobuf.NewGenerator() },
	)
}

func benchmarkGRPCStreamLBAsync(options core.Options) {
	benchmarkImpl(
		"GRPC/Stream/LBTimed/Async",
		options,
		func() core.Client { return &grpc_stream_lb_async.Client{} },
		func() core.Server { return &grpc_stream_lb_async.Server{} },
		func() core.Generator { return traceprotobuf.NewGenerator() },
	)
}

func benchmarkGRPCStreamNoLB(options core.Options) {
	benchmarkImpl(
		"GRPC/Stream/NoLB",
		options,
		func() core.Client { return &grpc_stream.Client{} },
		func() core.Server { return &grpc_stream.Server{} },
		func() core.Generator { return traceprotobuf.NewGenerator() },
	)
}

func benchmarkWSStreamSync(options core.Options) {
	benchmarkImpl(
		"WebSocket/Stream/Sync",
		options,
		func() core.Client { return &ws_stream_sync.Client{} },
		func() core.Server { return &ws_stream_sync.Server{} },
		func() core.Generator { return traceprotobuf.NewGenerator() },
	)
}

func benchmarkWSStreamAsync(options core.Options, compression traceprotobuf.CompressionMethod) {
	var suffix string
	switch compression {
	case traceprotobuf.CompressionMethod_NONE:
		suffix = ""
	case traceprotobuf.CompressionMethod_ZLIB:
		suffix = "/zlib"
	case traceprotobuf.CompressionMethod_LZ4:
		suffix = "/lz4"
	}

	benchmarkImpl(
		"WebSocket/Stream/Async"+suffix,
		options,
		func() core.Client { return &ws_stream_async.Client{Compression: compression} },
		func() core.Server { return &ws_stream_async.Server{} },
		func() core.Generator { return traceprotobuf.NewGenerator() },
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

	fmt.Printf("%-27s %7d spans, CPU time %5.1f sec, wall time %5.1f sec, %4.1f batches/cpusec, %4.1f batches/wallsec\n",
		name,
		options.Batches*options.SpansPerBatch,
		cpuSecs,
		wallSecs,
		float64(options.Batches)/cpuSecs,
		float64(options.Batches)/wallSecs)
}
