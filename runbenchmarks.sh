#!/usr/bin/env bash

echo ====================================================================================
echo Legend:
echo "GRPC/OpenCensus           - GRPC, OpenCensus protocol, stream without ack"
echo "GRPC/Unary                - GRPC, unary request per batch"
echo "GRPC/Stream/NoLB          - GRPC, streaming, not load balancer friendly"
echo "GRPC/Stream/LBAlways/Sync - GRPC, streaming, load balancer friendly, close stream after every batch"
echo "GRPC/Stream/LBTimed/Sync  - GRPC, streaming, load balancer friendly, close stream every 30 sec"
echo "GRPC/Stream/LBTimed/Async - GRPC, streaming, load balancer friendly, async ack, close stream every 30 sec"
echo

sudo tc qdisc delete dev lo root netem delay 100ms > /dev/null 2>&1

# Set MULTIPLIER to 1 for quick results and to 100 for more stable results.
MULTIPLIER=50

cd bin

let BATCHES=800*MULTIPLIER
SPANSPERBATCH=100
ATTRPERSPAN=2
echo Small batches
echo spans/batch=${SPANSPERBATCH}, attrs/span=${ATTRPERSPAN}

./benchmark -protocol opencensus -batches=${BATCHES} -spansperbatch=${SPANSPERBATCH} -attrperspan=${ATTRPERSPAN}
./benchmark -protocol unary -batches=${BATCHES} -spansperbatch=${SPANSPERBATCH} -attrperspan=${ATTRPERSPAN}
./benchmark -protocol streamsync -batches=${BATCHES} -spansperbatch=${SPANSPERBATCH} -attrperspan=${ATTRPERSPAN}
./benchmark -protocol streamlbalwayssync -batches=${BATCHES} -spansperbatch=${SPANSPERBATCH} -attrperspan=${ATTRPERSPAN}
./benchmark -protocol streamlbtimedsync -batches=${BATCHES} -spansperbatch=${SPANSPERBATCH} -attrperspan=${ATTRPERSPAN}
./benchmark -protocol streamlbasync -batches=${BATCHES} -spansperbatch=${SPANSPERBATCH} -attrperspan=${ATTRPERSPAN}

echo
let BATCHES=80*MULTIPLIER
SPANSPERBATCH=500
ATTRPERSPAN=5
echo Large batches
echo spans/batch=${SPANSPERBATCH}, attrs/span=${ATTRPERSPAN}

./benchmark -protocol opencensus -batches=${BATCHES} -spansperbatch=${SPANSPERBATCH} -attrperspan=${ATTRPERSPAN}
./benchmark -protocol unary -batches=${BATCHES} -spansperbatch=${SPANSPERBATCH} -attrperspan=${ATTRPERSPAN}
./benchmark -protocol streamsync -batches=${BATCHES} -spansperbatch=${SPANSPERBATCH} -attrperspan=${ATTRPERSPAN}
./benchmark -protocol streamlbalwayssync -batches=${BATCHES} -spansperbatch=${SPANSPERBATCH} -attrperspan=${ATTRPERSPAN}
./benchmark -protocol streamlbtimedsync -batches=${BATCHES} -spansperbatch=${SPANSPERBATCH} -attrperspan=${ATTRPERSPAN}
./benchmark -protocol streamlbasync -batches=${BATCHES} -spansperbatch=${SPANSPERBATCH} -attrperspan=${ATTRPERSPAN}

echo
let BATCHES=80*MULTIPLIER
SPANSPERBATCH=500
ATTRPERSPAN=5
echo Large batches, 2ms network roundtrip latency
echo spans/batch=${SPANSPERBATCH}, attrs/span=${ATTRPERSPAN}

sudo tc qdisc add dev lo root netem delay 1ms

./benchmark -protocol opencensus -batches=${BATCHES} -spansperbatch=${SPANSPERBATCH} -attrperspan=${ATTRPERSPAN}
./benchmark -protocol unary -batches=${BATCHES} -spansperbatch=${SPANSPERBATCH} -attrperspan=${ATTRPERSPAN}
./benchmark -protocol streamsync -batches=${BATCHES} -spansperbatch=${SPANSPERBATCH} -attrperspan=${ATTRPERSPAN}
./benchmark -protocol streamlbalwayssync -batches=${BATCHES} -spansperbatch=${SPANSPERBATCH} -attrperspan=${ATTRPERSPAN}
./benchmark -protocol streamlbtimedsync -batches=${BATCHES} -spansperbatch=${SPANSPERBATCH} -attrperspan=${ATTRPERSPAN}
./benchmark -protocol streamlbasync -batches=${BATCHES} -spansperbatch=${SPANSPERBATCH} -attrperspan=${ATTRPERSPAN}
sudo tc qdisc delete dev lo root netem delay 1ms

echo
let BATCHES=40*MULTIPLIER
SPANSPERBATCH=500
ATTRPERSPAN=5
echo Large batches, 20ms network roundtrip latency
echo spans/batch=${SPANSPERBATCH}, attrs/span=${ATTRPERSPAN}

sudo tc qdisc add dev lo root netem delay 10ms

./benchmark -protocol opencensus -batches=${BATCHES} -spansperbatch=${SPANSPERBATCH} -attrperspan=${ATTRPERSPAN}
./benchmark -protocol unary -batches=${BATCHES} -spansperbatch=${SPANSPERBATCH} -attrperspan=${ATTRPERSPAN}
./benchmark -protocol streamsync -batches=${BATCHES} -spansperbatch=${SPANSPERBATCH} -attrperspan=${ATTRPERSPAN}
./benchmark -protocol streamlbalwayssync -batches=${BATCHES} -spansperbatch=${SPANSPERBATCH} -attrperspan=${ATTRPERSPAN}
./benchmark -protocol streamlbtimedsync -batches=${BATCHES} -spansperbatch=${SPANSPERBATCH} -attrperspan=${ATTRPERSPAN}
./benchmark -protocol streamlbasync -batches=${BATCHES} -spansperbatch=${SPANSPERBATCH} -attrperspan=${ATTRPERSPAN}
sudo tc qdisc delete dev lo root netem delay 10ms

echo
let BATCHES=4*MULTIPLIER
SPANSPERBATCH=500
ATTRPERSPAN=5
echo Large batches, 200ms network roundtrip latency
echo spans/batch=${SPANSPERBATCH}, attrs/span=${ATTRPERSPAN}

sudo tc qdisc add dev lo root netem delay 100ms
let ASYNCBATCHES=10*${BATCHES}

./benchmark -protocol opencensus -batches=${BATCHES} -spansperbatch=${SPANSPERBATCH} -attrperspan=${ATTRPERSPAN}
./benchmark -protocol unary -batches=${BATCHES} -spansperbatch=${SPANSPERBATCH} -attrperspan=${ATTRPERSPAN}
./benchmark -protocol streamlbalwayssync -batches=${BATCHES} -spansperbatch=${SPANSPERBATCH} -attrperspan=${ATTRPERSPAN}
./benchmark -protocol streamsync -batches=${BATCHES} -spansperbatch=${SPANSPERBATCH} -attrperspan=${ATTRPERSPAN}
./benchmark -protocol streamlbtimedsync -batches=${BATCHES} -spansperbatch=${SPANSPERBATCH} -attrperspan=${ATTRPERSPAN}
./benchmark -protocol streamlbasync -batches=${ASYNCBATCHES} -spansperbatch=${SPANSPERBATCH} -attrperspan=${ATTRPERSPAN}
sudo tc qdisc delete dev lo root netem delay 100ms

echo ====================================================================================
