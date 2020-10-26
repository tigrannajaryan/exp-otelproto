package otlp_gogo2

import (
	"testing"

	otlpgogo "github.com/tigrannajaryan/exp-otelproto/encodings/otlp_gogo/collector/trace/v1"

	gogoproto "github.com/gogo/protobuf/proto"
	"github.com/golang/protobuf/proto"
	otlptrace "github.com/open-telemetry/opentelemetry-proto/gen/go/collector/trace/v1"
	"github.com/stretchr/testify/assert"
	"github.com/tigrannajaryan/exp-otelproto/encodings/otlp"
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
