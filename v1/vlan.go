package v1

import (
	"context"
	"fmt"

	"code.anexia.com/se/ks/go-anx-sdk/internal"
	"code.anexia.com/se/ks/go-anx-sdk/paging"
)

type VlanCreateRequest struct {
	LocationIdentifier  string `json:"location"`
	VmProvisioning      bool   `json:"vm_provisioning"`
	DescriptionCustomer string `json:"description_customer"`
}

type VlanCreateResponse struct {
	Identifier          string `json:"identifier"`
	Name                string `json:"name"`
	DescriptionCustomer string `json:"description_customer"`
}

type VlanGetResponse struct {
	Identifier          string                        `json:"identifier"`
	Name                string                        `json:"name"`
	DescriptionCustomer string                        `json:"description_customer"`
	DescriptionInternal string                        `json:"description_internal"`
	RoleText            string                        `json:"role_text"`
	Status              string                        `json:"status"`
	Locations           []VlanGetResponseLocationItem `json:"locations"`
	VmProvisioning      bool                          `json:"vm_provisioning"`
}

// TODO: maybe refactor this for other endpoitns to be reused
type VlanGetResponseLocationItem struct {
	Identifier string `json:"identifier"`
	Code       string `json:"code"`
	Name       string `json:"name"`
	Country    string `json:"country"`
	Lat        string `json:"lat"`
	Lon        string `json:"lon"`
	CityCode   string `json:"city_code"`
}

type VlanUpdateRequest struct {
	VmProvisioning      *bool   `json:"vm_provisioning,omitempty"`
	DescriptionCustomer *string `json:"description_customer,omitempty"`
}

type VlanUpdateResponse struct {
	Identifier          string `json:"identifier"`
	Name                string `json:"name"`
	DescriptionCustomer string `json:"description_customer"`
}

type VlanListParams struct {
	Page   int    `url:"page,omitempty"`
	Limit  int    `url:"limit,omitempty"`
	Search string `url:"search,omitempty"`
}

type VlanFilteredParams struct {
	Page                   int    `url:"page,omitempty"`
	Limit                  int    `url:"limit,omitempty"`
	Search                 string `url:"search,omitempty"`
	OrganizationIdentifier string `url:"organization_identifier,omitempty"`
	RoleText               string `url:"role_text,omitempty"`
	Status                 string `url:"status,omitempty"`
	Location               string `url:"location,omitempty"`
}

type VlanListItem struct {
	Identifier          string `json:"identifier"`
	Name                string `json:"name"`
	DescriptionCustomer string `json:"description_customer"`
}

type VlansClient struct {
	transport *internal.Transport
}

func NewVlansClient(transport *internal.Transport) *VlansClient {
	return &VlansClient{
		transport: transport,
	}
}

func (v *VlansClient) Create(ctx context.Context, request VlanCreateRequest) (VlanCreateResponse, error) {
	resp := VlanCreateResponse{}
	err := v.transport.Post(ctx, "/api/vlan/v1/vlan.json", request, &resp)
	return resp, err
}

func (v *VlansClient) Get(ctx context.Context, identifier string) (VlanGetResponse, error) {
	resp := VlanGetResponse{}
	err := v.transport.Get(ctx, fmt.Sprintf("api/vlan/v1/vlan.json/%s", identifier), &resp, nil)
	return resp, err
}

func (v *VlansClient) Update(ctx context.Context, identifier string, request VlanUpdateRequest) (VlanUpdateResponse, error) {
	resp := VlanUpdateResponse{}
	err := v.transport.Put(ctx, fmt.Sprintf("/api/vlan/v1/vlan.json/%s", identifier), request, &resp)
	return resp, err
}

func (v *VlansClient) List(ctx context.Context, params VlanListParams) (paging.PagedResponse[VlanListItem], error) {
	resp := internal.RequestWrapper[paging.PagedResponse[VlanListItem]]{}
	err := v.transport.Get(ctx, "/api/vlan/v1/vlan.json", &resp, params)
	return resp.Data, err
}

func (v *VlansClient) ListFiltered(ctx context.Context, params VlanFilteredParams) (paging.PagedResponse[VlanListItem], error) {
	resp := internal.RequestWrapper[paging.PagedResponse[VlanListItem]]{}
	err := v.transport.Get(ctx, "/api/vlan/v1/vlan/filtered.json", &resp, params)
	return resp.Data, err
}

func (v *VlansClient) Delete(ctx context.Context, identifier string) error {
	return v.transport.Delete(ctx, fmt.Sprintf("/api/vlan/v1/vlan.json/%s", identifier))
}
