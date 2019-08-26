package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
	"runtime/pprof"
	"strconv"
	"time"

	"github.com/tigrannajaryan/exp-otelproto/protoimpls/grpc_unary_async"

	"github.com/tigrannajaryan/exp-otelproto/protoimpls/grpc_stream_lb_srv"

	"github.com/tigrannajaryan/exp-otelproto/core"
	"github.com/tigrannajaryan/exp-otelproto/encodings/octraceprotobuf"
	"github.com/tigrannajaryan/exp-otelproto/encodings/otlp"
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

	ballastSizeBytes := uint64(4000) * 1024 * 1024
	ballast := make([]byte, ballastSizeBytes)

	var protocol string
	flag.StringVar(&protocol, "protocol", "",
		"protocol to benchmark (opencensus,ocack,unary,unaryasync,streamsync,streamlbtimedsync,"+
			"streamlbalwayssync,streamlbasync,streamlbconc,streamlbsrv,wsstreamsync,wsstreamasync,"+
			"wsstreamasynczlib)",
	)

	flag.IntVar(&options.Batches, "batches", 100, "total batches to send")
	flag.IntVar(&options.SpansPerBatch, "spansperbatch", 100, "spans per batch")
	flag.IntVar(&options.AttrPerSpan, "attrperspan", 2, "attributes per span")
	var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")
	var rebalancePeriodStr = flag.String("rebalance", "30s", "rebalance period (Valid time units are ns, us, ms, s, m, h)")
	var rebalanceRequestLimit = uint(*flag.Uint("rebalance-request", 1000, "rebalance after specified number of requests"))

	flag.Parse()

	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	rebalancePeriod, err := time.ParseDuration(*rebalancePeriodStr)
	if err != nil {
		log.Fatal(err)
	}

	switch protocol {
	case "opencensus":
		benchmarkGRPCOpenCensus(options)
	case "ocack":
		benchmarkGRPCOpenCensusWithAck(options)
	case "unary":
		benchmarkGRPCUnary(options)
	case "unaryasync":
		benchmarkGRPCUnaryAsync(options)
	case "streamsync":
		benchmarkGRPCStreamNoLB(options)
	case "streamlbtimedsync":
		benchmarkGRPCStreamLBTimedSync(options, rebalancePeriod)
	case "streamlbalwayssync":
		benchmarkGRPCStreamLBAlwaysSync(options)
	case "streamlbasync":
		benchmarkGRPCStreamLBAsync(options, rebalancePeriod, rebalanceRequestLimit, 1)
	case "streamlbconc":
		benchmarkGRPCStreamLBAsync(options, rebalancePeriod, rebalanceRequestLimit, 10)
	case "streamlbsrv":
		benchmarkGRPCStreamLBSrv(options, rebalancePeriod, rebalanceRequestLimit)
	case "wsstreamsync":
		benchmarkWSStreamSync(options)
	case "wsstreamasync":
		benchmarkWSStreamAsync(options, otlp.CompressionMethod_NONE)
	case "wsstreamasynczlib":
		benchmarkWSStreamAsync(options, otlp.CompressionMethod_ZLIB)
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
		func() core.Generator { return otlp.NewGenerator() },
	)
}

func benchmarkGRPCUnaryAsync(options core.Options) {
	benchmarkImpl(
		"GRPC/Unary/Async",
		options,
		func() core.Client { return &grpc_unary_async.Client{} },
		func() core.Server { return &grpc_unary_async.Server{} },
		func() core.Generator { return otlp.NewGenerator() },
	)
}

func benchmarkGRPCStreamLBTimedSync(options core.Options, streamReopenPeriod time.Duration) {
	benchmarkImpl(
		"GRPC/Stream/LBTimed/Sync",
		options,
		func() core.Client { return &grpc_stream_lb.Client{StreamReopenPeriod: streamReopenPeriod} },
		func() core.Server { return &grpc_stream_lb.Server{} },
		func() core.Generator { return otlp.NewGenerator() },
	)
}

func benchmarkGRPCStreamLBAlwaysSync(options core.Options) {
	benchmarkImpl(
		"GRPC/Stream/LBAlways/Sync",
		options,
		func() core.Client { return &grpc_stream_lb.Client{ReopenAfterEveryRequest: true} },
		func() core.Server { return &grpc_stream_lb.Server{} },
		func() core.Generator { return otlp.NewGenerator() },
	)
}

func benchmarkGRPCStreamLBAsync(
	options core.Options,
	streamReopenPeriod time.Duration,
	rebalanceRequestLimit uint,
	concurrency int,
) {
	benchmarkImpl(
		"GRPC/Stream/LBTimed/Async/"+strconv.Itoa(concurrency),
		options,
		func() core.Client {
			return &grpc_stream_lb_async.Client{
				Concurrency:              concurrency,
				StreamReopenPeriod:       streamReopenPeriod,
				StreamReopenRequestCount: uint32(rebalanceRequestLimit),
			}
		},
		func() core.Server { return &grpc_stream_lb_async.Server{} },
		func() core.Generator { return otlp.NewGenerator() },
	)
}

func benchmarkGRPCStreamLBSrv(options core.Options, streamReopenPeriod time.Duration, rebalanceRequestLimit uint) {
	benchmarkImpl(
		"GRPC/Stream/LBSrv/Async",
		options,
		func() core.Client { return &grpc_stream_lb_srv.Client{} },
		func() core.Server {
			return &grpc_stream_lb_srv.Server{
				StreamReopenPeriod:       streamReopenPeriod,
				StreamReopenRequestCount: rebalanceRequestLimit,
			}
		},
		func() core.Generator { return otlp.NewGenerator() },
	)
}

func benchmarkGRPCStreamNoLB(options core.Options) {
	benchmarkImpl(
		"GRPC/Stream/NoLB",
		options,
		func() core.Client { return &grpc_stream.Client{} },
		func() core.Server { return &grpc_stream.Server{} },
		func() core.Generator { return otlp.NewGenerator() },
	)
}

func benchmarkWSStreamSync(options core.Options) {
	benchmarkImpl(
		"WebSocket/Stream/Sync",
		options,
		func() core.Client { return &ws_stream_sync.Client{} },
		func() core.Server { return &ws_stream_sync.Server{} },
		func() core.Generator { return otlp.NewGenerator() },
	)
}

func benchmarkWSStreamAsync(options core.Options, compression otlp.CompressionMethod) {
	var suffix string
	switch compression {
	case otlp.CompressionMethod_NONE:
		suffix = ""
	case otlp.CompressionMethod_ZLIB:
		suffix = "/zlib"
	case otlp.CompressionMethod_LZ4:
		suffix = "/lz4"
	}

	benchmarkImpl(
		"WebSocket/Stream/Async"+suffix,
		options,
		func() core.Client { return &ws_stream_async.Client{Compression: compression} },
		func() core.Server { return &ws_stream_async.Server{} },
		func() core.Generator { return otlp.NewGenerator() },
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

	fmt.Printf("%-28s %7d spans, CPU time %5.1f sec, wall time %5.1f sec, %4.1f batches/cpusec, %4.1f batches/wallsec\n",
		name,
		options.Batches*options.SpansPerBatch,
		cpuSecs,
		wallSecs,
		float64(options.Batches)/cpuSecs,
		float64(options.Batches)/wallSecs)
}
