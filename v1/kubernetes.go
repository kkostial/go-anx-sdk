package v1

import (
	"context"

	"code.anexia.com/se/ks/go-anx-sdk/internal"
	"code.anexia.com/se/ks/go-anx-sdk/paging"
)

// ClusterListParams defines the available parameters for the cluster list endpoint.
type ClusterListParams struct {
	Page  int `url:"page,omitempty"`
	Limit int `url:"limit,omitempty"`
}

// ClusterListItem is an item in the cluster list response.
type ClusterListItem struct {
	Identifier string `json:"identifier"`
	Name       string `json:"name"`
}

// ClustersClient is an api client for managing clusters.
type ClustersClient struct {
	transport *internal.Transport
}

// NewClustersClient creates a new cluster client.
func NewClustersClient(transport *internal.Transport) *ClustersClient {
	return &ClustersClient{
		transport: transport,
	}
}

// List returns a list of paged clusters.
func (c *ClustersClient) List(ctx context.Context, params ClusterListParams) (paging.PagedResponse[ClusterListItem], error) {
	resp := internal.RequestWrapper[paging.PagedResponse[ClusterListItem]]{}
	err := c.transport.Get(ctx, "/api/kubernetes-dev/v1/cluster.json", &resp, params)
	return resp.Data, mapTransportError(err)
}
