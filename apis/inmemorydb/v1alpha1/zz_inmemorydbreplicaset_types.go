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

	// [string] The IP and subnet for your instance. Note the following unavailable IP ranges: 10.233.64.0/18, 10.233.0.0/18, 10.233.114.0/24.
	// The IP and subnet for your instance. Note the following unavailable IP ranges: 10.233.64.0/18, 10.233.0.0/18, 10.233.114.0/24
	Cidr *string `json:"cidr,omitempty" tf:"cidr,omitempty"`

	// [string] The datacenter to connect your instance to.
	// The datacenter to connect your instance to.
	// +crossplane:generate:reference:type=github.com/ionos-cloud/provider-upjet-ionoscloud/apis/compute/v1alpha1.Datacenter
	DatacenterID *string `json:"datacenterId,omitempty" tf:"datacenter_id,omitempty"`

	// Reference to a Datacenter in compute to populate datacenterId.
	// +kubebuilder:validation:Optional
	DatacenterIDRef *v1.Reference `json:"datacenterIdRef,omitempty" tf:"-"`

	// Selector for a Datacenter in compute to populate datacenterId.
	// +kubebuilder:validation:Optional
	DatacenterIDSelector *v1.Selector `json:"datacenterIdSelector,omitempty" tf:"-"`

	// [string] The numeric LAN ID to connect your instance to.
	// The numeric LAN ID to connect your instance to.
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

	// [string] The IP and subnet for your instance. Note the following unavailable IP ranges: 10.233.64.0/18, 10.233.0.0/18, 10.233.114.0/24.
	// The IP and subnet for your instance. Note the following unavailable IP ranges: 10.233.64.0/18, 10.233.0.0/18, 10.233.114.0/24
	Cidr *string `json:"cidr,omitempty" tf:"cidr,omitempty"`

	// [string] The datacenter to connect your instance to.
	// The datacenter to connect your instance to.
	DatacenterID *string `json:"datacenterId,omitempty" tf:"datacenter_id,omitempty"`

	// [string] The numeric LAN ID to connect your instance to.
	// The numeric LAN ID to connect your instance to.
	LanID *string `json:"lanId,omitempty" tf:"lan_id,omitempty"`
}

type ConnectionsParameters struct {

	// [string] The IP and subnet for your instance. Note the following unavailable IP ranges: 10.233.64.0/18, 10.233.0.0/18, 10.233.114.0/24.
	// The IP and subnet for your instance. Note the following unavailable IP ranges: 10.233.64.0/18, 10.233.0.0/18, 10.233.114.0/24
	// +kubebuilder:validation:Optional
	Cidr *string `json:"cidr" tf:"cidr,omitempty"`

	// [string] The datacenter to connect your instance to.
	// The datacenter to connect your instance to.
	// +crossplane:generate:reference:type=github.com/ionos-cloud/provider-upjet-ionoscloud/apis/compute/v1alpha1.Datacenter
	// +kubebuilder:validation:Optional
	DatacenterID *string `json:"datacenterId,omitempty" tf:"datacenter_id,omitempty"`

	// Reference to a Datacenter in compute to populate datacenterId.
	// +kubebuilder:validation:Optional
	DatacenterIDRef *v1.Reference `json:"datacenterIdRef,omitempty" tf:"-"`

	// Selector for a Datacenter in compute to populate datacenterId.
	// +kubebuilder:validation:Optional
	DatacenterIDSelector *v1.Selector `json:"datacenterIdSelector,omitempty" tf:"-"`

	// [string] The numeric LAN ID to connect your instance to.
	// The numeric LAN ID to connect your instance to.
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

type CredentialsInitParameters struct {

	// [object] The hashed password for a InMemoryDB user.
	// The hashed password for a InMemoryDB user.
	HashedPassword []HashedPasswordInitParameters `json:"hashedPassword,omitempty" tf:"hashed_password,omitempty"`

	// [string] The password for a InMemoryDB user, this is a field that is marked as Sensitive.
	// The password for a InMemoryDB user.
	PlainTextPasswordSecretRef *v1.SecretKeySelector `json:"plainTextPasswordSecretRef,omitempty" tf:"-"`

	// [string] The username for the initial InMemoryDB user. Some system usernames are restricted (e.g. 'admin', 'standby').
	// The username for the initial InMemoryDB user. Some system usernames are restricted (e.g. 'admin', 'standby').
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

type CredentialsObservation struct {

	// [object] The hashed password for a InMemoryDB user.
	// The hashed password for a InMemoryDB user.
	HashedPassword []HashedPasswordObservation `json:"hashedPassword,omitempty" tf:"hashed_password,omitempty"`

	// [string] The username for the initial InMemoryDB user. Some system usernames are restricted (e.g. 'admin', 'standby').
	// The username for the initial InMemoryDB user. Some system usernames are restricted (e.g. 'admin', 'standby').
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

type CredentialsParameters struct {

	// [object] The hashed password for a InMemoryDB user.
	// The hashed password for a InMemoryDB user.
	// +kubebuilder:validation:Optional
	HashedPassword []HashedPasswordParameters `json:"hashedPassword,omitempty" tf:"hashed_password,omitempty"`

	// [string] The password for a InMemoryDB user, this is a field that is marked as Sensitive.
	// The password for a InMemoryDB user.
	// +kubebuilder:validation:Optional
	PlainTextPasswordSecretRef *v1.SecretKeySelector `json:"plainTextPasswordSecretRef,omitempty" tf:"-"`

	// [string] The username for the initial InMemoryDB user. Some system usernames are restricted (e.g. 'admin', 'standby').
	// The username for the initial InMemoryDB user. Some system usernames are restricted (e.g. 'admin', 'standby').
	// +kubebuilder:validation:Optional
	Username *string `json:"username" tf:"username,omitempty"`
}

type HashedPasswordInitParameters struct {

	// [string] The value can be only: "SHA-256".
	Algorithm *string `json:"algorithm,omitempty" tf:"algorithm,omitempty"`

	// [string] The hashed password.
	Hash *string `json:"hash,omitempty" tf:"hash,omitempty"`
}

type HashedPasswordObservation struct {

	// [string] The value can be only: "SHA-256".
	Algorithm *string `json:"algorithm,omitempty" tf:"algorithm,omitempty"`

	// [string] The hashed password.
	Hash *string `json:"hash,omitempty" tf:"hash,omitempty"`
}

type HashedPasswordParameters struct {

	// [string] The value can be only: "SHA-256".
	// +kubebuilder:validation:Optional
	Algorithm *string `json:"algorithm" tf:"algorithm,omitempty"`

	// [string] The hashed password.
	// +kubebuilder:validation:Optional
	Hash *string `json:"hash" tf:"hash,omitempty"`
}

type InMemoryDBReplicasetInitParameters struct {

	// [object] The network connection for your replica set. Only one connection is allowed. Updates to the value of the fields force the replica set to be re-created.
	// The network connection for your replica set. Only one connection is allowed.
	Connections []ConnectionsInitParameters `json:"connections,omitempty" tf:"connections,omitempty"`

	// [object] Credentials for the InMemoryDB replicaset, only one type of password can be used since they are mutually exclusive. These values are used to create the initial InMemoryDB user, updating any of these will force recreation of the replica set resource.
	// Credentials for the InMemoryDB replicaset.
	Credentials []CredentialsInitParameters `json:"credentials,omitempty" tf:"credentials,omitempty"`

	// [string] The human-readable name of your replica set.
	// The human readable name of your replica set.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// [string] The eviction policy for the replica set, possible values are:
	// The eviction policy for the replica set.
	EvictionPolicy *string `json:"evictionPolicy,omitempty" tf:"eviction_policy,omitempty"`

	// [string] The ID of a snapshot to restore the replica set from. If set, the replica set will be created from the snapshot.
	// The ID of a snapshot to restore the replica set from. If set, the replica set will be created from the snapshot.
	InitialSnapshotID *string `json:"initialSnapshotId,omitempty" tf:"initial_snapshot_id,omitempty"`

	// [string] The location of your replica set. Updates to the value of the field force the replica set to be re-created.
	// The replica set location
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// (Computed) A weekly 4 hour-long window, during which maintenance might occur.
	// A weekly 4 hour-long window, during which maintenance might occur.
	MaintenanceWindow []MaintenanceWindowInitParameters `json:"maintenanceWindow,omitempty" tf:"maintenance_window,omitempty"`

	// [string] Specifies How and If data is persisted, possible values are:
	// Specifies How and If data is persisted.
	PersistenceMode *string `json:"persistenceMode,omitempty" tf:"persistence_mode,omitempty"`

	// [int] The total number of replicas in the replica set (one active and n-1 passive). In case of a standalone instance, the value is 1. In all other cases, the value is > 1. The replicas will not be available as read replicas, they are only standby for a failure of the active instance.
	// The total number of replicas in the replica set (one active and n-1 passive). In case of a standalone instance, the value is 1. In all other cases, the value is > 1. The replicas will not be available as read replicas, they are only standby for a failure of the active instance.
	Replicas *float64 `json:"replicas,omitempty" tf:"replicas,omitempty"`

	// [object] The resources of the individual replicas.
	// The resources of the individual replicas.
	Resources []ResourcesInitParameters `json:"resources,omitempty" tf:"resources,omitempty"`

	// [string] The InMemoryDB version of your replica set.
	// The InMemoryDB version of your replica set.
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type InMemoryDBReplicasetObservation struct {

	// [object] The network connection for your replica set. Only one connection is allowed. Updates to the value of the fields force the replica set to be re-created.
	// The network connection for your replica set. Only one connection is allowed.
	Connections []ConnectionsObservation `json:"connections,omitempty" tf:"connections,omitempty"`

	// [object] Credentials for the InMemoryDB replicaset, only one type of password can be used since they are mutually exclusive. These values are used to create the initial InMemoryDB user, updating any of these will force recreation of the replica set resource.
	// Credentials for the InMemoryDB replicaset.
	Credentials []CredentialsObservation `json:"credentials,omitempty" tf:"credentials,omitempty"`

	// (Computed)[string] The DNS name pointing to your replica set. Will be used to connect to the active/standalone instance.
	// The DNS name pointing to your replica set. Will be used to connect to the active/standalone instance.
	DNSName *string `json:"dnsName,omitempty" tf:"dns_name,omitempty"`

	// [string] The human-readable name of your replica set.
	// The human readable name of your replica set.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// [string] The eviction policy for the replica set, possible values are:
	// The eviction policy for the replica set.
	EvictionPolicy *string `json:"evictionPolicy,omitempty" tf:"eviction_policy,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// [string] The ID of a snapshot to restore the replica set from. If set, the replica set will be created from the snapshot.
	// The ID of a snapshot to restore the replica set from. If set, the replica set will be created from the snapshot.
	InitialSnapshotID *string `json:"initialSnapshotId,omitempty" tf:"initial_snapshot_id,omitempty"`

	// [string] The location of your replica set. Updates to the value of the field force the replica set to be re-created.
	// The replica set location
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// (Computed) A weekly 4 hour-long window, during which maintenance might occur.
	// A weekly 4 hour-long window, during which maintenance might occur.
	MaintenanceWindow []MaintenanceWindowObservation `json:"maintenanceWindow,omitempty" tf:"maintenance_window,omitempty"`

	// [string] Specifies How and If data is persisted, possible values are:
	// Specifies How and If data is persisted.
	PersistenceMode *string `json:"persistenceMode,omitempty" tf:"persistence_mode,omitempty"`

	// [int] The total number of replicas in the replica set (one active and n-1 passive). In case of a standalone instance, the value is 1. In all other cases, the value is > 1. The replicas will not be available as read replicas, they are only standby for a failure of the active instance.
	// The total number of replicas in the replica set (one active and n-1 passive). In case of a standalone instance, the value is 1. In all other cases, the value is > 1. The replicas will not be available as read replicas, they are only standby for a failure of the active instance.
	Replicas *float64 `json:"replicas,omitempty" tf:"replicas,omitempty"`

	// [object] The resources of the individual replicas.
	// The resources of the individual replicas.
	Resources []ResourcesObservation `json:"resources,omitempty" tf:"resources,omitempty"`

	// [string] The InMemoryDB version of your replica set.
	// The InMemoryDB version of your replica set.
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type InMemoryDBReplicasetParameters struct {

	// [object] The network connection for your replica set. Only one connection is allowed. Updates to the value of the fields force the replica set to be re-created.
	// The network connection for your replica set. Only one connection is allowed.
	// +kubebuilder:validation:Optional
	Connections []ConnectionsParameters `json:"connections,omitempty" tf:"connections,omitempty"`

	// [object] Credentials for the InMemoryDB replicaset, only one type of password can be used since they are mutually exclusive. These values are used to create the initial InMemoryDB user, updating any of these will force recreation of the replica set resource.
	// Credentials for the InMemoryDB replicaset.
	// +kubebuilder:validation:Optional
	Credentials []CredentialsParameters `json:"credentials,omitempty" tf:"credentials,omitempty"`

	// [string] The human-readable name of your replica set.
	// The human readable name of your replica set.
	// +kubebuilder:validation:Optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// [string] The eviction policy for the replica set, possible values are:
	// The eviction policy for the replica set.
	// +kubebuilder:validation:Optional
	EvictionPolicy *string `json:"evictionPolicy,omitempty" tf:"eviction_policy,omitempty"`

	// [string] The ID of a snapshot to restore the replica set from. If set, the replica set will be created from the snapshot.
	// The ID of a snapshot to restore the replica set from. If set, the replica set will be created from the snapshot.
	// +kubebuilder:validation:Optional
	InitialSnapshotID *string `json:"initialSnapshotId,omitempty" tf:"initial_snapshot_id,omitempty"`

	// [string] The location of your replica set. Updates to the value of the field force the replica set to be re-created.
	// The replica set location
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// (Computed) A weekly 4 hour-long window, during which maintenance might occur.
	// A weekly 4 hour-long window, during which maintenance might occur.
	// +kubebuilder:validation:Optional
	MaintenanceWindow []MaintenanceWindowParameters `json:"maintenanceWindow,omitempty" tf:"maintenance_window,omitempty"`

	// [string] Specifies How and If data is persisted, possible values are:
	// Specifies How and If data is persisted.
	// +kubebuilder:validation:Optional
	PersistenceMode *string `json:"persistenceMode,omitempty" tf:"persistence_mode,omitempty"`

	// [int] The total number of replicas in the replica set (one active and n-1 passive). In case of a standalone instance, the value is 1. In all other cases, the value is > 1. The replicas will not be available as read replicas, they are only standby for a failure of the active instance.
	// The total number of replicas in the replica set (one active and n-1 passive). In case of a standalone instance, the value is 1. In all other cases, the value is > 1. The replicas will not be available as read replicas, they are only standby for a failure of the active instance.
	// +kubebuilder:validation:Optional
	Replicas *float64 `json:"replicas,omitempty" tf:"replicas,omitempty"`

	// [object] The resources of the individual replicas.
	// The resources of the individual replicas.
	// +kubebuilder:validation:Optional
	Resources []ResourcesParameters `json:"resources,omitempty" tf:"resources,omitempty"`

	// [string] The InMemoryDB version of your replica set.
	// The InMemoryDB version of your replica set.
	// +kubebuilder:validation:Optional
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type MaintenanceWindowInitParameters struct {

	// [string] The name of the week day.
	// The name of the week day.
	DayOfTheWeek *string `json:"dayOfTheWeek,omitempty" tf:"day_of_the_week,omitempty"`

	// [string] Start of the maintenance window in UTC time.
	// Start of the maintenance window in UTC time.
	Time *string `json:"time,omitempty" tf:"time,omitempty"`
}

type MaintenanceWindowObservation struct {

	// [string] The name of the week day.
	// The name of the week day.
	DayOfTheWeek *string `json:"dayOfTheWeek,omitempty" tf:"day_of_the_week,omitempty"`

	// [string] Start of the maintenance window in UTC time.
	// Start of the maintenance window in UTC time.
	Time *string `json:"time,omitempty" tf:"time,omitempty"`
}

type MaintenanceWindowParameters struct {

	// [string] The name of the week day.
	// The name of the week day.
	// +kubebuilder:validation:Optional
	DayOfTheWeek *string `json:"dayOfTheWeek" tf:"day_of_the_week,omitempty"`

	// [string] Start of the maintenance window in UTC time.
	// Start of the maintenance window in UTC time.
	// +kubebuilder:validation:Optional
	Time *string `json:"time" tf:"time,omitempty"`
}

type ResourcesInitParameters struct {

	// [int] The number of CPU cores per instance.
	// The number of CPU cores per instance.
	Cores *float64 `json:"cores,omitempty" tf:"cores,omitempty"`

	// [int] The amount of memory per instance in gigabytes (GB).
	// The amount of memory per instance in gigabytes (GB).
	RAM *float64 `json:"ram,omitempty" tf:"ram,omitempty"`
}

type ResourcesObservation struct {

	// [int] The number of CPU cores per instance.
	// The number of CPU cores per instance.
	Cores *float64 `json:"cores,omitempty" tf:"cores,omitempty"`

	// [int] The amount of memory per instance in gigabytes (GB).
	// The amount of memory per instance in gigabytes (GB).
	RAM *float64 `json:"ram,omitempty" tf:"ram,omitempty"`

	// (Computed)[int] The size of the storage in GB. The size is derived from the amount of RAM and the persistence mode and is not configurable.
	// The size of the storage in GB. The size is derived from the amount of RAM and the persistence mode and is not configurable.
	Storage *float64 `json:"storage,omitempty" tf:"storage,omitempty"`
}

type ResourcesParameters struct {

	// [int] The number of CPU cores per instance.
	// The number of CPU cores per instance.
	// +kubebuilder:validation:Optional
	Cores *float64 `json:"cores" tf:"cores,omitempty"`

	// [int] The amount of memory per instance in gigabytes (GB).
	// The amount of memory per instance in gigabytes (GB).
	// +kubebuilder:validation:Optional
	RAM *float64 `json:"ram" tf:"ram,omitempty"`
}

// InMemoryDBReplicasetSpec defines the desired state of InMemoryDBReplicaset
type InMemoryDBReplicasetSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     InMemoryDBReplicasetParameters `json:"forProvider"`
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
	InitProvider InMemoryDBReplicasetInitParameters `json:"initProvider,omitempty"`
}

// InMemoryDBReplicasetStatus defines the observed state of InMemoryDBReplicaset.
type InMemoryDBReplicasetStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        InMemoryDBReplicasetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// InMemoryDBReplicaset is the Schema for the InMemoryDBReplicasets API. Creates and manages DBaaS InMemoryDB Replica Set objects.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,ionos}
type InMemoryDBReplicaset struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.connections) || (has(self.initProvider) && has(self.initProvider.connections))",message="spec.forProvider.connections is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.credentials) || (has(self.initProvider) && has(self.initProvider.credentials))",message="spec.forProvider.credentials is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.displayName) || (has(self.initProvider) && has(self.initProvider.displayName))",message="spec.forProvider.displayName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.evictionPolicy) || (has(self.initProvider) && has(self.initProvider.evictionPolicy))",message="spec.forProvider.evictionPolicy is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.location) || (has(self.initProvider) && has(self.initProvider.location))",message="spec.forProvider.location is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.persistenceMode) || (has(self.initProvider) && has(self.initProvider.persistenceMode))",message="spec.forProvider.persistenceMode is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.replicas) || (has(self.initProvider) && has(self.initProvider.replicas))",message="spec.forProvider.replicas is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.resources) || (has(self.initProvider) && has(self.initProvider.resources))",message="spec.forProvider.resources is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.version) || (has(self.initProvider) && has(self.initProvider.version))",message="spec.forProvider.version is a required parameter"
	Spec   InMemoryDBReplicasetSpec   `json:"spec"`
	Status InMemoryDBReplicasetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// InMemoryDBReplicasetList contains a list of InMemoryDBReplicasets
type InMemoryDBReplicasetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []InMemoryDBReplicaset `json:"items"`
}

// Repository type metadata.
var (
	InMemoryDBReplicaset_Kind             = "InMemoryDBReplicaset"
	InMemoryDBReplicaset_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: InMemoryDBReplicaset_Kind}.String()
	InMemoryDBReplicaset_KindAPIVersion   = InMemoryDBReplicaset_Kind + "." + CRDGroupVersion.String()
	InMemoryDBReplicaset_GroupVersionKind = CRDGroupVersion.WithKind(InMemoryDBReplicaset_Kind)
)

func init() {
	SchemeBuilder.Register(&InMemoryDBReplicaset{}, &InMemoryDBReplicasetList{})
}
