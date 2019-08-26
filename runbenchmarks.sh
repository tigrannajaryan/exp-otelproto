#!/usr/bin/env bash

# Set MULTIPLIER to 1 for quick results and to 100 for more stable results.
MULTIPLIER=5

echo ====================================================================================
echo Legend:
echo "GRPC/OpenCensus             - OpenCensus protocol, streaming, not load balancer friendly, without ack"
echo "GRPC/OpenCensusWithAck      - OpenCensus-like protocol, streaming, not load balancer friendly, with ack"
echo "GRPC/Unary                  - GRPC, unary request per batch, load balancer friendly, with ack"
echo "GRPC/Unary/Async            - GRPC, unary async request per batch, load balancer friendly, with ack"
echo "GRPC/Stream/NoLB            - GRPC, streaming, not load balancer friendly, with ack"
echo "GRPC/Stream/LBAlways/Sync   - GRPC, streaming, load balancer friendly, close stream after every batch, with ack"
echo "GRPC/Stream/LBTimed/Sync    - OTLP Synchronous. GRPC, streaming, load balancer friendly, close stream every 30 sec, with ack"
echo "GRPC/Stream/LBTimed/Async/N - OTLP Pipelined. GRPC, N streams, load balancer friendly, close stream every 30 sec, with async ack"
echo "GRPC/Stream/LBSrv/Async     - OTLP Pipelined. GRPC, streaming, load balancer friendly, server closes stream every 30 sec or 1000 batches, with async ack"
echo "WebSocket/Stream/Sync       - WebSocket, streaming, unknown load balancer friendliness, with sync ack"
echo "WebSocket/Stream/Async      - WebSocket, streaming, unknown load balancer friendliness, with async ack"
echo "WebSocket/Stream/Async/zlib - WebSocket, streaming, unknown load balancer friendliness, with async ack, zlib compression"
echo

benchmark() {
    nice -n -5 ./benchmark -protocol $1 -batches=${BATCHES} -spansperbatch=${SPANSPERBATCH} -attrperspan=${ATTRPERSPAN}
}

benchmark_all() {
    echo ${BATCHES} $1 batches, ${SPANSPERBATCH} spans per batch, ${ATTRPERSPAN} attrs per span
    benchmark streamlbasync
    benchmark streamlbconc
    benchmark unary
    benchmark unaryasync
    benchmark opencensus
    benchmark ocack
    benchmark streamsync
    benchmark streamlbalwayssync
    benchmark streamlbtimedsync
    benchmark streamlbsrv
    benchmark wsstreamsync
    benchmark wsstreamasync
    benchmark wsstreamasynczlib
    echo
}

benchmark_all_latency() {
    let roundtrip=$1*2
    echo ${roundtrip}ms network roundtrip latency
    tc qdisc add dev lo root netem delay ${1}ms
    benchmark_all large
    tc qdisc delete dev lo root netem delay ${1}ms
}

./beforebenchmarks.sh

tc qdisc delete dev lo root netem delay 100ms > /dev/null 2>&1
echo

cd bin

let BATCHES=800*MULTIPLIER
SPANSPERBATCH=100
ATTRPERSPAN=4
benchmark_all small


let BATCHES=80*MULTIPLIER
SPANSPERBATCH=500
ATTRPERSPAN=10
benchmark_all large

let BATCHES=80*MULTIPLIER
SPANSPERBATCH=500
ATTRPERSPAN=10
benchmark_all_latency 1

let BATCHES=40*MULTIPLIER
SPANSPERBATCH=500
ATTRPERSPAN=10
benchmark_all_latency 10

let BATCHES=4*MULTIPLIER
SPANSPERBATCH=500
ATTRPERSPAN=10
benchmark_all_latency 100

echo
let BATCHES=4*MULTIPLIER*10
SPANSPERBATCH=500
ATTRPERSPAN=10
echo ${BATCHES} large batches, ${SPANSPERBATCH} spans per batch, ${ATTRPERSPAN} attrs per span
echo 200ms network roundtrip latency

tc qdisc add dev lo root netem delay 100ms
benchmark opencensus
benchmark streamlbasync
benchmark streamlbconc
benchmark unaryasync
benchmark wsstreamasync
benchmark wsstreamasynczlib
tc qdisc delete dev lo root netem delay 100ms

echo ====================================================================================

./afterbenchmarks.sh