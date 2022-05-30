.PHONY: all gen genflatbuffers build build-image build-images publish-images

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

# Function to execute a command.
# Accepts command to execute as first parameter.
define exec-command
$(1)

endef

# Find all .proto files.
OTLP_PROTO_FILES := $(wildcard encodings/otlp_gogo/opentelemetry/proto/*/v1/*.proto encodings/otlp_gogo/opentelemetry/proto/collector/*/v1/*.proto)
OTLP_PROTO_FILES2 := $(wildcard encodings/otlp_gogo2/opentelemetry/proto/*/v1/*.proto encodings/otlp_gogo2/opentelemetry/proto/collector/*/v1/*.proto)
OTLP_PROTO_FILES3 := $(wildcard encodings/otlp_gogo3/opentelemetry/proto/*/v1/*.proto encodings/otlp_gogo3/opentelemetry/proto/collector/*/v1/*.proto)


all: build test

gen: gen-baseline gen-experimental gen-gogo

gen-gogo:
	$(foreach file,$(OTLP_PROTO_FILES),$(call exec-command,protoc -Iencodings/otlp_gogo/ -Ivendor --gogofaster_out=plugins=grpc:encodings/otlp_gogo $(file)))
	cp -R encodings/otlp_gogo/github.com/tigrannajaryan/exp-otelproto/encodings/otlp_gogo/* encodings/otlp_gogo/
	rm -rf encodings/otlp_gogo/github.com/

	$(foreach file,$(OTLP_PROTO_FILES2),$(call exec-command,protoc -Iencodings/otlp_gogo2/ -Ivendor --gogofaster_out=plugins=grpc:encodings/otlp_gogo2 $(file)))
	cp -R encodings/otlp_gogo2/github.com/tigrannajaryan/exp-otelproto/encodings/otlp_gogo2/* encodings/otlp_gogo2/
	rm -rf encodings/otlp_gogo2/github.com/

	$(foreach file,$(OTLP_PROTO_FILES3),$(call exec-command,protoc -Iencodings/otlp_gogo3/ -Ivendor --gogofaster_out=plugins=grpc:encodings/otlp_gogo3 $(file)))
	cp -R encodings/otlp_gogo3/github.com/tigrannajaryan/exp-otelproto/encodings/otlp_gogo3/* encodings/otlp_gogo3/
	rm -rf encodings/otlp_gogo3/github.com/

gen-traceprotobuf:
	protoc -I/usr/local/include -I encodings/traceprotobuf/ encodings/traceprotobuf/telemetry_data.proto --go_out=plugins=grpc:encodings/traceprotobuf
	protoc -I/usr/local/include -I encodings/traceprotobuf/ encodings/traceprotobuf/resource.proto --go_out=plugins=grpc:encodings/traceprotobuf
	protoc -I/usr/local/include -I encodings/traceprotobuf/ encodings/traceprotobuf/exchange.proto --go_out=plugins=grpc:encodings/traceprotobuf
	protoc -I/usr/local/include -I encodings/traceprotobuf/ encodings/traceprotobuf/grpc.proto --go_out=plugins=grpc:encodings/traceprotobuf
	protoc -I/usr/local/include -I encodings/traceprotobufb/ encodings/traceprotobufb/telemetry_data.proto --go_out=plugins=grpc:encodings/traceprotobufb
	protoc -I/usr/local/include -I encodings/traceprotobufb/ encodings/traceprotobufb/exchange.proto --go_out=plugins=grpc:encodings/traceprotobufb

#	protoc -I/usr/local/include -I encodings/otlp/ encodings/otlp/logs.proto --go_out=plugins=grpc:encodings/otlp

gen-experimental:
	protoc -I/usr/local/include -I encodings/experimental/ encodings/experimental/common.proto --go_out=plugins=grpc:encodings/experimental
	protoc -I/usr/local/include -I encodings/experimental/ encodings/experimental/metric_data.proto --go_out=plugins=grpc:encodings/experimental
	protoc -I/usr/local/include -I encodings/experimental/ encodings/experimental/telemetry_data.proto --go_out=plugins=grpc:encodings/experimental
	protoc -I/usr/local/include -I encodings/experimental/ encodings/experimental/exchange.proto --go_out=plugins=grpc:encodings/experimental
	protoc -I/usr/local/include -I encodings/experimental/ encodings/experimental/logs.proto --go_out=plugins=grpc:encodings/experimental
	protoc -I/usr/local/include -I encodings/experimental/ encodings/experimental/logs_service.proto --go_out=plugins=grpc:encodings/experimental
	protoc -I/usr/local/include -I encodings/experimental/ encodings/experimental/grpc.proto --go_out=plugins=grpc:encodings/experimental

gen-otelp2:
	protoc -I/usr/local/include -I encodings/otelp2/ encodings/otelp2/common.proto --go_out=.
	protoc -I/usr/local/include -I encodings/otelp2/ encodings/otelp2/metric_data.proto --go_out=.
	protoc -I/usr/local/include -I encodings/otelp2/ encodings/otelp2/telemetry_data.proto --go_out=.
	protoc -I/usr/local/include -I encodings/otelp2/ encodings/otelp2/exchange.proto --go_out=.
	protoc -I/usr/local/include -I encodings/otelp2/ encodings/otelp2/logs.proto --go_out=.
	protoc -I/usr/local/include -I encodings/otelp2/ encodings/otelp2/logs_service.proto --go_out=.
	protoc -I/usr/local/include -I encodings/otelp2/ encodings/otelp2/grpc.proto --go_out=.

gen-baseline:
	protoc -I/usr/local/include -I encodings/baseline/ encodings/baseline/common.proto --go_out=plugins=grpc:encodings/baseline
	protoc -I/usr/local/include -I encodings/baseline/ encodings/baseline/metric_data.proto --go_out=plugins=grpc:encodings/baseline
	protoc -I/usr/local/include -I encodings/baseline/ encodings/baseline/telemetry_data.proto --go_out=plugins=grpc:encodings/baseline
	protoc -I/usr/local/include -I encodings/baseline/ encodings/baseline/exchange.proto --go_out=plugins=grpc:encodings/baseline
	protoc -I/usr/local/include -I encodings/baseline/ encodings/baseline/logs.proto --go_out=plugins=grpc:encodings/baseline
	protoc -I/usr/local/include -I encodings/baseline/ encodings/baseline/logs_service.proto --go_out=plugins=grpc:encodings/baseline
	protoc -I/usr/local/include -I encodings/baseline/ encodings/baseline/grpc.proto --go_out=plugins=grpc:encodings/baseline

	protoc -I/usr/local/include -I encodings/baseline2/ encodings/baseline2/common.proto --go_out=plugins=grpc:encodings/baseline2
	protoc -I/usr/local/include -I encodings/baseline2/ encodings/baseline2/metric_data.proto --go_out=plugins=grpc:encodings/baseline2
	protoc -I/usr/local/include -I encodings/baseline2/ encodings/baseline2/telemetry_data.proto --go_out=plugins=grpc:encodings/baseline2
	protoc -I/usr/local/include -I encodings/baseline2/ encodings/baseline2/exchange.proto --go_out=plugins=grpc:encodings/baseline2
	protoc -I/usr/local/include -I encodings/baseline2/ encodings/baseline2/logs.proto --go_out=plugins=grpc:encodings/baseline2
	protoc -I/usr/local/include -I encodings/baseline2/ encodings/baseline2/logs_service.proto --go_out=plugins=grpc:encodings/baseline2
	protoc -I/usr/local/include -I encodings/baseline2/ encodings/baseline2/grpc.proto --go_out=plugins=grpc:encodings/baseline2

gen-otlptimewrapped:
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
	sudo nice -n -5 ${GO} test -bench . ./encodings -benchtime 5s -benchmem
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
