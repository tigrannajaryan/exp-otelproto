.PHONY: all genprotobuf genflatbuffers build build-image build-images publish-images

K8S_NAMESPACE?=otpexp
DOCKER_REGISTRY?=592865182265.dkr.ecr.us-west-2.amazonaws.com/
IMAGE_NAME=
PROTOCOL?=

export K8S_NAMESPACE
export DOCKER_REGISTRY
export IMAGE_NAME
export PROTOCOL

all: genprotobuf build

genprotobuf:
	protoc -I/usr/local/include -I encodings/traceprotobuf/ encodings/traceprotobuf/trace.proto --go_out=plugins=grpc:encodings/traceprotobuf
	protoc -I/usr/local/include -I encodings/traceprotobuf/ encodings/traceprotobuf/resource.proto --go_out=plugins=grpc:encodings/traceprotobuf
	protoc -I/usr/local/include -I encodings/traceprotobuf/ encodings/traceprotobuf/exchange.proto --go_out=plugins=grpc:encodings/traceprotobuf
	protoc -I/usr/local/include -I encodings/traceprotobuf/ encodings/traceprotobuf/grpc.proto --go_out=plugins=grpc:encodings/traceprotobuf

	protoc -I/usr/local/include -I encodings/octraceprotobuf/ encodings/octraceprotobuf/octrace.proto --go_out=plugins=grpc:encodings/octraceprotobuf
	protoc -I/usr/local/include -I encodings/octraceprotobuf/ encodings/octraceprotobuf/resource.proto --go_out=plugins=grpc:encodings/octraceprotobuf

genflatbuffers:
	protoc -I/usr/local/include -I encodings/traceflatbuffers/ encodings/traceflatbuffers/trace.proto --go_out=plugins=grpc:encodings/traceflatbuffers
	flatc --gen-object-api --go encodings/traceflatbuffers/trace.fbs

build:
	go build -o bin/benchmark cmd/benchmark/main.go
	go build -o bin/loadgen cmd/loadgen/main.go
	go build -o bin/server cmd/server/main.go

benchmark:
	./runbenchmarks.sh

run:
	go run cmd/grpc-protobuf.go

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
	docker tag otpexp-${IMAGE_NAME} ${DOCKER_REGISTRY}otpexp-${IMAGE_NAME}:latest
	docker push ${DOCKER_REGISTRY}otpexp-${IMAGE_NAME}:latest

deploy:
	envsubst < kubernetes/service.yaml | kubectl apply -n ${K8S_NAMESPACE} -f - 
	envsubst < kubernetes/deployment-server.yaml | kubectl apply -n ${K8S_NAMESPACE} -f - 
	envsubst < kubernetes/deployment-loadgen.yaml | kubectl apply -n ${K8S_NAMESPACE} -f - 