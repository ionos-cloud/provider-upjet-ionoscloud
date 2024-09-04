// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type ConnectionsInitParameters struct {

	// [string] The datacenter to connect your VPN Gateway to.
	// The datacenter to connect your VPN Gateway to.
	// +crossplane:generate:reference:type=github.com/ionos-cloud/provider-upjet-ionoscloud/apis/compute/v1alpha1.Datacenter
	DatacenterID *string `json:"datacenterId,omitempty" tf:"datacenter_id,omitempty"`

	// Reference to a Datacenter in compute to populate datacenterId.
	// +kubebuilder:validation:Optional
	DatacenterIDRef *v1.Reference `json:"datacenterIdRef,omitempty" tf:"-"`

	// Selector for a Datacenter in compute to populate datacenterId.
	// +kubebuilder:validation:Optional
	DatacenterIDSelector *v1.Selector `json:"datacenterIdSelector,omitempty" tf:"-"`

	// [string] Describes the private ipv4 subnet in your LAN that should be accessible by the
	// VPN Gateway. Note: this should be the subnet already assigned to the LAN
	// Describes the private ipv4 subnet in your LAN that should be accessible by the VPN Gateway. Note: this should be the subnet already assigned to the LAN
	IPv4Cidr *string `json:"ipv4Cidr,omitempty" tf:"ipv4_cidr,omitempty"`

	// [string] Describes the ipv6 subnet in your LAN that should be accessible by the VPN
	// Gateway. Note: this should be the subnet already assigned to the LAN
	// Describes the ipv6 subnet in your LAN that should be accessible by the VPN Gateway. Note: this should be the subnet already assigned to the LAN
	IPv6Cidr *string `json:"ipv6Cidr,omitempty" tf:"ipv6_cidr,omitempty"`

	// [string] The numeric LAN ID to connect your VPN Gateway to.
	// The numeric LAN ID to connect your VPN Gateway to.
	// +crossplane:generate:reference:type=github.com/ionos-cloud/provider-upjet-ionoscloud/apis/compute/v1alpha1.Lan
	LanID *string `json:"lanId,omitempty" tf:"lan_id,omitempty"`

	// Reference to a Lan in compute to populate lanId.
	// +kubebuilder:validation:Optional
	LanIDRef *v1.Reference `json:"lanIdRef,omitempty" tf:"-"`

	// Selector for a Lan in compute to populate lanId.
	// +kubebuilder:validation:Optional
	LanIDSelector *v1.Selector `json:"lanIdSelector,omitempty" tf:"-"`
}

type ConnectionsObservation struct {

	// [string] The datacenter to connect your VPN Gateway to.
	// The datacenter to connect your VPN Gateway to.
	DatacenterID *string `json:"datacenterId,omitempty" tf:"datacenter_id,omitempty"`

	// [string] Describes the private ipv4 subnet in your LAN that should be accessible by the
	// VPN Gateway. Note: this should be the subnet already assigned to the LAN
	// Describes the private ipv4 subnet in your LAN that should be accessible by the VPN Gateway. Note: this should be the subnet already assigned to the LAN
	IPv4Cidr *string `json:"ipv4Cidr,omitempty" tf:"ipv4_cidr,omitempty"`

	// [string] Describes the ipv6 subnet in your LAN that should be accessible by the VPN
	// Gateway. Note: this should be the subnet already assigned to the LAN
	// Describes the ipv6 subnet in your LAN that should be accessible by the VPN Gateway. Note: this should be the subnet already assigned to the LAN
	IPv6Cidr *string `json:"ipv6Cidr,omitempty" tf:"ipv6_cidr,omitempty"`

	// [string] The numeric LAN ID to connect your VPN Gateway to.
	// The numeric LAN ID to connect your VPN Gateway to.
	LanID *string `json:"lanId,omitempty" tf:"lan_id,omitempty"`
}

type ConnectionsParameters struct {

	// [string] The datacenter to connect your VPN Gateway to.
	// The datacenter to connect your VPN Gateway to.
	// +crossplane:generate:reference:type=github.com/ionos-cloud/provider-upjet-ionoscloud/apis/compute/v1alpha1.Datacenter
	// +kubebuilder:validation:Optional
	DatacenterID *string `json:"datacenterId,omitempty" tf:"datacenter_id,omitempty"`

	// Reference to a Datacenter in compute to populate datacenterId.
	// +kubebuilder:validation:Optional
	DatacenterIDRef *v1.Reference `json:"datacenterIdRef,omitempty" tf:"-"`

	// Selector for a Datacenter in compute to populate datacenterId.
	// +kubebuilder:validation:Optional
	DatacenterIDSelector *v1.Selector `json:"datacenterIdSelector,omitempty" tf:"-"`

	// [string] Describes the private ipv4 subnet in your LAN that should be accessible by the
	// VPN Gateway. Note: this should be the subnet already assigned to the LAN
	// Describes the private ipv4 subnet in your LAN that should be accessible by the VPN Gateway. Note: this should be the subnet already assigned to the LAN
	// +kubebuilder:validation:Optional
	IPv4Cidr *string `json:"ipv4Cidr" tf:"ipv4_cidr,omitempty"`

	// [string] Describes the ipv6 subnet in your LAN that should be accessible by the VPN
	// Gateway. Note: this should be the subnet already assigned to the LAN
	// Describes the ipv6 subnet in your LAN that should be accessible by the VPN Gateway. Note: this should be the subnet already assigned to the LAN
	// +kubebuilder:validation:Optional
	IPv6Cidr *string `json:"ipv6Cidr,omitempty" tf:"ipv6_cidr,omitempty"`

	// [string] The numeric LAN ID to connect your VPN Gateway to.
	// The numeric LAN ID to connect your VPN Gateway to.
	// +crossplane:generate:reference:type=github.com/ionos-cloud/provider-upjet-ionoscloud/apis/compute/v1alpha1.Lan
	// +kubebuilder:validation:Optional
	LanID *string `json:"lanId,omitempty" tf:"lan_id,omitempty"`

	// Reference to a Lan in compute to populate lanId.
	// +kubebuilder:validation:Optional
	LanIDRef *v1.Reference `json:"lanIdRef,omitempty" tf:"-"`

	// Selector for a Lan in compute to populate lanId.
	// +kubebuilder:validation:Optional
	LanIDSelector *v1.Selector `json:"lanIdSelector,omitempty" tf:"-"`
}

type VpnIpsecGatewayInitParameters struct {

	// [list] The network connection for your gateway. Note: all connections must belong to the
	// same datacenter. Minimum items: 1. Maximum items: 10.
	// The network connection for your gateway. Note: all connections must belong to the same datacenter.
	Connections []ConnectionsInitParameters `json:"connections,omitempty" tf:"connections,omitempty"`

	// [string] The human-readable description of the IPSec Gateway.
	// The human-readable description of your IPSec Gateway.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// [string] Public IP address to be assigned to the gateway. Note: This must be an IP address in
	// the same datacenter as the connections.
	// Public IP address to be assigned to the gateway. Note: This must be an IP address in the same datacenter as the connections.
	// +crossplane:generate:reference:type=github.com/ionos-cloud/provider-upjet-ionoscloud/apis/compute/v1alpha1.Ipblock
	// +crossplane:generate:reference:extractor=github.com/ionos-cloud/provider-upjet-ionoscloud/config/common.FirstIPBlockIP()
	GatewayIP *string `json:"gatewayIp,omitempty" tf:"gateway_ip,omitempty"`

	// Reference to a Ipblock in compute to populate gatewayIp.
	// +kubebuilder:validation:Optional
	GatewayIPRef *v1.Reference `json:"gatewayIpRef,omitempty" tf:"-"`

	// Selector for a Ipblock in compute to populate gatewayIp.
	// +kubebuilder:validation:Optional
	GatewayIPSelector *v1.Selector `json:"gatewayIpSelector,omitempty" tf:"-"`

	// [string] The location of the IPSec Gateway. Supported locations: de/fra, de/txl, es/vit,
	// gb/lhr, us/ewr, us/las, us/mci, fr/par
	// The location of the IPSec Gateway. Supported locations: de/fra, de/txl
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// [string] The name of the IPSec Gateway.
	// The human readable name of your IPSecGateway.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// [string] The IKE version that is permitted for the VPN tunnels. Default: IKEv2. Possible
	// values: IKEv2.
	// The IKE version that is permitted for the VPN tunnels.
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type VpnIpsecGatewayObservation struct {

	// [list] The network connection for your gateway. Note: all connections must belong to the
	// same datacenter. Minimum items: 1. Maximum items: 10.
	// The network connection for your gateway. Note: all connections must belong to the same datacenter.
	Connections []ConnectionsObservation `json:"connections,omitempty" tf:"connections,omitempty"`

	// [string] The human-readable description of the IPSec Gateway.
	// The human-readable description of your IPSec Gateway.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// [string] Public IP address to be assigned to the gateway. Note: This must be an IP address in
	// the same datacenter as the connections.
	// Public IP address to be assigned to the gateway. Note: This must be an IP address in the same datacenter as the connections.
	GatewayIP *string `json:"gatewayIp,omitempty" tf:"gateway_ip,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// [string] The location of the IPSec Gateway. Supported locations: de/fra, de/txl, es/vit,
	// gb/lhr, us/ewr, us/las, us/mci, fr/par
	// The location of the IPSec Gateway. Supported locations: de/fra, de/txl
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// [string] The name of the IPSec Gateway.
	// The human readable name of your IPSecGateway.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// [string] The IKE version that is permitted for the VPN tunnels. Default: IKEv2. Possible
	// values: IKEv2.
	// The IKE version that is permitted for the VPN tunnels.
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type VpnIpsecGatewayParameters struct {

	// [list] The network connection for your gateway. Note: all connections must belong to the
	// same datacenter. Minimum items: 1. Maximum items: 10.
	// The network connection for your gateway. Note: all connections must belong to the same datacenter.
	// +kubebuilder:validation:Optional
	Connections []ConnectionsParameters `json:"connections,omitempty" tf:"connections,omitempty"`

	// [string] The human-readable description of the IPSec Gateway.
	// The human-readable description of your IPSec Gateway.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// [string] Public IP address to be assigned to the gateway. Note: This must be an IP address in
	// the same datacenter as the connections.
	// Public IP address to be assigned to the gateway. Note: This must be an IP address in the same datacenter as the connections.
	// +crossplane:generate:reference:type=github.com/ionos-cloud/provider-upjet-ionoscloud/apis/compute/v1alpha1.Ipblock
	// +crossplane:generate:reference:extractor=github.com/ionos-cloud/provider-upjet-ionoscloud/config/common.FirstIPBlockIP()
	// +kubebuilder:validation:Optional
	GatewayIP *string `json:"gatewayIp,omitempty" tf:"gateway_ip,omitempty"`

	// Reference to a Ipblock in compute to populate gatewayIp.
	// +kubebuilder:validation:Optional
	GatewayIPRef *v1.Reference `json:"gatewayIpRef,omitempty" tf:"-"`

	// Selector for a Ipblock in compute to populate gatewayIp.
	// +kubebuilder:validation:Optional
	GatewayIPSelector *v1.Selector `json:"gatewayIpSelector,omitempty" tf:"-"`

	// [string] The location of the IPSec Gateway. Supported locations: de/fra, de/txl, es/vit,
	// gb/lhr, us/ewr, us/las, us/mci, fr/par
	// The location of the IPSec Gateway. Supported locations: de/fra, de/txl
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// [string] The name of the IPSec Gateway.
	// The human readable name of your IPSecGateway.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// [string] The IKE version that is permitted for the VPN tunnels. Default: IKEv2. Possible
	// values: IKEv2.
	// The IKE version that is permitted for the VPN tunnels.
	// +kubebuilder:validation:Optional
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

// VpnIpsecGatewaySpec defines the desired state of VpnIpsecGateway
type VpnIpsecGatewaySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VpnIpsecGatewayParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider VpnIpsecGatewayInitParameters `json:"initProvider,omitempty"`
}

// VpnIpsecGatewayStatus defines the observed state of VpnIpsecGateway.
type VpnIpsecGatewayStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VpnIpsecGatewayObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// VpnIpsecGateway is the Schema for the VpnIpsecGateways API. IPSec Gateway
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,ionos}
type VpnIpsecGateway struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.connections) || (has(self.initProvider) && has(self.initProvider.connections))",message="spec.forProvider.connections is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.location) || (has(self.initProvider) && has(self.initProvider.location))",message="spec.forProvider.location is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   VpnIpsecGatewaySpec   `json:"spec"`
	Status VpnIpsecGatewayStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VpnIpsecGatewayList contains a list of VpnIpsecGateways
type VpnIpsecGatewayList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VpnIpsecGateway `json:"items"`
}

// Repository type metadata.
var (
	VpnIpsecGateway_Kind             = "VpnIpsecGateway"
	VpnIpsecGateway_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VpnIpsecGateway_Kind}.String()
	VpnIpsecGateway_KindAPIVersion   = VpnIpsecGateway_Kind + "." + CRDGroupVersion.String()
	VpnIpsecGateway_GroupVersionKind = CRDGroupVersion.WithKind(VpnIpsecGateway_Kind)
)

func init() {
	SchemeBuilder.Register(&VpnIpsecGateway{}, &VpnIpsecGatewayList{})
}
