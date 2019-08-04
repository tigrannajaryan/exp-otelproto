package core

// Server allows to listen on a port and receive Batches of spans.
type Server interface {
	Listen(endpoint string, onReceive func(batch ExportRequest, spanCount int)) error
	Stop()
}
