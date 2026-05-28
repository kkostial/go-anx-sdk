package v1

import (
	"context"

	"github.com/kkostial/go-anx-sdk/internal"
	"github.com/kkostial/go-anx-sdk/paging"
)

// LocationListParams defines the available parameters for the location list endpoint.
type LocationListParams struct {
	Page   int    `url:"page,omitempty"`
	Limit  int    `url:"limit,omitempty"`
	Search string `url:"search,omitempty"`
}

// LocationListItem is an item in the location list response.
type LocationListItem struct {
	Identifier string  `json:"identifier"`
	Code       string  `json:"code"`
	Name       string  `json:"name"`
	CityCode   *string `json:"city_code"`
	Country    *string `json:"country"`
	Lat        *string `json:"lat"`
	Lon        *string `json:"lon"`
}

// LocationsClient is an api client for managing locations.
type LocationsClient struct {
	transport *internal.Transport
}

// NewLocationsClient creates a new location client.
func NewLocationsClient(transport *internal.Transport) *LocationsClient {
	return &LocationsClient{
		transport,
	}
}

// List returns a list of paged locations.
func (v *LocationsClient) List(ctx context.Context, params LocationListParams) (paging.PagedResponse[LocationListItem], error) {
	resp := internal.RequestWrapper[paging.PagedResponse[LocationListItem]]{}
	err := v.transport.Get(ctx, "/api/core/v1/location.json", &resp, params)
	return resp.Data, mapTransportError(err)
}
