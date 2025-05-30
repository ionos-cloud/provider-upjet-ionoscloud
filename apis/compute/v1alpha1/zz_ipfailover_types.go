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

type IpfailoverInitParameters struct {

	// [string] The ID of a Virtual Data Center.
	// +crossplane:generate:reference:type=github.com/ionos-cloud/provider-upjet-ionoscloud/apis/compute/v1alpha1.Datacenter
	DatacenterID *string `json:"datacenterId,omitempty" tf:"datacenter_id,omitempty"`

	// Reference to a Datacenter in compute to populate datacenterId.
	// +kubebuilder:validation:Optional
	DatacenterIDRef *v1.Reference `json:"datacenterIdRef,omitempty" tf:"-"`

	// Selector for a Datacenter in compute to populate datacenterId.
	// +kubebuilder:validation:Optional
	DatacenterIDSelector *v1.Selector `json:"datacenterIdSelector,omitempty" tf:"-"`

	// [string] The reserved IP address to be used in the IP failover group.
	// Failover IP
	// +crossplane:generate:reference:type=github.com/ionos-cloud/provider-upjet-ionoscloud/apis/compute/v1alpha1.Ipblock
	// +crossplane:generate:reference:extractor=github.com/ionos-cloud/provider-upjet-ionoscloud/config/common.FirstIPBlockIP()
	IP *string `json:"ip,omitempty" tf:"ip,omitempty"`

	// Reference to a Ipblock in compute to populate ip.
	// +kubebuilder:validation:Optional
	IPRef *v1.Reference `json:"ipRef,omitempty" tf:"-"`

	// Selector for a Ipblock in compute to populate ip.
	// +kubebuilder:validation:Optional
	IPSelector *v1.Selector `json:"ipSelector,omitempty" tf:"-"`

	// [string] The ID of a LAN.
	// +crossplane:generate:reference:type=github.com/ionos-cloud/provider-upjet-ionoscloud/apis/compute/v1alpha1.Lan
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	LanID *string `json:"lanId,omitempty" tf:"lan_id,omitempty"`

	// Reference to a Lan in compute to populate lanId.
	// +kubebuilder:validation:Optional
	LanIDRef *v1.Reference `json:"lanIdRef,omitempty" tf:"-"`

	// Selector for a Lan in compute to populate lanId.
	// +kubebuilder:validation:Optional
	LanIDSelector *v1.Selector `json:"lanIdSelector,omitempty" tf:"-"`

	// [string] The ID of a NIC.
	// The UUID of the master NIC
	// +crossplane:generate:reference:type=github.com/ionos-cloud/provider-upjet-ionoscloud/apis/compute/v1alpha1.Server
	// +crossplane:generate:reference:extractor=github.com/ionos-cloud/provider-upjet-ionoscloud/config/common.ServerPrimaryNIC()
	Nicuuid *string `json:"nicuuid,omitempty" tf:"nicuuid,omitempty"`

	// Reference to a Server in compute to populate nicuuid.
	// +kubebuilder:validation:Optional
	NicuuidRef *v1.Reference `json:"nicuuidRef,omitempty" tf:"-"`

	// Selector for a Server in compute to populate nicuuid.
	// +kubebuilder:validation:Optional
	NicuuidSelector *v1.Selector `json:"nicuuidSelector,omitempty" tf:"-"`
}

type IpfailoverObservation struct {

	// [string] The ID of a Virtual Data Center.
	DatacenterID *string `json:"datacenterId,omitempty" tf:"datacenter_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// [string] The reserved IP address to be used in the IP failover group.
	// Failover IP
	IP *string `json:"ip,omitempty" tf:"ip,omitempty"`

	// [string] The ID of a LAN.
	LanID *string `json:"lanId,omitempty" tf:"lan_id,omitempty"`

	// [string] The ID of a NIC.
	// The UUID of the master NIC
	Nicuuid *string `json:"nicuuid,omitempty" tf:"nicuuid,omitempty"`
}

type IpfailoverParameters struct {

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

	// [string] The reserved IP address to be used in the IP failover group.
	// Failover IP
	// +crossplane:generate:reference:type=github.com/ionos-cloud/provider-upjet-ionoscloud/apis/compute/v1alpha1.Ipblock
	// +crossplane:generate:reference:extractor=github.com/ionos-cloud/provider-upjet-ionoscloud/config/common.FirstIPBlockIP()
	// +kubebuilder:validation:Optional
	IP *string `json:"ip,omitempty" tf:"ip,omitempty"`

	// Reference to a Ipblock in compute to populate ip.
	// +kubebuilder:validation:Optional
	IPRef *v1.Reference `json:"ipRef,omitempty" tf:"-"`

	// Selector for a Ipblock in compute to populate ip.
	// +kubebuilder:validation:Optional
	IPSelector *v1.Selector `json:"ipSelector,omitempty" tf:"-"`

	// [string] The ID of a LAN.
	// +crossplane:generate:reference:type=github.com/ionos-cloud/provider-upjet-ionoscloud/apis/compute/v1alpha1.Lan
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	LanID *string `json:"lanId,omitempty" tf:"lan_id,omitempty"`

	// Reference to a Lan in compute to populate lanId.
	// +kubebuilder:validation:Optional
	LanIDRef *v1.Reference `json:"lanIdRef,omitempty" tf:"-"`

	// Selector for a Lan in compute to populate lanId.
	// +kubebuilder:validation:Optional
	LanIDSelector *v1.Selector `json:"lanIdSelector,omitempty" tf:"-"`

	// [string] The ID of a NIC.
	// The UUID of the master NIC
	// +crossplane:generate:reference:type=github.com/ionos-cloud/provider-upjet-ionoscloud/apis/compute/v1alpha1.Server
	// +crossplane:generate:reference:extractor=github.com/ionos-cloud/provider-upjet-ionoscloud/config/common.ServerPrimaryNIC()
	// +kubebuilder:validation:Optional
	Nicuuid *string `json:"nicuuid,omitempty" tf:"nicuuid,omitempty"`

	// Reference to a Server in compute to populate nicuuid.
	// +kubebuilder:validation:Optional
	NicuuidRef *v1.Reference `json:"nicuuidRef,omitempty" tf:"-"`

	// Selector for a Server in compute to populate nicuuid.
	// +kubebuilder:validation:Optional
	NicuuidSelector *v1.Selector `json:"nicuuidSelector,omitempty" tf:"-"`
}

// IpfailoverSpec defines the desired state of Ipfailover
type IpfailoverSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     IpfailoverParameters `json:"forProvider"`
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
	InitProvider IpfailoverInitParameters `json:"initProvider,omitempty"`
}

// IpfailoverStatus defines the observed state of Ipfailover.
type IpfailoverStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        IpfailoverObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Ipfailover is the Schema for the Ipfailovers API. Creates and manages ipfailover objects.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,ionos}
type Ipfailover struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              IpfailoverSpec   `json:"spec"`
	Status            IpfailoverStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IpfailoverList contains a list of Ipfailovers
type IpfailoverList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Ipfailover `json:"items"`
}

// Repository type metadata.
var (
	Ipfailover_Kind             = "Ipfailover"
	Ipfailover_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Ipfailover_Kind}.String()
	Ipfailover_KindAPIVersion   = Ipfailover_Kind + "." + CRDGroupVersion.String()
	Ipfailover_GroupVersionKind = CRDGroupVersion.WithKind(Ipfailover_Kind)
)

func init() {
	SchemeBuilder.Register(&Ipfailover{}, &IpfailoverList{})
}
