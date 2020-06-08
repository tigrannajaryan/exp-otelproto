package otlp_gogo3

import (
	"testing"

	gogoproto "github.com/gogo/protobuf/proto"
	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/assert"

	"github.com/tigrannajaryan/exp-otelproto/encodings/experimental2"
	otlpgogo3 "github.com/tigrannajaryan/exp-otelproto/encodings/otlp_gogo3/trace/v1"
)

func TestCompatibility(t *testing.T) {
	//t.SkipNow()
	gen := experimental2.NewGenerator()
	batch := gen.GenerateSpanBatch(3, 10, 5)
	request := batch.(*experimental2.TraceExportRequest)
	rs := request.ResourceSpans
	wire, err := proto.Marshal(rs[0])
	assert.NotNil(t, wire)
	assert.NoError(t, err)

	var gogoRequest otlpgogo3.ResourceSpans
	err = gogoproto.Unmarshal(wire, &gogoRequest)
	assert.NoError(t, err)

	wire2, err := gogoproto.Marshal(&gogoRequest)
	assert.NoError(t, err)

	assert.EqualValues(t, wire, wire2)
}
