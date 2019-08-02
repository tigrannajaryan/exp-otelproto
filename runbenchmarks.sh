#!/usr/bin/env bash

echo ====================================================================================
echo Legend:
echo "WebSocket/Stream/Sync       - WebSocket, streaming, unknown load balancer friendliness, with sync ack"
echo "WebSocket/Stream/Async      - WebSocket, streaming, unknown load balancer friendliness, with async ack"
echo "WebSocket/Stream/Async/zlib - WebSocket, streaming, unknown load balancer friendliness, with async ack, zlib compression"
echo "GRPC/OpenCensus             - OpenCensus protocol, streaming, not load balancer friendly, without ack"
echo "GRPC/OpenCensusWithAck      - OpenCensus-like protocol, streaming, not load balancer friendly, with ack"
echo "GRPC/Unary                  - GRPC, unary request per batch, load balancer friendly, with ack"
echo "GRPC/Stream/NoLB            - GRPC, streaming, not load balancer friendly, with ack"
echo "GRPC/Stream/LBAlways/Sync   - GRPC, streaming, load balancer friendly, close stream after every batch, with ack"
echo "GRPC/Stream/LBTimed/Sync    - GRPC, streaming, load balancer friendly, close stream every 30 sec, with ack"
echo "GRPC/Stream/LBTimed/Async   - GRPC, streaming, load balancer friendly, close stream every 30 sec, with async ack"
echo

tc qdisc delete dev lo root netem delay 100ms > /dev/null 2>&1

# Set MULTIPLIER to 1 for quick results and to 100 for more stable results.
MULTIPLIER=1

cd bin

let BATCHES=800*MULTIPLIER
SPANSPERBATCH=100
ATTRPERSPAN=4
echo Small batches
echo spans/batch=${SPANSPERBATCH}, attrs/span=${ATTRPERSPAN}

./benchmark -protocol opencensus -batches=${BATCHES} -spansperbatch=${SPANSPERBATCH} -attrperspan=${ATTRPERSPAN}
./benchmark -protocol ocack -batches=${BATCHES} -spansperbatch=${SPANSPERBATCH} -attrperspan=${ATTRPERSPAN}
./benchmark -protocol unary -batches=${BATCHES} -spansperbatch=${SPANSPERBATCH} -attrperspan=${ATTRPERSPAN}
./benchmark -protocol streamsync -batches=${BATCHES} -spansperbatch=${SPANSPERBATCH} -attrperspan=${ATTRPERSPAN}
./benchmark -protocol streamlbalwayssync -batches=${BATCHES} -spansperbatch=${SPANSPERBATCH} -attrperspan=${ATTRPERSPAN}
./benchmark -protocol streamlbtimedsync -batches=${BATCHES} -spansperbatch=${SPANSPERBATCH} -attrperspan=${ATTRPERSPAN}
./benchmark -protocol streamlbasync -batches=${BATCHES} -spansperbatch=${SPANSPERBATCH} -attrperspan=${ATTRPERSPAN}
./benchmark -protocol wsstreamsync -batches=${BATCHES} -spansperbatch=${SPANSPERBATCH} -attrperspan=${ATTRPERSPAN}
./benchmark -protocol wsstreamasync -batches=${BATCHES} -spansperbatch=${SPANSPERBATCH} -attrperspan=${ATTRPERSPAN}
./benchmark -protocol wsstreamasynczlib -batches=${BATCHES} -spansperbatch=${SPANSPERBATCH} -attrperspan=${ATTRPERSPAN}

echo
let BATCHES=80*MULTIPLIER
SPANSPERBATCH=500
ATTRPERSPAN=10
echo Large batches
echo spans/batch=${SPANSPERBATCH}, attrs/span=${ATTRPERSPAN}

./benchmark -protocol opencensus -batches=${BATCHES} -spansperbatch=${SPANSPERBATCH} -attrperspan=${ATTRPERSPAN}
./benchmark -protocol ocack -batches=${BATCHES} -spansperbatch=${SPANSPERBATCH} -attrperspan=${ATTRPERSPAN}
./benchmark -protocol unary -batches=${BATCHES} -spansperbatch=${SPANSPERBATCH} -attrperspan=${ATTRPERSPAN}
./benchmark -protocol streamsync -batches=${BATCHES} -spansperbatch=${SPANSPERBATCH} -attrperspan=${ATTRPERSPAN}
./benchmark -protocol streamlbalwayssync -batches=${BATCHES} -spansperbatch=${SPANSPERBATCH} -attrperspan=${ATTRPERSPAN}
./benchmark -protocol streamlbtimedsync -batches=${BATCHES} -spansperbatch=${SPANSPERBATCH} -attrperspan=${ATTRPERSPAN}
./benchmark -protocol streamlbasync -batches=${BATCHES} -spansperbatch=${SPANSPERBATCH} -attrperspan=${ATTRPERSPAN}
./benchmark -protocol wsstreamsync -batches=${BATCHES} -spansperbatch=${SPANSPERBATCH} -attrperspan=${ATTRPERSPAN}
./benchmark -protocol wsstreamasync -batches=${BATCHES} -spansperbatch=${SPANSPERBATCH} -attrperspan=${ATTRPERSPAN}
./benchmark -protocol wsstreamasynczlib -batches=${BATCHES} -spansperbatch=${SPANSPERBATCH} -attrperspan=${ATTRPERSPAN}

echo
let BATCHES=80*MULTIPLIER
SPANSPERBATCH=500
ATTRPERSPAN=10
echo Large batches, 2ms network roundtrip latency
echo spans/batch=${SPANSPERBATCH}, attrs/span=${ATTRPERSPAN}

tc qdisc add dev lo root netem delay 1ms
./benchmark -protocol opencensus -batches=${BATCHES} -spansperbatch=${SPANSPERBATCH} -attrperspan=${ATTRPERSPAN}
./benchmark -protocol ocack -batches=${BATCHES} -spansperbatch=${SPANSPERBATCH} -attrperspan=${ATTRPERSPAN}
./benchmark -protocol unary -batches=${BATCHES} -spansperbatch=${SPANSPERBATCH} -attrperspan=${ATTRPERSPAN}
./benchmark -protocol streamsync -batches=${BATCHES} -spansperbatch=${SPANSPERBATCH} -attrperspan=${ATTRPERSPAN}
./benchmark -protocol streamlbalwayssync -batches=${BATCHES} -spansperbatch=${SPANSPERBATCH} -attrperspan=${ATTRPERSPAN}
./benchmark -protocol streamlbtimedsync -batches=${BATCHES} -spansperbatch=${SPANSPERBATCH} -attrperspan=${ATTRPERSPAN}
./benchmark -protocol streamlbasync -batches=${BATCHES} -spansperbatch=${SPANSPERBATCH} -attrperspan=${ATTRPERSPAN}
./benchmark -protocol wsstreamsync -batches=${BATCHES} -spansperbatch=${SPANSPERBATCH} -attrperspan=${ATTRPERSPAN}
./benchmark -protocol wsstreamasync -batches=${BATCHES} -spansperbatch=${SPANSPERBATCH} -attrperspan=${ATTRPERSPAN}
./benchmark -protocol wsstreamasynczlib -batches=${BATCHES} -spansperbatch=${SPANSPERBATCH} -attrperspan=${ATTRPERSPAN}
tc qdisc delete dev lo root netem delay 1ms

echo
let BATCHES=40*MULTIPLIER
SPANSPERBATCH=500
ATTRPERSPAN=10
echo Large batches, 20ms network roundtrip latency
echo spans/batch=${SPANSPERBATCH}, attrs/span=${ATTRPERSPAN}

tc qdisc add dev lo root netem delay 10ms
./benchmark -protocol opencensus -batches=${BATCHES} -spansperbatch=${SPANSPERBATCH} -attrperspan=${ATTRPERSPAN}
./benchmark -protocol ocack -batches=${BATCHES} -spansperbatch=${SPANSPERBATCH} -attrperspan=${ATTRPERSPAN}
./benchmark -protocol unary -batches=${BATCHES} -spansperbatch=${SPANSPERBATCH} -attrperspan=${ATTRPERSPAN}
./benchmark -protocol streamsync -batches=${BATCHES} -spansperbatch=${SPANSPERBATCH} -attrperspan=${ATTRPERSPAN}
./benchmark -protocol streamlbalwayssync -batches=${BATCHES} -spansperbatch=${SPANSPERBATCH} -attrperspan=${ATTRPERSPAN}
./benchmark -protocol streamlbtimedsync -batches=${BATCHES} -spansperbatch=${SPANSPERBATCH} -attrperspan=${ATTRPERSPAN}
./benchmark -protocol streamlbasync -batches=${BATCHES} -spansperbatch=${SPANSPERBATCH} -attrperspan=${ATTRPERSPAN}
./benchmark -protocol wsstreamsync -batches=${BATCHES} -spansperbatch=${SPANSPERBATCH} -attrperspan=${ATTRPERSPAN}
./benchmark -protocol wsstreamasync -batches=${BATCHES} -spansperbatch=${SPANSPERBATCH} -attrperspan=${ATTRPERSPAN}
./benchmark -protocol wsstreamasynczlib -batches=${BATCHES} -spansperbatch=${SPANSPERBATCH} -attrperspan=${ATTRPERSPAN}
tc qdisc delete dev lo root netem delay 10ms

echo
let BATCHES=4*MULTIPLIER
SPANSPERBATCH=500
ATTRPERSPAN=10
echo Large batches, 200ms network roundtrip latency
echo spans/batch=${SPANSPERBATCH}, attrs/span=${ATTRPERSPAN}

tc qdisc add dev lo root netem delay 100ms
let ASYNCBATCHES=10*${BATCHES}

./benchmark -protocol opencensus -batches=${ASYNCBATCHES} -spansperbatch=${SPANSPERBATCH} -attrperspan=${ATTRPERSPAN}
./benchmark -protocol ocack -batches=${BATCHES} -spansperbatch=${SPANSPERBATCH} -attrperspan=${ATTRPERSPAN}
./benchmark -protocol unary -batches=${BATCHES} -spansperbatch=${SPANSPERBATCH} -attrperspan=${ATTRPERSPAN}
./benchmark -protocol streamlbalwayssync -batches=${BATCHES} -spansperbatch=${SPANSPERBATCH} -attrperspan=${ATTRPERSPAN}
./benchmark -protocol streamsync -batches=${BATCHES} -spansperbatch=${SPANSPERBATCH} -attrperspan=${ATTRPERSPAN}
./benchmark -protocol streamlbtimedsync -batches=${BATCHES} -spansperbatch=${SPANSPERBATCH} -attrperspan=${ATTRPERSPAN}
./benchmark -protocol streamlbasync -batches=${ASYNCBATCHES} -spansperbatch=${SPANSPERBATCH} -attrperspan=${ATTRPERSPAN}
./benchmark -protocol wsstreamsync -batches=${BATCHES} -spansperbatch=${SPANSPERBATCH} -attrperspan=${ATTRPERSPAN}
./benchmark -protocol wsstreamasync -batches=${ASYNCBATCHES} -spansperbatch=${SPANSPERBATCH} -attrperspan=${ATTRPERSPAN}
./benchmark -protocol wsstreamasynczlib -batches=${ASYNCBATCHES} -spansperbatch=${SPANSPERBATCH} -attrperspan=${ATTRPERSPAN}
tc qdisc delete dev lo root netem delay 100ms

echo ====================================================================================
