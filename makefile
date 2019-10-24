.PHONY: all genprotobuf genflatbuffers build build-image build-images publish-images

GO=$(shell which go)

K8S_NAMESPACE?=otpexp
DOCKER_REGISTRY?=592865182265.dkr.ecr.us-west-2.amazonaws.com/
IMAGE_NAME=
PROTOCOL?=
VERSION=1.1

export VERSION
export K8S_NAMESPACE
export DOCKER_REGISTRY
export IMAGE_NAME
export PROTOCOL

all: genprotobuf build test

genprotobuf:
	protoc -I/usr/local/include -I encodings/traceprotobuf/ encodings/traceprotobuf/telemetry_data.proto --go_out=plugins=grpc:encodings/traceprotobuf
	protoc -I/usr/local/include -I encodings/traceprotobuf/ encodings/traceprotobuf/resource.proto --go_out=plugins=grpc:encodings/traceprotobuf
	protoc -I/usr/local/include -I encodings/traceprotobuf/ encodings/traceprotobuf/exchange.proto --go_out=plugins=grpc:encodings/traceprotobuf
	protoc -I/usr/local/include -I encodings/traceprotobuf/ encodings/traceprotobuf/grpc.proto --go_out=plugins=grpc:encodings/traceprotobuf

	protoc -I/usr/local/include -I encodings/traceprotobufb/ encodings/traceprotobufb/telemetry_data.proto --go_out=plugins=grpc:encodings/traceprotobufb
	protoc -I/usr/local/include -I encodings/traceprotobufb/ encodings/traceprotobufb/exchange.proto --go_out=plugins=grpc:encodings/traceprotobufb

	protoc -I/usr/local/include -I encodings/otlp/ encodings/otlp/metric_data.proto --go_out=plugins=grpc:encodings/otlp
	protoc -I/usr/local/include -I encodings/otlp/ encodings/otlp/telemetry_data.proto --go_out=plugins=grpc:encodings/otlp
	protoc -I/usr/local/include -I encodings/otlp/ encodings/otlp/exchange.proto --go_out=plugins=grpc:encodings/otlp
	protoc -I/usr/local/include -I encodings/otlp/ encodings/otlp/grpc.proto --go_out=plugins=grpc:encodings/otlp

	protoc -I/usr/local/include -I encodings/otlptimewrapped/ encodings/otlptimewrapped/telemetry_data.proto --go_out=plugins=grpc:encodings/otlptimewrapped
	protoc -I/usr/local/include -I encodings/otlptimewrapped/ encodings/otlptimewrapped/exchange.proto --go_out=plugins=grpc:encodings/otlptimewrapped

	protoc -I/usr/local/include -I encodings/octraceprotobuf/ encodings/octraceprotobuf/octrace.proto --go_out=plugins=grpc:encodings/octraceprotobuf
	protoc -I/usr/local/include -I encodings/octraceprotobuf/ encodings/octraceprotobuf/resource.proto --go_out=plugins=grpc:encodings/octraceprotobuf
	protoc -I/usr/local/include -I encodings/octraceprotobuf/ encodings/octraceprotobuf/metrics.proto --go_out=plugins=grpc:encodings/octraceprotobuf

genflatbuffers:
	protoc -I/usr/local/include -I encodings/traceflatbuffers/ encodings/traceflatbuffers/trace.proto --go_out=plugins=grpc:encodings/traceflatbuffers
	flatc --gen-object-api --go encodings/traceflatbuffers/trace.fbs

build:
	go build -o bin/benchmark cmd/benchmark/main.go
	go build -o bin/loadgen cmd/loadgen/main.go
	go build -o bin/server cmd/server/main.go

benchmark:
	./runbenchmarks.sh

benchmark-encoding:
	#sudo ./beforebenchmarks.sh
	sudo nice -n -5 ${GO} test -bench . ./encodings -benchtime 5s
	#sudo ./afterbenchmarks.sh

run:
	go run cmd/grpc-protobuf.go

test:
	go test -v ./...

build-images: build
	$(MAKE) build-image IMAGE_NAME=server
	$(MAKE) build-image IMAGE_NAME=loadgen

publish-images: build-images
	$(MAKE) publish-image IMAGE_NAME=server
	$(MAKE) publish-image IMAGE_NAME=loadgen

build-image:
	cp bin/${IMAGE_NAME} ./cmd/${IMAGE_NAME}
	docker build -t otpexp-${IMAGE_NAME} ./cmd/${IMAGE_NAME}

publish-image:
	docker tag otpexp-${IMAGE_NAME} ${DOCKER_REGISTRY}otpexp-${IMAGE_NAME}:${VERSION}
	docker push ${DOCKER_REGISTRY}otpexp-${IMAGE_NAME}:${VERSION}

deploy:
	envsubst < kubernetes/service.yaml | kubectl apply -n ${K8S_NAMESPACE} -f - 
	envsubst < kubernetes/deployment-server.yaml | kubectl apply -n ${K8S_NAMESPACE} -f - 
	envsubst < kubernetes/deployment-loadgen.yaml | kubectl apply -n ${K8S_NAMESPACE} -f - 

deploy-all: publish-images deploy
