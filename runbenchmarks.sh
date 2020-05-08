#!/usr/bin/env bash

# Set MULTIPLIER to 1 for quick results and to 100 for more stable results.
MULTIPLIER=10

echo ====================================================================================
echo Legend:
echo "OTLP/GRPC-Unary/Sequential  - OTLP, Unary, sequential. One request per batch, load balancer friendly, with ack"
echo "OTLP/GRPC-Unary/Concurrent  - OTLP, Unary, 10 concurrent requests, load balancer friendly, with ack"
echo "GRPC/Stream/LBTimed/Sync    - OTLP ProtoBuf,GRPC, streaming, load balancer friendly, close stream every 30 sec, with ack"
echo "GRPC/Stream/LBTimed/Async/N - OTLP ProtoBuf,GRPC, streaming. N streams, load balancer friendly, close stream every 30 sec, with async ack"
echo "GRPC/OpenCensus             - OpenCensus protocol, streaming, not load balancer friendly, without ack"
echo "GRPC/OpenCensusWithAck      - OpenCensus-like protocol, streaming, not load balancer friendly, with ack"
echo "GRPC/Stream/NoLB            - OTLP ProtoBuf, GRPC, streaming, not load balancer friendly, with ack"
echo "GRPC/Stream/LBAlways/Sync   - OTLP ProtoBuf,GRPC, streaming, load balancer friendly, close stream after every batch, with ack"
echo "GRPC/Stream/LBSrv/Async     - OTLP ProtoBuf,GRPC Streaming. Load balancer friendly, server closes stream every 30 sec or 1000 batches, with async ack"
echo "WebSocket/Stream/Sync       - OTLP ProtoBuf,WebSocket, streaming, unknown load balancer friendliness, with sync ack"
echo "WebSocket/Stream/Async/N    - OTLP ProtoBuf,WebSocket, N streams, unknown load balancer friendliness, with async ack"
echo "WebSocket/Stream/Async/zlib - OTLP ProtoBuf,WebSocket, streaming, unknown load balancer friendliness, with async ack, zlib compression"
echo "OTLP/HTTP1.1/N              - OTLP ProtoBuf,HTTP 1.1, N concurrent requests. Load balancer friendly."
echo "SAPM/N                      - SAPM, N concurrent requests. Load balacner friendly."
echo

benchmark() {
    nice -n -5 ./benchmark -protocol $1 -batches=${BATCHES} -spansperbatch=${SPANSPERBATCH} -attrperspan=${ATTRPERSPAN}
}

benchmark_all() {
    echo ${BATCHES} $1 batches, ${SPANSPERBATCH} spans per batch, ${ATTRPERSPAN} attrs per span
    #benchmark sapm
    benchmark http11
    benchmark http11conc
    #benchmark wsstreamsync
    #benchmark wsstreamasync
    #benchmark wsstreamasyncconc
    #benchmark wsstreamasynczlib
    benchmark unary
    benchmark unaryasync
    #benchmark streamlbasync
    #benchmark streamlbconc
    #benchmark opencensus
    #benchmark ocack
    #benchmark streamsync
    #benchmark streamlbalwayssync
    #benchmark streamlbtimedsync
    #benchmark streamlbsrv
    echo
}

benchmark_all_latency() {
    let roundtrip=$1*2
    echo ${roundtrip}ms network roundtrip latency
    tc qdisc add dev lo root netem delay ${1}ms
    benchmark_all large
    tc qdisc delete dev lo root netem delay ${1}ms
}

benchmark_some_latency() {
    echo 200ms network roundtrip latency
    echo ${BATCHES} large batches, ${SPANSPERBATCH} spans per batch, ${ATTRPERSPAN} attrs per span

    tc qdisc add dev lo root netem delay 100ms
    benchmark http11
    #benchmark unaryasync
    #benchmark opencensus
    #benchmark streamlbasync
    #benchmark streamlbconc
    benchmark wsstreamasync
    #benchmark wsstreamasyncconc
    #benchmark wsstreamasynczlib
    tc qdisc delete dev lo root netem delay 100ms
}


./beforebenchmarks.sh

tc qdisc delete dev lo root netem delay 100ms > /dev/null 2>&1
echo

cd bin

let BATCHES=6400*MULTIPLIER
SPANSPERBATCH=1
ATTRPERSPAN=10
benchmark_all nano

let BATCHES=1600*MULTIPLIER
SPANSPERBATCH=10
ATTRPERSPAN=10
benchmark_all tiny


let BATCHES=800*MULTIPLIER
SPANSPERBATCH=100
ATTRPERSPAN=10
benchmark_all small


let BATCHES=80*MULTIPLIER
SPANSPERBATCH=500
ATTRPERSPAN=10
benchmark_all large

let BATCHES=10*MULTIPLIER
SPANSPERBATCH=5000
ATTRPERSPAN=10
benchmark_all "very large"



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

let BATCHES=4*MULTIPLIER*10
SPANSPERBATCH=500
ATTRPERSPAN=10
benchmark_some_latency

echo ====================================================================================

cd ..

./afterbenchmarks.sh