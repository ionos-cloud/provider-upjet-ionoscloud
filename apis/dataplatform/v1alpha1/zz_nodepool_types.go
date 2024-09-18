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

type NodePoolInitParameters struct {

	// [map] Key-value pairs attached to node pool resource as Kubernetes annotations.
	// Key-value pairs attached to node pool resource as [Kubernetes annotations](https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/)
	// +mapType=granular
	Annotations map[string]*string `json:"annotations,omitempty" tf:"annotations,omitempty"`

	// [string] The availability zone of the virtual datacenter region where the node pool resources should be provisioned. Must be set with one of the values AUTO, ZONE_1 or ZONE_2. The default value is AUTO.
	// The availability zone of the virtual datacenter region where the node pool resources should be provisioned.
	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone,omitempty"`

	// [string] A valid CPU family name or AUTO if the platform shall choose the best fitting option. Available CPU architectures can be retrieved from the datacenter resource. The default value is AUTO.
	// A valid CPU family name or `AUTO` if the platform shall choose the best fitting option. Available CPU architectures can be retrieved from the datacenter resource.
	CPUFamily *string `json:"cpuFamily,omitempty" tf:"cpu_family,omitempty"`

	// [string] The UUID of an existing Dataplatform cluster.
	// The UUID of an existing Dataplatform cluster.
	// +crossplane:generate:reference:type=github.com/ionos-cloud/provider-upjet-ionoscloud/apis/dataplatform/v1alpha1.Cluster
	ClusterID *string `json:"clusterId,omitempty" tf:"cluster_id,omitempty"`

	// Reference to a Cluster in dataplatform to populate clusterId.
	// +kubebuilder:validation:Optional
	ClusterIDRef *v1.Reference `json:"clusterIdRef,omitempty" tf:"-"`

	// Selector for a Cluster in dataplatform to populate clusterId.
	// +kubebuilder:validation:Optional
	ClusterIDSelector *v1.Selector `json:"clusterIdSelector,omitempty" tf:"-"`

	// [int] The number of CPU cores per node. Must be set with a minimum value of 1. The default value is 4.
	// The number of CPU cores per node.
	CoresCount *float64 `json:"coresCount,omitempty" tf:"cores_count,omitempty"`

	// [map] Key-value pairs attached to the node pool resource as Kubernetes labels.
	// Key-value pairs attached to the node pool resource as [Kubernetes labels](https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/)
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// [string] Starting time of a weekly 4 hour-long window, during which maintenance might occur in hh:mm:ss format
	// Starting time of a weekly 4 hour-long window, during which maintenance might occur in hh:mm:ss format
	MaintenanceWindow []NodePoolMaintenanceWindowInitParameters `json:"maintenanceWindow,omitempty" tf:"maintenance_window,omitempty"`

	// [string] The name of your node pool. Must be 63 characters or less and must be empty or begin and end with an alphanumeric character ([a-z0-9A-Z]). It can contain dashes (-), underscores (_), dots (.), and alphanumerics in-between.
	// The name of your node pool. Must be 63 characters or less and must be empty or begin and end with an alphanumeric character ([a-z0-9A-Z]). It can contain dashes (-), underscores (_), dots (.), and alphanumerics in-between.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// [int] The number of nodes that make up the node pool. Must be set with a minimum value of 1.
	// The number of nodes that make up the node pool.
	NodeCount *float64 `json:"nodeCount,omitempty" tf:"node_count,omitempty"`

	// [int] The RAM size for one node in MB. Must be set in multiples of 1024MB, with a minimum size is of 2048MB. The default value is 4096.
	// The RAM size for one node in MB. Must be set in multiples of 1024 MB, with a minimum size is of 2048 MB.
	RAMSize *float64 `json:"ramSize,omitempty" tf:"ram_size,omitempty"`

	// [int] The size of the volume in GB. The size must be greater than 10GB. The default value is 20.
	// The size of the volume in GB. The size must be greater than 10GB.
	StorageSize *float64 `json:"storageSize,omitempty" tf:"storage_size,omitempty"`

	// [int] The type of hardware for the volume. Must be set with one of the values HDD or SSD. The default value is SSD.
	// The type of hardware for the volume.
	StorageType *string `json:"storageType,omitempty" tf:"storage_type,omitempty"`
}

type NodePoolMaintenanceWindowInitParameters struct {

	// [string] Must be set with one the values Monday, Tuesday, Wednesday, Thursday, Friday, Saturday or Sunday.
	DayOfTheWeek *string `json:"dayOfTheWeek,omitempty" tf:"day_of_the_week,omitempty"`

	// [string] Time at which the maintenance should start. Must conform to the 'HH:MM:SS' 24-hour format. This pattern matches the "HH:MM:SS 24-hour format with leading 0" format. For more information take a look at this link.
	// Time at which the maintenance should start.
	Time *string `json:"time,omitempty" tf:"time,omitempty"`
}

type NodePoolMaintenanceWindowObservation struct {

	// [string] Must be set with one the values Monday, Tuesday, Wednesday, Thursday, Friday, Saturday or Sunday.
	DayOfTheWeek *string `json:"dayOfTheWeek,omitempty" tf:"day_of_the_week,omitempty"`

	// [string] Time at which the maintenance should start. Must conform to the 'HH:MM:SS' 24-hour format. This pattern matches the "HH:MM:SS 24-hour format with leading 0" format. For more information take a look at this link.
	// Time at which the maintenance should start.
	Time *string `json:"time,omitempty" tf:"time,omitempty"`
}

type NodePoolMaintenanceWindowParameters struct {

	// [string] Must be set with one the values Monday, Tuesday, Wednesday, Thursday, Friday, Saturday or Sunday.
	// +kubebuilder:validation:Optional
	DayOfTheWeek *string `json:"dayOfTheWeek" tf:"day_of_the_week,omitempty"`

	// [string] Time at which the maintenance should start. Must conform to the 'HH:MM:SS' 24-hour format. This pattern matches the "HH:MM:SS 24-hour format with leading 0" format. For more information take a look at this link.
	// Time at which the maintenance should start.
	// +kubebuilder:validation:Optional
	Time *string `json:"time" tf:"time,omitempty"`
}

type NodePoolObservation struct {

	// [map] Key-value pairs attached to node pool resource as Kubernetes annotations.
	// Key-value pairs attached to node pool resource as [Kubernetes annotations](https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/)
	// +mapType=granular
	Annotations map[string]*string `json:"annotations,omitempty" tf:"annotations,omitempty"`

	// [string] The availability zone of the virtual datacenter region where the node pool resources should be provisioned. Must be set with one of the values AUTO, ZONE_1 or ZONE_2. The default value is AUTO.
	// The availability zone of the virtual datacenter region where the node pool resources should be provisioned.
	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone,omitempty"`

	// [string] A valid CPU family name or AUTO if the platform shall choose the best fitting option. Available CPU architectures can be retrieved from the datacenter resource. The default value is AUTO.
	// A valid CPU family name or `AUTO` if the platform shall choose the best fitting option. Available CPU architectures can be retrieved from the datacenter resource.
	CPUFamily *string `json:"cpuFamily,omitempty" tf:"cpu_family,omitempty"`

	// [string] The UUID of an existing Dataplatform cluster.
	// The UUID of an existing Dataplatform cluster.
	ClusterID *string `json:"clusterId,omitempty" tf:"cluster_id,omitempty"`

	// [int] The number of CPU cores per node. Must be set with a minimum value of 1. The default value is 4.
	// The number of CPU cores per node.
	CoresCount *float64 `json:"coresCount,omitempty" tf:"cores_count,omitempty"`

	// The UUID of the virtual data center (VDC) in which the nodepool is provisioned
	DatacenterID *string `json:"datacenterId,omitempty" tf:"datacenter_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// [map] Key-value pairs attached to the node pool resource as Kubernetes labels.
	// Key-value pairs attached to the node pool resource as [Kubernetes labels](https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/)
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// [string] Starting time of a weekly 4 hour-long window, during which maintenance might occur in hh:mm:ss format
	// Starting time of a weekly 4 hour-long window, during which maintenance might occur in hh:mm:ss format
	MaintenanceWindow []NodePoolMaintenanceWindowObservation `json:"maintenanceWindow,omitempty" tf:"maintenance_window,omitempty"`

	// [string] The name of your node pool. Must be 63 characters or less and must be empty or begin and end with an alphanumeric character ([a-z0-9A-Z]). It can contain dashes (-), underscores (_), dots (.), and alphanumerics in-between.
	// The name of your node pool. Must be 63 characters or less and must be empty or begin and end with an alphanumeric character ([a-z0-9A-Z]). It can contain dashes (-), underscores (_), dots (.), and alphanumerics in-between.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// [int] The number of nodes that make up the node pool. Must be set with a minimum value of 1.
	// The number of nodes that make up the node pool.
	NodeCount *float64 `json:"nodeCount,omitempty" tf:"node_count,omitempty"`

	// [int] The RAM size for one node in MB. Must be set in multiples of 1024MB, with a minimum size is of 2048MB. The default value is 4096.
	// The RAM size for one node in MB. Must be set in multiples of 1024 MB, with a minimum size is of 2048 MB.
	RAMSize *float64 `json:"ramSize,omitempty" tf:"ram_size,omitempty"`

	// [int] The size of the volume in GB. The size must be greater than 10GB. The default value is 20.
	// The size of the volume in GB. The size must be greater than 10GB.
	StorageSize *float64 `json:"storageSize,omitempty" tf:"storage_size,omitempty"`

	// [int] The type of hardware for the volume. Must be set with one of the values HDD or SSD. The default value is SSD.
	// The type of hardware for the volume.
	StorageType *string `json:"storageType,omitempty" tf:"storage_type,omitempty"`

	// The version of the Data Platform.
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type NodePoolParameters struct {

	// [map] Key-value pairs attached to node pool resource as Kubernetes annotations.
	// Key-value pairs attached to node pool resource as [Kubernetes annotations](https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/)
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Annotations map[string]*string `json:"annotations,omitempty" tf:"annotations,omitempty"`

	// [string] The availability zone of the virtual datacenter region where the node pool resources should be provisioned. Must be set with one of the values AUTO, ZONE_1 or ZONE_2. The default value is AUTO.
	// The availability zone of the virtual datacenter region where the node pool resources should be provisioned.
	// +kubebuilder:validation:Optional
	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone,omitempty"`

	// [string] A valid CPU family name or AUTO if the platform shall choose the best fitting option. Available CPU architectures can be retrieved from the datacenter resource. The default value is AUTO.
	// A valid CPU family name or `AUTO` if the platform shall choose the best fitting option. Available CPU architectures can be retrieved from the datacenter resource.
	// +kubebuilder:validation:Optional
	CPUFamily *string `json:"cpuFamily,omitempty" tf:"cpu_family,omitempty"`

	// [string] The UUID of an existing Dataplatform cluster.
	// The UUID of an existing Dataplatform cluster.
	// +crossplane:generate:reference:type=github.com/ionos-cloud/provider-upjet-ionoscloud/apis/dataplatform/v1alpha1.Cluster
	// +kubebuilder:validation:Optional
	ClusterID *string `json:"clusterId,omitempty" tf:"cluster_id,omitempty"`

	// Reference to a Cluster in dataplatform to populate clusterId.
	// +kubebuilder:validation:Optional
	ClusterIDRef *v1.Reference `json:"clusterIdRef,omitempty" tf:"-"`

	// Selector for a Cluster in dataplatform to populate clusterId.
	// +kubebuilder:validation:Optional
	ClusterIDSelector *v1.Selector `json:"clusterIdSelector,omitempty" tf:"-"`

	// [int] The number of CPU cores per node. Must be set with a minimum value of 1. The default value is 4.
	// The number of CPU cores per node.
	// +kubebuilder:validation:Optional
	CoresCount *float64 `json:"coresCount,omitempty" tf:"cores_count,omitempty"`

	// [map] Key-value pairs attached to the node pool resource as Kubernetes labels.
	// Key-value pairs attached to the node pool resource as [Kubernetes labels](https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/)
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// [string] Starting time of a weekly 4 hour-long window, during which maintenance might occur in hh:mm:ss format
	// Starting time of a weekly 4 hour-long window, during which maintenance might occur in hh:mm:ss format
	// +kubebuilder:validation:Optional
	MaintenanceWindow []NodePoolMaintenanceWindowParameters `json:"maintenanceWindow,omitempty" tf:"maintenance_window,omitempty"`

	// [string] The name of your node pool. Must be 63 characters or less and must be empty or begin and end with an alphanumeric character ([a-z0-9A-Z]). It can contain dashes (-), underscores (_), dots (.), and alphanumerics in-between.
	// The name of your node pool. Must be 63 characters or less and must be empty or begin and end with an alphanumeric character ([a-z0-9A-Z]). It can contain dashes (-), underscores (_), dots (.), and alphanumerics in-between.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// [int] The number of nodes that make up the node pool. Must be set with a minimum value of 1.
	// The number of nodes that make up the node pool.
	// +kubebuilder:validation:Optional
	NodeCount *float64 `json:"nodeCount,omitempty" tf:"node_count,omitempty"`

	// [int] The RAM size for one node in MB. Must be set in multiples of 1024MB, with a minimum size is of 2048MB. The default value is 4096.
	// The RAM size for one node in MB. Must be set in multiples of 1024 MB, with a minimum size is of 2048 MB.
	// +kubebuilder:validation:Optional
	RAMSize *float64 `json:"ramSize,omitempty" tf:"ram_size,omitempty"`

	// [int] The size of the volume in GB. The size must be greater than 10GB. The default value is 20.
	// The size of the volume in GB. The size must be greater than 10GB.
	// +kubebuilder:validation:Optional
	StorageSize *float64 `json:"storageSize,omitempty" tf:"storage_size,omitempty"`

	// [int] The type of hardware for the volume. Must be set with one of the values HDD or SSD. The default value is SSD.
	// The type of hardware for the volume.
	// +kubebuilder:validation:Optional
	StorageType *string `json:"storageType,omitempty" tf:"storage_type,omitempty"`
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

// NodePool is the Schema for the NodePools API. Creates and manages Dataplatform Node Pool objects.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,ionos}
type NodePool struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.nodeCount) || (has(self.initProvider) && has(self.initProvider.nodeCount))",message="spec.forProvider.nodeCount is a required parameter"
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