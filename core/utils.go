package core

import (
	"encoding/binary"
	"math/rand"
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

var ExampleAttributeNames = []string{
	"enduser.id",
	"enduser.role",
	"enduser.scope",
	"net.host.ip",
	"net.host.name",
	"net.host.port",
	"net.peer.ip",
	"net.peer.name",
	"net.peer.port",
	"net.transport",
	"peer.service",
	"http",
	"grpc",
	"http.client_ip",
	"http.flavor",
	"http.host",
	"host.name",
	"host.port",
	"http.method",
	"http.request_content_length",
	"http.request_content_length_uncompressed",
	"http.response_content_length",
	"http.response_content_length_uncompressed",
	"http.route",
	"http.scheme",
	"http.server_name",
	"http.status_code",
	"http.status_text",
	"http.target",
	"http.url",
	"http.user_agent",
	"db.connection_string",
	"db.cassandra.keyspace",
	"db.hbase.namespace",
	"db.jdbc.driver_classname",
	"db.mongodb.collection",
	"db.mssql.instance_name",
	"db.name",
	"db.operation",
	"db.redis.database_index",
	"db.statement",
	"db.system",
	"db.user",
	"message.compressed_size",
	"message.id",
	"message.type",
	"message.uncompressed_size",
	"rpc.method",
	"rpc.service",
	"rpc.system",
	"message",
	"RECEIVED",
	"SENT",
	"faas.cron",
	"faas.document.collection",
	"faas.document.name",
	"faas.document.operation",
	"faas.document.time",
	"faas.execution",
	"faas.time",
	"faas.trigger",
	"datasource",
	"http",
	"other",
	"pubsub",
	"timer",
	"messaging.conversation_id",
	"messaging.destination",
	"messaging.destination_kind",
	"messaging.message_id",
	"messaging.operation",
	"messaging.message_payload_compressed_size_bytes",
	"messaging.message_payload_size_bytes",
	"messaging.protocol",
	"messaging.protocol_version",
	"messaging.system",
	"messaging.temp_destination",
	"messaging.url",
	"exception",
	"exception.message",
	"exception.stacktrace",
	"exception.type",
}

func GenRandAttrName(random *rand.Rand) string {
	i := random.Intn(len(ExampleAttributeNames))
	return ExampleAttributeNames[i]
}
