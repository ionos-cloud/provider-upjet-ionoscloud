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

type ConnectableDatacentersInitParameters struct {
}

type ConnectableDatacentersObservation struct {

	// The UUID of the connectable datacenter
	// The UUID of the connectable datacenter
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The physical location of the connectable datacenter
	// The physical location of the connectable datacenter
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// [string] The name of the cross-connection.
	// The name of the connectable datacenter
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type ConnectableDatacentersParameters struct {
}

type CrossconnectInitParameters struct {

	// A list containing all the connectable datacenters
	ConnectableDatacenters []ConnectableDatacentersInitParameters `json:"connectableDatacenters,omitempty" tf:"connectable_datacenters,omitempty"`

	// [string] A short description for the cross-connection.
	// The desired description
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// [string] The name of the cross-connection.
	// The desired name
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Computed) Lists LAN's joined to this cross connect
	// A list containing the details of all cross-connected datacenters
	Peers []PeersInitParameters `json:"peers,omitempty" tf:"peers,omitempty"`
}

type CrossconnectObservation struct {

	// A list containing all the connectable datacenters
	ConnectableDatacenters []ConnectableDatacentersObservation `json:"connectableDatacenters,omitempty" tf:"connectable_datacenters,omitempty"`

	// [string] A short description for the cross-connection.
	// The desired description
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The UUID of the connectable datacenter
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// [string] The name of the cross-connection.
	// The desired name
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Computed) Lists LAN's joined to this cross connect
	// A list containing the details of all cross-connected datacenters
	Peers []PeersObservation `json:"peers,omitempty" tf:"peers,omitempty"`
}

type CrossconnectParameters struct {

	// A list containing all the connectable datacenters
	// +kubebuilder:validation:Optional
	ConnectableDatacenters []ConnectableDatacentersParameters `json:"connectableDatacenters,omitempty" tf:"connectable_datacenters,omitempty"`

	// [string] A short description for the cross-connection.
	// The desired description
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// [string] The name of the cross-connection.
	// The desired name
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Computed) Lists LAN's joined to this cross connect
	// A list containing the details of all cross-connected datacenters
	// +kubebuilder:validation:Optional
	Peers []PeersParameters `json:"peers,omitempty" tf:"peers,omitempty"`
}

type PeersInitParameters struct {
}

type PeersObservation struct {

	// The id of the cross-connected datacenter
	// The id of the cross-connected datacenter
	DatacenterID *string `json:"datacenterId,omitempty" tf:"datacenter_id,omitempty"`

	// The name of the cross-connected datacenter
	// The name of the cross-connected datacenter
	DatacenterName *string `json:"datacenterName,omitempty" tf:"datacenter_name,omitempty"`

	// The id of the cross-connected LAN
	// The id of the cross-connected LAN
	LanID *string `json:"lanId,omitempty" tf:"lan_id,omitempty"`

	// The name of the cross-connected LAN
	// The name of the cross-connected LAN
	LanName *string `json:"lanName,omitempty" tf:"lan_name,omitempty"`

	// The physical location of the connectable datacenter
	// The location of the cross-connected datacenter
	Location *string `json:"location,omitempty" tf:"location,omitempty"`
}

type PeersParameters struct {
}

// CrossconnectSpec defines the desired state of Crossconnect
type CrossconnectSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CrossconnectParameters `json:"forProvider"`
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
	InitProvider CrossconnectInitParameters `json:"initProvider,omitempty"`
}

// CrossconnectStatus defines the observed state of Crossconnect.
type CrossconnectStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        CrossconnectObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Crossconnect is the Schema for the Crossconnects API. Creates and manages Cross Connections between virtual datacenters.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,ionos}
type Crossconnect struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   CrossconnectSpec   `json:"spec"`
	Status CrossconnectStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CrossconnectList contains a list of Crossconnects
type CrossconnectList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Crossconnect `json:"items"`
}

// Repository type metadata.
var (
	Crossconnect_Kind             = "Crossconnect"
	Crossconnect_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Crossconnect_Kind}.String()
	Crossconnect_KindAPIVersion   = Crossconnect_Kind + "." + CRDGroupVersion.String()
	Crossconnect_GroupVersionKind = CRDGroupVersion.WithKind(Crossconnect_Kind)
)

func init() {
	SchemeBuilder.Register(&Crossconnect{}, &CrossconnectList{})
}
