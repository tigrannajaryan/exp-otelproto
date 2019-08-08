package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/tigrannajaryan/exp-otelproto/core"
	"github.com/tigrannajaryan/exp-otelproto/protoimpls/grpc_oc"
	"github.com/tigrannajaryan/exp-otelproto/protoimpls/grpc_stream"
	"github.com/tigrannajaryan/exp-otelproto/protoimpls/grpc_stream_lb"
	"github.com/tigrannajaryan/exp-otelproto/protoimpls/grpc_stream_lb_async"
	"github.com/tigrannajaryan/exp-otelproto/protoimpls/grpc_unary"
	"github.com/tigrannajaryan/exp-otelproto/protoimpls/ws_stream_async"
	"github.com/tigrannajaryan/exp-otelproto/protoimpls/ws_stream_sync"
)

var (
	batchesReceived = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "otlp_server_batches_received",
		Help: "The total number of received batches",
	}, []string{"protocol"})
	spansReceived = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "otlp_server_spans_received",
		Help: "The total number of received spans",
	}, []string{"protocol"})
)

func newOnReceive(protocol string) func(int) {
	return func(spanCount int) {
		batchesReceived.With(prometheus.Labels{"protocol": protocol}).Inc()
		spansReceived.With(prometheus.Labels{"protocol": protocol}).Add(float64(spanCount))
	}
}

func main() {
	log.Println("Load Generator.")

	var listenAddress string
	flag.StringVar(&listenAddress, "listen", "0.0.0.0:3465", "local address to listen on")

	var protocol string
	flag.StringVar(&protocol, "protocol", "",
		"protocol to benchmark (opencensus,ocack,unary,streamsync,streamlbtimedsync,streamlbalwayssync,streamlbasync,wsstreamsync,wsstreamasync,wsstreamasynczlib)")

	flag.Parse()

	http.Handle("/metrics", promhttp.Handler())
	go http.ListenAndServe(":2112", nil)

	onReceive := newOnReceive(protocol)
	switch protocol {
	case "opencensus":
		core.RunServer(&grpc_oc.Server{}, listenAddress, onReceive)
	case "ocack":
		core.RunServer(&grpc_oc.Server{SendAck: true}, listenAddress, onReceive)
	case "unary":
		core.RunServer(&grpc_unary.Server{}, listenAddress, onReceive)
	case "streamsync":
		core.RunServer(&grpc_stream.Server{}, listenAddress, onReceive)
	case "streamlbtimedsync":
		core.RunServer(&grpc_stream_lb.Server{}, listenAddress, onReceive)
	case "streamlbalwayssync":
		core.RunServer(&grpc_stream_lb.Server{}, listenAddress, onReceive)
	case "streamlbasync":
		core.RunServer(&grpc_stream_lb_async.Server{}, listenAddress, onReceive)
	case "wsstreamsync":
		core.RunServer(&ws_stream_sync.Server{}, listenAddress, onReceive)
	case "wsstreamasync":
		core.RunServer(&ws_stream_async.Server{}, listenAddress, onReceive)
	case "wsstreamasynczlib":
		core.RunServer(&ws_stream_async.Server{}, listenAddress, onReceive)
	default:
		flag.Usage()
	}

}
