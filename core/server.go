package core

// Server allows to listen on a port and receive batches of spans.
type Server interface {
	Listen(endpoint string, onReceive func(batch SpanBatch)) error
	Stop()
}
