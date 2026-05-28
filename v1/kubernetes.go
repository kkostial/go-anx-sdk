package v1

import (
	"context"
	"fmt"

	"github.com/kkostial/go-anx-sdk/internal"
	"github.com/kkostial/go-anx-sdk/paging"
)

type ClusterState string

const (
	ClusterStateOk                     ClusterState = "0"
	ClusterStateError                  ClusterState = "1"
	CLusterStatePending                ClusterState = "2"
	ClusterStateAnxDevNoGa             ClusterState = "3"
	ClusterStateWaitingForNetworks     ClusterState = "4"
	ClusterStateWaitingForServiceVMs   ClusterState = "5"
	ClusterStateWaitingForControlPlane ClusterState = "6"
	ClusterStateUpdatingControlPlane   ClusterState = "7"
	ClusterStateUpdatingNodes          ClusterState = "8"
)

type ClusterVersion string

const (
	ClusterVersion1_32 ClusterVersion = "1.32"
	ClusterVersion1_33 ClusterVersion = "1.33"
	ClusterVersion1_34 ClusterVersion = "1.34"
	ClusterVersion1_35 ClusterVersion = "1.35"
)

type ClusterCniPlugin string

const (
	ClusterCniPluginCanal ClusterCniPlugin = "Canal"
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
type State struct {
	Text  string       `json:"text"`
	Title string       `json:"title"`
	Id    ClusterState `json:"id"`
	Type  int          `json:"type"`
}

// ClusterGetResponse represents the response of the cluster get endpoint.
type ClusterGetResponse struct {
	CustomerIdentifier             string           `json:"customer_identifier"`
	ResellerIdentifier             string           `json:"reseller_identifier"`
	Identifier                     string           `json:"identifier"`
	Name                           string           `json:"name"`
	State                          State            `json:"state"`
	Location                       Resource         `json:"location"`
	Version                        ClusterVersion   `json:"version"`
	PatchVersion                   string           `json:"patch_version"`
	Kubeconfig                     string           `json:"kubeconfig"`
	Autoscaling                    bool             `json:"autoscaling"`
	EnablePersistentStorage        bool             `json:"enable_persistent_storage"`
	CniPlugin                      ClusterCniPlugin `json:"cni_plugin"`
	ApiServerAllowlist             string           `json:"apiserver_allowlist"`
	BackendName                    string           `json:"backend_name"`
	Backend                        string           `json:"backend"`
	MaintenanceWindowStartTime     string           `json:"maintenance_window_start_time"`
	MaintenanceWindowDuration      string           `json:"maintenance_window_duration"`
	ServiceUser                    Resource         `json:"service_user"`
	ManageInternalIpv4Prefix       bool             `json:"manage_internal_ipv4_prefix"`
	InternalIpv4Prefix             Resource         `json:"internal_ipv4_prefix"`
	NeedsServiceVms                bool             `json:"needs_service_vms"`
	EnableNatGateways              bool             `json:"enable_nat_gateways"`
	EnableLbaas                    bool             `json:"enable_lbaas"`
	ExternalIpFamilies             string           `json:"external_ip_families"`
	ManageExternalIpv4Prefix       bool             `json:"manage_external_ipv4_prefix"`
	ExternalIpv4Prefix             Resource         `json:"external_ipv4_prefix"`
	ManageExternalIpv6Prefix       bool             `json:"manage_external_ipv6_prefix"`
	ExternalIpv6Prefix             Resource         `json:"external_ipv6_prefix"`
	ServiceVm01                    Resource         `json:"service_vm_01"`
	ServiceVm02                    Resource         `json:"service_vm_02"`
	ServiceVm01InternalIpv4Address Resource         `json:"service_vm_01_internal_ipv4_address"`
	ServiceVm02InternalIpv4Address Resource         `json:"service_vm_02_internal_ipv4_address"`
	ServiceVm01ExternalIpv4Address Resource         `json:"service_vm_01_external_ipv4_address"`
	ServiceVm02ExternalIpv4Address Resource         `json:"service_vm_02_external_ipv4_address"`
	ServiceVm01ExternalIpv6Address Resource         `json:"service_vm_01_external_ipv6_address"`
	ServiceVm02ExternalIpv6Address Resource         `json:"service_vm_02_external_ipv6_address"`
	ServiceLb01                    Resource         `json:"service_lb_01"`
	ServiceLb02                    Resource         `json:"service_lb_02"`
	ExternalIpv4Vip                Resource         `json:"external_ipv4_vip"`
	ExternalIpv6Vip                Resource         `json:"external_ipv6_vip"`
	KkpApiLbaasBackend01           Resource         `json:"kkp_api_lbaas_backend_01"`
	KkpApiLbaasBackend02           Resource         `json:"kkp_api_lbaas_backend_02"`
	KkpVpnLbaasBackend01           Resource         `json:"kkp_vpn_lbaas_backend_01"`
	KkpVpnLbaasBackend02           Resource         `json:"kkp_vpn_lbaas_backend_02"`
	StorageServerInterfaceAddress  Resource         `json:"storage_server_interface_address"`
	KkpProjectId                   string           `json:"kkp_project_id"`
	KkpClusterId                   string           `json:"kkp_cluster_id"`
	EnableOidcAuthentication       bool             `json:"enable_oidc_authentication"`
	OidcClientId                   string           `json:"oidc_client_id"`
	OidcIssuerUrl                  string           `json:"oidc_issuer_url"`
	OidcGroupsClaim                string           `json:"oidc_groups_claim"`
	OidcUsernameClaim              string           `json:"oidc_username_claim"`
	OidcExtraScopes                string           `json:"oidc_extra_scopes"`
	OidcGroupsPrefix               string           `json:"oidc_groups_prefix"`
	OidcRequiredClaim              string           `json:"oidc_required_claim"`
	OidcUsernamePrefix             string           `json:"oidc_username_prefix"`
	AutomationRules                []Resource       `json:"automation_rules"`
}

type ClusterUpdateRequest struct {
	Name                       *string           `json:"name,omitempty"`
	State                      *ClusterState     `json:"state,omitempty"`
	LocationIdentifier         *string           `json:"location,omitempty"`
	Version                    *ClusterVersion   `json:"version,omitempty"`
	PatchVersion               *string           `json:"patch_version,omitempty"`
	KubeConfig                 *string           `json:"kubeconfig,omitempty"`
	Autoscaling                *bool             `json:"autoscaling,omitempty"`
	EnablePersistentStorage    *bool             `json:"enable_persistent_storage,omitempty"`
	CniPlugin                  *ClusterCniPlugin `json:"cni_plugin,omitempty"`
	ApiServerAllowList         *string           `json:"api_server_allow_list,omitempty"`
	BackendName                *string           `json:"backend_name,omitempty"`
	Backend                    *string           `json:"backend,omitempty"`
	MaintenanceWindowStartTime *string           `json:"maintenance_window_start_time,omitempty"`
	MaintenanceWindowDuration  *string           `json:"maintenance_window_duration,omitempty"`
	ServiceUserIdentifier      *string           `json:"service_user,omitempty"`
	// manage_internal_ipv4_prefix
}

type ClusterUpdateResponse struct {
	CustomerIdentifier string `json:"customer_identifier"`
	ResellerIdentifier string `json:"reseller_identifier"`
	Identifier         string `json:"identifier"`
	Name               string `json:"name"`
	State              struct {
		Text  string `json:"text"`
		Title string `json:"title"`
		Id    string `json:"id"`
		Type  int    `json:"type"`
	} `json:"state"`
	Location struct {
		Identifier string `json:"identifier"`
		Name       string `json:"name"`
	} `json:"location"`
	Version                    string      `json:"version"`
	PatchVersion               interface{} `json:"patch_version"`
	Kubeconfig                 string      `json:"kubeconfig"`
	Autoscaling                bool        `json:"autoscaling"`
	EnablePersistentStorage    bool        `json:"enable_persistent_storage"`
	CniPlugin                  string      `json:"cni_plugin"`
	ApiserverAllowlist         interface{} `json:"apiserver_allowlist"`
	BackendName                string      `json:"backend_name"`
	Backend                    string      `json:"backend"`
	MaintenanceWindowStartTime string      `json:"maintenance_window_start_time"`
	MaintenanceWindowDuration  string      `json:"maintenance_window_duration"`
	ServiceUser                struct {
		Identifier string `json:"identifier"`
		Name       string `json:"name"`
	} `json:"service_user"`
	ManageInternalIpv4Prefix bool `json:"manage_internal_ipv4_prefix"`
	InternalIpv4Prefix       struct {
		Identifier string `json:"identifier"`
		Name       string `json:"name"`
	} `json:"internal_ipv4_prefix"`
	NeedsServiceVms          bool   `json:"needs_service_vms"`
	EnableNatGateways        bool   `json:"enable_nat_gateways"`
	EnableLbaas              bool   `json:"enable_lbaas"`
	ExternalIpFamilies       string `json:"external_ip_families"`
	ManageExternalIpv4Prefix bool   `json:"manage_external_ipv4_prefix"`
	ExternalIpv4Prefix       struct {
		Identifier string `json:"identifier"`
		Name       string `json:"name"`
	} `json:"external_ipv4_prefix"`
	ManageExternalIpv6Prefix bool `json:"manage_external_ipv6_prefix"`
	ExternalIpv6Prefix       struct {
		Identifier string `json:"identifier"`
		Name       string `json:"name"`
	} `json:"external_ipv6_prefix"`
	ServiceVm01 struct {
		Identifier string `json:"identifier"`
		Name       string `json:"name"`
	} `json:"service_vm_01"`
	ServiceVm02 struct {
		Identifier string `json:"identifier"`
		Name       string `json:"name"`
	} `json:"service_vm_02"`
	ServiceVm01InternalIpv4Address struct {
		Identifier string `json:"identifier"`
		Name       string `json:"name"`
	} `json:"service_vm_01_internal_ipv4_address"`
	ServiceVm02InternalIpv4Address struct {
		Identifier string `json:"identifier"`
		Name       string `json:"name"`
	} `json:"service_vm_02_internal_ipv4_address"`
	ServiceVm01ExternalIpv4Address struct {
		Identifier string `json:"identifier"`
		Name       string `json:"name"`
	} `json:"service_vm_01_external_ipv4_address"`
	ServiceVm02ExternalIpv4Address struct {
		Identifier string `json:"identifier"`
		Name       string `json:"name"`
	} `json:"service_vm_02_external_ipv4_address"`
	ServiceVm01ExternalIpv6Address struct {
		Identifier string `json:"identifier"`
		Name       string `json:"name"`
	} `json:"service_vm_01_external_ipv6_address"`
	ServiceVm02ExternalIpv6Address struct {
		Identifier string `json:"identifier"`
		Name       string `json:"name"`
	} `json:"service_vm_02_external_ipv6_address"`
	ServiceLb01 struct {
		Identifier string `json:"identifier"`
		Name       string `json:"name"`
	} `json:"service_lb_01"`
	ServiceLb02 struct {
		Identifier string `json:"identifier"`
		Name       string `json:"name"`
	} `json:"service_lb_02"`
	ExternalIpv4Vip struct {
		Identifier string `json:"identifier"`
		Name       string `json:"name"`
	} `json:"external_ipv4_vip"`
	ExternalIpv6Vip struct {
		Identifier string `json:"identifier"`
		Name       string `json:"name"`
	} `json:"external_ipv6_vip"`
	KkpApiLbaasBackend01 struct {
		Identifier string `json:"identifier"`
		Name       string `json:"name"`
	} `json:"kkp_api_lbaas_backend_01"`
	KkpApiLbaasBackend02 struct {
		Identifier string `json:"identifier"`
		Name       string `json:"name"`
	} `json:"kkp_api_lbaas_backend_02"`
	KkpVpnLbaasBackend01 struct {
		Identifier string `json:"identifier"`
		Name       string `json:"name"`
	} `json:"kkp_vpn_lbaas_backend_01"`
	KkpVpnLbaasBackend02 struct {
		Identifier string `json:"identifier"`
		Name       string `json:"name"`
	} `json:"kkp_vpn_lbaas_backend_02"`
	StorageServerInterfaceAddress interface{} `json:"storage_server_interface_address"`
	KkpProjectId                  string      `json:"kkp_project_id"`
	KkpClusterId                  string      `json:"kkp_cluster_id"`
	InternalVlan                  interface{} `json:"internal_vlan"`
	ExternalVlan                  interface{} `json:"external_vlan"`
	EnableOidcAuthentication      bool        `json:"enable_oidc_authentication"`
	OidcClientId                  interface{} `json:"oidc_client_id"`
	OidcIssuerUrl                 interface{} `json:"oidc_issuer_url"`
	OidcGroupsClaim               string      `json:"oidc_groups_claim"`
	OidcUsernameClaim             string      `json:"oidc_username_claim"`
	OidcExtraScopes               interface{} `json:"oidc_extra_scopes"`
	OidcGroupsPrefix              interface{} `json:"oidc_groups_prefix"`
	OidcRequiredClaim             interface{} `json:"oidc_required_claim"`
	OidcUsernamePrefix            interface{} `json:"oidc_username_prefix"`
	AutomationRules               []struct {
		Identifier string `json:"identifier"`
		Name       string `json:"name"`
	} `json:"automation_rules"`
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

func (c *ClustersClient) Update(ctx context.Context, identifier string, request ClusterUpdateRequest) (ClusterUpdateResponse, error) {
	resp := ClusterUpdateResponse{}
	err := c.transport.Put(ctx, fmt.Sprintf("/api/kubernetes-dev/v1/cluster.json/%s", identifier), request, &resp)
	return resp, mapTransportError(err)
}

// List returns a list of paged clusters.
func (c *ClustersClient) List(ctx context.Context, params ClusterListParams) (paging.PagedResponse[ClusterListItem], error) {
	resp := internal.RequestWrapper[paging.PagedResponse[ClusterListItem]]{}
	err := c.transport.Get(ctx, "/api/kubernetes-dev/v1/cluster.json", &resp, params)
	return resp.Data, mapTransportError(err)
}
