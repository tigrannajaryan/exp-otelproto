package sapm

import (
	"context"
	"log"

	jaegerpb "github.com/jaegertracing/jaeger/model"
	sapmclient "github.com/signalfx/sapm-proto/client"

	"github.com/tigrannajaryan/exp-otelproto/core"
)

type Client struct {
	Concurrency int
	client      *sapmclient.Client
}

func (c *Client) Connect(server string) error {
	var err error
	opts := []sapmclient.Option{
		sapmclient.WithEndpoint("http://" + server + "/v2/trace"),
		sapmclient.WithDisabledCompression(),
	}
	if c.Concurrency > 0 {
		opts = append(opts, sapmclient.WithMaxConnections(uint(c.Concurrency)))
	}

	c.client, err = sapmclient.New(opts...)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) Export(batch core.ExportRequest) {
	if err := c.client.Export(context.Background(), []*jaegerpb.Batch{batch.(*jaegerpb.Batch)}); err != nil {
		log.Fatal(err)
	}
}

func (c *Client) Shutdown() {
	c.client.Stop()
}
