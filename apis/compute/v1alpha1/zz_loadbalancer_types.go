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

type LoadbalancerInitParameters struct {

	// [Boolean] Indicates if the load balancer will reserve an IP using DHCP.
	DHCP *bool `json:"dhcp,omitempty" tf:"dhcp,omitempty"`

	// [string] The ID of a Virtual Data Center.
	// +crossplane:generate:reference:type=github.com/ionos-cloud/provider-upjet-ionoscloud/apis/compute/v1alpha1.Datacenter
	DatacenterID *string `json:"datacenterId,omitempty" tf:"datacenter_id,omitempty"`

	// Reference to a Datacenter in compute to populate datacenterId.
	// +kubebuilder:validation:Optional
	DatacenterIDRef *v1.Reference `json:"datacenterIdRef,omitempty" tf:"-"`

	// Selector for a Datacenter in compute to populate datacenterId.
	// +kubebuilder:validation:Optional
	DatacenterIDSelector *v1.Selector `json:"datacenterIdSelector,omitempty" tf:"-"`

	// [string] IPv4 address of the load balancer.
	IP *string `json:"ip,omitempty" tf:"ip,omitempty"`

	// [string] The name of the load balancer.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// [list] A list of NIC IDs that are part of the load balancer.
	NicIds []*string `json:"nicIds,omitempty" tf:"nic_ids,omitempty"`
}

type LoadbalancerObservation struct {

	// [Boolean] Indicates if the load balancer will reserve an IP using DHCP.
	DHCP *bool `json:"dhcp,omitempty" tf:"dhcp,omitempty"`

	// [string] The ID of a Virtual Data Center.
	DatacenterID *string `json:"datacenterId,omitempty" tf:"datacenter_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// [string] IPv4 address of the load balancer.
	IP *string `json:"ip,omitempty" tf:"ip,omitempty"`

	// [string] The name of the load balancer.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// [list] A list of NIC IDs that are part of the load balancer.
	NicIds []*string `json:"nicIds,omitempty" tf:"nic_ids,omitempty"`
}

type LoadbalancerParameters struct {

	// [Boolean] Indicates if the load balancer will reserve an IP using DHCP.
	// +kubebuilder:validation:Optional
	DHCP *bool `json:"dhcp,omitempty" tf:"dhcp,omitempty"`

	// [string] The ID of a Virtual Data Center.
	// +crossplane:generate:reference:type=github.com/ionos-cloud/provider-upjet-ionoscloud/apis/compute/v1alpha1.Datacenter
	// +kubebuilder:validation:Optional
	DatacenterID *string `json:"datacenterId,omitempty" tf:"datacenter_id,omitempty"`

	// Reference to a Datacenter in compute to populate datacenterId.
	// +kubebuilder:validation:Optional
	DatacenterIDRef *v1.Reference `json:"datacenterIdRef,omitempty" tf:"-"`

	// Selector for a Datacenter in compute to populate datacenterId.
	// +kubebuilder:validation:Optional
	DatacenterIDSelector *v1.Selector `json:"datacenterIdSelector,omitempty" tf:"-"`

	// [string] IPv4 address of the load balancer.
	// +kubebuilder:validation:Optional
	IP *string `json:"ip,omitempty" tf:"ip,omitempty"`

	// [string] The name of the load balancer.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// [list] A list of NIC IDs that are part of the load balancer.
	// +kubebuilder:validation:Optional
	NicIds []*string `json:"nicIds,omitempty" tf:"nic_ids,omitempty"`
}

// LoadbalancerSpec defines the desired state of Loadbalancer
type LoadbalancerSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LoadbalancerParameters `json:"forProvider"`
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
	InitProvider LoadbalancerInitParameters `json:"initProvider,omitempty"`
}

// LoadbalancerStatus defines the observed state of Loadbalancer.
type LoadbalancerStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LoadbalancerObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Loadbalancer is the Schema for the Loadbalancers API. Creates and manages Load Balancers
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,ionos}
type Loadbalancer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.nicIds) || (has(self.initProvider) && has(self.initProvider.nicIds))",message="spec.forProvider.nicIds is a required parameter"
	Spec   LoadbalancerSpec   `json:"spec"`
	Status LoadbalancerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LoadbalancerList contains a list of Loadbalancers
type LoadbalancerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Loadbalancer `json:"items"`
}

// Repository type metadata.
var (
	Loadbalancer_Kind             = "Loadbalancer"
	Loadbalancer_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Loadbalancer_Kind}.String()
	Loadbalancer_KindAPIVersion   = Loadbalancer_Kind + "." + CRDGroupVersion.String()
	Loadbalancer_GroupVersionKind = CRDGroupVersion.WithKind(Loadbalancer_Kind)
)

func init() {
	SchemeBuilder.Register(&Loadbalancer{}, &LoadbalancerList{})
}