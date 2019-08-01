package traceflatbuffers

import (
	"encoding/binary"
	"sync/atomic"
	"time"

	flatbuffers "github.com/google/flatbuffers/go"

	"github.com/golang/protobuf/ptypes/timestamp"

	"github.com/tigrannajaryan/exp-otelproto/core"
)

// Generator allows to generate a ExportRequest.
type Generator struct {
	tracesSent uint64
	spansSent  uint64
}

func (g *Generator) GenerateBatch(spansPerBatch int, attrsPerSpan int) core.ExportRequest {
	traceID := atomic.AddUint64(&g.tracesSent, 1)

	builder := flatbuffers.NewBuilder(1024)
	const N = 100

	var spans [N]flatbuffers.UOffsetT

	for i := 0; i < N; i++ {
		startTime := time.Now()

		spanID := atomic.AddUint64(&g.spansSent, 1)

		name := builder.CreateString("load-generator-span")

		//TraceId := generateTraceID(traceID)
		//SpanId := generateSpanID(spanID)

		seqNumKey := builder.CreateString("load_generator.span_seq_num")

		Int64ValueStart(builder)
		Int64ValueAddInt64Value(builder, int64(spanID))
		seqNumVal := Int64ValueEnd(builder)

		AttributeStart(builder)
		AttributeAddKey(builder, seqNumKey)
		AttributeAddValueType(builder, AttributeValueInt64Value)
		AttributeAddValue(builder, seqNumVal)
		attr1 := AttributeEnd(builder)

		seqNumKey = builder.CreateString("load_generator.trace_seq_num")

		Int64ValueStart(builder)
		Int64ValueAddInt64Value(builder, int64(traceID))
		seqNumVal = Int64ValueEnd(builder)

		AttributeStart(builder)
		AttributeAddKey(builder, seqNumKey)
		AttributeAddValueType(builder, AttributeValueInt64Value)
		AttributeAddValue(builder, seqNumVal)
		attr2 := AttributeEnd(builder)

		SpanStartAttributesVector(builder, 2)
		builder.PrependUOffsetT(attr1)
		builder.PrependUOffsetT(attr2)
		attrs := builder.EndVector(2)

		SpanStart(builder)

		SpanAddTraceIdLo(builder, traceID)
		SpanAddTraceIdHi(builder, 0)
		SpanAddSpanId(builder, spanID)
		SpanAddName(builder, name)
		SpanAddKind(builder, SpanKindCLIENT)
		SpanAddStartTime(builder, startTime.UnixNano())
		SpanAddEndTime(builder, startTime.Add(time.Duration(time.Millisecond)).UnixNano())

		SpanAddAttributes(builder, attrs)

		spans[i] = SpanEnd(builder)

		// Create a span.
		//span := &Span{
		//	Attributes: &Span_Attributes{
		//		AttributeMap: map[string]*AttributeValue{
		//			"load_generator.span_seq_num":  &AttributeValue{Value: &AttributeValue_IntValu
		//			e{IntValue: int64(spanID)}},
		//			"load_generator.trace_seq_num": &AttributeValue{Value: &AttributeValue_IntValue{IntValue: int64(traceID)}},
		//		},
		//	},
		//}

		// Append attributes.
		//for k, v := range g.options.Attributes {
		//	span.Attributes[k] = v
		//}

		//batch.Spans = append(batch.Spans, span)
	}

	SpanBatchStartSpansVector(builder, N)

	for i := N - 1; i >= 0; i-- {
		builder.PrependUOffsetT(spans[i])
	}

	spansField := builder.EndVector(N)

	SpanBatchStart(builder)
	SpanBatchAddSpans(builder, spansField)

	spanBatch := SpanBatchEnd(builder)
	builder.Finish(spanBatch)
	buf := builder.FinishedBytes()

	return &BatchRequest{EncodedSpans: buf}
}

func generateTraceID(id uint64) []byte {
	var traceID [16]byte
	binary.PutUvarint(traceID[:], id)
	return traceID[:]
}

func generateSpanID(id uint64) []byte {
	var spanID [8]byte
	binary.PutUvarint(spanID[:], id)
	return spanID[:]
}

func timeToTimestamp(t time.Time) *timestamp.Timestamp {
	nanoTime := t.UnixNano()
	return &timestamp.Timestamp{
		Seconds: nanoTime / 1e9,
		Nanos:   int32(nanoTime % 1e9),
	}
}
