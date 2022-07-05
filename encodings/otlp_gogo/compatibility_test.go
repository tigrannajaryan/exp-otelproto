package otlp_gogo

import (
	"testing"

	otlpgogo "github.com/tigrannajaryan/exp-otelproto/encodings/otlp_gogo/collector/trace/v1"

	gogoproto "github.com/gogo/protobuf/proto"
	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/assert"
	"github.com/tigrannajaryan/exp-otelproto/encodings/otlp"
	otlptrace "go.opentelemetry.io/proto/otlp/collector/trace/v1"
)

func TestCompatibility(t *testing.T) {
	t.SkipNow()

	gen := otlp.NewGenerator()
	batch := gen.GenerateSpanBatch(3, 10, 5)
	request := batch.(*otlptrace.ExportTraceServiceRequest)
	wire, err := proto.Marshal(request)
	assert.NotNil(t, wire)
	assert.NoError(t, err)

	var gogoRequest otlpgogo.ExportTraceServiceRequest
	err = gogoproto.Unmarshal(wire, &gogoRequest)
	assert.NoError(t, err)

	wire2, err := gogoproto.Marshal(&gogoRequest)
	assert.NoError(t, err)

	assert.EqualValues(t, wire, wire2)
}
