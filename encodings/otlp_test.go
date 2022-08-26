package encodings

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"os"
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/require"
	v1 "go.opentelemetry.io/proto/otlp/collector/trace/v1"

	"github.com/tigrannajaryan/exp-otelproto/encodings/otlp"
)

var gzipWriter = gzip.NewWriter(nil)
var gzipReader = createGzipReader()

func createGzipReader() *gzip.Reader {
	gzipReader, err := gzip.NewReader(bytes.NewBuffer(compressGzip(nil)))
	if err != nil {
		log.Fatal(err)
	}
	return gzipReader
}

func compressGzip(input []byte) []byte {
	var b bytes.Buffer
	gzipWriter.Reset(&b)
	gzipWriter.Write(input)
	gzipWriter.Close()
	return b.Bytes()
}

func decompressGzip(input []byte) []byte {
	b := bytes.NewBuffer(input)
	err := gzipReader.Reset(b)
	if err != nil {
		log.Fatal(err)
	}
	ub, err := io.ReadAll(gzipReader)
	if err != nil {
		log.Fatal(err)
	}
	return ub
}

func TestOTLPCompression(t *testing.T) {

	var tests = []struct {
		name       string
		compress   func(input []byte) []byte
		decompress func(input []byte) []byte
	}{
		{
			name:     "gzip",
			compress: compressGzip,
		},
		{
			name:     "zstd",
			compress: compressZstd,
		},
	}

	fmt.Println("===== Encoded sizes")

	for _, test := range tests {
		t.Run(
			test.name, func(t *testing.T) {
				f, err := os.Open("testdata/traces.protobuf")
				require.NoError(t, err)

				messageCount := 0
				uncompressedSize := 0
				compressedSize := 0

				for {
					msg := otlp.ReadTraceMessage(f)
					if msg == nil {
						break
					}
					messageCount++

					bodyBytes, err := proto.Marshal(msg)
					if err != nil {
						log.Fatal(err)
					}

					compressedBytes := test.compress(bodyBytes)

					uncompressedSize += len(bodyBytes)
					compressedSize += len(compressedBytes)
				}

				compressedRatioStr := fmt.Sprintf(
					"[%1.3f]", float64(uncompressedSize)/float64(compressedSize),
				)

				fmt.Printf(
					"%-7v %4d messages, %6d bytes, compressed %5d bytes%8s\n",
					test.name,
					messageCount,
					uncompressedSize,
					compressedSize,
					compressedRatioStr,
				)

			},
		)
		fmt.Println("")
	}
}

func BenchmarkOTLPCompression(b *testing.B) {

	var tests = []struct {
		name       string
		compress   func(input []byte) []byte
		decompress func(input []byte) []byte
	}{
		{
			name:       "gzip",
			compress:   compressGzip,
			decompress: decompressGzip,
		},
		{
			name:       "zstd",
			compress:   compressZstd,
			decompress: decompressZstd,
		},
	}

	f, err := os.Open("testdata/traces.protobuf")
	require.NoError(b, err)

	var msgs []proto.Message

	for {
		msg := otlp.ReadTraceMessage(f)
		if msg == nil {
			break
		}
		msgs = append(msgs, msg)
	}

	b.Run(
		"Marshal Protobuf only", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				for _, msg := range msgs {
					_, err := proto.Marshal(msg)
					if err != nil {
						log.Fatal(err)
					}
				}
			}
		},
	)

	for _, test := range tests {
		b.Run(
			"Marshal+Compress:"+test.name, func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					for _, msg := range msgs {
						bodyBytes, err := proto.Marshal(msg)
						if err != nil {
							log.Fatal(err)
						}
						test.compress(bodyBytes)
					}
				}
			},
		)
	}

	for _, test := range tests {
		b.Run(
			"Decompress:"+test.name+"+Unmarshal", func(b *testing.B) {

				var compressedBytes [][]byte

				for _, msg := range msgs {

					bodyBytes, err := proto.Marshal(msg)
					if err != nil {
						log.Fatal(err)
					}

					cb := test.compress(bodyBytes)
					compressedBytes = append(compressedBytes, cb)
				}

				b.ResetTimer()
				for i := 0; i < b.N; i++ {
					for _, cb := range compressedBytes {
						bodyBytes := test.decompress(cb)
						var msg v1.ExportTraceServiceRequest
						err := proto.Unmarshal(bodyBytes, &msg)
						if err != nil {
							log.Fatal(err)
						}
					}
				}
			},
		)
	}

	b.Run(
		"Unmarshal Protobuf only", func(b *testing.B) {

			var uncompressedBytes [][]byte

			for _, msg := range msgs {
				bodyBytes, err := proto.Marshal(msg)
				if err != nil {
					log.Fatal(err)
				}
				uncompressedBytes = append(uncompressedBytes, bodyBytes)
			}

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				for _, bodyBytes := range uncompressedBytes {
					var msg v1.ExportTraceServiceRequest
					err := proto.Unmarshal(bodyBytes, &msg)
					if err != nil {
						log.Fatal(err)
					}
				}
			}
		},
	)
}
