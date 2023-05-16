package encodings

import (
	"bytes"
	"compress/zlib"
	"encoding/binary"
	"fmt"
	"log"
	"os"
	"runtime"
	"testing"
	"unsafe"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"github.com/klauspost/compress/zstd"
	"github.com/stretchr/testify/assert"
	v1 "go.opentelemetry.io/proto/otlp/collector/trace/v1"

	v15 "github.com/tigrannajaryan/exp-otelproto/encodings/experimental/collector/trace/v1"
	v12 "github.com/tigrannajaryan/exp-otelproto/encodings/experimental/common/v1"
	v14 "github.com/tigrannajaryan/exp-otelproto/encodings/experimental/logs/v1"
	"github.com/tigrannajaryan/exp-otelproto/encodings/experimental2"
	"github.com/tigrannajaryan/exp-otelproto/encodings/otelp2"
	v13 "github.com/tigrannajaryan/exp-otelproto/encodings/otlp_gogo/trace/v1"

	"github.com/tigrannajaryan/exp-otelproto/core"
	experimental "github.com/tigrannajaryan/exp-otelproto/encodings/experimental"
	"github.com/tigrannajaryan/exp-otelproto/encodings/otlp"
)

const spansPerBatch = 100
const metricsPerBatch = spansPerBatch
const logsPerBatch = spansPerBatch

const attrsPerSpans = 3
const eventsPerSpan = 3
const attrsPerLog = attrsPerSpans

var tests = []struct {
	name string
	gen  func() core.Generator
}{
	//{
	//	name: "SepAnyExtValue",
	//	gen:  func() core.Generator { return baseline2.NewGenerator() },
	//},
	{
		name: "OTLP 0.19",
		gen:  func() core.Generator { return otlp.NewGenerator() },
	},
	//{
	//	name: "OTLP",
	//	gen:  func() core.Generator { return baseline.NewGenerator() },
	//},
	{
		name: "OTLP DICT",
		gen:  func() core.Generator { return otelp2.NewGenerator() },
	},
	{
		name: "OTLP MDICT",
		gen:  func() core.Generator { return experimental2.NewGenerator() },
	},
	//{
	//	name: "Proposed",
	//	gen:  func() core.Generator { return baseline.NewGenerator() },
	//},
	//{
	//	name: "Alternate",
	//	gen:  func() core.Generator { return experimental.NewGenerator() },
	//},
	//{
	//	name: "Current(Gogo)",
	//	gen:  func() core.Generator { return otlp_gogo.NewGenerator() },
	//},
	//{
	//	name: "gogoCustom",
	//	gen:  func() core.Generator { return otlp_gogo2.NewGenerator() },
	//},
	//{
	//	name: "Proposed(Gogo)",
	//	gen:  func() core.Generator { return otlp_gogo3.NewGenerator() },
	//},
	//{
	//	name: "OpenCensus",
	//	gen:  func() core.Generator { return octraceprotobuf.NewGenerator() },
	//},
	//// These are historical experiments. Uncomment if interested to see results.
	//{
	//	name: "OC+AttrAsMap",
	//	gen:  func() core.Generator { return traceprotobuf.NewGenerator() },
	//},
	//{
	//	name: "OC+AttrAsList+TimeWrapped",
	//	gen:  func() core.Generator { return otlptimewrapped.NewGenerator() },
	//},
}

var batchTypes = []struct {
	name     string
	batchGen func(gen core.Generator) []core.ExportRequest
}{
	{name: "Logs", batchGen: generateLogBatches},
	{name: "Trace/Attribs", batchGen: generateAttrBatches},
	//{name: "Trace/Events", batchGen: generateTimedEventBatches},
	//{name: "Metric/Int64", batchGen: generateMetricInt64Batches},
	//{name: "Metric/Summary", batchGen: generateMetricSummaryBatches},
	//{name: "Metric/Histogram", batchGen: generateMetricHistogramBatches},
	//{name: "Metric/HistogramSeries", batchGen: generateMetricHistogramSeriesBatches},
	//{name: "Metric/Mix", batchGen: generateMetricOneBatches},
	{name: "Metric/MixSeries", batchGen: generateMetricSeriesBatches},
}

const BatchCount = 1

func BenchmarkGenerate(b *testing.B) {
	b.SkipNow()

	for _, batchType := range batchTypes {
		for _, test := range tests {
			b.Run(
				test.name+"/"+batchType.name, func(b *testing.B) {
					gen := test.gen()
					for i := 0; i < b.N; i++ {
						batches := batchType.batchGen(gen)
						if batches == nil {
							// Unsupported test type and batch type combination.
							b.SkipNow()
							return
						}
					}
				},
			)
		}
		fmt.Println("")
	}
}

func BenchmarkEncode(b *testing.B) {

	for _, batchType := range batchTypes {
		for _, test := range tests {
			b.Run(
				test.name+"/"+batchType.name, func(b *testing.B) {
					b.StopTimer()
					gen := test.gen()
					batches := batchType.batchGen(gen)
					if batches == nil {
						// Unsupported test type and batch type combination.
						b.SkipNow()
						return
					}

					runtime.GC()
					b.StartTimer()
					for i := 0; i < b.N; i++ {
						for _, batch := range batches {
							encode(batch)
						}
					}
				},
			)
		}
		fmt.Println("")
	}
}

func BenchmarkDecode(b *testing.B) {
	for _, batchType := range batchTypes {
		for _, test := range tests {
			b.Run(
				test.name+"/"+batchType.name, func(b *testing.B) {
					b.StopTimer()
					batches := batchType.batchGen(test.gen())
					if batches == nil {
						// Unsupported test type and batch type combination.
						b.SkipNow()
						return
					}

					var encodedBytes [][]byte
					for _, batch := range batches {
						encodedBytes = append(encodedBytes, encode(batch))
					}

					runtime.GC()
					b.StartTimer()
					for i := 0; i < b.N; i++ {
						for j, bytes := range encodedBytes {
							decode(bytes, batches[j].(proto.Message))
						}
					}
				},
			)
		}
		fmt.Println("")
	}
}

/*
func BenchmarkEncodeInternalToOtlp2Step(b *testing.B) {

	b.StopTimer()
	g := otlp.NewGenerator()
	batches := generateAttrBatches(g)
	if batches == nil {
		// Unsupported test type and batch type combination.
		b.SkipNow()
		return
	}

	var intbatch []*internal.TraceExportRequest
	for _, b := range batches {
		intbatch = append(intbatch, internal.FromOtlp(b.(*otlp.TraceExportRequest)))
	}

	runtime.GC()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for _, ib := range intbatch {
			ot := internal.ToOtlp(ib)
			encode(ot)
		}
	}

}

func BenchmarkEncodeIntOtlpToOtlp(b *testing.B) {

	b.StopTimer()
	g := otlp.NewGenerator()
	batches := generateAttrBatches(g)
	if batches == nil {
		// Unsupported test type and batch type combination.
		b.SkipNow()
		return
	}

	var intbatch []*intotlp.TraceExportRequest
	for _, b := range batches {
		intbatch = append(intbatch, intotlp.FromOtlp(b.(*otlp.TraceExportRequest)))
	}

	runtime.GC()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for _, ib := range intbatch {
			ot := intotlp.ToOtlp(ib)
			encode(ot)
		}
	}

}

func BenchmarkEncodeInternalDirectToOtlp(b *testing.B) {

	b.StopTimer()
	g := otlp.NewGenerator()
	batches := generateAttrBatches(g)
	if batches == nil {
		// Unsupported test type and batch type combination.
		b.SkipNow()
		return
	}

	var intbatch []*internal.TraceExportRequest
	for _, b := range batches {
		intbatch = append(intbatch, internal.FromOtlp(b.(*otlp.TraceExportRequest)))
	}

	runtime.GC()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for _, ib := range intbatch {
			_, err := internal.Marshal(ib)
			assert.NoError(b, err)
		}
	}

}

func BenchmarkDecodeOtlpToInternal2Step(b *testing.B) {
	b.StopTimer()
	g := otlp.NewGenerator()
	batches := generateAttrBatches(g)
	if batches == nil {
		// Unsupported test type and batch type combination.
		b.SkipNow()
		return
	}

	var encodedBytes [][]byte
	for _, batch := range batches {
		encodedBytes = append(encodedBytes, encode(batch))
	}

	runtime.GC()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for _, bytes := range encodedBytes {
			var tep otlp.TraceExportRequest
			decode(bytes, &tep)
			internal.FromOtlp(&tep)
		}
	}
}

func BenchmarkDecodeOtlpToIntOtlp(b *testing.B) {
	b.StopTimer()
	g := otlp.NewGenerator()
	batches := generateAttrBatches(g)
	if batches == nil {
		// Unsupported test type and batch type combination.
		b.SkipNow()
		return
	}

	var encodedBytes [][]byte
	for _, batch := range batches {
		encodedBytes = append(encodedBytes, encode(batch))
	}

	runtime.GC()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for _, bytes := range encodedBytes {
			var tep otlp.TraceExportRequest
			decode(bytes, &tep)
			intotlp.FromOtlp(&tep)
		}
	}
}
*/

func generateAttrBatches(gen core.Generator) []core.ExportRequest {
	var batches []core.ExportRequest
	for i := 0; i < BatchCount; i++ {
		batches = append(batches, gen.GenerateSpanBatch(spansPerBatch, attrsPerSpans, 0))
	}
	return batches
}

func generateTimedEventBatches(gen core.Generator) []core.ExportRequest {
	var batches []core.ExportRequest
	for i := 0; i < BatchCount; i++ {
		batches = append(batches, gen.GenerateSpanBatch(spansPerBatch, 3, eventsPerSpan))
	}
	return batches
}

func generateLogBatches(gen core.Generator) []core.ExportRequest {
	var batches []core.ExportRequest
	for i := 0; i < BatchCount; i++ {
		batch := gen.GenerateLogBatch(logsPerBatch, attrsPerLog)
		if batch == nil {
			return nil
		}
		batches = append(batches, batch)
	}
	return batches
}

func generateMetricOneBatches(gen core.Generator) []core.ExportRequest {
	var batches []core.ExportRequest
	for i := 0; i < BatchCount; i++ {
		batch := gen.GenerateMetricBatch(metricsPerBatch, 1, true, true, true)
		if batch == nil {
			return nil
		}
		batches = append(batches, batch)
	}
	return batches
}

func generateMetricSeriesBatches(gen core.Generator) []core.ExportRequest {
	var batches []core.ExportRequest
	for i := 0; i < BatchCount; i++ {
		batch := gen.GenerateMetricBatch(metricsPerBatch, 5, true, true, true)
		if batch == nil {
			return nil
		}
		batches = append(batches, batch)
	}
	return batches
}

func generateMetricInt64Batches(gen core.Generator) []core.ExportRequest {
	var batches []core.ExportRequest
	for i := 0; i < BatchCount; i++ {
		batch := gen.GenerateMetricBatch(metricsPerBatch, 1, true, false, false)
		if batch == nil {
			return nil
		}
		batches = append(batches, batch)
	}
	return batches
}

func generateMetricHistogramBatches(gen core.Generator) []core.ExportRequest {
	var batches []core.ExportRequest
	for i := 0; i < BatchCount; i++ {
		batch := gen.GenerateMetricBatch(metricsPerBatch, 1, false, true, false)
		if batch == nil {
			return nil
		}
		batches = append(batches, batch)
	}
	return batches
}

func generateMetricHistogramSeriesBatches(gen core.Generator) []core.ExportRequest {
	var batches []core.ExportRequest
	for i := 0; i < BatchCount; i++ {
		batch := gen.GenerateMetricBatch(metricsPerBatch, 5, false, true, false)
		if batch == nil {
			return nil
		}
		batches = append(batches, batch)
	}
	return batches
}

func generateMetricSummaryBatches(gen core.Generator) []core.ExportRequest {
	var batches []core.ExportRequest
	for i := 0; i < BatchCount; i++ {
		batch := gen.GenerateMetricBatch(metricsPerBatch, 1, false, false, true)
		if batch == nil {
			return nil
		}
		batches = append(batches, batch)
	}
	return batches
}

func encode(request core.ExportRequest) []byte {
	bytes, err := proto.Marshal(request.(proto.Message))
	if err != nil {
		log.Fatal(err)
	}
	return bytes
}

func decode(bytes []byte, pb proto.Message) {
	err := proto.Unmarshal(bytes, pb)
	if err != nil {
		log.Fatal(err)
	}
}

func TestEncodeSize(t *testing.T) {

	const batchSize = spansPerBatch

	variation := []struct {
		name                     string
		genFunc                  func(gen core.Generator) core.ExportRequest
		firstUncompessedSize     int
		firstUncompessedJSONSize int
		firstZlibedSize          int
		firstZlibedJSONSize      int
		firstZstdedSize          int
	}{
		{
			name: "Logs",
			genFunc: func(gen core.Generator) core.ExportRequest {
				return gen.GenerateLogBatch(batchSize, 4)
			},
		},
		{
			name: "Trace/Attribs",
			genFunc: func(gen core.Generator) core.ExportRequest {
				return gen.GenerateSpanBatch(batchSize, attrsPerSpans, 0)
			},
		},
		{
			name: "Trace/Events",
			genFunc: func(gen core.Generator) core.ExportRequest {
				return gen.GenerateSpanBatch(batchSize, 0, eventsPerSpan)
			},
		},
		{
			name: "Metric/Histogram",
			genFunc: func(gen core.Generator) core.ExportRequest {
				return gen.GenerateMetricBatch(batchSize, 1, false, true, false)
			},
		},
		{
			name: "Metric/MixOne",
			genFunc: func(gen core.Generator) core.ExportRequest {
				return gen.GenerateMetricBatch(batchSize, 1, true, true, true)
			},
		},
		{
			name: "Metric/MixSeries",
			genFunc: func(gen core.Generator) core.ExportRequest {
				return gen.GenerateMetricBatch(batchSize, 5, true, true, true)
			},
		},
	}

	fmt.Println("===== Encoded sizes")

	for _, v := range variation {
		fmt.Println("Encoding                       Uncomp/bin[Improved] Uncomp/json[Improved]  zlib/bin[Improved]  zlib/json[Improved] zstd/bin[Improved]")
		for _, test := range tests {
			t.Run(
				test.name, func(t *testing.T) {
					gen := test.gen()

					batch := v.genFunc(gen)
					if batch == nil {
						// Skip this case.
						return
					}

					bodyBytes, err := proto.Marshal(batch.(proto.Message))
					if err != nil {
						log.Fatal(err)
					}

					zlibedBytes := doZlib(bodyBytes)
					zstdedBytes := compressZstd(bodyBytes)

					uncompressedSize := len(bodyBytes)
					zlibedSize := len(zlibedBytes)
					zstdedSize := len(zstdedBytes)

					uncompressedRatioStr := "[1.000]"
					zlibedRatioStr := "[1.000]"
					zstdedRatioStr := "[1.000]"

					if v.firstUncompessedSize == 0 {
						v.firstUncompessedSize = uncompressedSize
					} else {
						uncompressedRatioStr = fmt.Sprintf(
							"[%1.3f]", float64(v.firstUncompessedSize)/float64(uncompressedSize),
						)
					}

					if v.firstZlibedSize == 0 {
						v.firstZlibedSize = zlibedSize
					} else {
						zlibedRatioStr = fmt.Sprintf(
							"[%1.3f]", float64(v.firstZlibedSize)/float64(zlibedSize),
						)
					}

					if v.firstZstdedSize == 0 {
						v.firstZstdedSize = zstdedSize
					} else {
						zstdedRatioStr = fmt.Sprintf(
							"[%1.3f]", float64(v.firstZstdedSize)/float64(zstdedSize),
						)
					}

					m := jsonpb.Marshaler{}
					str, err := m.MarshalToString(batch.(proto.Message))

					uncompressedJSONSize := len(str)

					uncompressedJSONRatioStr := "[1.000]"

					if v.firstUncompessedJSONSize == 0 {
						v.firstUncompessedJSONSize = uncompressedJSONSize
					} else {
						uncompressedJSONRatioStr = fmt.Sprintf(
							"[%1.3f]",
							float64(v.firstUncompessedJSONSize)/float64(uncompressedJSONSize),
						)
					}

					zlibedJSON := doZlib([]byte(str))
					zlibedJSONSize := len(zlibedJSON)
					zlibedJSONRatioStr := "[1.000]"
					if v.firstZlibedJSONSize == 0 {
						v.firstZlibedJSONSize = zlibedJSONSize
					} else {
						zlibedJSONRatioStr = fmt.Sprintf(
							"[%1.3f]", float64(v.firstZlibedJSONSize)/float64(zlibedJSONSize),
						)
					}

					fmt.Printf(
						"%-31v%6d by%8s  %7d by  %8s    %5d by%8s    %5d by%8s    %5d by%8s\n",
						test.name+"/"+v.name,
						uncompressedSize,
						uncompressedRatioStr,
						uncompressedJSONSize,
						uncompressedJSONRatioStr,
						zlibedSize,
						zlibedRatioStr,
						zlibedJSONSize,
						zlibedJSONRatioStr,
						zstdedSize,
						zstdedRatioStr,
					)

				},
			)
		}
		fmt.Println("")
	}
}

func countTracesAndSpans(msg *v1.ExportTraceServiceRequest) (int, int) {
	traces := map[string]bool{}
	spans := 0
	for _, rss := range msg.ResourceSpans {
		for _, ss := range rss.ScopeSpans {
			spans += len(ss.Spans)
			for _, s := range ss.Spans {
				traces[string(s.TraceId)] = true
			}
		}
	}
	return len(traces), spans
}

func TestEncodeSizeFromFile(t *testing.T) {

	var tests = []struct {
		name       string
		translator func() core.SpanTranslator
	}{
		//{
		//	name: "SepAnyExtValue",
		//	gen:  func() core.Generator { return baseline2.NewGenerator() },
		//},
		//{
		//	name: "OTLP 0.4",
		//	gen:  func() core.Generator { return otlp.NewGenerator() },
		//},
		{
			name:       "OTLP",
			translator: func() core.SpanTranslator { return &otlp.SpanTranslator{} },
		},
		{
			name:       "OTLP DICT",
			translator: func() core.SpanTranslator { return otelp2.NewSpanTranslator() },
		},
		//{
		//	name: "MoreFieldsinAKV",
		//	gen:  func() core.Generator { return experimental.NewGenerator() },
		//},
		//{
		//	name: "Proposed",
		//	gen:  func() core.Generator { return baseline.NewGenerator() },
		//},
		//{
		//	name: "Current(Gogo)",
		//	gen:  func() core.Generator { return otlp_gogo.NewGenerator() },
		//},
		//{
		//	name: "gogoCustom",
		//	gen:  func() core.Generator { return otlp_gogo2.NewGenerator() },
		//},
		//{
		//	name: "Proposed(Gogo)",
		//	gen:  func() core.Generator { return otlp_gogo3.NewGenerator() },
		//},
		//{
		//	name: "OpenCensus",
		//	gen:  func() core.Generator { return octraceprotobuf.NewGenerator() },
		//},
		//// These are historical experiments. Uncomment if interested to see results.
		//{
		//	name: "OC+AttrAsMap",
		//	gen:  func() core.Generator { return traceprotobuf.NewGenerator() },
		//},
		//{
		//	name: "OC+AttrAsList+TimeWrapped",
		//	gen:  func() core.Generator { return otlptimewrapped.NewGenerator() },
		//},
	}

	fmt.Println("===== Encoded sizes")

	firstUncompessedSize := 0
	firstZlibedSize := 0
	firstZstdedSize := 0

	for _, test := range tests {
		fmt.Println("Encoding                       Uncompressed  Improved      Compressed  Improved      Compressed  Improved")
		t.Run(
			test.name, func(t *testing.T) {
				translator := test.translator()

				f, err := os.Open("testdata/traces.protobuf")
				assert.NoError(t, err)

				uncompressedSize := 0
				zlibedSize := 0
				zstdedSize := 0
				totalTraces := 0
				totalSpans := 0

				for {
					msg := otlp.ReadTraceMessage(f)
					if msg == nil {
						break
					}
					traces, spans := countTracesAndSpans(msg)
					totalTraces += traces
					totalSpans += spans

					batch := translator.TranslateSpans(msg)
					if batch == nil {
						// Skip this case.
						return
					}

					bodyBytes, err := proto.Marshal(batch.(proto.Message))
					if err != nil {
						log.Fatal(err)
					}

					zlibedBytes := doZlib(bodyBytes)
					zstdedBytes := compressZstd(bodyBytes)

					uncompressedSize += len(bodyBytes)
					zlibedSize += len(zlibedBytes)
					zstdedSize += len(zstdedBytes)
				}

				uncompressedRatioStr := "[1.000]"
				zlibedRatioStr := "[1.000]"
				zstdedRatioStr := "[1.000]"

				if firstUncompessedSize == 0 {
					firstUncompessedSize = uncompressedSize
				} else {
					uncompressedRatioStr = fmt.Sprintf(
						"[%1.3f]", float64(firstUncompessedSize)/float64(uncompressedSize),
					)
				}

				if firstZlibedSize == 0 {
					firstZlibedSize = zlibedSize
				} else {
					zlibedRatioStr = fmt.Sprintf(
						"[%1.3f]", float64(firstZlibedSize)/float64(zlibedSize),
					)
				}

				if firstZstdedSize == 0 {
					firstZstdedSize = zstdedSize
				} else {
					zstdedRatioStr = fmt.Sprintf(
						"[%1.3f]", float64(firstZstdedSize)/float64(zstdedSize),
					)
				}

				fmt.Printf(
					"%-31v %6d bytes%8s, zlib %5d bytes%8s, zstd %5d bytes%8s\n",
					test.name,
					uncompressedSize,
					uncompressedRatioStr,
					zlibedSize,
					zlibedRatioStr,
					zstdedSize,
					zstdedRatioStr,
					//totalTraces,
					//totalSpans,
				)

			},
		)
		fmt.Println("")
	}
}

func doZlib(input []byte) []byte {
	var b bytes.Buffer
	w := zlib.NewWriter(&b)
	w.Write(input)
	w.Close()
	return b.Bytes()
}

// Create a writer that caches compressors.
// For this operation type we supply a nil Reader.
var zstdEncoder, _ = zstd.NewWriter(nil, zstd.WithEncoderLevel(zstd.SpeedBestCompression))

// Create a reader that caches decompressors.
// For this operation type we supply a nil Reader.
var zstdDecoder, _ = zstd.NewReader(nil, zstd.WithDecoderConcurrency(1))

func compressZstd(input []byte) []byte {
	return zstdEncoder.EncodeAll(input, make([]byte, 0, len(input)))
}

func decompressZstd(input []byte) []byte {
	b, err := zstdDecoder.DecodeAll(input, nil)
	if err != nil {
		log.Fatal(err)
	}
	return b
}

func BenchmarkEndianness(b *testing.B) {
	var tests = []struct {
		name  string
		order binary.ByteOrder
	}{
		{
			name:  "Little",
			order: binary.LittleEndian,
		},
		{
			name:  "Big",
			order: binary.BigEndian,
		},
	}

	for _, test := range tests {
		b.Run(
			test.name, func(b *testing.B) {
				b.StartTimer()
				var spanID [8]byte
				for i := 0; i < b.N; i++ {
					test.order.PutUint64(spanID[:], uint64(i))
				}
			},
		)
	}
}

func TestSizes(t *testing.T) {
	akv := v12.KeyValue{}
	log.Printf("AttributeKeyValue is %d bytes", unsafe.Sizeof(akv))
	log.Printf("AttributeKeyValue.Key is %d bytes", unsafe.Sizeof(akv.Key))

	log.Printf("Span is %d bytes", unsafe.Sizeof(v13.Span{}))
	log.Printf("LogRecord is %d bytes", unsafe.Sizeof(v14.LogRecord{}))
}

func createAKV() *v12.KeyValue {
	for i := 0; i < 1; i++ {
		return &v12.KeyValue{}
	}
	return nil
}

func createAV() *v12.KeyValue {
	for i := 0; i < 1; i++ {
		return &v12.KeyValue{}
	}
	return nil
}

func BenchmarkAttributeKeyValueSize(b *testing.B) {
	for i := 0; i < b.N; i++ {
		createAKV()
	}
}

func BenchmarkAttributeValueSize(b *testing.B) {
	for i := 0; i < b.N; i++ {
		createAV()
	}
}

func TestJson(t *testing.T) {
	g := experimental.NewGenerator()
	b := g.GenerateSpanBatch(1, 1, 1)

	m := jsonpb.Marshaler{}
	str, err := m.MarshalToString(b.(*v15.ExportTraceServiceRequest))
	assert.NoError(t, err)
	fmt.Println("ShortKeys, OrigName=false")
	fmt.Println(str)

	err = jsonpb.UnmarshalString(str, b.(proto.Message))
	assert.NoError(t, err)

	m.OrigName = true
	str, err = m.MarshalToString(b.(*v15.ExportTraceServiceRequest))
	assert.NoError(t, err)
	fmt.Println("ShortKeys, OrigName=true")
	fmt.Println(str)

	err = jsonpb.UnmarshalString(str, b.(proto.Message))
	assert.NoError(t, err)

	g2 := otlp.NewGenerator()
	b2 := g2.GenerateSpanBatch(1, 1, 1)

	m2 := jsonpb.Marshaler{}
	str, err = m2.MarshalToString(b2.(*v1.ExportTraceServiceRequest))
	assert.NoError(t, err)
	fmt.Println("OTLP, OrigName=false")
	fmt.Println(str)

	err = jsonpb.UnmarshalString(str, b2.(proto.Message))
	assert.NoError(t, err)

	m2.OrigName = true
	str, err = m2.MarshalToString(b2.(*v1.ExportTraceServiceRequest))
	assert.NoError(t, err)
	fmt.Println("OTLP, OrigName=true")
	fmt.Println(str)

	err = jsonpb.UnmarshalString(str, b2.(proto.Message))
	assert.NoError(t, err)
}
