package otlp

import (
	"encoding/binary"
	"io"

	"github.com/golang/protobuf/proto"
	v1 "github.com/open-telemetry/opentelemetry-proto/gen/go/collector/trace/v1"
)

func ReadTraceMessage(reader io.Reader) *v1.ExportTraceServiceRequest {
	var lenBytes [8]byte
	n, err := reader.Read(lenBytes[:])
	if n!=8 || err!=nil {
		return nil
	}

	len := binary.LittleEndian.Uint64(lenBytes[:])
	bytes := make([]byte,len)
	n, err = reader.Read(bytes)
	if n!=int(len) || err!=nil {
		return nil
	}

	var msg v1.ExportTraceServiceRequest
	err = proto.Unmarshal(bytes, &msg)
	if err!=nil {
		return nil
	}

	return &msg
}
