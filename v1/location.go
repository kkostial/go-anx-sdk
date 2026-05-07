package v1

import (
	"context"

	"code.anexia.com/se/ks/go-anx-sdk/internal"
	"code.anexia.com/se/ks/go-anx-sdk/paging"
)

type LocationListParams struct {
	Page   int    `url:"page,omitempty"`
	Limit  int    `url:"limit,omitempty"`
	Search string `url:"search,omitempty"`
}

type LocationListItem struct {
	Identifier string  `json:"identifier"`
	Code       string  `json:"code"`
	Name       string  `json:"name"`
	CityCode   *string `json:"city_code"`
	Country    *string `json:"country"`
	Lat        *string `json:"lat"`
	Lon        *string `json:"lon"`
}

type LocationsClient struct {
	transport *internal.Transport
}

func NewLocationsClient(transport *internal.Transport) *LocationsClient {
	return &LocationsClient{
		transport,
	}
}

func (v *LocationsClient) List(ctx context.Context, params LocationListParams) (paging.PagedResponse[LocationListItem], error) {
	resp := internal.RequestWrapper[paging.PagedResponse[LocationListItem]]{}
	err := v.transport.Get(ctx, "/api/core/v1/location.json", &resp, params)
	return resp.Data, err
}
