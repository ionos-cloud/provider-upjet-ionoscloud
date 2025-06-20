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

type ClusterInitParameters struct {

	// [string] The UUID of the virtual data center (VDC) the cluster is provisioned.
	// The UUID of the virtual data center (VDC) in which the cluster is provisioned
	// +crossplane:generate:reference:type=github.com/ionos-cloud/provider-upjet-ionoscloud/apis/compute/v1alpha1.Datacenter
	DatacenterID *string `json:"datacenterId,omitempty" tf:"datacenter_id,omitempty"`

	// Reference to a Datacenter in compute to populate datacenterId.
	// +kubebuilder:validation:Optional
	DatacenterIDRef *v1.Reference `json:"datacenterIdRef,omitempty" tf:"-"`

	// Selector for a Datacenter in compute to populate datacenterId.
	// +kubebuilder:validation:Optional
	DatacenterIDSelector *v1.Selector `json:"datacenterIdSelector,omitempty" tf:"-"`

	// [list] A list of LANs you want this node pool to be part of.
	// A list of LANs you want this node pool to be part of
	Lans []LansInitParameters `json:"lans,omitempty" tf:"lans,omitempty"`

	// Starting time of a weekly 4 hour-long window, during which maintenance might occur in hh:mm:ss format
	// Starting time of a weekly 4 hour-long window, during which maintenance might occur in hh:mm:ss format
	MaintenanceWindow *MaintenanceWindowInitParameters `json:"maintenanceWindow,omitempty" tf:"maintenance_window,omitempty"`

	// [string] The name of your cluster. Must be 63 characters or less and must be empty or begin and end with an alphanumeric character ([a-z0-9A-Z]). It can contain dashes (-), underscores (_), dots (.), and alphanumerics in-between.
	// The name of your cluster. Must be 63 characters or less and must be empty or begin and end with an alphanumeric character ([a-z0-9A-Z]). It can contain dashes (-), underscores (_), dots (.), and alphanumerics in-between.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// [int] The version of the Data Platform.
	// The version of the Data Platform.
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type ClusterObservation struct {

	// [string] The UUID of the virtual data center (VDC) the cluster is provisioned.
	// The UUID of the virtual data center (VDC) in which the cluster is provisioned
	DatacenterID *string `json:"datacenterId,omitempty" tf:"datacenter_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// [list] A list of LANs you want this node pool to be part of.
	// A list of LANs you want this node pool to be part of
	Lans []LansObservation `json:"lans,omitempty" tf:"lans,omitempty"`

	// Starting time of a weekly 4 hour-long window, during which maintenance might occur in hh:mm:ss format
	// Starting time of a weekly 4 hour-long window, during which maintenance might occur in hh:mm:ss format
	MaintenanceWindow *MaintenanceWindowObservation `json:"maintenanceWindow,omitempty" tf:"maintenance_window,omitempty"`

	// [string] The name of your cluster. Must be 63 characters or less and must be empty or begin and end with an alphanumeric character ([a-z0-9A-Z]). It can contain dashes (-), underscores (_), dots (.), and alphanumerics in-between.
	// The name of your cluster. Must be 63 characters or less and must be empty or begin and end with an alphanumeric character ([a-z0-9A-Z]). It can contain dashes (-), underscores (_), dots (.), and alphanumerics in-between.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// [int] The version of the Data Platform.
	// The version of the Data Platform.
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type ClusterParameters struct {

	// [string] The UUID of the virtual data center (VDC) the cluster is provisioned.
	// The UUID of the virtual data center (VDC) in which the cluster is provisioned
	// +crossplane:generate:reference:type=github.com/ionos-cloud/provider-upjet-ionoscloud/apis/compute/v1alpha1.Datacenter
	// +kubebuilder:validation:Optional
	DatacenterID *string `json:"datacenterId,omitempty" tf:"datacenter_id,omitempty"`

	// Reference to a Datacenter in compute to populate datacenterId.
	// +kubebuilder:validation:Optional
	DatacenterIDRef *v1.Reference `json:"datacenterIdRef,omitempty" tf:"-"`

	// Selector for a Datacenter in compute to populate datacenterId.
	// +kubebuilder:validation:Optional
	DatacenterIDSelector *v1.Selector `json:"datacenterIdSelector,omitempty" tf:"-"`

	// [list] A list of LANs you want this node pool to be part of.
	// A list of LANs you want this node pool to be part of
	// +kubebuilder:validation:Optional
	Lans []LansParameters `json:"lans,omitempty" tf:"lans,omitempty"`

	// Starting time of a weekly 4 hour-long window, during which maintenance might occur in hh:mm:ss format
	// Starting time of a weekly 4 hour-long window, during which maintenance might occur in hh:mm:ss format
	// +kubebuilder:validation:Optional
	MaintenanceWindow *MaintenanceWindowParameters `json:"maintenanceWindow,omitempty" tf:"maintenance_window,omitempty"`

	// [string] The name of your cluster. Must be 63 characters or less and must be empty or begin and end with an alphanumeric character ([a-z0-9A-Z]). It can contain dashes (-), underscores (_), dots (.), and alphanumerics in-between.
	// The name of your cluster. Must be 63 characters or less and must be empty or begin and end with an alphanumeric character ([a-z0-9A-Z]). It can contain dashes (-), underscores (_), dots (.), and alphanumerics in-between.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// [int] The version of the Data Platform.
	// The version of the Data Platform.
	// +kubebuilder:validation:Optional
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type LansInitParameters struct {

	// [bool] Indicates if the Kubernetes node pool LAN will reserve an IP using DHCP. The default value is 'true'.
	// Indicates if the Kubernetes node pool LAN will reserve an IP using DHCP. The default value is 'true'
	DHCP *bool `json:"dhcp,omitempty" tf:"dhcp,omitempty"`

	// [string] The LAN ID of an existing LAN at the related data center.
	// The LAN ID of an existing LAN at the related data center
	// +crossplane:generate:reference:type=github.com/ionos-cloud/provider-upjet-ionoscloud/apis/compute/v1alpha1.Lan
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	LanID *string `json:"lanId,omitempty" tf:"lan_id,omitempty"`

	// Reference to a Lan in compute to populate lanId.
	// +kubebuilder:validation:Optional
	LanIDRef *v1.Reference `json:"lanIdRef,omitempty" tf:"-"`

	// Selector for a Lan in compute to populate lanId.
	// +kubebuilder:validation:Optional
	LanIDSelector *v1.Selector `json:"lanIdSelector,omitempty" tf:"-"`

	// [list] An array of additional LANs attached to worker nodes.
	// An array of additional LANs attached to worker nodes
	Routes []RoutesInitParameters `json:"routes,omitempty" tf:"routes,omitempty"`
}

type LansObservation struct {

	// [bool] Indicates if the Kubernetes node pool LAN will reserve an IP using DHCP. The default value is 'true'.
	// Indicates if the Kubernetes node pool LAN will reserve an IP using DHCP. The default value is 'true'
	DHCP *bool `json:"dhcp,omitempty" tf:"dhcp,omitempty"`

	// [string] The LAN ID of an existing LAN at the related data center.
	// The LAN ID of an existing LAN at the related data center
	LanID *string `json:"lanId,omitempty" tf:"lan_id,omitempty"`

	// [list] An array of additional LANs attached to worker nodes.
	// An array of additional LANs attached to worker nodes
	Routes []RoutesObservation `json:"routes,omitempty" tf:"routes,omitempty"`
}

type LansParameters struct {

	// [bool] Indicates if the Kubernetes node pool LAN will reserve an IP using DHCP. The default value is 'true'.
	// Indicates if the Kubernetes node pool LAN will reserve an IP using DHCP. The default value is 'true'
	// +kubebuilder:validation:Optional
	DHCP *bool `json:"dhcp,omitempty" tf:"dhcp,omitempty"`

	// [string] The LAN ID of an existing LAN at the related data center.
	// The LAN ID of an existing LAN at the related data center
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

	// [list] An array of additional LANs attached to worker nodes.
	// An array of additional LANs attached to worker nodes
	// +kubebuilder:validation:Optional
	Routes []RoutesParameters `json:"routes,omitempty" tf:"routes,omitempty"`
}

type MaintenanceWindowInitParameters struct {

	// [string] Must be set with one the values Monday, Tuesday, Wednesday, Thursday, Friday, Saturday or Sunday.
	DayOfTheWeek *string `json:"dayOfTheWeek,omitempty" tf:"day_of_the_week,omitempty"`

	// [string] Time at which the maintenance should start. Must conform to the 'HH:MM:SS' 24-hour format. This pattern matches the "HH:MM:SS 24-hour format with leading 0" format. For more information take a look at this link.
	// Time at which the maintenance should start.
	Time *string `json:"time,omitempty" tf:"time,omitempty"`
}

type MaintenanceWindowObservation struct {

	// [string] Must be set with one the values Monday, Tuesday, Wednesday, Thursday, Friday, Saturday or Sunday.
	DayOfTheWeek *string `json:"dayOfTheWeek,omitempty" tf:"day_of_the_week,omitempty"`

	// [string] Time at which the maintenance should start. Must conform to the 'HH:MM:SS' 24-hour format. This pattern matches the "HH:MM:SS 24-hour format with leading 0" format. For more information take a look at this link.
	// Time at which the maintenance should start.
	Time *string `json:"time,omitempty" tf:"time,omitempty"`
}

type MaintenanceWindowParameters struct {

	// [string] Must be set with one the values Monday, Tuesday, Wednesday, Thursday, Friday, Saturday or Sunday.
	// +kubebuilder:validation:Optional
	DayOfTheWeek *string `json:"dayOfTheWeek" tf:"day_of_the_week,omitempty"`

	// [string] Time at which the maintenance should start. Must conform to the 'HH:MM:SS' 24-hour format. This pattern matches the "HH:MM:SS 24-hour format with leading 0" format. For more information take a look at this link.
	// Time at which the maintenance should start.
	// +kubebuilder:validation:Optional
	Time *string `json:"time" tf:"time,omitempty"`
}

type RoutesInitParameters struct {

	// [string] IPv4 or IPv6 gateway IP for the route.
	// IPv4 or IPv6 gateway IP for the route
	Gateway *string `json:"gateway,omitempty" tf:"gateway,omitempty"`

	// [string] IPv4 or IPv6 CIDR to be routed via the interface.
	// IPv4 or IPv6 CIDR to be routed via the interface
	Network *string `json:"network,omitempty" tf:"network,omitempty"`
}

type RoutesObservation struct {

	// [string] IPv4 or IPv6 gateway IP for the route.
	// IPv4 or IPv6 gateway IP for the route
	Gateway *string `json:"gateway,omitempty" tf:"gateway,omitempty"`

	// [string] IPv4 or IPv6 CIDR to be routed via the interface.
	// IPv4 or IPv6 CIDR to be routed via the interface
	Network *string `json:"network,omitempty" tf:"network,omitempty"`
}

type RoutesParameters struct {

	// [string] IPv4 or IPv6 gateway IP for the route.
	// IPv4 or IPv6 gateway IP for the route
	// +kubebuilder:validation:Optional
	Gateway *string `json:"gateway" tf:"gateway,omitempty"`

	// [string] IPv4 or IPv6 CIDR to be routed via the interface.
	// IPv4 or IPv6 CIDR to be routed via the interface
	// +kubebuilder:validation:Optional
	Network *string `json:"network" tf:"network,omitempty"`
}

// ClusterSpec defines the desired state of Cluster
type ClusterSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ClusterParameters `json:"forProvider"`
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
	InitProvider ClusterInitParameters `json:"initProvider,omitempty"`
}

// ClusterStatus defines the observed state of Cluster.
type ClusterStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ClusterObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Cluster is the Schema for the Clusters API. Creates and manages Dataplatform Cluster objects.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,ionos}
type Cluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   ClusterSpec   `json:"spec"`
	Status ClusterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ClusterList contains a list of Clusters
type ClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Cluster `json:"items"`
}

// Repository type metadata.
var (
	Cluster_Kind             = "Cluster"
	Cluster_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Cluster_Kind}.String()
	Cluster_KindAPIVersion   = Cluster_Kind + "." + CRDGroupVersion.String()
	Cluster_GroupVersionKind = CRDGroupVersion.WithKind(Cluster_Kind)
)

func init() {
	SchemeBuilder.Register(&Cluster{}, &ClusterList{})
}
