package internal

const BatchCount = 1000

/*
func BenchmarkFromOtlpToInternal(b *testing.B) {
	b.StopTimer()
	g := otlp.NewGenerator()

	var batch []*otlp.TraceExportRequest
	for i := 0; i < BatchCount; i++ {
		batch = append(batch,
			g.GenerateSpanBatch(100, 5, 0).(*otlp.TraceExportRequest))
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for i := 0; i < BatchCount; i++ {
			FromOtlp(batch[i])
		}
	}
}

func TestDecodeFromProto(t *testing.T) {
	rs := &otlp.TraceExportRequest{
		ResourceSpans: []*otlp.ResourceSpans{
			{
				Resource: &otlp.Resource{},
				Spans: []*otlp.Span{
					{
						Name: "spanA",
					},
					{},
				},
			},
		},
	}
	b, err := proto.Marshal(rs)
	//fmt.Printf("%x", b)
	require.NoError(t, err)
	buf := proto.NewBuffer(b)

	_, err = ResourceSpansFromBuf(buf, 0)
	assert.NoError(t, err)
}

func TestEncodeFromProto(t *testing.T) {
	tes := &TraceExportRequest{
		ResourceSpans: []*ResourceSpans{
			{
				Resource: &Resource{
					Labels: map[string]*AttributeValue{
						"label1": {stringValue: "val1"},
						"label2": {stringValue: "val2"},
						"label3": {stringValue: "val3"},
						"label4": {stringValue: "val4"},
					},
				},
				Spans: []*Span{
					{
						Name: "spanA",
						Attributes: map[string]*AttributeValue{
							"attribute1": {stringValue: "value1"},
							"attribute2": {stringValue: "value2"},
							"attribute3": {stringValue: "value3"},
							"attribute4": {stringValue: "value4"},
							"attribute5": {stringValue: "value5"},
						},
					},
					{},
				},
			},
		},
	}
	g := otlp.NewGenerator()
	otlpTes := g.GenerateSpanBatch(100, 3, 0).(*otlp.TraceExportRequest)
	//b1, err := proto.Marshal(otlpTes)
	//log.Printf("%x", b1)
	//require.NoError(t, err)
	//
	tes = FromOtlp(otlpTes)

	b2, err := Marshal(tes)
	//log.Printf("%x", b2)
	require.NoError(t, err)

	var tes2 experimental.TraceExportRequest
	err = proto.Unmarshal(b2, &tes2)
	assert.NoError(t, err)
}
*/
