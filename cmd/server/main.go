package main

import (
	"flag"
	"log"

	"github.com/tigrannajaryan/exp-otelproto/protoimpls/ws_stream_async"
	"github.com/tigrannajaryan/exp-otelproto/protoimpls/ws_stream_sync"

	"github.com/tigrannajaryan/exp-otelproto/protoimpls/grpc_stream_lb_async"

	"github.com/tigrannajaryan/exp-otelproto/protoimpls/grpc_stream_lb"

	"github.com/tigrannajaryan/exp-otelproto/protoimpls/grpc_stream"

	"github.com/tigrannajaryan/exp-otelproto/protoimpls/grpc_unary"

	"github.com/tigrannajaryan/exp-otelproto/core"
	"github.com/tigrannajaryan/exp-otelproto/protoimpls/grpc_oc"
)

func main() {
	log.Println("Load Generator.")

	var listenAddress string
	flag.StringVar(&listenAddress, "listen", "0.0.0.0:3465", "local address to listen on")

	var protocol string
	flag.StringVar(&protocol, "protocol", "",
		"protocol to benchmark (opencensus,ocack,unary,streamsync,streamlbtimedsync,streamlbalwayssync,streamlbasync,wsstreamsync,wsstreamasync,wsstreamasynczlib)")

	flag.Parse()

	switch protocol {
	case "opencensus":
		core.RunServer(&grpc_oc.Server{}, listenAddress)
	case "ocack":
		core.RunServer(&grpc_oc.Server{SendAck: true}, listenAddress)
	case "unary":
		core.RunServer(&grpc_unary.Server{}, listenAddress)
	case "streamsync":
		core.RunServer(&grpc_stream.Server{}, listenAddress)
	case "streamlbtimedsync":
		core.RunServer(&grpc_stream_lb.Server{}, listenAddress)
	case "streamlbalwayssync":
		core.RunServer(&grpc_stream_lb.Server{}, listenAddress)
	case "streamlbasync":
		core.RunServer(&grpc_stream_lb_async.Server{}, listenAddress)
	case "wsstreamsync":
		core.RunServer(&ws_stream_sync.Server{}, listenAddress)
	case "wsstreamasync":
		core.RunServer(&ws_stream_async.Server{}, listenAddress)
	case "wsstreamasynczlib":
		core.RunServer(&ws_stream_async.Server{}, listenAddress)
	default:
		flag.Usage()
	}

}
