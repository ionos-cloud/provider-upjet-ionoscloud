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

type FeaturesInitParameters struct {

	// [bool] Enables or disables the Vulnerability Scanning feature for the Container Registry. To disable this feature, set the attribute to false when creating the CR resource.
	// Enables vulnerability scanning for images in the container registry. Note: this feature can incur additional charges
	VulnerabilityScanning *bool `json:"vulnerabilityScanning,omitempty" tf:"vulnerability_scanning,omitempty"`
}

type FeaturesObservation struct {

	// [bool] Enables or disables the Vulnerability Scanning feature for the Container Registry. To disable this feature, set the attribute to false when creating the CR resource.
	// Enables vulnerability scanning for images in the container registry. Note: this feature can incur additional charges
	VulnerabilityScanning *bool `json:"vulnerabilityScanning,omitempty" tf:"vulnerability_scanning,omitempty"`
}

type FeaturesParameters struct {

	// [bool] Enables or disables the Vulnerability Scanning feature for the Container Registry. To disable this feature, set the attribute to false when creating the CR resource.
	// Enables vulnerability scanning for images in the container registry. Note: this feature can incur additional charges
	// +kubebuilder:validation:Optional
	VulnerabilityScanning *bool `json:"vulnerabilityScanning,omitempty" tf:"vulnerability_scanning,omitempty"`
}

type GarbageCollectionScheduleInitParameters struct {

	// [list] Elements of list must have one of the values: Saturday, Sunday, Monday, Tuesday,  Wednesday,  Thursday,  Friday
	Days []*string `json:"days,omitempty" tf:"days,omitempty"`

	// [string]
	// UTC time of day e.g. 01:00:00 - as defined by partial-time - RFC3339
	Time *string `json:"time,omitempty" tf:"time,omitempty"`
}

type GarbageCollectionScheduleObservation struct {

	// [list] Elements of list must have one of the values: Saturday, Sunday, Monday, Tuesday,  Wednesday,  Thursday,  Friday
	Days []*string `json:"days,omitempty" tf:"days,omitempty"`

	// [string]
	// UTC time of day e.g. 01:00:00 - as defined by partial-time - RFC3339
	Time *string `json:"time,omitempty" tf:"time,omitempty"`
}

type GarbageCollectionScheduleParameters struct {

	// [list] Elements of list must have one of the values: Saturday, Sunday, Monday, Tuesday,  Wednesday,  Thursday,  Friday
	// +kubebuilder:validation:Optional
	Days []*string `json:"days" tf:"days,omitempty"`

	// [string]
	// UTC time of day e.g. 01:00:00 - as defined by partial-time - RFC3339
	// +kubebuilder:validation:Optional
	Time *string `json:"time" tf:"time,omitempty"`
}

type RegistryInitParameters struct {

	// [list] The subnet CIDRs that are allowed to connect to the registry.  Specify "a.b.c.d/32" for an individual IP address. Note: If this list is empty or not set, there are no restrictions.
	// The subnet CIDRs that are allowed to connect to the registry. Specify 'a.b.c.d/32' for an individual IP address. __Note__: If this list is empty or not set, there are no restrictions.
	APISubnetAllowList []*string `json:"apiSubnetAllowList,omitempty" tf:"api_subnet_allow_list,omitempty"`

	// [Map]
	Features *FeaturesInitParameters `json:"features,omitempty" tf:"features,omitempty"`

	// [Map]
	GarbageCollectionSchedule *GarbageCollectionScheduleInitParameters `json:"garbageCollectionSchedule,omitempty" tf:"garbage_collection_schedule,omitempty"`

	// [string] Immutable, update forces re-creation of the resource.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The name of the container registry. Immutable, update forces re-creation of the resource.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type RegistryObservation struct {

	// [list] The subnet CIDRs that are allowed to connect to the registry.  Specify "a.b.c.d/32" for an individual IP address. Note: If this list is empty or not set, there are no restrictions.
	// The subnet CIDRs that are allowed to connect to the registry. Specify 'a.b.c.d/32' for an individual IP address. __Note__: If this list is empty or not set, there are no restrictions.
	APISubnetAllowList []*string `json:"apiSubnetAllowList,omitempty" tf:"api_subnet_allow_list,omitempty"`

	// [Map]
	Features *FeaturesObservation `json:"features,omitempty" tf:"features,omitempty"`

	// [Map]
	GarbageCollectionSchedule *GarbageCollectionScheduleObservation `json:"garbageCollectionSchedule,omitempty" tf:"garbage_collection_schedule,omitempty"`

	// The name of the container registry. Immutable, update forces re-creation of the resource.
	Hostname *string `json:"hostname,omitempty" tf:"hostname,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// [string] Immutable, update forces re-creation of the resource.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The name of the container registry. Immutable, update forces re-creation of the resource.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	StorageUsage []StorageUsageObservation `json:"storageUsage,omitempty" tf:"storage_usage,omitempty"`
}

type RegistryParameters struct {

	// [list] The subnet CIDRs that are allowed to connect to the registry.  Specify "a.b.c.d/32" for an individual IP address. Note: If this list is empty or not set, there are no restrictions.
	// The subnet CIDRs that are allowed to connect to the registry. Specify 'a.b.c.d/32' for an individual IP address. __Note__: If this list is empty or not set, there are no restrictions.
	// +kubebuilder:validation:Optional
	APISubnetAllowList []*string `json:"apiSubnetAllowList,omitempty" tf:"api_subnet_allow_list,omitempty"`

	// [Map]
	// +kubebuilder:validation:Optional
	Features *FeaturesParameters `json:"features,omitempty" tf:"features,omitempty"`

	// [Map]
	// +kubebuilder:validation:Optional
	GarbageCollectionSchedule *GarbageCollectionScheduleParameters `json:"garbageCollectionSchedule,omitempty" tf:"garbage_collection_schedule,omitempty"`

	// [string] Immutable, update forces re-creation of the resource.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The name of the container registry. Immutable, update forces re-creation of the resource.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type StorageUsageInitParameters struct {
}

type StorageUsageObservation struct {
	Bytes *float64 `json:"bytes,omitempty" tf:"bytes,omitempty"`

	UpdatedAt *string `json:"updatedAt,omitempty" tf:"updated_at,omitempty"`
}

type StorageUsageParameters struct {
}

// RegistrySpec defines the desired state of Registry
type RegistrySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RegistryParameters `json:"forProvider"`
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
	InitProvider RegistryInitParameters `json:"initProvider,omitempty"`
}

// RegistryStatus defines the observed state of Registry.
type RegistryStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RegistryObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Registry is the Schema for the Registrys API. Creates and manages IonosCloud Container Registry.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,ionos}
type Registry struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.location) || (has(self.initProvider) && has(self.initProvider.location))",message="spec.forProvider.location is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   RegistrySpec   `json:"spec"`
	Status RegistryStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RegistryList contains a list of Registrys
type RegistryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Registry `json:"items"`
}

// Repository type metadata.
var (
	Registry_Kind             = "Registry"
	Registry_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Registry_Kind}.String()
	Registry_KindAPIVersion   = Registry_Kind + "." + CRDGroupVersion.String()
	Registry_GroupVersionKind = CRDGroupVersion.WithKind(Registry_Kind)
)

func init() {
	SchemeBuilder.Register(&Registry{}, &RegistryList{})
}