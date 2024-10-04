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

	// [list] Access to the K8s API server is restricted to these CIDRs. Cluster-internal traffic is not affected by this restriction. If no allowlist is specified, access is not restricted. If an IP without subnet mask is provided, the default value will be used: 32 for IPv4 and 128 for IPv6.
	// Access to the K8s API server is restricted to these CIDRs. Cluster-internal traffic is not affected by this restriction. If no allowlist is specified, access is not restricted. If an IP without subnet mask is provided, the default value will be used: 32 for IPv4 and 128 for IPv6.
	APISubnetAllowList []*string `json:"apiSubnetAllowList,omitempty" tf:"api_subnet_allow_list,omitempty"`

	// [bool] When set to true, allows the update of immutable fields by first destroying and then re-creating the cluster.
	// When set to true, allows the update of immutable fields by destroying and re-creating the cluster.
	AllowReplace *bool `json:"allowReplace,omitempty" tf:"allow_replace,omitempty"`

	// [string] The desired Kubernetes Version. For supported values, please check the API documentation. Downgrades are not supported. The provider will ignore downgrades of patch level.
	// The desired Kubernetes Version. For supported values, please check the API documentation. Downgrades are not supported. The provider will ignore downgrades of patch level.
	K8SVersion *string `json:"k8sVersion,omitempty" tf:"k8s_version,omitempty"`

	// [string] This attribute is mandatory if the cluster is private. The location must be enabled for your contract, or you must have a data center at that location. This property is not adjustable.
	// This attribute is mandatory if the cluster is private. The location must be enabled for your contract, or you must have a data center at that location. This attribute is immutable.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// A maintenance window comprise of a day of the week and a time for maintenance to be allowed
	// A maintenance window comprise of a day of the week and a time for maintenance to be allowed
	MaintenanceWindow *MaintenanceWindowInitParameters `json:"maintenanceWindow,omitempty" tf:"maintenance_window,omitempty"`

	// [string] The NAT gateway IP of the cluster if the cluster is private. This attribute is immutable. Must be a reserved IP in the same location as the cluster's location. This attribute is mandatory if the cluster is private.
	// The NAT gateway IP of the cluster if the cluster is private. This attribute is immutable. Must be a reserved IP in the same location as the cluster's location. This attribute is mandatory if the cluster is private.
	NATGatewayIP *string `json:"natGatewayIp,omitempty" tf:"nat_gateway_ip,omitempty"`

	// [string] The name of the Kubernetes Cluster.
	// The desired name for the cluster
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// [string] The node subnet of the cluster, if the cluster is private. This attribute is optional and immutable. Must be a valid CIDR notation for an IPv4 network prefix of 16 bits length.
	// The node subnet of the cluster, if the cluster is private. This attribute is optional and immutable. Must be a valid CIDR notation for an IPv4 network prefix of 16 bits length.
	NodeSubnet *string `json:"nodeSubnet,omitempty" tf:"node_subnet,omitempty"`

	// [boolean] Indicates if the cluster is public or private. This attribute is immutable.
	// The indicator if the cluster is public or private.
	Public *bool `json:"public,omitempty" tf:"public,omitempty"`

	// [list] List of S3 bucket configured for K8s usage. For now it contains only an S3 bucket used to store K8s API audit logs.
	// List of S3 bucket configured for K8s usage. For now it contains only an S3 bucket used to store K8s API audit logs.
	S3Buckets []S3BucketsInitParameters `json:"s3Buckets,omitempty" tf:"s3_buckets,omitempty"`
}

type ClusterObservation struct {

	// [list] Access to the K8s API server is restricted to these CIDRs. Cluster-internal traffic is not affected by this restriction. If no allowlist is specified, access is not restricted. If an IP without subnet mask is provided, the default value will be used: 32 for IPv4 and 128 for IPv6.
	// Access to the K8s API server is restricted to these CIDRs. Cluster-internal traffic is not affected by this restriction. If no allowlist is specified, access is not restricted. If an IP without subnet mask is provided, the default value will be used: 32 for IPv4 and 128 for IPv6.
	APISubnetAllowList []*string `json:"apiSubnetAllowList,omitempty" tf:"api_subnet_allow_list,omitempty"`

	// [bool] When set to true, allows the update of immutable fields by first destroying and then re-creating the cluster.
	// When set to true, allows the update of immutable fields by destroying and re-creating the cluster.
	AllowReplace *bool `json:"allowReplace,omitempty" tf:"allow_replace,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// [string] The desired Kubernetes Version. For supported values, please check the API documentation. Downgrades are not supported. The provider will ignore downgrades of patch level.
	// The desired Kubernetes Version. For supported values, please check the API documentation. Downgrades are not supported. The provider will ignore downgrades of patch level.
	K8SVersion *string `json:"k8sVersion,omitempty" tf:"k8s_version,omitempty"`

	// [string] This attribute is mandatory if the cluster is private. The location must be enabled for your contract, or you must have a data center at that location. This property is not adjustable.
	// This attribute is mandatory if the cluster is private. The location must be enabled for your contract, or you must have a data center at that location. This attribute is immutable.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// A maintenance window comprise of a day of the week and a time for maintenance to be allowed
	// A maintenance window comprise of a day of the week and a time for maintenance to be allowed
	MaintenanceWindow *MaintenanceWindowObservation `json:"maintenanceWindow,omitempty" tf:"maintenance_window,omitempty"`

	// [string] The NAT gateway IP of the cluster if the cluster is private. This attribute is immutable. Must be a reserved IP in the same location as the cluster's location. This attribute is mandatory if the cluster is private.
	// The NAT gateway IP of the cluster if the cluster is private. This attribute is immutable. Must be a reserved IP in the same location as the cluster's location. This attribute is mandatory if the cluster is private.
	NATGatewayIP *string `json:"natGatewayIp,omitempty" tf:"nat_gateway_ip,omitempty"`

	// [string] The name of the Kubernetes Cluster.
	// The desired name for the cluster
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// [string] The node subnet of the cluster, if the cluster is private. This attribute is optional and immutable. Must be a valid CIDR notation for an IPv4 network prefix of 16 bits length.
	// The node subnet of the cluster, if the cluster is private. This attribute is optional and immutable. Must be a valid CIDR notation for an IPv4 network prefix of 16 bits length.
	NodeSubnet *string `json:"nodeSubnet,omitempty" tf:"node_subnet,omitempty"`

	// [boolean] Indicates if the cluster is public or private. This attribute is immutable.
	// The indicator if the cluster is public or private.
	Public *bool `json:"public,omitempty" tf:"public,omitempty"`

	// [list] List of S3 bucket configured for K8s usage. For now it contains only an S3 bucket used to store K8s API audit logs.
	// List of S3 bucket configured for K8s usage. For now it contains only an S3 bucket used to store K8s API audit logs.
	S3Buckets []S3BucketsObservation `json:"s3Buckets,omitempty" tf:"s3_buckets,omitempty"`

	// (Computed)[list] List of versions that may be used for node pools under this cluster
	// List of versions that may be used for node pools under this cluster
	ViableNodePoolVersions []*string `json:"viableNodePoolVersions,omitempty" tf:"viable_node_pool_versions,omitempty"`
}

type ClusterParameters struct {

	// [list] Access to the K8s API server is restricted to these CIDRs. Cluster-internal traffic is not affected by this restriction. If no allowlist is specified, access is not restricted. If an IP without subnet mask is provided, the default value will be used: 32 for IPv4 and 128 for IPv6.
	// Access to the K8s API server is restricted to these CIDRs. Cluster-internal traffic is not affected by this restriction. If no allowlist is specified, access is not restricted. If an IP without subnet mask is provided, the default value will be used: 32 for IPv4 and 128 for IPv6.
	// +kubebuilder:validation:Optional
	APISubnetAllowList []*string `json:"apiSubnetAllowList,omitempty" tf:"api_subnet_allow_list,omitempty"`

	// [bool] When set to true, allows the update of immutable fields by first destroying and then re-creating the cluster.
	// When set to true, allows the update of immutable fields by destroying and re-creating the cluster.
	// +kubebuilder:validation:Optional
	AllowReplace *bool `json:"allowReplace,omitempty" tf:"allow_replace,omitempty"`

	// [string] The desired Kubernetes Version. For supported values, please check the API documentation. Downgrades are not supported. The provider will ignore downgrades of patch level.
	// The desired Kubernetes Version. For supported values, please check the API documentation. Downgrades are not supported. The provider will ignore downgrades of patch level.
	// +kubebuilder:validation:Optional
	K8SVersion *string `json:"k8sVersion,omitempty" tf:"k8s_version,omitempty"`

	// [string] This attribute is mandatory if the cluster is private. The location must be enabled for your contract, or you must have a data center at that location. This property is not adjustable.
	// This attribute is mandatory if the cluster is private. The location must be enabled for your contract, or you must have a data center at that location. This attribute is immutable.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// A maintenance window comprise of a day of the week and a time for maintenance to be allowed
	// A maintenance window comprise of a day of the week and a time for maintenance to be allowed
	// +kubebuilder:validation:Optional
	MaintenanceWindow *MaintenanceWindowParameters `json:"maintenanceWindow,omitempty" tf:"maintenance_window,omitempty"`

	// [string] The NAT gateway IP of the cluster if the cluster is private. This attribute is immutable. Must be a reserved IP in the same location as the cluster's location. This attribute is mandatory if the cluster is private.
	// The NAT gateway IP of the cluster if the cluster is private. This attribute is immutable. Must be a reserved IP in the same location as the cluster's location. This attribute is mandatory if the cluster is private.
	// +kubebuilder:validation:Optional
	NATGatewayIP *string `json:"natGatewayIp,omitempty" tf:"nat_gateway_ip,omitempty"`

	// [string] The name of the Kubernetes Cluster.
	// The desired name for the cluster
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// [string] The node subnet of the cluster, if the cluster is private. This attribute is optional and immutable. Must be a valid CIDR notation for an IPv4 network prefix of 16 bits length.
	// The node subnet of the cluster, if the cluster is private. This attribute is optional and immutable. Must be a valid CIDR notation for an IPv4 network prefix of 16 bits length.
	// +kubebuilder:validation:Optional
	NodeSubnet *string `json:"nodeSubnet,omitempty" tf:"node_subnet,omitempty"`

	// [boolean] Indicates if the cluster is public or private. This attribute is immutable.
	// The indicator if the cluster is public or private.
	// +kubebuilder:validation:Optional
	Public *bool `json:"public,omitempty" tf:"public,omitempty"`

	// [list] List of S3 bucket configured for K8s usage. For now it contains only an S3 bucket used to store K8s API audit logs.
	// List of S3 bucket configured for K8s usage. For now it contains only an S3 bucket used to store K8s API audit logs.
	// +kubebuilder:validation:Optional
	S3Buckets []S3BucketsParameters `json:"s3Buckets,omitempty" tf:"s3_buckets,omitempty"`
}

type MaintenanceWindowInitParameters struct {

	// [string] Day of the week when maintenance is allowed
	// Day of the week when maintenance is allowed
	DayOfTheWeek *string `json:"dayOfTheWeek,omitempty" tf:"day_of_the_week,omitempty"`

	// [string] A clock time in the day when maintenance is allowed
	// A clock time in the day when maintenance is allowed
	Time *string `json:"time,omitempty" tf:"time,omitempty"`
}

type MaintenanceWindowObservation struct {

	// [string] Day of the week when maintenance is allowed
	// Day of the week when maintenance is allowed
	DayOfTheWeek *string `json:"dayOfTheWeek,omitempty" tf:"day_of_the_week,omitempty"`

	// [string] A clock time in the day when maintenance is allowed
	// A clock time in the day when maintenance is allowed
	Time *string `json:"time,omitempty" tf:"time,omitempty"`
}

type MaintenanceWindowParameters struct {

	// [string] Day of the week when maintenance is allowed
	// Day of the week when maintenance is allowed
	// +kubebuilder:validation:Optional
	DayOfTheWeek *string `json:"dayOfTheWeek" tf:"day_of_the_week,omitempty"`

	// [string] A clock time in the day when maintenance is allowed
	// A clock time in the day when maintenance is allowed
	// +kubebuilder:validation:Optional
	Time *string `json:"time" tf:"time,omitempty"`
}

type S3BucketsInitParameters struct {

	// [string] The name of the Kubernetes Cluster.
	// Name of the S3 bucket
	// +crossplane:generate:reference:type=github.com/ionos-cloud/provider-upjet-ionoscloud/apis/s3/v1alpha1.Bucket
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Reference to a Bucket in s3 to populate name.
	// +kubebuilder:validation:Optional
	NameRef *v1.Reference `json:"nameRef,omitempty" tf:"-"`

	// Selector for a Bucket in s3 to populate name.
	// +kubebuilder:validation:Optional
	NameSelector *v1.Selector `json:"nameSelector,omitempty" tf:"-"`
}

type S3BucketsObservation struct {

	// [string] The name of the Kubernetes Cluster.
	// Name of the S3 bucket
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type S3BucketsParameters struct {

	// [string] The name of the Kubernetes Cluster.
	// Name of the S3 bucket
	// +crossplane:generate:reference:type=github.com/ionos-cloud/provider-upjet-ionoscloud/apis/s3/v1alpha1.Bucket
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Reference to a Bucket in s3 to populate name.
	// +kubebuilder:validation:Optional
	NameRef *v1.Reference `json:"nameRef,omitempty" tf:"-"`

	// Selector for a Bucket in s3 to populate name.
	// +kubebuilder:validation:Optional
	NameSelector *v1.Selector `json:"nameSelector,omitempty" tf:"-"`
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

// Cluster is the Schema for the Clusters API. Creates and manages IonosCloud Kubernetes Clusters.
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
