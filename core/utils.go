package core

import (
	"encoding/binary"
	"time"
)

func GenerateTraceID(id uint64) []byte {
	var traceID [16]byte
	binary.PutUvarint(traceID[:], id)
	return traceID[:]
}

func GenerateSpanID(id uint64) []byte {
	var spanID [8]byte
	binary.PutUvarint(spanID[:], id)
	return spanID[:]
}

func TimeToTimestamp(t time.Time) int64 {
	return t.UnixNano()
}
