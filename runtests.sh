#!/usr/bin/env bash

for i in {1..5}
do
    echo =============
    go test -bench BenchmarkGRPCUnary -benchmem -timeout 600m github.com/tigrannajaryan/exp-otelproto/tests
    go test -bench BenchmarkGRPCStreamLB -benchmem -timeout 600m github.com/tigrannajaryan/exp-otelproto/tests
    go test -bench BenchmarkGRPCStreamNoLB -benchmem -timeout 600m github.com/tigrannajaryan/exp-otelproto/tests
done