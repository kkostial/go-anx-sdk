package v1

import "code.anexia.com/se/ks/go-anx-sdk/internal"

type Client struct {
	transport *internal.Transport
}

func NewClient(transport *internal.Transport) *Client {
	return &Client{
		transport: transport,
	}
}

func (c *Client) Vlans() *VlansClient {
	return NewVlansClient(c.transport)
}

func (c *Client) Locations() *LocationsClient {
	return NewLocationsClient(c.transport)
}
