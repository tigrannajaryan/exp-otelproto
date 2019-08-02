package wsframing

import (
	"log"

	"github.com/golang/protobuf/proto"
	"github.com/tigrannajaryan/exp-otelproto/encodings/traceprotobuf"
)

func Encode(request *traceprotobuf.ExportRequest) []byte {
	body, err := proto.Marshal(request)
	if err != nil {
		log.Fatal("cannot encode:", err)
	}
	wsRequest := &traceprotobuf.WSExportRequest{
		Compression: traceprotobuf.WSExportRequest_NO_COMPRESSION,
		Body:        body,
	}

	wsBytes, err := proto.Marshal(wsRequest)
	if err != nil {
		log.Fatal("cannot encode:", err)
	}

	return wsBytes
}

func Decode(bytes []byte) *traceprotobuf.ExportRequest {
	var wsRequest traceprotobuf.WSExportRequest
	err := proto.Unmarshal(bytes, &wsRequest)
	if err != nil {
		log.Fatal("cannot decode:", err)
	}

	var request traceprotobuf.ExportRequest
	err = proto.Unmarshal(wsRequest.Body, &request)
	if err != nil {
		log.Fatal("cannot decode:", err)
	}

	return &request
}
