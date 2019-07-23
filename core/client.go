package core

// Client can connect to a server and send a batch of spans.
type Client interface {
	Connect(server string) error
	SendBatch(batch SpanBatch)
}