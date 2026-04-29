package go_anx_sdk

import (
	"code.anexia.com/se/ks/go-anx-sdk/config"
	"code.anexia.com/se/ks/go-anx-sdk/internal"
	v1 "code.anexia.com/se/ks/go-anx-sdk/v1"
)

type Client struct {
	transport *internal.Transport
}

func NewClient(opts ...config.ClientOption) *Client {
	var cfg config.Config

	for _, opt := range opts {
		opt(&cfg)
	}

	return &Client{
		transport: cfg.CreateTransport(),
	}
}

func (c *Client) V1() *v1.Client {
	return v1.NewClient(c.transport)
}
