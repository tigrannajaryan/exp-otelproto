package otlp_gogo2

import (
	"testing"

	"github.com/tigrannajaryan/exp-otelproto/encodings/baseline"
	otlpgogo2 "github.com/tigrannajaryan/exp-otelproto/encodings/otlp_gogo2/trace/v1"

	gogoproto "github.com/gogo/protobuf/proto"
	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/assert"
)

func TestCompatibility(t *testing.T) {
	//t.SkipNow()

	gen := baseline.NewGenerator()
	batch := gen.GenerateSpanBatch(3, 10, 5)
	request := batch.(*baseline.TraceExportRequest)
	rs := request.ResourceSpans
	wire, err := proto.Marshal(rs[0])
	assert.NotNil(t, wire)
	assert.NoError(t, err)

	var gogoRequest otlpgogo2.ResourceSpans
	err = gogoproto.Unmarshal(wire, &gogoRequest)
	assert.NoError(t, err)

	wire2, err := gogoproto.Marshal(&gogoRequest)
	assert.NoError(t, err)

	assert.EqualValues(t, wire, wire2)
}
