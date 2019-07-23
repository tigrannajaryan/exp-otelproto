package main

import (
	"flag"
	"log"

	"github.com/tigrannajaryan/exp-otelproto/core"
	"github.com/tigrannajaryan/exp-otelproto/grpcimpl"
)

func main() {
	log.Println("GRPC/Protobuf Agent.")

	var destination string
	flag.StringVar(&destination, "dest", "localhost:3465", "destination endpoint to forward to")

	var listenAddress string
	flag.StringVar(&listenAddress, "listen", "0.0.0.0:3465", "local address to listen on")

	flag.Parse()

	core.RunAgent(&grpcimpl.Client{}, &grpcimpl.Server{}, listenAddress, destination)
}
