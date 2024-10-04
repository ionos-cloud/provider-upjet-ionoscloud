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

type ShareInitParameters struct {

	// [Boolean] The group has permission to edit privileges on this resource.
	EditPrivilege *bool `json:"editPrivilege,omitempty" tf:"edit_privilege,omitempty"`

	// [string] The ID of the specific group containing the resource to update.
	// +crossplane:generate:reference:type=github.com/ionos-cloud/provider-upjet-ionoscloud/apis/compute/v1alpha1.Group
	GroupID *string `json:"groupId,omitempty" tf:"group_id,omitempty"`

	// Reference to a Group in compute to populate groupId.
	// +kubebuilder:validation:Optional
	GroupIDRef *v1.Reference `json:"groupIdRef,omitempty" tf:"-"`

	// Selector for a Group in compute to populate groupId.
	// +kubebuilder:validation:Optional
	GroupIDSelector *v1.Selector `json:"groupIdSelector,omitempty" tf:"-"`

	// [string] The ID of the specific resource to update.
	// +crossplane:generate:reference:type=github.com/ionos-cloud/provider-upjet-ionoscloud/apis/compute/v1alpha1.Datacenter
	ResourceID *string `json:"resourceId,omitempty" tf:"resource_id,omitempty"`

	// Reference to a Datacenter in compute to populate resourceId.
	// +kubebuilder:validation:Optional
	ResourceIDRef *v1.Reference `json:"resourceIdRef,omitempty" tf:"-"`

	// Selector for a Datacenter in compute to populate resourceId.
	// +kubebuilder:validation:Optional
	ResourceIDSelector *v1.Selector `json:"resourceIdSelector,omitempty" tf:"-"`

	// [Boolean] The group has permission to share this resource.
	SharePrivilege *bool `json:"sharePrivilege,omitempty" tf:"share_privilege,omitempty"`
}

type ShareObservation struct {

	// [Boolean] The group has permission to edit privileges on this resource.
	EditPrivilege *bool `json:"editPrivilege,omitempty" tf:"edit_privilege,omitempty"`

	// [string] The ID of the specific group containing the resource to update.
	GroupID *string `json:"groupId,omitempty" tf:"group_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// [string] The ID of the specific resource to update.
	ResourceID *string `json:"resourceId,omitempty" tf:"resource_id,omitempty"`

	// [Boolean] The group has permission to share this resource.
	SharePrivilege *bool `json:"sharePrivilege,omitempty" tf:"share_privilege,omitempty"`
}

type ShareParameters struct {

	// [Boolean] The group has permission to edit privileges on this resource.
	// +kubebuilder:validation:Optional
	EditPrivilege *bool `json:"editPrivilege,omitempty" tf:"edit_privilege,omitempty"`

	// [string] The ID of the specific group containing the resource to update.
	// +crossplane:generate:reference:type=github.com/ionos-cloud/provider-upjet-ionoscloud/apis/compute/v1alpha1.Group
	// +kubebuilder:validation:Optional
	GroupID *string `json:"groupId,omitempty" tf:"group_id,omitempty"`

	// Reference to a Group in compute to populate groupId.
	// +kubebuilder:validation:Optional
	GroupIDRef *v1.Reference `json:"groupIdRef,omitempty" tf:"-"`

	// Selector for a Group in compute to populate groupId.
	// +kubebuilder:validation:Optional
	GroupIDSelector *v1.Selector `json:"groupIdSelector,omitempty" tf:"-"`

	// [string] The ID of the specific resource to update.
	// +crossplane:generate:reference:type=github.com/ionos-cloud/provider-upjet-ionoscloud/apis/compute/v1alpha1.Datacenter
	// +kubebuilder:validation:Optional
	ResourceID *string `json:"resourceId,omitempty" tf:"resource_id,omitempty"`

	// Reference to a Datacenter in compute to populate resourceId.
	// +kubebuilder:validation:Optional
	ResourceIDRef *v1.Reference `json:"resourceIdRef,omitempty" tf:"-"`

	// Selector for a Datacenter in compute to populate resourceId.
	// +kubebuilder:validation:Optional
	ResourceIDSelector *v1.Selector `json:"resourceIdSelector,omitempty" tf:"-"`

	// [Boolean] The group has permission to share this resource.
	// +kubebuilder:validation:Optional
	SharePrivilege *bool `json:"sharePrivilege,omitempty" tf:"share_privilege,omitempty"`
}

// ShareSpec defines the desired state of Share
type ShareSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ShareParameters `json:"forProvider"`
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
	InitProvider ShareInitParameters `json:"initProvider,omitempty"`
}

// ShareStatus defines the observed state of Share.
type ShareStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ShareObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Share is the Schema for the Shares API. Creates and manages share objects.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,ionos}
type Share struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ShareSpec   `json:"spec"`
	Status            ShareStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ShareList contains a list of Shares
type ShareList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Share `json:"items"`
}

// Repository type metadata.
var (
	Share_Kind             = "Share"
	Share_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Share_Kind}.String()
	Share_KindAPIVersion   = Share_Kind + "." + CRDGroupVersion.String()
	Share_GroupVersionKind = CRDGroupVersion.WithKind(Share_Kind)
)

func init() {
	SchemeBuilder.Register(&Share{}, &ShareList{})
}
