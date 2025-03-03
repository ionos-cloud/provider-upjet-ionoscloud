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

type AutoScalingInitParameters struct {

	// [int] The maximum number of worker nodes that the node pool can scale to. Should be greater than min_node_count
	// The maximum number of worker nodes that the node pool can scale to. Should be greater than min_node_count
	MaxNodeCount *float64 `json:"maxNodeCount,omitempty" tf:"max_node_count,omitempty"`

	// [int] The minimum number of worker nodes the node pool can scale down to. Should be less than max_node_count
	// The minimum number of worker nodes the node pool can scale down to. Should be less than max_node_count
	MinNodeCount *float64 `json:"minNodeCount,omitempty" tf:"min_node_count,omitempty"`
}

type AutoScalingObservation struct {

	// [int] The maximum number of worker nodes that the node pool can scale to. Should be greater than min_node_count
	// The maximum number of worker nodes that the node pool can scale to. Should be greater than min_node_count
	MaxNodeCount *float64 `json:"maxNodeCount,omitempty" tf:"max_node_count,omitempty"`

	// [int] The minimum number of worker nodes the node pool can scale down to. Should be less than max_node_count
	// The minimum number of worker nodes the node pool can scale down to. Should be less than max_node_count
	MinNodeCount *float64 `json:"minNodeCount,omitempty" tf:"min_node_count,omitempty"`
}

type AutoScalingParameters struct {

	// [int] The maximum number of worker nodes that the node pool can scale to. Should be greater than min_node_count
	// The maximum number of worker nodes that the node pool can scale to. Should be greater than min_node_count
	// +kubebuilder:validation:Optional
	MaxNodeCount *float64 `json:"maxNodeCount" tf:"max_node_count,omitempty"`

	// [int] The minimum number of worker nodes the node pool can scale down to. Should be less than max_node_count
	// The minimum number of worker nodes the node pool can scale down to. Should be less than max_node_count
	// +kubebuilder:validation:Optional
	MinNodeCount *float64 `json:"minNodeCount" tf:"min_node_count,omitempty"`
}

type LansInitParameters struct {

	// [boolean] Indicates if the Kubernetes Node Pool LAN will reserve an IP using DHCP. Default value is true
	// Indicates if the Kubernetes Node Pool LAN will reserve an IP using DHCP
	DHCP *bool `json:"dhcp,omitempty" tf:"dhcp,omitempty"`

	// [int] The LAN ID of an existing LAN at the related datacenter
	// The LAN ID of an existing LAN at the related datacenter
	// +crossplane:generate:reference:type=github.com/ionos-cloud/provider-upjet-ionoscloud/apis/compute/v1alpha1.Lan
	ID *float64 `json:"id,omitempty" tf:"id,omitempty"`

	// Reference to a Lan in compute to populate id.
	// +kubebuilder:validation:Optional
	IDRef *v1.Reference `json:"idRef,omitempty" tf:"-"`

	// Selector for a Lan in compute to populate id.
	// +kubebuilder:validation:Optional
	IDSelector *v1.Selector `json:"idSelector,omitempty" tf:"-"`

	// An array of additional LANs attached to worker nodes
	// An array of additional LANs attached to worker nodes
	Routes []RoutesInitParameters `json:"routes,omitempty" tf:"routes,omitempty"`
}

type LansObservation struct {

	// [boolean] Indicates if the Kubernetes Node Pool LAN will reserve an IP using DHCP. Default value is true
	// Indicates if the Kubernetes Node Pool LAN will reserve an IP using DHCP
	DHCP *bool `json:"dhcp,omitempty" tf:"dhcp,omitempty"`

	// [int] The LAN ID of an existing LAN at the related datacenter
	// The LAN ID of an existing LAN at the related datacenter
	ID *float64 `json:"id,omitempty" tf:"id,omitempty"`

	// An array of additional LANs attached to worker nodes
	// An array of additional LANs attached to worker nodes
	Routes []RoutesObservation `json:"routes,omitempty" tf:"routes,omitempty"`
}

type LansParameters struct {

	// [boolean] Indicates if the Kubernetes Node Pool LAN will reserve an IP using DHCP. Default value is true
	// Indicates if the Kubernetes Node Pool LAN will reserve an IP using DHCP
	// +kubebuilder:validation:Optional
	DHCP *bool `json:"dhcp,omitempty" tf:"dhcp,omitempty"`

	// [int] The LAN ID of an existing LAN at the related datacenter
	// The LAN ID of an existing LAN at the related datacenter
	// +crossplane:generate:reference:type=github.com/ionos-cloud/provider-upjet-ionoscloud/apis/compute/v1alpha1.Lan
	// +kubebuilder:validation:Optional
	ID *float64 `json:"id,omitempty" tf:"id,omitempty"`

	// Reference to a Lan in compute to populate id.
	// +kubebuilder:validation:Optional
	IDRef *v1.Reference `json:"idRef,omitempty" tf:"-"`

	// Selector for a Lan in compute to populate id.
	// +kubebuilder:validation:Optional
	IDSelector *v1.Selector `json:"idSelector,omitempty" tf:"-"`

	// An array of additional LANs attached to worker nodes
	// An array of additional LANs attached to worker nodes
	// +kubebuilder:validation:Optional
	Routes []RoutesParameters `json:"routes,omitempty" tf:"routes,omitempty"`
}

type NodePoolInitParameters struct {

	// [bool] When set to true, allows the update of immutable fields by first destroying and then re-creating the node pool.
	// When set to true, allows the update of immutable fields by destroying and re-creating the node pool
	AllowReplace *bool `json:"allowReplace,omitempty" tf:"allow_replace,omitempty"`

	// [map] A key/value map of annotations
	// +mapType=granular
	Annotations map[string]*string `json:"annotations,omitempty" tf:"annotations,omitempty"`

	// [string] Whether the Node Pool should autoscale. For more details, please check the API documentation
	// The range defining the minimum and maximum number of worker nodes that the managed node group can scale in
	AutoScaling *AutoScalingInitParameters `json:"autoScaling,omitempty" tf:"auto_scaling,omitempty"`

	// [string] - The desired Compute availability zone - See the API documentation for more information. This attribute is immutable.
	// The compute availability zone in which the nodes should exist
	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone,omitempty"`

	// [string] The desired CPU Family - See the API documentation for more information. This attribute is immutable.
	// CPU Family
	CPUFamily *string `json:"cpuFamily,omitempty" tf:"cpu_family,omitempty"`

	// [int] - The CPU cores count for each node of the node pool. This attribute is immutable.
	// CPU cores count
	CoresCount *float64 `json:"coresCount,omitempty" tf:"cores_count,omitempty"`

	// [string] A Datacenter's UUID
	// The UUID of the VDC
	// +crossplane:generate:reference:type=github.com/ionos-cloud/provider-upjet-ionoscloud/apis/compute/v1alpha1.Datacenter
	DatacenterID *string `json:"datacenterId,omitempty" tf:"datacenter_id,omitempty"`

	// Reference to a Datacenter in compute to populate datacenterId.
	// +kubebuilder:validation:Optional
	DatacenterIDRef *v1.Reference `json:"datacenterIdRef,omitempty" tf:"-"`

	// Selector for a Datacenter in compute to populate datacenterId.
	// +kubebuilder:validation:Optional
	DatacenterIDSelector *v1.Selector `json:"datacenterIdSelector,omitempty" tf:"-"`

	// [string] A k8s cluster's UUID
	// The UUID of an existing kubernetes cluster
	// +crossplane:generate:reference:type=github.com/ionos-cloud/provider-upjet-ionoscloud/apis/k8s/v1alpha1.Cluster
	K8SClusterID *string `json:"k8sClusterId,omitempty" tf:"k8s_cluster_id,omitempty"`

	// Reference to a Cluster in k8s to populate k8sClusterId.
	// +kubebuilder:validation:Optional
	K8SClusterIDRef *v1.Reference `json:"k8sClusterIdRef,omitempty" tf:"-"`

	// Selector for a Cluster in k8s to populate k8sClusterId.
	// +kubebuilder:validation:Optional
	K8SClusterIDSelector *v1.Selector `json:"k8sClusterIdSelector,omitempty" tf:"-"`

	// [string] The desired Kubernetes Version. For supported values, please check the API documentation. Downgrades are not supported. The provider will ignore downgrades of patch level.
	// The desired Kubernetes Version. For supported values, please check the API documentation. Downgrades are not supported. The provider will ignore downgrades of patch level.
	// +crossplane:generate:reference:type=github.com/ionos-cloud/provider-upjet-ionoscloud/apis/k8s/v1alpha1.Cluster
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("k8s_version",false)
	K8SVersion *string `json:"k8sVersion,omitempty" tf:"k8s_version,omitempty"`

	// Reference to a Cluster in k8s to populate k8sVersion.
	// +kubebuilder:validation:Optional
	K8SVersionRef *v1.Reference `json:"k8sVersionRef,omitempty" tf:"-"`

	// Selector for a Cluster in k8s to populate k8sVersion.
	// +kubebuilder:validation:Optional
	K8SVersionSelector *v1.Selector `json:"k8sVersionSelector,omitempty" tf:"-"`

	// [map] A key/value map of labels
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// [list] A list of numeric LAN id's you want this node pool to be part of. For more details, please check the API documentation, as well as the example above
	// A list of Local Area Networks the node pool should be part of
	Lans []LansInitParameters `json:"lans,omitempty" tf:"lans,omitempty"`

	// See the maintenance_window section in the example above
	// A maintenance window comprise of a day of the week and a time for maintenance to be allowed
	MaintenanceWindow *NodePoolMaintenanceWindowInitParameters `json:"maintenanceWindow,omitempty" tf:"maintenance_window,omitempty"`

	// [string] The name of the Kubernetes Cluster. This attribute is immutable.
	// The desired name for the node pool
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// [int] - The desired number of nodes in the node pool
	// The number of nodes in this node pool
	NodeCount *float64 `json:"nodeCount,omitempty" tf:"node_count,omitempty"`

	// [list] A list of public IPs associated with the node pool; must have at least node_count + 1 elements
	// A list of fixed IPs. Cannot be set on private clusters.
	PublicIps []*string `json:"publicIps,omitempty" tf:"public_ips,omitempty"`

	// [int] - The desired amount of RAM, in MB. This attribute is immutable.
	// The amount of RAM in MB
	RAMSize *float64 `json:"ramSize,omitempty" tf:"ram_size,omitempty"`

	// [int] - The size of the volume in GB. The size should be greater than 10GB. This attribute is immutable.
	// The total allocated storage capacity of a node in GB
	StorageSize *float64 `json:"storageSize,omitempty" tf:"storage_size,omitempty"`

	// [string] - The desired storage type - SSD/HDD. This attribute is immutable.
	// Storage type to use
	StorageType *string `json:"storageType,omitempty" tf:"storage_type,omitempty"`
}

type NodePoolMaintenanceWindowInitParameters struct {

	// [string] Day of the week when maintenance is allowed
	// Day of the week when maintenance is allowed
	DayOfTheWeek *string `json:"dayOfTheWeek,omitempty" tf:"day_of_the_week,omitempty"`

	// [string] A clock time in the day when maintenance is allowed
	// A clock time in the day when maintenance is allowed
	Time *string `json:"time,omitempty" tf:"time,omitempty"`
}

type NodePoolMaintenanceWindowObservation struct {

	// [string] Day of the week when maintenance is allowed
	// Day of the week when maintenance is allowed
	DayOfTheWeek *string `json:"dayOfTheWeek,omitempty" tf:"day_of_the_week,omitempty"`

	// [string] A clock time in the day when maintenance is allowed
	// A clock time in the day when maintenance is allowed
	Time *string `json:"time,omitempty" tf:"time,omitempty"`
}

type NodePoolMaintenanceWindowParameters struct {

	// [string] Day of the week when maintenance is allowed
	// Day of the week when maintenance is allowed
	// +kubebuilder:validation:Optional
	DayOfTheWeek *string `json:"dayOfTheWeek" tf:"day_of_the_week,omitempty"`

	// [string] A clock time in the day when maintenance is allowed
	// A clock time in the day when maintenance is allowed
	// +kubebuilder:validation:Optional
	Time *string `json:"time" tf:"time,omitempty"`
}

type NodePoolObservation struct {

	// [bool] When set to true, allows the update of immutable fields by first destroying and then re-creating the node pool.
	// When set to true, allows the update of immutable fields by destroying and re-creating the node pool
	AllowReplace *bool `json:"allowReplace,omitempty" tf:"allow_replace,omitempty"`

	// [map] A key/value map of annotations
	// +mapType=granular
	Annotations map[string]*string `json:"annotations,omitempty" tf:"annotations,omitempty"`

	// [string] Whether the Node Pool should autoscale. For more details, please check the API documentation
	// The range defining the minimum and maximum number of worker nodes that the managed node group can scale in
	AutoScaling *AutoScalingObservation `json:"autoScaling,omitempty" tf:"auto_scaling,omitempty"`

	// [string] - The desired Compute availability zone - See the API documentation for more information. This attribute is immutable.
	// The compute availability zone in which the nodes should exist
	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone,omitempty"`

	// [string] The desired CPU Family - See the API documentation for more information. This attribute is immutable.
	// CPU Family
	CPUFamily *string `json:"cpuFamily,omitempty" tf:"cpu_family,omitempty"`

	// [int] - The CPU cores count for each node of the node pool. This attribute is immutable.
	// CPU cores count
	CoresCount *float64 `json:"coresCount,omitempty" tf:"cores_count,omitempty"`

	// [string] A Datacenter's UUID
	// The UUID of the VDC
	DatacenterID *string `json:"datacenterId,omitempty" tf:"datacenter_id,omitempty"`

	// [int] The LAN ID of an existing LAN at the related datacenter
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// [string] A k8s cluster's UUID
	// The UUID of an existing kubernetes cluster
	K8SClusterID *string `json:"k8sClusterId,omitempty" tf:"k8s_cluster_id,omitempty"`

	// [string] The desired Kubernetes Version. For supported values, please check the API documentation. Downgrades are not supported. The provider will ignore downgrades of patch level.
	// The desired Kubernetes Version. For supported values, please check the API documentation. Downgrades are not supported. The provider will ignore downgrades of patch level.
	K8SVersion *string `json:"k8sVersion,omitempty" tf:"k8s_version,omitempty"`

	// [map] A key/value map of labels
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// [list] A list of numeric LAN id's you want this node pool to be part of. For more details, please check the API documentation, as well as the example above
	// A list of Local Area Networks the node pool should be part of
	Lans []LansObservation `json:"lans,omitempty" tf:"lans,omitempty"`

	// See the maintenance_window section in the example above
	// A maintenance window comprise of a day of the week and a time for maintenance to be allowed
	MaintenanceWindow *NodePoolMaintenanceWindowObservation `json:"maintenanceWindow,omitempty" tf:"maintenance_window,omitempty"`

	// [string] The name of the Kubernetes Cluster. This attribute is immutable.
	// The desired name for the node pool
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// [int] - The desired number of nodes in the node pool
	// The number of nodes in this node pool
	NodeCount *float64 `json:"nodeCount,omitempty" tf:"node_count,omitempty"`

	// [list] A list of public IPs associated with the node pool; must have at least node_count + 1 elements
	// A list of fixed IPs. Cannot be set on private clusters.
	PublicIps []*string `json:"publicIps,omitempty" tf:"public_ips,omitempty"`

	// [int] - The desired amount of RAM, in MB. This attribute is immutable.
	// The amount of RAM in MB
	RAMSize *float64 `json:"ramSize,omitempty" tf:"ram_size,omitempty"`

	// [int] - The size of the volume in GB. The size should be greater than 10GB. This attribute is immutable.
	// The total allocated storage capacity of a node in GB
	StorageSize *float64 `json:"storageSize,omitempty" tf:"storage_size,omitempty"`

	// [string] - The desired storage type - SSD/HDD. This attribute is immutable.
	// Storage type to use
	StorageType *string `json:"storageType,omitempty" tf:"storage_type,omitempty"`
}

type NodePoolParameters struct {

	// [bool] When set to true, allows the update of immutable fields by first destroying and then re-creating the node pool.
	// When set to true, allows the update of immutable fields by destroying and re-creating the node pool
	// +kubebuilder:validation:Optional
	AllowReplace *bool `json:"allowReplace,omitempty" tf:"allow_replace,omitempty"`

	// [map] A key/value map of annotations
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Annotations map[string]*string `json:"annotations,omitempty" tf:"annotations,omitempty"`

	// [string] Whether the Node Pool should autoscale. For more details, please check the API documentation
	// The range defining the minimum and maximum number of worker nodes that the managed node group can scale in
	// +kubebuilder:validation:Optional
	AutoScaling *AutoScalingParameters `json:"autoScaling,omitempty" tf:"auto_scaling,omitempty"`

	// [string] - The desired Compute availability zone - See the API documentation for more information. This attribute is immutable.
	// The compute availability zone in which the nodes should exist
	// +kubebuilder:validation:Optional
	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone,omitempty"`

	// [string] The desired CPU Family - See the API documentation for more information. This attribute is immutable.
	// CPU Family
	// +kubebuilder:validation:Optional
	CPUFamily *string `json:"cpuFamily,omitempty" tf:"cpu_family,omitempty"`

	// [int] - The CPU cores count for each node of the node pool. This attribute is immutable.
	// CPU cores count
	// +kubebuilder:validation:Optional
	CoresCount *float64 `json:"coresCount,omitempty" tf:"cores_count,omitempty"`

	// [string] A Datacenter's UUID
	// The UUID of the VDC
	// +crossplane:generate:reference:type=github.com/ionos-cloud/provider-upjet-ionoscloud/apis/compute/v1alpha1.Datacenter
	// +kubebuilder:validation:Optional
	DatacenterID *string `json:"datacenterId,omitempty" tf:"datacenter_id,omitempty"`

	// Reference to a Datacenter in compute to populate datacenterId.
	// +kubebuilder:validation:Optional
	DatacenterIDRef *v1.Reference `json:"datacenterIdRef,omitempty" tf:"-"`

	// Selector for a Datacenter in compute to populate datacenterId.
	// +kubebuilder:validation:Optional
	DatacenterIDSelector *v1.Selector `json:"datacenterIdSelector,omitempty" tf:"-"`

	// [string] A k8s cluster's UUID
	// The UUID of an existing kubernetes cluster
	// +crossplane:generate:reference:type=github.com/ionos-cloud/provider-upjet-ionoscloud/apis/k8s/v1alpha1.Cluster
	// +kubebuilder:validation:Optional
	K8SClusterID *string `json:"k8sClusterId,omitempty" tf:"k8s_cluster_id,omitempty"`

	// Reference to a Cluster in k8s to populate k8sClusterId.
	// +kubebuilder:validation:Optional
	K8SClusterIDRef *v1.Reference `json:"k8sClusterIdRef,omitempty" tf:"-"`

	// Selector for a Cluster in k8s to populate k8sClusterId.
	// +kubebuilder:validation:Optional
	K8SClusterIDSelector *v1.Selector `json:"k8sClusterIdSelector,omitempty" tf:"-"`

	// [string] The desired Kubernetes Version. For supported values, please check the API documentation. Downgrades are not supported. The provider will ignore downgrades of patch level.
	// The desired Kubernetes Version. For supported values, please check the API documentation. Downgrades are not supported. The provider will ignore downgrades of patch level.
	// +crossplane:generate:reference:type=github.com/ionos-cloud/provider-upjet-ionoscloud/apis/k8s/v1alpha1.Cluster
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("k8s_version",false)
	// +kubebuilder:validation:Optional
	K8SVersion *string `json:"k8sVersion,omitempty" tf:"k8s_version,omitempty"`

	// Reference to a Cluster in k8s to populate k8sVersion.
	// +kubebuilder:validation:Optional
	K8SVersionRef *v1.Reference `json:"k8sVersionRef,omitempty" tf:"-"`

	// Selector for a Cluster in k8s to populate k8sVersion.
	// +kubebuilder:validation:Optional
	K8SVersionSelector *v1.Selector `json:"k8sVersionSelector,omitempty" tf:"-"`

	// [map] A key/value map of labels
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// [list] A list of numeric LAN id's you want this node pool to be part of. For more details, please check the API documentation, as well as the example above
	// A list of Local Area Networks the node pool should be part of
	// +kubebuilder:validation:Optional
	Lans []LansParameters `json:"lans,omitempty" tf:"lans,omitempty"`

	// See the maintenance_window section in the example above
	// A maintenance window comprise of a day of the week and a time for maintenance to be allowed
	// +kubebuilder:validation:Optional
	MaintenanceWindow *NodePoolMaintenanceWindowParameters `json:"maintenanceWindow,omitempty" tf:"maintenance_window,omitempty"`

	// [string] The name of the Kubernetes Cluster. This attribute is immutable.
	// The desired name for the node pool
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// [int] - The desired number of nodes in the node pool
	// The number of nodes in this node pool
	// +kubebuilder:validation:Optional
	NodeCount *float64 `json:"nodeCount,omitempty" tf:"node_count,omitempty"`

	// [list] A list of public IPs associated with the node pool; must have at least node_count + 1 elements
	// A list of fixed IPs. Cannot be set on private clusters.
	// +kubebuilder:validation:Optional
	PublicIps []*string `json:"publicIps,omitempty" tf:"public_ips,omitempty"`

	// [int] - The desired amount of RAM, in MB. This attribute is immutable.
	// The amount of RAM in MB
	// +kubebuilder:validation:Optional
	RAMSize *float64 `json:"ramSize,omitempty" tf:"ram_size,omitempty"`

	// [int] - The size of the volume in GB. The size should be greater than 10GB. This attribute is immutable.
	// The total allocated storage capacity of a node in GB
	// +kubebuilder:validation:Optional
	StorageSize *float64 `json:"storageSize,omitempty" tf:"storage_size,omitempty"`

	// [string] - The desired storage type - SSD/HDD. This attribute is immutable.
	// Storage type to use
	// +kubebuilder:validation:Optional
	StorageType *string `json:"storageType,omitempty" tf:"storage_type,omitempty"`
}

type RoutesInitParameters struct {

	// [string] IPv4 or IPv6 Gateway IP for the route
	// IPv4 or IPv6 Gateway IP for the route
	GatewayIP *string `json:"gatewayIp,omitempty" tf:"gateway_ip,omitempty"`

	// [string] IPv4 or IPv6 CIDR to be routed via the interface
	// IPv4 or IPv6 CIDR to be routed via the interface
	Network *string `json:"network,omitempty" tf:"network,omitempty"`
}

type RoutesObservation struct {

	// [string] IPv4 or IPv6 Gateway IP for the route
	// IPv4 or IPv6 Gateway IP for the route
	GatewayIP *string `json:"gatewayIp,omitempty" tf:"gateway_ip,omitempty"`

	// [string] IPv4 or IPv6 CIDR to be routed via the interface
	// IPv4 or IPv6 CIDR to be routed via the interface
	Network *string `json:"network,omitempty" tf:"network,omitempty"`
}

type RoutesParameters struct {

	// [string] IPv4 or IPv6 Gateway IP for the route
	// IPv4 or IPv6 Gateway IP for the route
	// +kubebuilder:validation:Optional
	GatewayIP *string `json:"gatewayIp" tf:"gateway_ip,omitempty"`

	// [string] IPv4 or IPv6 CIDR to be routed via the interface
	// IPv4 or IPv6 CIDR to be routed via the interface
	// +kubebuilder:validation:Optional
	Network *string `json:"network" tf:"network,omitempty"`
}

// NodePoolSpec defines the desired state of NodePool
type NodePoolSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     NodePoolParameters `json:"forProvider"`
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
	InitProvider NodePoolInitParameters `json:"initProvider,omitempty"`
}

// NodePoolStatus defines the observed state of NodePool.
type NodePoolStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        NodePoolObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// NodePool is the Schema for the NodePools API. Creates and manages IonosCloud Kubernetes Node Pools.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,ionos}
type NodePool struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.availabilityZone) || (has(self.initProvider) && has(self.initProvider.availabilityZone))",message="spec.forProvider.availabilityZone is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.coresCount) || (has(self.initProvider) && has(self.initProvider.coresCount))",message="spec.forProvider.coresCount is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.cpuFamily) || (has(self.initProvider) && has(self.initProvider.cpuFamily))",message="spec.forProvider.cpuFamily is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.nodeCount) || (has(self.initProvider) && has(self.initProvider.nodeCount))",message="spec.forProvider.nodeCount is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.ramSize) || (has(self.initProvider) && has(self.initProvider.ramSize))",message="spec.forProvider.ramSize is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.storageSize) || (has(self.initProvider) && has(self.initProvider.storageSize))",message="spec.forProvider.storageSize is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.storageType) || (has(self.initProvider) && has(self.initProvider.storageType))",message="spec.forProvider.storageType is a required parameter"
	Spec   NodePoolSpec   `json:"spec"`
	Status NodePoolStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NodePoolList contains a list of NodePools
type NodePoolList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NodePool `json:"items"`
}

// Repository type metadata.
var (
	NodePool_Kind             = "NodePool"
	NodePool_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: NodePool_Kind}.String()
	NodePool_KindAPIVersion   = NodePool_Kind + "." + CRDGroupVersion.String()
	NodePool_GroupVersionKind = CRDGroupVersion.WithKind(NodePool_Kind)
)

func init() {
	SchemeBuilder.Register(&NodePool{}, &NodePoolList{})
}
