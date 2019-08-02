package wsframing

import (
	"bytes"
	"compress/zlib"
	"io/ioutil"
	"log"

	"github.com/golang/protobuf/proto"
	"github.com/tigrannajaryan/exp-otelproto/encodings/traceprotobuf"
)

func Encode(
	request *traceprotobuf.ExportRequest,
	compression traceprotobuf.WSExportRequest_CompressionMethod,
) []byte {
	body, err := proto.Marshal(request)
	if err != nil {
		log.Fatal("cannot encode:", err)
	}

	switch compression {
	case traceprotobuf.WSExportRequest_NO_COMPRESSION:
		break
	case traceprotobuf.WSExportRequest_ZLIB_COMPRESSION:
		var b bytes.Buffer
		w := zlib.NewWriter(&b)
		w.Write(body)
		w.Close()
		body = b.Bytes()
	}

	wsRequest := &traceprotobuf.WSExportRequest{
		Compression: compression,
		Body:        body,
	}
	frameBytes, err := proto.Marshal(wsRequest)
	if err != nil {
		log.Fatal("cannot encode:", err)
	}

	return frameBytes
}

func Decode(frameBytes []byte) *traceprotobuf.ExportRequest {
	var wsRequest traceprotobuf.WSExportRequest
	err := proto.Unmarshal(frameBytes, &wsRequest)
	if err != nil {
		log.Fatal("cannot decode:", err)
	}

	var bodyBytes []byte
	switch wsRequest.Compression {
	case traceprotobuf.WSExportRequest_NO_COMPRESSION:
		bodyBytes = wsRequest.Body

	case traceprotobuf.WSExportRequest_ZLIB_COMPRESSION:
		b := bytes.NewBuffer(wsRequest.Body)
		r, err := zlib.NewReader(b)
		if err != nil {
			log.Fatal("cannot decode:", err)
		}

		bodyBytes, err = ioutil.ReadAll(r)
		if err != nil {
			log.Fatal("cannot decode:", err)
		}
	}

	var request traceprotobuf.ExportRequest
	err = proto.Unmarshal(bodyBytes, &request)
	if err != nil {
		log.Fatal("cannot decode:", err)
	}

	return &request
}
