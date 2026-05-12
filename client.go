package go_anx_sdk

import (
	"code.anexia.com/se/ks/go-anx-sdk/config"
	"code.anexia.com/se/ks/go-anx-sdk/internal"
	v1 "code.anexia.com/se/ks/go-anx-sdk/v1"
)

// Client is the entry point to the anexia api.
type Client struct {
	transport *internal.Transport
}

// NewClient creates a new anexia go sdk client with the provided options.
func NewClient(opts ...config.ClientOption) *Client {
	var cfg config.Config

	for _, opt := range opts {
		opt(&cfg)
	}

	return &Client{
		transport: cfg.CreateTransport(),
	}
}

// V1 returns an entry point to anexia api v1 api clients.
func (c *Client) V1() *v1.Client {
	return v1.NewClient(c.transport)
}
