package v1

import (
	"context"
	"fmt"

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

// ClusterState represents a clusters current state.
// Maybe extract out as soon as others need it too.
type ClusterState struct {
	Text  string `json:"text"`
	Title string `json:"title"`
	Id    string `json:"id"`
	Type  int    `json:"type"`
}

// ClusterGetResponse represents the response of the cluster get endpoint.
type ClusterGetResponse struct {
	CustomerIdentifier             string       `json:"customer_identifier"`
	ResellerIdentifier             string       `json:"reseller_identifier"`
	Identifier                     string       `json:"identifier"`
	Name                           string       `json:"name"`
	State                          ClusterState `json:"state"`
	Location                       Resource     `json:"location"`
	Version                        string       `json:"version"`
	PatchVersion                   string       `json:"patch_version"`
	Kubeconfig                     string       `json:"kubeconfig"`
	Autoscaling                    bool         `json:"autoscaling"`
	EnablePersistentStorage        bool         `json:"enable_persistent_storage"`
	CniPlugin                      string       `json:"cni_plugin"`
	ApiServerAllowlist             string       `json:"apiserver_allowlist"`
	BackendName                    string       `json:"backend_name"`
	Backend                        string       `json:"backend"`
	MaintenanceWindowStartTime     string       `json:"maintenance_window_start_time"`
	MaintenanceWindowDuration      string       `json:"maintenance_window_duration"`
	ServiceUser                    Resource     `json:"service_user"`
	ManageInternalIpv4Prefix       bool         `json:"manage_internal_ipv4_prefix"`
	InternalIpv4Prefix             Resource     `json:"internal_ipv4_prefix"`
	NeedsServiceVms                bool         `json:"needs_service_vms"`
	EnableNatGateways              bool         `json:"enable_nat_gateways"`
	EnableLbaas                    bool         `json:"enable_lbaas"`
	ExternalIpFamilies             string       `json:"external_ip_families"`
	ManageExternalIpv4Prefix       bool         `json:"manage_external_ipv4_prefix"`
	ExternalIpv4Prefix             Resource     `json:"external_ipv4_prefix"`
	ManageExternalIpv6Prefix       bool         `json:"manage_external_ipv6_prefix"`
	ExternalIpv6Prefix             Resource     `json:"external_ipv6_prefix"`
	ServiceVm01                    Resource     `json:"service_vm_01"`
	ServiceVm02                    Resource     `json:"service_vm_02"`
	ServiceVm01InternalIpv4Address Resource     `json:"service_vm_01_internal_ipv4_address"`
	ServiceVm02InternalIpv4Address Resource     `json:"service_vm_02_internal_ipv4_address"`
	ServiceVm01ExternalIpv4Address Resource     `json:"service_vm_01_external_ipv4_address"`
	ServiceVm02ExternalIpv4Address Resource     `json:"service_vm_02_external_ipv4_address"`
	ServiceVm01ExternalIpv6Address Resource     `json:"service_vm_01_external_ipv6_address"`
	ServiceVm02ExternalIpv6Address Resource     `json:"service_vm_02_external_ipv6_address"`
	ServiceLb01                    Resource     `json:"service_lb_01"`
	ServiceLb02                    Resource     `json:"service_lb_02"`
	ExternalIpv4Vip                Resource     `json:"external_ipv4_vip"`
	ExternalIpv6Vip                Resource     `json:"external_ipv6_vip"`
	KkpApiLbaasBackend01           Resource     `json:"kkp_api_lbaas_backend_01"`
	KkpApiLbaasBackend02           Resource     `json:"kkp_api_lbaas_backend_02"`
	KkpVpnLbaasBackend01           Resource     `json:"kkp_vpn_lbaas_backend_01"`
	KkpVpnLbaasBackend02           Resource     `json:"kkp_vpn_lbaas_backend_02"`
	StorageServerInterfaceAddress  Resource     `json:"storage_server_interface_address"`
	KkpProjectId                   string       `json:"kkp_project_id"`
	KkpClusterId                   string       `json:"kkp_cluster_id"`
	EnableOidcAuthentication       bool         `json:"enable_oidc_authentication"`
	OidcClientId                   string       `json:"oidc_client_id"`
	OidcIssuerUrl                  string       `json:"oidc_issuer_url"`
	OidcGroupsClaim                string       `json:"oidc_groups_claim"`
	OidcUsernameClaim              string       `json:"oidc_username_claim"`
	OidcExtraScopes                string       `json:"oidc_extra_scopes"`
	OidcGroupsPrefix               string       `json:"oidc_groups_prefix"`
	OidcRequiredClaim              string       `json:"oidc_required_claim"`
	OidcUsernamePrefix             string       `json:"oidc_username_prefix"`
	AutomationRules                []Resource   `json:"automation_rules"`
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

// Get returns a single cluster by its id.
func (c *ClustersClient) Get(ctx context.Context, identifier string) (ClusterGetResponse, error) {
	resp := ClusterGetResponse{}
	err := c.transport.Get(ctx, fmt.Sprintf("/api/kubernetes-dev/v1/cluster.json/%s", identifier), &resp, nil)
	return resp, mapTransportError(err)
}

// List returns a list of paged clusters.
func (c *ClustersClient) List(ctx context.Context, params ClusterListParams) (paging.PagedResponse[ClusterListItem], error) {
	resp := internal.RequestWrapper[paging.PagedResponse[ClusterListItem]]{}
	err := c.transport.Get(ctx, "/api/kubernetes-dev/v1/cluster.json", &resp, params)
	return resp.Data, mapTransportError(err)
}
