package encodings

import (
	"bytes"
	"compress/zlib"
	"io/ioutil"
	"log"

	"github.com/golang/protobuf/proto"
	"github.com/tigrannajaryan/exp-otelproto/encodings/experimental"
)

type RequestHeader byte

const RequestHeaderSize = 1

const (
	RequestHeader_CompressionMethodMask  = 0x07
	RequestHeader_CompressionMethod_NONE = 0x00
	RequestHeader_CompressionMethod_ZLIB = 0x01
	RequestHeader_CompressionMethod_LZ4  = 0x02
)

// Encode request body into a message of continuous bytes. The message starts with one by
// specifying the length of the RequestHeader, followed by RequestHeader encoded in
// Protobuf format, followed by RequestBody encoded in Protobuf format.
// +--------------------+-------------------------------------------+-----------------------------------------+
// + Header Length Byte | Variable length Header (Protobuf-encoded) | Variable length Body (Protobuf-encoded) |
// +--------------------+-------------------------------------------+-----------------------------------------+
func Encode(
	requestBody *experimental.RequestBody,
	compression experimental.CompressionMethod,
) []byte {
	bodyBytes, err := proto.Marshal(requestBody)
	if err != nil {
		log.Fatal("cannot encode:", err)
	}

	var header RequestHeader
	switch compression {
	case experimental.CompressionMethod_NONE:
		header |= RequestHeader(RequestHeader_CompressionMethod_NONE)
		break
	case experimental.CompressionMethod_ZLIB:
		header |= RequestHeader(RequestHeader_CompressionMethod_NONE)
		var b bytes.Buffer
		w := zlib.NewWriter(&b)
		w.Write(bodyBytes)
		w.Close()
		bodyBytes = b.Bytes()
	}

	b := bytes.NewBuffer(make([]byte, 0, RequestHeaderSize+len(bodyBytes)))
	b.WriteByte(byte(header))
	b.Write(bodyBytes)

	return b.Bytes()
}

// Decode a continuous message of bytes into a RequestBody. This function perform the
// reverse of Encode operation.
func Decode(messageBytes []byte) *experimental.RequestBody {
	header := RequestHeader(messageBytes[0])
	bodyBytes := messageBytes[RequestHeaderSize:]

	switch header & RequestHeader_CompressionMethodMask {
	case RequestHeader_CompressionMethod_NONE:
		break

	case RequestHeader_CompressionMethod_ZLIB:
		b := bytes.NewBuffer(bodyBytes)
		r, err := zlib.NewReader(b)
		if err != nil {
			log.Fatal("cannot decode:", err)
		}

		bodyBytes, err = ioutil.ReadAll(r)
		if err != nil {
			log.Fatal("cannot decode:", err)
		}
	}

	var body experimental.RequestBody
	err := proto.Unmarshal(bodyBytes, &body)
	if err != nil {
		log.Fatal("cannot decode:", err)
	}

	return &body
}
