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

type SnapshotInitParameters struct {

	// (Computed)[string] Is capable of CPU hot plug (no reboot required). Can only be updated.
	CPUHotPlug *bool `json:"cpuHotPlug,omitempty" tf:"cpu_hot_plug,omitempty"`

	// [string] The ID of the Virtual Data Center.
	// +crossplane:generate:reference:type=github.com/ionos-cloud/provider-upjet-ionoscloud/apis/compute/v1alpha1.Datacenter
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	DatacenterID *string `json:"datacenterId,omitempty" tf:"datacenter_id,omitempty"`

	// Reference to a Datacenter in compute to populate datacenterId.
	// +kubebuilder:validation:Optional
	DatacenterIDRef *v1.Reference `json:"datacenterIdRef,omitempty" tf:"-"`

	// Selector for a Datacenter in compute to populate datacenterId.
	// +kubebuilder:validation:Optional
	DatacenterIDSelector *v1.Selector `json:"datacenterIdSelector,omitempty" tf:"-"`

	// (Computed)[string] Human readable description
	// Human readable description
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (Computed)[string] Is capable of Virt-IO drive hot plug (no reboot required). Can only be updated.
	DiscVirtioHotPlug *bool `json:"discVirtioHotPlug,omitempty" tf:"disc_virtio_hot_plug,omitempty"`

	// (Computed)[string] Is capable of Virt-IO drive hot unplug (no reboot required). This works only for non-Windows virtual Machines. Can only be updated.
	DiscVirtioHotUnplug *bool `json:"discVirtioHotUnplug,omitempty" tf:"disc_virtio_hot_unplug,omitempty"`

	// (Computed)[string] OS type of this Snapshot
	// OS type of this Snapshot
	LicenceType *string `json:"licenceType,omitempty" tf:"licence_type,omitempty"`

	// [string] The name of the snapshot.
	// A name of that resource
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Computed)[string] Is capable of nic hot plug (no reboot required). Can only be updated.
	NicHotPlug *bool `json:"nicHotPlug,omitempty" tf:"nic_hot_plug,omitempty"`

	// (Computed)[string] Is capable of nic hot unplug (no reboot required). Can only be updated.
	NicHotUnplug *bool `json:"nicHotUnplug,omitempty" tf:"nic_hot_unplug,omitempty"`

	// (Computed)[string] Is capable of memory hot plug (no reboot required). Can only be updated.
	RAMHotPlug *bool `json:"ramHotPlug,omitempty" tf:"ram_hot_plug,omitempty"`

	// Boolean value representing if the snapshot requires extra protection e.g. two factor protection
	// Boolean value representing if the snapshot requires extra protection e.g. two factor protection
	SecAuthProtection *bool `json:"secAuthProtection,omitempty" tf:"sec_auth_protection,omitempty"`

	// [string] The ID of the specific volume to take the snapshot from.
	// +crossplane:generate:reference:type=github.com/ionos-cloud/provider-upjet-ionoscloud/apis/compute/v1alpha1.Volume
	VolumeID *string `json:"volumeId,omitempty" tf:"volume_id,omitempty"`

	// Reference to a Volume in compute to populate volumeId.
	// +kubebuilder:validation:Optional
	VolumeIDRef *v1.Reference `json:"volumeIdRef,omitempty" tf:"-"`

	// Selector for a Volume in compute to populate volumeId.
	// +kubebuilder:validation:Optional
	VolumeIDSelector *v1.Selector `json:"volumeIdSelector,omitempty" tf:"-"`
}

type SnapshotObservation struct {

	// (Computed)[string] Is capable of CPU hot plug (no reboot required). Can only be updated.
	CPUHotPlug *bool `json:"cpuHotPlug,omitempty" tf:"cpu_hot_plug,omitempty"`

	// Is capable of CPU hot unplug (no reboot required)
	CPUHotUnplug *bool `json:"cpuHotUnplug,omitempty" tf:"cpu_hot_unplug,omitempty"`

	// [string] The ID of the Virtual Data Center.
	DatacenterID *string `json:"datacenterId,omitempty" tf:"datacenter_id,omitempty"`

	// (Computed)[string] Human readable description
	// Human readable description
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Is capable of SCSI drive hot plug (no reboot required)
	DiscScsiHotPlug *bool `json:"discScsiHotPlug,omitempty" tf:"disc_scsi_hot_plug,omitempty"`

	// Is capable of SCSI drive hot unplug (no reboot required). This works only for non-Windows virtual Machines.
	DiscScsiHotUnplug *bool `json:"discScsiHotUnplug,omitempty" tf:"disc_scsi_hot_unplug,omitempty"`

	// (Computed)[string] Is capable of Virt-IO drive hot plug (no reboot required). Can only be updated.
	DiscVirtioHotPlug *bool `json:"discVirtioHotPlug,omitempty" tf:"disc_virtio_hot_plug,omitempty"`

	// (Computed)[string] Is capable of Virt-IO drive hot unplug (no reboot required). This works only for non-Windows virtual Machines. Can only be updated.
	DiscVirtioHotUnplug *bool `json:"discVirtioHotUnplug,omitempty" tf:"disc_virtio_hot_unplug,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (Computed)[string] OS type of this Snapshot
	// OS type of this Snapshot
	LicenceType *string `json:"licenceType,omitempty" tf:"licence_type,omitempty"`

	// Location of that image/snapshot
	// Location of that image/snapshot
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// [string] The name of the snapshot.
	// A name of that resource
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Computed)[string] Is capable of nic hot plug (no reboot required). Can only be updated.
	NicHotPlug *bool `json:"nicHotPlug,omitempty" tf:"nic_hot_plug,omitempty"`

	// (Computed)[string] Is capable of nic hot unplug (no reboot required). Can only be updated.
	NicHotUnplug *bool `json:"nicHotUnplug,omitempty" tf:"nic_hot_unplug,omitempty"`

	// (Computed)[string] Is capable of memory hot plug (no reboot required). Can only be updated.
	RAMHotPlug *bool `json:"ramHotPlug,omitempty" tf:"ram_hot_plug,omitempty"`

	// Is capable of memory hot unplug (no reboot required)
	RAMHotUnplug *bool `json:"ramHotUnplug,omitempty" tf:"ram_hot_unplug,omitempty"`

	// Boolean value representing if the snapshot requires extra protection e.g. two factor protection
	// Boolean value representing if the snapshot requires extra protection e.g. two factor protection
	SecAuthProtection *bool `json:"secAuthProtection,omitempty" tf:"sec_auth_protection,omitempty"`

	// The size of the image in GB
	// The size of the image in GB
	Size *float64 `json:"size,omitempty" tf:"size,omitempty"`

	// [string] The ID of the specific volume to take the snapshot from.
	VolumeID *string `json:"volumeId,omitempty" tf:"volume_id,omitempty"`
}

type SnapshotParameters struct {

	// (Computed)[string] Is capable of CPU hot plug (no reboot required). Can only be updated.
	// +kubebuilder:validation:Optional
	CPUHotPlug *bool `json:"cpuHotPlug,omitempty" tf:"cpu_hot_plug,omitempty"`

	// [string] The ID of the Virtual Data Center.
	// +crossplane:generate:reference:type=github.com/ionos-cloud/provider-upjet-ionoscloud/apis/compute/v1alpha1.Datacenter
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	DatacenterID *string `json:"datacenterId,omitempty" tf:"datacenter_id,omitempty"`

	// Reference to a Datacenter in compute to populate datacenterId.
	// +kubebuilder:validation:Optional
	DatacenterIDRef *v1.Reference `json:"datacenterIdRef,omitempty" tf:"-"`

	// Selector for a Datacenter in compute to populate datacenterId.
	// +kubebuilder:validation:Optional
	DatacenterIDSelector *v1.Selector `json:"datacenterIdSelector,omitempty" tf:"-"`

	// (Computed)[string] Human readable description
	// Human readable description
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (Computed)[string] Is capable of Virt-IO drive hot plug (no reboot required). Can only be updated.
	// +kubebuilder:validation:Optional
	DiscVirtioHotPlug *bool `json:"discVirtioHotPlug,omitempty" tf:"disc_virtio_hot_plug,omitempty"`

	// (Computed)[string] Is capable of Virt-IO drive hot unplug (no reboot required). This works only for non-Windows virtual Machines. Can only be updated.
	// +kubebuilder:validation:Optional
	DiscVirtioHotUnplug *bool `json:"discVirtioHotUnplug,omitempty" tf:"disc_virtio_hot_unplug,omitempty"`

	// (Computed)[string] OS type of this Snapshot
	// OS type of this Snapshot
	// +kubebuilder:validation:Optional
	LicenceType *string `json:"licenceType,omitempty" tf:"licence_type,omitempty"`

	// [string] The name of the snapshot.
	// A name of that resource
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Computed)[string] Is capable of nic hot plug (no reboot required). Can only be updated.
	// +kubebuilder:validation:Optional
	NicHotPlug *bool `json:"nicHotPlug,omitempty" tf:"nic_hot_plug,omitempty"`

	// (Computed)[string] Is capable of nic hot unplug (no reboot required). Can only be updated.
	// +kubebuilder:validation:Optional
	NicHotUnplug *bool `json:"nicHotUnplug,omitempty" tf:"nic_hot_unplug,omitempty"`

	// (Computed)[string] Is capable of memory hot plug (no reboot required). Can only be updated.
	// +kubebuilder:validation:Optional
	RAMHotPlug *bool `json:"ramHotPlug,omitempty" tf:"ram_hot_plug,omitempty"`

	// Boolean value representing if the snapshot requires extra protection e.g. two factor protection
	// Boolean value representing if the snapshot requires extra protection e.g. two factor protection
	// +kubebuilder:validation:Optional
	SecAuthProtection *bool `json:"secAuthProtection,omitempty" tf:"sec_auth_protection,omitempty"`

	// [string] The ID of the specific volume to take the snapshot from.
	// +crossplane:generate:reference:type=github.com/ionos-cloud/provider-upjet-ionoscloud/apis/compute/v1alpha1.Volume
	// +kubebuilder:validation:Optional
	VolumeID *string `json:"volumeId,omitempty" tf:"volume_id,omitempty"`

	// Reference to a Volume in compute to populate volumeId.
	// +kubebuilder:validation:Optional
	VolumeIDRef *v1.Reference `json:"volumeIdRef,omitempty" tf:"-"`

	// Selector for a Volume in compute to populate volumeId.
	// +kubebuilder:validation:Optional
	VolumeIDSelector *v1.Selector `json:"volumeIdSelector,omitempty" tf:"-"`
}

// SnapshotSpec defines the desired state of Snapshot
type SnapshotSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SnapshotParameters `json:"forProvider"`
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
	InitProvider SnapshotInitParameters `json:"initProvider,omitempty"`
}

// SnapshotStatus defines the observed state of Snapshot.
type SnapshotStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SnapshotObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Snapshot is the Schema for the Snapshots API. Creates and manages snapshot objects.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,ionos}
type Snapshot struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   SnapshotSpec   `json:"spec"`
	Status SnapshotStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SnapshotList contains a list of Snapshots
type SnapshotList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Snapshot `json:"items"`
}

// Repository type metadata.
var (
	Snapshot_Kind             = "Snapshot"
	Snapshot_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Snapshot_Kind}.String()
	Snapshot_KindAPIVersion   = Snapshot_Kind + "." + CRDGroupVersion.String()
	Snapshot_GroupVersionKind = CRDGroupVersion.WithKind(Snapshot_Kind)
)

func init() {
	SchemeBuilder.Register(&Snapshot{}, &SnapshotList{})
}
