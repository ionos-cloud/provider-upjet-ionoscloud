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

	// [String] The ID of the datacenter where the WireGuard Gateway is located.
	// +crossplane:generate:reference:type=github.com/ionos-cloud/provider-upjet-ionoscloud/apis/compute/v1alpha1.Datacenter
	DatacenterID *string `json:"datacenterId,omitempty" tf:"datacenter_id,omitempty"`

	// Reference to a Datacenter in compute to populate datacenterId.
	// +kubebuilder:validation:Optional
	DatacenterIDRef *v1.Reference `json:"datacenterIdRef,omitempty" tf:"-"`

	// Selector for a Datacenter in compute to populate datacenterId.
	// +kubebuilder:validation:Optional
	DatacenterIDSelector *v1.Selector `json:"datacenterIdSelector,omitempty" tf:"-"`

	// [String] The IPv4 CIDR for the WireGuard Gateway connection.
	IPv4Cidr *string `json:"ipv4Cidr,omitempty" tf:"ipv4_cidr,omitempty"`

	// [String] The IPv6 CIDR for the WireGuard Gateway connection.
	IPv6Cidr *string `json:"ipv6Cidr,omitempty" tf:"ipv6_cidr,omitempty"`

	// [String] The ID of the LAN where the WireGuard Gateway is connected.
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

	// [String] The ID of the datacenter where the WireGuard Gateway is located.
	DatacenterID *string `json:"datacenterId,omitempty" tf:"datacenter_id,omitempty"`

	// [String] The IPv4 CIDR for the WireGuard Gateway connection.
	IPv4Cidr *string `json:"ipv4Cidr,omitempty" tf:"ipv4_cidr,omitempty"`

	// [String] The IPv6 CIDR for the WireGuard Gateway connection.
	IPv6Cidr *string `json:"ipv6Cidr,omitempty" tf:"ipv6_cidr,omitempty"`

	// [String] The ID of the LAN where the WireGuard Gateway is connected.
	LanID *string `json:"lanId,omitempty" tf:"lan_id,omitempty"`
}

type ConnectionsParameters struct {

	// [String] The ID of the datacenter where the WireGuard Gateway is located.
	// +crossplane:generate:reference:type=github.com/ionos-cloud/provider-upjet-ionoscloud/apis/compute/v1alpha1.Datacenter
	// +kubebuilder:validation:Optional
	DatacenterID *string `json:"datacenterId,omitempty" tf:"datacenter_id,omitempty"`

	// Reference to a Datacenter in compute to populate datacenterId.
	// +kubebuilder:validation:Optional
	DatacenterIDRef *v1.Reference `json:"datacenterIdRef,omitempty" tf:"-"`

	// Selector for a Datacenter in compute to populate datacenterId.
	// +kubebuilder:validation:Optional
	DatacenterIDSelector *v1.Selector `json:"datacenterIdSelector,omitempty" tf:"-"`

	// [String] The IPv4 CIDR for the WireGuard Gateway connection.
	// +kubebuilder:validation:Optional
	IPv4Cidr *string `json:"ipv4Cidr,omitempty" tf:"ipv4_cidr,omitempty"`

	// [String] The IPv6 CIDR for the WireGuard Gateway connection.
	// +kubebuilder:validation:Optional
	IPv6Cidr *string `json:"ipv6Cidr,omitempty" tf:"ipv6_cidr,omitempty"`

	// [String] The ID of the LAN where the WireGuard Gateway is connected.
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

type MaintenanceWindowInitParameters struct {

	// [string] The name of the week day.
	// The name of the week day
	DayOfTheWeek *string `json:"dayOfTheWeek,omitempty" tf:"day_of_the_week,omitempty"`

	// [string] Start of the maintenance window in UTC time.
	// Start of the maintenance window in UTC time.
	Time *string `json:"time,omitempty" tf:"time,omitempty"`
}

type MaintenanceWindowObservation struct {

	// [string] The name of the week day.
	// The name of the week day
	DayOfTheWeek *string `json:"dayOfTheWeek,omitempty" tf:"day_of_the_week,omitempty"`

	// [string] Start of the maintenance window in UTC time.
	// Start of the maintenance window in UTC time.
	Time *string `json:"time,omitempty" tf:"time,omitempty"`
}

type MaintenanceWindowParameters struct {

	// [string] The name of the week day.
	// The name of the week day
	// +kubebuilder:validation:Optional
	DayOfTheWeek *string `json:"dayOfTheWeek" tf:"day_of_the_week,omitempty"`

	// [string] Start of the maintenance window in UTC time.
	// Start of the maintenance window in UTC time.
	// +kubebuilder:validation:Optional
	Time *string `json:"time" tf:"time,omitempty"`
}

type VpnWireguardGatewayInitParameters struct {

	// [Block] The connection configuration for the WireGuard Gateway. This block supports fields documented below.
	Connections []ConnectionsInitParameters `json:"connections,omitempty" tf:"connections,omitempty"`

	// [String] A description of the WireGuard Gateway.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// [String] The IP address of the WireGuard Gateway.
	// +crossplane:generate:reference:type=github.com/ionos-cloud/provider-upjet-ionoscloud/apis/compute/v1alpha1.Ipblock
	// +crossplane:generate:reference:extractor=github.com/ionos-cloud/provider-upjet-ionoscloud/config/common.FirstIPBlockIP()
	GatewayIP *string `json:"gatewayIp,omitempty" tf:"gateway_ip,omitempty"`

	// Reference to a Ipblock in compute to populate gatewayIp.
	// +kubebuilder:validation:Optional
	GatewayIPRef *v1.Reference `json:"gatewayIpRef,omitempty" tf:"-"`

	// Selector for a Ipblock in compute to populate gatewayIp.
	// +kubebuilder:validation:Optional
	GatewayIPSelector *v1.Selector `json:"gatewayIpSelector,omitempty" tf:"-"`

	// [String] The IPv4 CIDR for the WireGuard Gateway interface.
	// The IPV4 address (with CIDR mask) to be assigned to the WireGuard interface.
	// __Note__: either interfaceIPv4CIDR or interfaceIPv6CIDR is __required__.
	InterfaceIPv4Cidr *string `json:"interfaceIpv4Cidr,omitempty" tf:"interface_ipv4_cidr,omitempty"`

	// [String] The IPv6 CIDR for the WireGuard Gateway interface.
	// The IPV6 address (with CIDR mask) to be assigned to the WireGuard interface.
	// __Note__: either interfaceIPv6CIDR or interfaceIPv4CIDR is __required__.
	InterfaceIPv6Cidr *string `json:"interfaceIpv6Cidr,omitempty" tf:"interface_ipv6_cidr,omitempty"`

	ListenPort *float64 `json:"listenPort,omitempty" tf:"listen_port,omitempty"`

	// [String] The location of the WireGuard Gateway. Supported locations: de/fra, de/txl, es/vit,
	// gb/bhx, gb/lhr, us/ewr, us/las, us/mci, fr/par.
	// The location of the WireGuard Gateway. Supported locations: de/fra, de/txl, es/vit, gb/bhx, gb/lhr, us/ewr, us/las, us/mci, fr/par
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// (Computed) A weekly 4 hour-long window, during which maintenance might occur.
	// A weekly 4 hour-long window, during which maintenance might occur
	MaintenanceWindow *MaintenanceWindowInitParameters `json:"maintenanceWindow,omitempty" tf:"maintenance_window,omitempty"`

	// [String] The name of the WireGuard Gateway.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// [String] The private key for the WireGuard Gateway. To be created with the wg utility.
	// PrivateKey used for WireGuard Server
	PrivateKeySecretRef v1.SecretKeySelector `json:"privateKeySecretRef" tf:"-"`

	// (Computed)[string] Gateway performance options.  See product documentation for full details. Options: STANDARD, STANDARD_HA, ENHANCED, ENHANCED_HA, PREMIUM, PREMIUM_HA.
	// Gateway performance options. See the documentation for the available options
	Tier *string `json:"tier,omitempty" tf:"tier,omitempty"`
}

type VpnWireguardGatewayObservation struct {

	// [Block] The connection configuration for the WireGuard Gateway. This block supports fields documented below.
	Connections []ConnectionsObservation `json:"connections,omitempty" tf:"connections,omitempty"`

	// [String] A description of the WireGuard Gateway.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// [String] The IP address of the WireGuard Gateway.
	GatewayIP *string `json:"gatewayIp,omitempty" tf:"gateway_ip,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// [String] The IPv4 CIDR for the WireGuard Gateway interface.
	// The IPV4 address (with CIDR mask) to be assigned to the WireGuard interface.
	// __Note__: either interfaceIPv4CIDR or interfaceIPv6CIDR is __required__.
	InterfaceIPv4Cidr *string `json:"interfaceIpv4Cidr,omitempty" tf:"interface_ipv4_cidr,omitempty"`

	// [String] The IPv6 CIDR for the WireGuard Gateway interface.
	// The IPV6 address (with CIDR mask) to be assigned to the WireGuard interface.
	// __Note__: either interfaceIPv6CIDR or interfaceIPv4CIDR is __required__.
	InterfaceIPv6Cidr *string `json:"interfaceIpv6Cidr,omitempty" tf:"interface_ipv6_cidr,omitempty"`

	ListenPort *float64 `json:"listenPort,omitempty" tf:"listen_port,omitempty"`

	// [String] The location of the WireGuard Gateway. Supported locations: de/fra, de/txl, es/vit,
	// gb/bhx, gb/lhr, us/ewr, us/las, us/mci, fr/par.
	// The location of the WireGuard Gateway. Supported locations: de/fra, de/txl, es/vit, gb/bhx, gb/lhr, us/ewr, us/las, us/mci, fr/par
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// (Computed) A weekly 4 hour-long window, during which maintenance might occur.
	// A weekly 4 hour-long window, during which maintenance might occur
	MaintenanceWindow *MaintenanceWindowObservation `json:"maintenanceWindow,omitempty" tf:"maintenance_window,omitempty"`

	// [String] The name of the WireGuard Gateway.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Computed)[String] The public key for the WireGuard Gateway.
	// PublicKey used for WireGuard Server. Received in response from API
	PublicKey *string `json:"publicKey,omitempty" tf:"public_key,omitempty"`

	// (Computed)[String] The current status of the WireGuard Gateway.
	// The status of the WireGuard Gateway
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// (Computed)[string] Gateway performance options.  See product documentation for full details. Options: STANDARD, STANDARD_HA, ENHANCED, ENHANCED_HA, PREMIUM, PREMIUM_HA.
	// Gateway performance options. See the documentation for the available options
	Tier *string `json:"tier,omitempty" tf:"tier,omitempty"`
}

type VpnWireguardGatewayParameters struct {

	// [Block] The connection configuration for the WireGuard Gateway. This block supports fields documented below.
	// +kubebuilder:validation:Optional
	Connections []ConnectionsParameters `json:"connections,omitempty" tf:"connections,omitempty"`

	// [String] A description of the WireGuard Gateway.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// [String] The IP address of the WireGuard Gateway.
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

	// [String] The IPv4 CIDR for the WireGuard Gateway interface.
	// The IPV4 address (with CIDR mask) to be assigned to the WireGuard interface.
	// __Note__: either interfaceIPv4CIDR or interfaceIPv6CIDR is __required__.
	// +kubebuilder:validation:Optional
	InterfaceIPv4Cidr *string `json:"interfaceIpv4Cidr,omitempty" tf:"interface_ipv4_cidr,omitempty"`

	// [String] The IPv6 CIDR for the WireGuard Gateway interface.
	// The IPV6 address (with CIDR mask) to be assigned to the WireGuard interface.
	// __Note__: either interfaceIPv6CIDR or interfaceIPv4CIDR is __required__.
	// +kubebuilder:validation:Optional
	InterfaceIPv6Cidr *string `json:"interfaceIpv6Cidr,omitempty" tf:"interface_ipv6_cidr,omitempty"`

	// +kubebuilder:validation:Optional
	ListenPort *float64 `json:"listenPort,omitempty" tf:"listen_port,omitempty"`

	// [String] The location of the WireGuard Gateway. Supported locations: de/fra, de/txl, es/vit,
	// gb/bhx, gb/lhr, us/ewr, us/las, us/mci, fr/par.
	// The location of the WireGuard Gateway. Supported locations: de/fra, de/txl, es/vit, gb/bhx, gb/lhr, us/ewr, us/las, us/mci, fr/par
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// (Computed) A weekly 4 hour-long window, during which maintenance might occur.
	// A weekly 4 hour-long window, during which maintenance might occur
	// +kubebuilder:validation:Optional
	MaintenanceWindow *MaintenanceWindowParameters `json:"maintenanceWindow,omitempty" tf:"maintenance_window,omitempty"`

	// [String] The name of the WireGuard Gateway.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// [String] The private key for the WireGuard Gateway. To be created with the wg utility.
	// PrivateKey used for WireGuard Server
	// +kubebuilder:validation:Optional
	PrivateKeySecretRef v1.SecretKeySelector `json:"privateKeySecretRef" tf:"-"`

	// (Computed)[string] Gateway performance options.  See product documentation for full details. Options: STANDARD, STANDARD_HA, ENHANCED, ENHANCED_HA, PREMIUM, PREMIUM_HA.
	// Gateway performance options. See the documentation for the available options
	// +kubebuilder:validation:Optional
	Tier *string `json:"tier,omitempty" tf:"tier,omitempty"`
}

// VpnWireguardGatewaySpec defines the desired state of VpnWireguardGateway
type VpnWireguardGatewaySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VpnWireguardGatewayParameters `json:"forProvider"`
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
	InitProvider VpnWireguardGatewayInitParameters `json:"initProvider,omitempty"`
}

// VpnWireguardGatewayStatus defines the observed state of VpnWireguardGateway.
type VpnWireguardGatewayStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VpnWireguardGatewayObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// VpnWireguardGateway is the Schema for the VpnWireguardGateways API. Creates and manages IonosCloud VPN Wireguard Gateway objects.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,ionos}
type VpnWireguardGateway struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.connections) || (has(self.initProvider) && has(self.initProvider.connections))",message="spec.forProvider.connections is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.privateKeySecretRef)",message="spec.forProvider.privateKeySecretRef is a required parameter"
	Spec   VpnWireguardGatewaySpec   `json:"spec"`
	Status VpnWireguardGatewayStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VpnWireguardGatewayList contains a list of VpnWireguardGateways
type VpnWireguardGatewayList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VpnWireguardGateway `json:"items"`
}

// Repository type metadata.
var (
	VpnWireguardGateway_Kind             = "VpnWireguardGateway"
	VpnWireguardGateway_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VpnWireguardGateway_Kind}.String()
	VpnWireguardGateway_KindAPIVersion   = VpnWireguardGateway_Kind + "." + CRDGroupVersion.String()
	VpnWireguardGateway_GroupVersionKind = CRDGroupVersion.WithKind(VpnWireguardGateway_Kind)
)

func init() {
	SchemeBuilder.Register(&VpnWireguardGateway{}, &VpnWireguardGatewayList{})
}
