# Building and Running

Run `make` to build.

Run `make test` to run benchmarks.

# Source code

Interfaces `core.Client` and `core.Server` define a client and a server that can send and receive a batch of spans.

`grpc_unary` and `grpc_stream` packages have implementations of these interfaces.

experiment.go contains generic `BenchmarkLocalDelivery` test that accepts factories to create client, server and batch generator and runs the test using local connection.

The test is executed twice by `BenchmarkGRPC` in grpc_test.go, once for Unary, once for Stream implementations.  

I execucted it on my local machine and saved the results to `perf.txt`. It shows that Stream uses about 1.2 times less CPU when sending 60000 batches (each batch containing 100 spans). It also does 1.025 times less memory allocation (which is probably not significant).
 