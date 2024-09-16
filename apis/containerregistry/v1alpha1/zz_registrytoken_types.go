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

type CredentialsInitParameters struct {
}

type CredentialsObservation struct {
	Password *string `json:"password,omitempty" tf:"password,omitempty"`

	// [string] The name of the container registry token. Immutable, update forces re-creation of the resource.
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

type CredentialsParameters struct {
}

type RegistryTokenInitParameters struct {
	ExpiryDate *string `json:"expiryDate,omitempty" tf:"expiry_date,omitempty"`

	// [string] The name of the container registry token. Immutable, update forces re-creation of the resource.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// +crossplane:generate:reference:type=github.com/ionos-cloud/provider-upjet-ionoscloud/apis/containerregistry/v1alpha1.Registry
	RegistryID *string `json:"registryId,omitempty" tf:"registry_id,omitempty"`

	// Reference to a Registry in containerregistry to populate registryId.
	// +kubebuilder:validation:Optional
	RegistryIDRef *v1.Reference `json:"registryIdRef,omitempty" tf:"-"`

	// Selector for a Registry in containerregistry to populate registryId.
	// +kubebuilder:validation:Optional
	RegistryIDSelector *v1.Selector `json:"registryIdSelector,omitempty" tf:"-"`

	// [string] Saves token password to file. Only works on create. Takes as argument a file name, or a file path
	// Saves password to file. Only works on create. Takes as argument a file name, or a file path
	SavePasswordToFile *string `json:"savePasswordToFile,omitempty" tf:"save_password_to_file,omitempty"`

	// [map]
	Scopes []ScopesInitParameters `json:"scopes,omitempty" tf:"scopes,omitempty"`

	// [string] Must have on of the values: enabled, disabled
	// Can be one of enabled, disabled
	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type RegistryTokenObservation struct {
	Credentials []CredentialsObservation `json:"credentials,omitempty" tf:"credentials,omitempty"`

	ExpiryDate *string `json:"expiryDate,omitempty" tf:"expiry_date,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// [string] The name of the container registry token. Immutable, update forces re-creation of the resource.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	RegistryID *string `json:"registryId,omitempty" tf:"registry_id,omitempty"`

	// [string] Saves token password to file. Only works on create. Takes as argument a file name, or a file path
	// Saves password to file. Only works on create. Takes as argument a file name, or a file path
	SavePasswordToFile *string `json:"savePasswordToFile,omitempty" tf:"save_password_to_file,omitempty"`

	// [map]
	Scopes []ScopesObservation `json:"scopes,omitempty" tf:"scopes,omitempty"`

	// [string] Must have on of the values: enabled, disabled
	// Can be one of enabled, disabled
	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type RegistryTokenParameters struct {

	// +kubebuilder:validation:Optional
	ExpiryDate *string `json:"expiryDate,omitempty" tf:"expiry_date,omitempty"`

	// [string] The name of the container registry token. Immutable, update forces re-creation of the resource.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// +crossplane:generate:reference:type=github.com/ionos-cloud/provider-upjet-ionoscloud/apis/containerregistry/v1alpha1.Registry
	// +kubebuilder:validation:Optional
	RegistryID *string `json:"registryId,omitempty" tf:"registry_id,omitempty"`

	// Reference to a Registry in containerregistry to populate registryId.
	// +kubebuilder:validation:Optional
	RegistryIDRef *v1.Reference `json:"registryIdRef,omitempty" tf:"-"`

	// Selector for a Registry in containerregistry to populate registryId.
	// +kubebuilder:validation:Optional
	RegistryIDSelector *v1.Selector `json:"registryIdSelector,omitempty" tf:"-"`

	// [string] Saves token password to file. Only works on create. Takes as argument a file name, or a file path
	// Saves password to file. Only works on create. Takes as argument a file name, or a file path
	// +kubebuilder:validation:Optional
	SavePasswordToFile *string `json:"savePasswordToFile,omitempty" tf:"save_password_to_file,omitempty"`

	// [map]
	// +kubebuilder:validation:Optional
	Scopes []ScopesParameters `json:"scopes,omitempty" tf:"scopes,omitempty"`

	// [string] Must have on of the values: enabled, disabled
	// Can be one of enabled, disabled
	// +kubebuilder:validation:Optional
	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type ScopesInitParameters struct {

	// [string] Example: ["pull", "push", "delete"]
	// Example: ["pull", "push", "delete"]
	Actions []*string `json:"actions,omitempty" tf:"actions,omitempty"`

	// [string] The name of the container registry token. Immutable, update forces re-creation of the resource.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// [string]
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type ScopesObservation struct {

	// [string] Example: ["pull", "push", "delete"]
	// Example: ["pull", "push", "delete"]
	Actions []*string `json:"actions,omitempty" tf:"actions,omitempty"`

	// [string] The name of the container registry token. Immutable, update forces re-creation of the resource.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// [string]
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type ScopesParameters struct {

	// [string] Example: ["pull", "push", "delete"]
	// Example: ["pull", "push", "delete"]
	// +kubebuilder:validation:Optional
	Actions []*string `json:"actions" tf:"actions,omitempty"`

	// [string] The name of the container registry token. Immutable, update forces re-creation of the resource.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// [string]
	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`
}

// RegistryTokenSpec defines the desired state of RegistryToken
type RegistryTokenSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RegistryTokenParameters `json:"forProvider"`
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
	InitProvider RegistryTokenInitParameters `json:"initProvider,omitempty"`
}

// RegistryTokenStatus defines the observed state of RegistryToken.
type RegistryTokenStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RegistryTokenObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// RegistryToken is the Schema for the RegistryTokens API. Creates and manages IonosCloud Container Registry Token.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,ionos}
type RegistryToken struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   RegistryTokenSpec   `json:"spec"`
	Status RegistryTokenStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RegistryTokenList contains a list of RegistryTokens
type RegistryTokenList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RegistryToken `json:"items"`
}

// Repository type metadata.
var (
	RegistryToken_Kind             = "RegistryToken"
	RegistryToken_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RegistryToken_Kind}.String()
	RegistryToken_KindAPIVersion   = RegistryToken_Kind + "." + CRDGroupVersion.String()
	RegistryToken_GroupVersionKind = CRDGroupVersion.WithKind(RegistryToken_Kind)
)

func init() {
	SchemeBuilder.Register(&RegistryToken{}, &RegistryTokenList{})
}