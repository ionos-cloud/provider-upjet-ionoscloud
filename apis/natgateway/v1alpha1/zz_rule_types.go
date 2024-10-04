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

type RuleInitParameters struct {

	// [string] A Datacenter's UUID.
	// +crossplane:generate:reference:type=github.com/ionos-cloud/provider-upjet-ionoscloud/apis/compute/v1alpha1.Datacenter
	DatacenterID *string `json:"datacenterId,omitempty" tf:"datacenter_id,omitempty"`

	// Reference to a Datacenter in compute to populate datacenterId.
	// +kubebuilder:validation:Optional
	DatacenterIDRef *v1.Reference `json:"datacenterIdRef,omitempty" tf:"-"`

	// Selector for a Datacenter in compute to populate datacenterId.
	// +kubebuilder:validation:Optional
	DatacenterIDSelector *v1.Selector `json:"datacenterIdSelector,omitempty" tf:"-"`

	// [string] Name of the NAT gateway rule.
	// Name of the NAT gateway rule
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// [string] Nat Gateway's UUID.
	// +crossplane:generate:reference:type=github.com/ionos-cloud/provider-upjet-ionoscloud/apis/natgateway/v1alpha1.Natgateway
	NatgatewayID *string `json:"natgatewayId,omitempty" tf:"natgateway_id,omitempty"`

	// Reference to a Natgateway in natgateway to populate natgatewayId.
	// +kubebuilder:validation:Optional
	NatgatewayIDRef *v1.Reference `json:"natgatewayIdRef,omitempty" tf:"-"`

	// Selector for a Natgateway in natgateway to populate natgatewayId.
	// +kubebuilder:validation:Optional
	NatgatewayIDSelector *v1.Selector `json:"natgatewayIdSelector,omitempty" tf:"-"`

	// [string] Protocol of the NAT gateway rule. Defaults to ALL. If protocol is 'ICMP' then targetPortRange start and end cannot be set.
	// Protocol of the NAT gateway rule. Defaults to ALL. If protocol is 'ICMP' then targetPortRange start and end cannot be set.
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// [string] Public IP address of the NAT gateway rule. Specifies the address used for masking outgoing packets source address field. Should be one of the customer reserved IP address already configured on the NAT gateway resource.
	// Public IP address of the NAT gateway rule. Specifies the address used for masking outgoing packets source address field. Should be one of the customer reserved IP address already configured on the NAT gateway resource
	PublicIP *string `json:"publicIp,omitempty" tf:"public_ip,omitempty"`

	// [string] Source subnet of the NAT gateway rule. For SNAT rules it specifies which packets this translation rule applies to based on the packets source IP address.
	// Source subnet of the NAT gateway rule. For SNAT rules it specifies which packets this translation rule applies to based on the packets source IP address.
	SourceSubnet *string `json:"sourceSubnet,omitempty" tf:"source_subnet,omitempty"`

	// Target port range of the NAT gateway rule. For SNAT rules it specifies which packets this translation rule applies to based on destination port. If none is provided, rule will match any port.
	// Target port range of the NAT gateway rule. For SNAT rules it specifies which packets this translation rule applies to based on destination port. If none is provided, rule will match any port
	TargetPortRange *TargetPortRangeInitParameters `json:"targetPortRange,omitempty" tf:"target_port_range,omitempty"`

	// [string] Target or destination subnet of the NAT gateway rule. For SNAT rules it specifies which packets this translation rule applies to based on the packets destination IP address. If none is provided, rule will match any address.
	// Target or destination subnet of the NAT gateway rule. For SNAT rules it specifies which packets this translation rule applies to based on the packets destination IP address. If none is provided, rule will match any address.
	TargetSubnet *string `json:"targetSubnet,omitempty" tf:"target_subnet,omitempty"`

	// [string] Type of the NAT gateway rule.
	// Type of the NAT gateway rule.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type RuleObservation struct {

	// [string] A Datacenter's UUID.
	DatacenterID *string `json:"datacenterId,omitempty" tf:"datacenter_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// [string] Name of the NAT gateway rule.
	// Name of the NAT gateway rule
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// [string] Nat Gateway's UUID.
	NatgatewayID *string `json:"natgatewayId,omitempty" tf:"natgateway_id,omitempty"`

	// [string] Protocol of the NAT gateway rule. Defaults to ALL. If protocol is 'ICMP' then targetPortRange start and end cannot be set.
	// Protocol of the NAT gateway rule. Defaults to ALL. If protocol is 'ICMP' then targetPortRange start and end cannot be set.
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// [string] Public IP address of the NAT gateway rule. Specifies the address used for masking outgoing packets source address field. Should be one of the customer reserved IP address already configured on the NAT gateway resource.
	// Public IP address of the NAT gateway rule. Specifies the address used for masking outgoing packets source address field. Should be one of the customer reserved IP address already configured on the NAT gateway resource
	PublicIP *string `json:"publicIp,omitempty" tf:"public_ip,omitempty"`

	// [string] Source subnet of the NAT gateway rule. For SNAT rules it specifies which packets this translation rule applies to based on the packets source IP address.
	// Source subnet of the NAT gateway rule. For SNAT rules it specifies which packets this translation rule applies to based on the packets source IP address.
	SourceSubnet *string `json:"sourceSubnet,omitempty" tf:"source_subnet,omitempty"`

	// Target port range of the NAT gateway rule. For SNAT rules it specifies which packets this translation rule applies to based on destination port. If none is provided, rule will match any port.
	// Target port range of the NAT gateway rule. For SNAT rules it specifies which packets this translation rule applies to based on destination port. If none is provided, rule will match any port
	TargetPortRange *TargetPortRangeObservation `json:"targetPortRange,omitempty" tf:"target_port_range,omitempty"`

	// [string] Target or destination subnet of the NAT gateway rule. For SNAT rules it specifies which packets this translation rule applies to based on the packets destination IP address. If none is provided, rule will match any address.
	// Target or destination subnet of the NAT gateway rule. For SNAT rules it specifies which packets this translation rule applies to based on the packets destination IP address. If none is provided, rule will match any address.
	TargetSubnet *string `json:"targetSubnet,omitempty" tf:"target_subnet,omitempty"`

	// [string] Type of the NAT gateway rule.
	// Type of the NAT gateway rule.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type RuleParameters struct {

	// [string] A Datacenter's UUID.
	// +crossplane:generate:reference:type=github.com/ionos-cloud/provider-upjet-ionoscloud/apis/compute/v1alpha1.Datacenter
	// +kubebuilder:validation:Optional
	DatacenterID *string `json:"datacenterId,omitempty" tf:"datacenter_id,omitempty"`

	// Reference to a Datacenter in compute to populate datacenterId.
	// +kubebuilder:validation:Optional
	DatacenterIDRef *v1.Reference `json:"datacenterIdRef,omitempty" tf:"-"`

	// Selector for a Datacenter in compute to populate datacenterId.
	// +kubebuilder:validation:Optional
	DatacenterIDSelector *v1.Selector `json:"datacenterIdSelector,omitempty" tf:"-"`

	// [string] Name of the NAT gateway rule.
	// Name of the NAT gateway rule
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// [string] Nat Gateway's UUID.
	// +crossplane:generate:reference:type=github.com/ionos-cloud/provider-upjet-ionoscloud/apis/natgateway/v1alpha1.Natgateway
	// +kubebuilder:validation:Optional
	NatgatewayID *string `json:"natgatewayId,omitempty" tf:"natgateway_id,omitempty"`

	// Reference to a Natgateway in natgateway to populate natgatewayId.
	// +kubebuilder:validation:Optional
	NatgatewayIDRef *v1.Reference `json:"natgatewayIdRef,omitempty" tf:"-"`

	// Selector for a Natgateway in natgateway to populate natgatewayId.
	// +kubebuilder:validation:Optional
	NatgatewayIDSelector *v1.Selector `json:"natgatewayIdSelector,omitempty" tf:"-"`

	// [string] Protocol of the NAT gateway rule. Defaults to ALL. If protocol is 'ICMP' then targetPortRange start and end cannot be set.
	// Protocol of the NAT gateway rule. Defaults to ALL. If protocol is 'ICMP' then targetPortRange start and end cannot be set.
	// +kubebuilder:validation:Optional
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// [string] Public IP address of the NAT gateway rule. Specifies the address used for masking outgoing packets source address field. Should be one of the customer reserved IP address already configured on the NAT gateway resource.
	// Public IP address of the NAT gateway rule. Specifies the address used for masking outgoing packets source address field. Should be one of the customer reserved IP address already configured on the NAT gateway resource
	// +kubebuilder:validation:Optional
	PublicIP *string `json:"publicIp,omitempty" tf:"public_ip,omitempty"`

	// [string] Source subnet of the NAT gateway rule. For SNAT rules it specifies which packets this translation rule applies to based on the packets source IP address.
	// Source subnet of the NAT gateway rule. For SNAT rules it specifies which packets this translation rule applies to based on the packets source IP address.
	// +kubebuilder:validation:Optional
	SourceSubnet *string `json:"sourceSubnet,omitempty" tf:"source_subnet,omitempty"`

	// Target port range of the NAT gateway rule. For SNAT rules it specifies which packets this translation rule applies to based on destination port. If none is provided, rule will match any port.
	// Target port range of the NAT gateway rule. For SNAT rules it specifies which packets this translation rule applies to based on destination port. If none is provided, rule will match any port
	// +kubebuilder:validation:Optional
	TargetPortRange *TargetPortRangeParameters `json:"targetPortRange,omitempty" tf:"target_port_range,omitempty"`

	// [string] Target or destination subnet of the NAT gateway rule. For SNAT rules it specifies which packets this translation rule applies to based on the packets destination IP address. If none is provided, rule will match any address.
	// Target or destination subnet of the NAT gateway rule. For SNAT rules it specifies which packets this translation rule applies to based on the packets destination IP address. If none is provided, rule will match any address.
	// +kubebuilder:validation:Optional
	TargetSubnet *string `json:"targetSubnet,omitempty" tf:"target_subnet,omitempty"`

	// [string] Type of the NAT gateway rule.
	// Type of the NAT gateway rule.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type TargetPortRangeInitParameters struct {

	// [int] Target port range end associated with the NAT gateway rule.
	// Target port range end associated with the NAT gateway rule.
	End *float64 `json:"end,omitempty" tf:"end,omitempty"`

	// [int] Target port range start associated with the NAT gateway rule.
	// Target port range start associated with the NAT gateway rule.
	Start *float64 `json:"start,omitempty" tf:"start,omitempty"`
}

type TargetPortRangeObservation struct {

	// [int] Target port range end associated with the NAT gateway rule.
	// Target port range end associated with the NAT gateway rule.
	End *float64 `json:"end,omitempty" tf:"end,omitempty"`

	// [int] Target port range start associated with the NAT gateway rule.
	// Target port range start associated with the NAT gateway rule.
	Start *float64 `json:"start,omitempty" tf:"start,omitempty"`
}

type TargetPortRangeParameters struct {

	// [int] Target port range end associated with the NAT gateway rule.
	// Target port range end associated with the NAT gateway rule.
	// +kubebuilder:validation:Optional
	End *float64 `json:"end,omitempty" tf:"end,omitempty"`

	// [int] Target port range start associated with the NAT gateway rule.
	// Target port range start associated with the NAT gateway rule.
	// +kubebuilder:validation:Optional
	Start *float64 `json:"start,omitempty" tf:"start,omitempty"`
}

// RuleSpec defines the desired state of Rule
type RuleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RuleParameters `json:"forProvider"`
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
	InitProvider RuleInitParameters `json:"initProvider,omitempty"`
}

// RuleStatus defines the observed state of Rule.
type RuleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RuleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Rule is the Schema for the Rules API. Creates and manages Nat Gateway Rule objects.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,ionos}
type Rule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.publicIp) || (has(self.initProvider) && has(self.initProvider.publicIp))",message="spec.forProvider.publicIp is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.sourceSubnet) || (has(self.initProvider) && has(self.initProvider.sourceSubnet))",message="spec.forProvider.sourceSubnet is a required parameter"
	Spec   RuleSpec   `json:"spec"`
	Status RuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RuleList contains a list of Rules
type RuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Rule `json:"items"`
}

// Repository type metadata.
var (
	Rule_Kind             = "Rule"
	Rule_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Rule_Kind}.String()
	Rule_KindAPIVersion   = Rule_Kind + "." + CRDGroupVersion.String()
	Rule_GroupVersionKind = CRDGroupVersion.WithKind(Rule_Kind)
)

func init() {
	SchemeBuilder.Register(&Rule{}, &RuleList{})
}
