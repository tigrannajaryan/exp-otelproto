package core

import (
	"encoding/binary"
	"time"
)

func GenerateTraceID(id uint64) []byte {
	var traceID [16]byte
	binary.LittleEndian.PutUint64(traceID[:], id)
	binary.LittleEndian.PutUint64(traceID[8:], 0x123456780abcdef0)
	return traceID[:]
}

func GenerateTraceIDLow(id uint64) uint64 {
	return id
}

func GenerateTraceIDHigh(id uint64) uint64 {
	return 0
}

func GenerateSpanID(id uint64) []byte {
	var spanID [8]byte
	binary.LittleEndian.PutUint64(spanID[:], id)
	return spanID[:]
}

func TimeToTimestamp(t time.Time) uint64 {
	return uint64(t.UnixNano())
}
