package main

import (
	"flag"
	"fmt"
	ws_async_worker "github.com/tigrannajaryan/exp-otelproto/protoimpls/ws_async_workers"
	"log"
	"os"
	"runtime"
	"runtime/pprof"
	"strconv"
	"time"

	"github.com/tigrannajaryan/exp-otelproto/encodings/experimental"
	sapmenc "github.com/tigrannajaryan/exp-otelproto/encodings/sapm"

	"github.com/tigrannajaryan/exp-otelproto/protoimpls/sapm"

	"github.com/tigrannajaryan/exp-otelproto/core"
	"github.com/tigrannajaryan/exp-otelproto/encodings/octraceprotobuf"
	"github.com/tigrannajaryan/exp-otelproto/encodings/otlp"
	"github.com/tigrannajaryan/exp-otelproto/protoimpls/grpc_oc"
	"github.com/tigrannajaryan/exp-otelproto/protoimpls/grpc_stream"
	"github.com/tigrannajaryan/exp-otelproto/protoimpls/grpc_stream_lb"
	"github.com/tigrannajaryan/exp-otelproto/protoimpls/grpc_stream_lb_async"
	"github.com/tigrannajaryan/exp-otelproto/protoimpls/grpc_stream_lb_srv"
	"github.com/tigrannajaryan/exp-otelproto/protoimpls/grpc_unary"
	"github.com/tigrannajaryan/exp-otelproto/protoimpls/grpc_unary_async"
	"github.com/tigrannajaryan/exp-otelproto/protoimpls/http11"
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
			"streamlbalwayssync,streamlbasync,streamlbconc,streamlbsrv,wsstreamsync,wsstreamasync,wsstreamasyncconc,"+
			"wsstreamasynczlib,http11,http11conc,sapm,wsasyncworker,wsasyncworkerconc)",
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

	concurrency := 4 // runtime.GOMAXPROCS(0)
	if concurrency < 1 {
		concurrency = 1
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
		benchmarkGRPCStreamLBAsync(options, rebalancePeriod, rebalanceRequestLimit, concurrency)
	case "streamlbsrv":
		benchmarkGRPCStreamLBSrv(options, rebalancePeriod, rebalanceRequestLimit)
	case "wsstreamsync":
		benchmarkWSStreamSync(options, experimental.CompressionMethod_NONE)
	case "wsstreamsynczlib":
		benchmarkWSStreamSync(options, experimental.CompressionMethod_ZLIB)
	case "wsstreamsynczstd":
		benchmarkWSStreamSync(options, experimental.CompressionMethod_ZSTD)
	case "wsstreamasync":
		benchmarkWSStreamAsync(options, experimental.CompressionMethod_NONE, 1)
	case "wsasyncworker":
		benchmarkWSAsyncWorker(options, 1)
	case "wsasyncworkerconc":
		benchmarkWSAsyncWorker(options, concurrency)
	case "wsstreamasyncconc":
		benchmarkWSStreamAsync(options, experimental.CompressionMethod_NONE, concurrency)
	case "wsstreamasynczlib":
		benchmarkWSStreamAsync(options, experimental.CompressionMethod_ZLIB, 1)
	case "http11":
		benchmarkHttp11(options, 1)
	case "http11conc":
		benchmarkHttp11(options, concurrency)
	case "sapm":
		benchmarkSAPM(options, 1)
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
		func() core.SpanGenerator { return octraceprotobuf.NewGenerator() },
	)
}

func benchmarkGRPCOpenCensusWithAck(options core.Options) {
	benchmarkImpl(
		"GRPC/OpenCensusWithAck",
		options,
		func() core.Client { return &grpc_oc.Client{WaitForAck: true} },
		func() core.Server { return &grpc_oc.Server{SendAck: true} },
		func() core.SpanGenerator { return octraceprotobuf.NewGenerator() },
	)
}

func benchmarkGRPCUnary(options core.Options) {
	benchmarkImpl(
		"OTLP/GRPC-Unary/Sequential",
		options,
		func() core.Client { return &grpc_unary.Client{} },
		func() core.Server { return &grpc_unary.Server{} },
		func() core.SpanGenerator { return otlp.NewGenerator() },
	)
}

func benchmarkGRPCUnaryAsync(options core.Options) {
	benchmarkImpl(
		"OTLP/GRPC-Unary/Concurrent",
		options,
		func() core.Client { return &grpc_unary_async.Client{} },
		func() core.Server { return &grpc_unary_async.Server{} },
		func() core.SpanGenerator { return otlp.NewGenerator() },
	)
}

func benchmarkGRPCStreamLBTimedSync(options core.Options, streamReopenPeriod time.Duration) {
	benchmarkImpl(
		"GRPC/Stream/LBTimed/Sync",
		options,
		func() core.Client { return &grpc_stream_lb.Client{StreamReopenPeriod: streamReopenPeriod} },
		func() core.Server { return &grpc_stream_lb.Server{} },
		func() core.SpanGenerator { return experimental.NewGenerator() },
	)
}

func benchmarkGRPCStreamLBAlwaysSync(options core.Options) {
	benchmarkImpl(
		"GRPC/Stream/LBAlways/Sync",
		options,
		func() core.Client { return &grpc_stream_lb.Client{ReopenAfterEveryRequest: true} },
		func() core.Server { return &grpc_stream_lb.Server{} },
		func() core.SpanGenerator { return experimental.NewGenerator() },
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
		func() core.SpanGenerator { return experimental.NewGenerator() },
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
		func() core.SpanGenerator { return experimental.NewGenerator() },
	)
}

func benchmarkGRPCStreamNoLB(options core.Options) {
	benchmarkImpl(
		"GRPC/Stream/NoLB",
		options,
		func() core.Client { return &grpc_stream.Client{} },
		func() core.Server { return &grpc_stream.Server{} },
		func() core.SpanGenerator { return experimental.NewGenerator() },
	)
}

func getCompSuffix(compression experimental.CompressionMethod) string {
	var suffix string
	switch compression {
	case experimental.CompressionMethod_NONE:
		suffix = ""
	case experimental.CompressionMethod_ZLIB:
		suffix = "/zlib"
	case experimental.CompressionMethod_LZ4:
		suffix = "/lz4"
	case experimental.CompressionMethod_ZSTD:
		suffix = "/zstd"
	}
	return suffix
}

func benchmarkWSStreamSync(options core.Options, compression experimental.CompressionMethod) {
	benchmarkImpl(
		"WebSocket/Stream/Sync"+getCompSuffix(compression),
		options,
		func() core.Client { return &ws_stream_sync.Client{Compression: compression} },
		func() core.Server { return &ws_stream_sync.Server{} },
		func() core.SpanGenerator { return otlp.NewGenerator() },
	)
}

func benchmarkWSStreamAsync(options core.Options, compression experimental.CompressionMethod, concurrency int) {
	benchmarkImpl(
		"WebSocket/Stream/Async/"+strconv.Itoa(concurrency)+getCompSuffix(compression),
		options,
		func() core.Client { return &ws_stream_async.Client{Compression: compression, Concurrency: concurrency} },
		func() core.Server { return &ws_stream_async.Server{} },
		func() core.SpanGenerator { return otlp.NewGenerator() },
	)
}

func benchmarkWSAsyncWorker(options core.Options, concurrency int) {
	benchmarkImpl(
		"WebSocket/AsyncWorker/"+strconv.Itoa(concurrency),
		options,
		func() core.Client { return &ws_async_worker.Client{Concurrency: concurrency} },
		func() core.Server { return &ws_async_worker.Server{} },
		func() core.SpanGenerator { return otlp.NewGenerator() },
	)
}

func benchmarkHttp11(options core.Options, concurrency int) {
	benchmarkImpl(
		"OTLP/HTTP1.1/"+strconv.Itoa(concurrency),
		options,
		func() core.Client { return &http11.Client{Concurrency: concurrency} },
		func() core.Server { return &http11.Server{} },
		func() core.SpanGenerator { return otlp.NewGenerator() },
	)
}

func benchmarkSAPM(options core.Options, concurrency int) {
	benchmarkImpl(
		"SAPM/"+strconv.Itoa(concurrency),
		options,
		func() core.Client { return &sapm.Client{Concurrency: concurrency} },
		func() core.Server { return &sapm.Server{} },
		func() core.SpanGenerator { return sapmenc.NewGenerator() },
	)
}

func benchmarkImpl(
	name string,
	options core.Options,
	clientFactory func() core.Client,
	serverFactory func() core.Server,
	generatorFactory func() core.SpanGenerator,
) {
	cpuSecs, wallSecs := core.BenchmarkLocalDelivery(
		clientFactory,
		serverFactory,
		generatorFactory,
		options,
	)

	fmt.Printf("%-28s %7d spans|%5.1f cpusec|%5.1f wallsec|%7.1f batches/cpusec|%8.1f batches/wallsec|%5.1f cpuμs/span\n",
		name,
		options.Batches*options.SpansPerBatch,
		cpuSecs,
		wallSecs,
		float64(options.Batches)/cpuSecs,
		float64(options.Batches)/wallSecs,
		cpuSecs*1e6/float64(options.Batches*options.SpansPerBatch),
	)
}
