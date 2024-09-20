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

type UnitInitParameters struct {

	// [string] The email address assigned to the backup unit
	// The e-mail address you want assigned to the backup unit.
	Email *string `json:"email,omitempty" tf:"email,omitempty"`

	// [string] The name of the Backup Unit. This argument is immutable.
	// Alphanumeric name you want assigned to the backup unit.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// [string] The desired password for the Backup Unit
	// The password you want assigned to the backup unit.
	PasswordSecretRef v1.SecretKeySelector `json:"passwordSecretRef" tf:"-"`
}

type UnitObservation struct {

	// [string] The email address assigned to the backup unit
	// The e-mail address you want assigned to the backup unit.
	Email *string `json:"email,omitempty" tf:"email,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (Computed) The login associated with the backup unit. Derived from the contract number
	// The login associated with the backup unit. Derived from the contract number
	Login *string `json:"login,omitempty" tf:"login,omitempty"`

	// [string] The name of the Backup Unit. This argument is immutable.
	// Alphanumeric name you want assigned to the backup unit.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type UnitParameters struct {

	// [string] The email address assigned to the backup unit
	// The e-mail address you want assigned to the backup unit.
	// +kubebuilder:validation:Optional
	Email *string `json:"email,omitempty" tf:"email,omitempty"`

	// [string] The name of the Backup Unit. This argument is immutable.
	// Alphanumeric name you want assigned to the backup unit.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// [string] The desired password for the Backup Unit
	// The password you want assigned to the backup unit.
	// +kubebuilder:validation:Optional
	PasswordSecretRef v1.SecretKeySelector `json:"passwordSecretRef" tf:"-"`
}

// UnitSpec defines the desired state of Unit
type UnitSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     UnitParameters `json:"forProvider"`
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
	InitProvider UnitInitParameters `json:"initProvider,omitempty"`
}

// UnitStatus defines the observed state of Unit.
type UnitStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        UnitObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Unit is the Schema for the Units API. Creates and manages IonosCloud Backup Units.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,ionos}
type Unit struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.email) || (has(self.initProvider) && has(self.initProvider.email))",message="spec.forProvider.email is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.passwordSecretRef)",message="spec.forProvider.passwordSecretRef is a required parameter"
	Spec   UnitSpec   `json:"spec"`
	Status UnitStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// UnitList contains a list of Units
type UnitList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Unit `json:"items"`
}

// Repository type metadata.
var (
	Unit_Kind             = "Unit"
	Unit_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Unit_Kind}.String()
	Unit_KindAPIVersion   = Unit_Kind + "." + CRDGroupVersion.String()
	Unit_GroupVersionKind = CRDGroupVersion.WithKind(Unit_Kind)
)

func init() {
	SchemeBuilder.Register(&Unit{}, &UnitList{})
}
